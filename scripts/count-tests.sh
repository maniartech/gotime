#!/bin/bash

# GoTime Test Counter Script
# Counts all test functions and test executions in the GoTime library

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
BOLD='\033[1m'
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
            test_functions=$(grep -h "^func Test\|^func Example\|^func Benchmark" *_test.go 2>/dev/null | wc -l || echo "0")
            echo "  Test functions: $test_functions"

            # Count test executions
            test_executions=$(go test -v . 2>&1 | grep -c "=== RUN" || echo "0")
            echo "  Test executions: $test_executions"

            # Count each type separately
            test_count=$(grep -h "^func Test" *_test.go 2>/dev/null | wc -l || echo "0")
            example_count=$(grep -h "^func Example" *_test.go 2>/dev/null | wc -l || echo "0")
            benchmark_count=$(grep -h "^func Benchmark" *_test.go 2>/dev/null | wc -l || echo "0")

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

            # Return both values using global variables
            PACKAGE_FUNCTIONS=$test_functions
            PACKAGE_EXECUTIONS=$test_executions
            return 0
        else
            echo "  No test files found"
            cd "$PROJECT_ROOT"
            PACKAGE_FUNCTIONS=0
            PACKAGE_EXECUTIONS=0
            return 0
        fi
    else
        echo "  Package not found: $pkg_path"
        PACKAGE_FUNCTIONS=0
        PACKAGE_EXECUTIONS=0
        return 0
    fi
    echo
}

echo -e "${BLUE}=== Package Analysis ===${NC}"

# Initialize counters
total_functions=0
total_executions=0

# Initialize global variables for package counting
PACKAGE_FUNCTIONS=0
PACKAGE_EXECUTIONS=0

# Count main package tests
echo -e "${YELLOW}Main Package:${NC}"
count_package_tests "." "Main Package"
main_functions=$PACKAGE_FUNCTIONS
main_executions=$PACKAGE_EXECUTIONS
total_functions=$((total_functions + main_functions))
total_executions=$((total_executions + main_executions))
echo

# Count internal package tests
echo -e "${YELLOW}Internal Packages:${NC}"

# Cache package
cache_functions=0
cache_executions=0
if [ -d "internal/cache" ]; then
    count_package_tests "internal/cache" "Cache Package"
    cache_functions=$PACKAGE_FUNCTIONS
    cache_executions=$PACKAGE_EXECUTIONS
    total_functions=$((total_functions + cache_functions))
    total_executions=$((total_executions + cache_executions))
    echo
fi

# NITES package
nites_functions=0
nites_executions=0
if [ -d "internal/nites" ]; then
    count_package_tests "internal/nites" "NITES Package"
    nites_functions=$PACKAGE_FUNCTIONS
    nites_executions=$PACKAGE_EXECUTIONS
    total_functions=$((total_functions + nites_functions))
    total_executions=$((total_executions + nites_executions))
    echo
fi

# Utils package
utils_functions=0
utils_executions=0
if [ -d "internal/utils" ]; then
    count_package_tests "internal/utils" "Utils Package"
    utils_functions=$PACKAGE_FUNCTIONS
    utils_executions=$PACKAGE_EXECUTIONS
    total_functions=$((total_functions + utils_functions))
    total_executions=$((total_executions + utils_executions))
    echo
fi

echo -e "${BLUE}=== Summary Statistics ===${NC}"
echo

# Calculate the exact padding needed for the box
functions_text="TOTAL TEST FUNCTIONS:  $(printf "%3d" $total_functions)"
executions_text="TOTAL TEST EXECUTIONS: $(printf "%3d" $total_executions)"

# Box width is 64 characters (including borders), so content width is 62
box_content_width=62

# Calculate padding for each line (account for the 2 spaces at start)
functions_padding=$((box_content_width - 2 - ${#functions_text}))
executions_padding=$((box_content_width - 2 - ${#executions_text}))

# Create padding strings
functions_spaces=$(printf "%*s" $functions_padding "")
executions_spaces=$(printf "%*s" $executions_padding "")

echo -e "${BOLD}${CYAN}╔════════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BOLD}${CYAN}║                          TEST SUMMARY                          ║${NC}"
echo -e "${BOLD}${CYAN}╠════════════════════════════════════════════════════════════════╣${NC}"
echo -e "${BOLD}${CYAN}║  ${BOLD}${GREEN}${functions_text}${BOLD}${CYAN}${functions_spaces}║${NC}"
echo -e "${BOLD}${CYAN}║  ${BOLD}${GREEN}${executions_text}${BOLD}${CYAN}${executions_spaces}║${NC}"
echo -e "${BOLD}${CYAN}╚════════════════════════════════════════════════════════════════╝${NC}"
echo
echo -e "${BOLD}${BLUE}Breakdown by package:${NC}"
echo -e "  ${BOLD}Main package:${NC}  $main_functions functions, $main_executions executions"
if [ -d "internal/cache" ]; then
    echo -e "  ${BOLD}Cache package:${NC} $cache_functions functions, $cache_executions executions"
fi
if [ -d "internal/nites" ]; then
    echo -e "  ${BOLD}NITES package:${NC} $nites_functions functions, $nites_executions executions"
fi
if [ -d "internal/utils" ]; then
    echo -e "  ${BOLD}Utils package:${NC} $utils_functions functions, $utils_executions executions"
fi

echo
echo -e "${BLUE}=== Verification with go test ./... ===${NC}"
combined_executions=$(go test -v ./... 2>&1 | grep -c "=== RUN" || echo "0")
echo "Combined test executions: $combined_executions"

if [ "$combined_executions" -eq "$total_executions" ]; then
    echo -e "${BOLD}${GREEN}✓ Execution counts match perfectly!${NC}"
else
    echo -e "${BOLD}${YELLOW}⚠ Execution count difference: $combined_executions vs $total_executions${NC}"
fi

echo
# Final summary box with proper padding (same width as first box: 64 chars)
final_summary_text="Total Functions: $(printf "%3d" $total_functions)  Total Executions: $(printf "%3d" $total_executions)"
final_padding=$((box_content_width - 2 - ${#final_summary_text}))
final_spaces=$(printf "%*s" $final_padding "")

echo -e "${BOLD}${GREEN}╔════════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BOLD}${GREEN}║                    TEST COUNTING COMPLETE                     ║${NC}"
echo -e "${BOLD}${GREEN}║  ${BOLD}${CYAN}${final_summary_text}${BOLD}${GREEN}${final_spaces}║${NC}"
echo -e "${BOLD}${GREEN}╚════════════════════════════════════════════════════════════════╝${NC}"
