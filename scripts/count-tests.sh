#!/bin/bash

# GoTime Test Counter Script
# Counts all test functions and test executions in the GoTime library

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

echo -e "${BLUE}=== GoTime Test Counter ===${NC}"
echo

# Function to count tests in a package
count_package_tests() {
    local pkg_path="$1"
    local pkg_name="$2"

    echo -e "${YELLOW}Analyzing $pkg_name...${NC}"

    if [ -d "$pkg_path" ] || [ "$pkg_path" = "." ]; then
        cd "$PROJECT_ROOT"
        if [ "$pkg_path" != "." ]; then
            cd "$pkg_path"
        fi

        if ls *_test.go 1> /dev/null 2>&1; then
            # Count test functions
            test_functions=$(grep -c "^func Test\|^func Example\|^func Benchmark" *_test.go 2>/dev/null || echo "0")
            echo "  Test functions: $test_functions"

            # Count test executions
            test_executions=$(go test -v . 2>&1 | grep -c "=== RUN" || echo "0")
            echo "  Test executions: $test_executions"

            # Count each type separately
            test_count=$(grep -c "^func Test" *_test.go 2>/dev/null || echo "0")
            example_count=$(grep -c "^func Example" *_test.go 2>/dev/null || echo "0")
            benchmark_count=$(grep -c "^func Benchmark" *_test.go 2>/dev/null || echo "0")

            echo "    - Test functions: $test_count"
            echo "    - Example functions: $example_count"
            echo "    - Benchmark functions: $benchmark_count"

            # List test files
            echo "  Test files:"
            for file in *_test.go; do
                if [ -f "$file" ]; then
                    local_tests=$(grep -c "^func Test\|^func Example\|^func Benchmark" "$file" 2>/dev/null || echo "0")
                    echo "    - $file ($local_tests functions)"
                fi
            done

            cd "$PROJECT_ROOT"
            echo "  Package totals: $test_functions functions, $test_executions executions"
            return $test_functions
        else
            echo "  No test files found"
            cd "$PROJECT_ROOT"
            return 0
        fi
    else
        echo "  Package not found: $pkg_path"
        return 0
    fi
    echo
}

echo -e "${BLUE}=== Package Analysis ===${NC}"

# Initialize counters
total_functions=0
total_executions=0

# Count main package tests
echo -e "${YELLOW}Main Package:${NC}"
count_package_tests "." "Main Package"
main_functions=$?
main_executions=$(go test -v . 2>&1 | grep -c "=== RUN" || echo "0")
total_functions=$((total_functions + main_functions))
total_executions=$((total_executions + main_executions))
echo

# Count internal package tests
echo -e "${YELLOW}Internal Packages:${NC}"

# Cache package
if [ -d "internal/cache" ]; then
    count_package_tests "internal/cache" "Cache Package"
    cache_functions=$?
    cache_executions=$(cd "internal/cache" && go test -v . 2>&1 | grep -c "=== RUN" || echo "0")
    total_functions=$((total_functions + cache_functions))
    total_executions=$((total_executions + cache_executions))
    echo
fi

# NITES package
if [ -d "internal/nites" ]; then
    count_package_tests "internal/nites" "NITES Package"
    nites_functions=$?
    nites_executions=$(cd "internal/nites" && go test -v . 2>&1 | grep -c "=== RUN" || echo "0")
    total_functions=$((total_functions + nites_functions))
    total_executions=$((total_executions + nites_executions))
    echo
fi

# Utils package
if [ -d "internal/utils" ]; then
    count_package_tests "internal/utils" "Utils Package"
    utils_functions=$?
    utils_executions=$(cd "internal/utils" && go test -v . 2>&1 | grep -c "=== RUN" 2>/dev/null || echo "0")
    total_functions=$((total_functions + utils_functions))
    total_executions=$((total_executions + utils_executions))
    echo
fi

echo -e "${BLUE}=== Summary Statistics ===${NC}"
echo -e "${GREEN}Total Test Functions: $total_functions${NC}"
echo -e "${GREEN}Total Test Executions: $total_executions${NC}"
echo
echo "Breakdown by package:"
echo "  Main package: $main_functions functions, $main_executions executions"
if [ -d "internal/cache" ]; then
    echo "  Cache package: $cache_functions functions, $cache_executions executions"
fi
if [ -d "internal/nites" ]; then
    echo "  NITES package: $nites_functions functions, $nites_executions executions"
fi
if [ -d "internal/utils" ]; then
    echo "  Utils package: $utils_functions functions, $utils_executions executions"
fi

echo
echo -e "${BLUE}=== Verification with go test ./... ===${NC}"
combined_executions=$(go test -v ./... 2>&1 | grep -c "=== RUN" || echo "0")
echo "Combined test executions: $combined_executions"

if [ "$combined_executions" -eq "$total_executions" ]; then
    echo -e "${GREEN}✓ Execution counts match!${NC}"
else
    echo -e "${YELLOW}⚠ Execution count difference: $combined_executions vs $total_executions${NC}"
fi

echo
echo -e "${GREEN}=== Test Counting Complete ===${NC}"
