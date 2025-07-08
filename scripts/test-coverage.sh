#!/bin/bash

# GoTime Test Coverage Script
# Comprehensive testing and coverage analysis for the GoTime library

set -e

echo "=== GoTime Test Coverage Analysis ==="
echo

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get project root directory
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PROJECT_ROOT"

echo -e "${BLUE}Project Root:${NC} $PROJECT_ROOT"
echo

# Function to run tests for a package
run_package_tests() {
    local package_path="$1"
    local package_name="$2"

    echo -e "${YELLOW}Testing $package_name...${NC}"

    if [ -d "$package_path" ]; then
        cd "$package_path"

        # Check if test files exist
        if ls *_test.go 1> /dev/null 2>&1; then
            # Count test functions
            test_count=$(grep -c "^func Test\|^func Example\|^func Benchmark" *_test.go 2>/dev/null || echo "0")
            echo "  Test functions: $test_count"

            # Run tests with coverage
            if go test -v -cover -coverprofile=coverage.out .; then
                echo -e "  ${GREEN}✓ Tests passed${NC}"

                # Show coverage percentage
                coverage=$(go tool cover -func=coverage.out | tail -1 | awk '{print $3}')
                echo "  Coverage: $coverage"

                # Count test executions
                executions=$(go test -v . 2>&1 | grep -c "=== RUN" || echo "0")
                echo "  Test executions: $executions"
            else
                echo -e "  ${RED}✗ Tests failed${NC}"
                return 1
            fi
        else
            echo "  No test files found"
        fi

        cd "$PROJECT_ROOT"
    else
        echo "  Package not found: $package_path"
    fi
    echo
}

# Run tests for main package
echo -e "${BLUE}=== Main Package Tests ===${NC}"
run_package_tests "." "Main Package"

# Run tests for internal packages
echo -e "${BLUE}=== Internal Package Tests ===${NC}"

# Cache package
run_package_tests "internal/cache" "Cache Package"

# NITES package
run_package_tests "internal/nites" "NITES Package"

# Utils package (if it has tests)
run_package_tests "internal/utils" "Utils Package"

echo -e "${BLUE}=== Combined Coverage Analysis ===${NC}"

# Generate combined coverage report
echo "Generating combined coverage report..."

# Run all tests with coverage
go test -v -cover -coverprofile=coverage.out ./...

if [ -f "coverage.out" ]; then
    echo
    echo -e "${GREEN}Combined Coverage Report:${NC}"
    go tool cover -func=coverage.out | tail -1

    echo
    echo "Detailed coverage by file:"
    go tool cover -func=coverage.out | head -20

    # Generate HTML coverage report
    echo
    echo "Generating HTML coverage report..."
    go tool cover -html=coverage.out -o coverage.html
    echo -e "${GREEN}HTML coverage report generated: coverage.html${NC}"
fi

echo
echo -e "${GREEN}=== Test Coverage Analysis Complete ===${NC}"

# Summary statistics
echo
echo -e "${BLUE}Summary Statistics:${NC}"
total_test_functions=$(find . -name "*_test.go" -exec grep -c "^func Test\|^func Example\|^func Benchmark" {} \; | awk '{sum+=$1} END {print sum}')
total_executions=$(go test -v ./... 2>&1 | grep -c "=== RUN" || echo "0")

echo "Total test functions: $total_test_functions"
echo "Total test executions: $total_executions"
echo "Main package functions: $(grep -c "^func Test\|^func Example\|^func Benchmark" *_test.go 2>/dev/null || echo "0")"
echo "Internal package functions: $((total_test_functions - $(grep -c "^func Test\|^func Example\|^func Benchmark" *_test.go 2>/dev/null || echo "0")))"
