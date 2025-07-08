#!/bin/bash

# GoTime Test Runner Script
# Fast test execution for the GoTime library

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get project root directory
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PROJECT_ROOT"

echo -e "${BLUE}=== GoTime Test Runner ===${NC}"
echo

# Parse command line arguments
VERBOSE=false
COVERAGE=false
PACKAGE=""
BENCHMARK=false

while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        -c|--coverage)
            COVERAGE=true
            shift
            ;;
        -b|--benchmark)
            BENCHMARK=true
            shift
            ;;
        -p|--package)
            PACKAGE="$2"
            shift 2
            ;;
        -h|--help)
            echo "Usage: $0 [OPTIONS]"
            echo "Options:"
            echo "  -v, --verbose    Run tests with verbose output"
            echo "  -c, --coverage   Run tests with coverage analysis"
            echo "  -b, --benchmark  Run benchmark tests"
            echo "  -p, --package    Run tests for specific package (main, cache, idfs, utils, all)"
            echo "  -h, --help       Show this help message"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            echo "Use -h or --help for usage information"
            exit 1
            ;;
    esac
done

# Build test command
TEST_CMD="go test"

if [ "$VERBOSE" = true ]; then
    TEST_CMD="$TEST_CMD -v"
fi

if [ "$COVERAGE" = true ]; then
    TEST_CMD="$TEST_CMD -cover -coverprofile=coverage.out"
fi

if [ "$BENCHMARK" = true ]; then
    TEST_CMD="$TEST_CMD -bench=."
fi

# Function to run tests for a specific package
run_tests() {
    local pkg_path="$1"
    local pkg_name="$2"

    echo -e "${YELLOW}Running tests for $pkg_name...${NC}"

    if [ -d "$pkg_path" ] || [ "$pkg_path" = "." ]; then
        cd "$PROJECT_ROOT"
        if [ "$pkg_path" != "." ]; then
            cd "$pkg_path"
        fi

        # Check if test files exist
        if ls *_test.go 1> /dev/null 2>&1; then
            echo "Command: $TEST_CMD ."
            if eval "$TEST_CMD ."; then
                echo -e "${GREEN}✓ $pkg_name tests passed${NC}"

                # Show test statistics
                if [ "$VERBOSE" = true ]; then
                    test_count=$(grep -c "^func Test\|^func Example\|^func Benchmark" *_test.go 2>/dev/null || echo "0")
                    echo "Test functions: $test_count"
                fi
            else
                echo -e "${RED}✗ $pkg_name tests failed${NC}"
                return 1
            fi
        else
            echo "No test files found in $pkg_name"
        fi

        cd "$PROJECT_ROOT"
    else
        echo "Package not found: $pkg_path"
        return 1
    fi
    echo
}

# Run tests based on package selection
case "$PACKAGE" in
    "main"|"")
        run_tests "." "Main Package"
        ;;
    "cache")
        run_tests "internal/cache" "Cache Package"
        ;;
    "idfs")
        run_tests "internal/idfs" "IDFS Package"
        ;;
    "utils")
        run_tests "internal/utils" "Utils Package"
        ;;
    "all")
        echo -e "${BLUE}Running all tests...${NC}"
        cd "$PROJECT_ROOT"

        if [ "$COVERAGE" = true ]; then
            echo "Command: go test -v -cover -coverprofile=coverage.out ./..."
            go test -v -cover -coverprofile=coverage.out ./...
        else
            echo "Command: go test ./..."
            if [ "$VERBOSE" = true ]; then
                go test -v ./...
            else
                go test ./...
            fi
        fi

        if [ "$BENCHMARK" = true ]; then
            echo
            echo -e "${YELLOW}Running benchmarks...${NC}"
            go test -bench=. ./...
        fi
        ;;
    *)
        echo "Unknown package: $PACKAGE"
        echo "Available packages: main, cache, idfs, utils, all"
        exit 1
        ;;
esac

# Show coverage report if generated
if [ "$COVERAGE" = true ] && [ -f "coverage.out" ]; then
    echo
    echo -e "${GREEN}Coverage Summary:${NC}"
    go tool cover -func=coverage.out | tail -1
fi

echo
echo -e "${GREEN}=== Test Run Complete ===${NC}"