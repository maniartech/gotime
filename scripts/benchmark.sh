#!/bin/bash

# GoTime Quick Benchmark Script
# Runs benchmark tests for performance analysis

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

echo -e "${BLUE}=== GoTime Benchmark Runner ===${NC}"
echo

# Parse command line arguments
BENCHTIME="1s"
COUNT="3"
PACKAGE="all"
MEMPROFILE=""
CPUPROFILE=""

while [[ $# -gt 0 ]]; do
    case $1 in
        -t|--benchtime)
            BENCHTIME="$2"
            shift 2
            ;;
        -c|--count)
            COUNT="$2"
            shift 2
            ;;
        -p|--package)
            PACKAGE="$2"
            shift 2
            ;;
        --memprofile)
            MEMPROFILE="$2"
            shift 2
            ;;
        --cpuprofile)
            CPUPROFILE="$2"
            shift 2
            ;;
        -h|--help)
            echo "Usage: $0 [OPTIONS]"
            echo "Options:"
            echo "  -t, --benchtime  Benchmark duration (default: 1s)"
            echo "  -c, --count      Number of benchmark runs (default: 3)"
            echo "  -p, --package    Package to benchmark (main, all)"
            echo "  --memprofile     Generate memory profile"
            echo "  --cpuprofile     Generate CPU profile"
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

# Build benchmark command
BENCH_CMD="go test -bench=."
BENCH_CMD="$BENCH_CMD -benchtime=$BENCHTIME"
BENCH_CMD="$BENCH_CMD -count=$COUNT"

if [ -n "$MEMPROFILE" ]; then
    BENCH_CMD="$BENCH_CMD -memprofile=$MEMPROFILE"
fi

if [ -n "$CPUPROFILE" ]; then
    BENCH_CMD="$BENCH_CMD -cpuprofile=$CPUPROFILE"
fi

# Function to run benchmarks for a package
run_benchmarks() {
    local pkg_path="$1"
    local pkg_name="$2"

    echo -e "${YELLOW}Running benchmarks for $pkg_name...${NC}"

    if [ -d "$pkg_path" ] || [ "$pkg_path" = "." ]; then
        cd "$PROJECT_ROOT"
        if [ "$pkg_path" != "." ]; then
            cd "$pkg_path"
        fi

        # Check if benchmark functions exist
        if grep -q "^func Benchmark" *_test.go 2>/dev/null; then
            echo "Command: $BENCH_CMD ."
            eval "$BENCH_CMD ."
            echo -e "${GREEN}✓ $pkg_name benchmarks completed${NC}"
        else
            echo "No benchmark functions found in $pkg_name"
        fi

        cd "$PROJECT_ROOT"
    else
        echo "Package not found: $pkg_path"
    fi
    echo
}

# Run benchmarks based on package selection
case "$PACKAGE" in
    "main")
        run_benchmarks "." "Main Package"
        ;;
    "all")
        echo -e "${BLUE}Running all benchmarks...${NC}"
        echo

        # Main package
        run_benchmarks "." "Main Package"

        # Internal packages (if they have benchmarks)
        if [ -d "internal/cache" ]; then
            run_benchmarks "internal/cache" "Cache Package"
        fi

        if [ -d "internal/nites" ]; then
            run_benchmarks "internal/nites" "NITES Package"
        fi
        ;;
    *)
        echo "Unknown package: $PACKAGE"
        echo "Available packages: main, all"
        exit 1
        ;;
esac

echo
echo -e "${GREEN}=== Benchmark Run Complete ===${NC}"

# Show profile information if generated
if [ -n "$MEMPROFILE" ] && [ -f "$MEMPROFILE" ]; then
    echo "Memory profile generated: $MEMPROFILE"
    echo "View with: go tool pprof $MEMPROFILE"
fi

if [ -n "$CPUPROFILE" ] && [ -f "$CPUPROFILE" ]; then
    echo "CPU profile generated: $CPUPROFILE"
    echo "View with: go tool pprof $CPUPROFILE"
fi
