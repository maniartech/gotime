# GoTime Scripts

This directory contains utility scripts for testing, coverage analysis, and benchmarking the GoTime library.

## Available Scripts

### `test-runner.sh`
Fast test execution script with flexible options.

**Usage:**
```bash
# Run all tests
./scripts/test-runner.sh -p all

# Run main package tests with verbose output
./scripts/test-runner.sh -v

# Run specific package tests with coverage
./scripts/test-runner.sh -p cache -c

# Run benchmarks
./scripts/test-runner.sh -b

# Show help
./scripts/test-runner.sh -h
```

**Options:**
- `-v, --verbose`: Run tests with verbose output
- `-c, --coverage`: Run tests with coverage analysis
- `-b, --benchmark`: Run benchmark tests
- `-p, --package`: Run tests for specific package (main, cache, nites, utils, all)
- `-h, --help`: Show help message

### `test-coverage.sh`
Comprehensive test coverage analysis script.

**Usage:**
```bash
# Run full coverage analysis
./scripts/test-coverage.sh
```

**Features:**
- Tests all packages (main + internal)
- Generates combined coverage report
- Creates HTML coverage report (coverage.html)
- Shows detailed coverage statistics
- Counts test functions and executions

### `count-tests.sh`
Test counting and analysis script.

**Usage:**
```bash
# Count all tests and show statistics
./scripts/count-tests.sh
```

**Features:**
- Counts test functions in all packages
- Shows test executions per package
- Breaks down by test type (Test, Example, Benchmark)
- Lists all test files
- Verifies counts with combined test run

### `benchmark.sh`
Focused benchmark execution script.

**Usage:**
```bash
# Run all benchmarks
./scripts/benchmark.sh

# Run benchmarks with custom duration and count
./scripts/benchmark.sh -t 5s -c 5

# Generate memory profile
./scripts/benchmark.sh --memprofile mem.prof

# Generate CPU profile
./scripts/benchmark.sh --cpuprofile cpu.prof

# Show help
./scripts/benchmark.sh -h
```

**Options:**
- `-t, --benchtime`: Benchmark duration (default: 1s)
- `-c, --count`: Number of benchmark runs (default: 3)
- `-p, --package`: Package to benchmark (main, all)
- `--memprofile`: Generate memory profile
- `--cpuprofile`: Generate CPU profile
- `-h, --help`: Show help message

## Test Statistics

The GoTime library has comprehensive test coverage with:

- **Total test executions**: 238 across all packages
  - Main package: 225 test executions
  - Internal cache package: 2 test executions
  - Internal nites package: 11 test executions
- **Test coverage**: 100% across all packages
- **Test types**: Unit tests, table-driven tests, examples, and benchmarks

## Quick Commands

```bash
# Quick test run
./scripts/test-runner.sh

# Full coverage analysis
./scripts/test-coverage.sh

# Count all tests
./scripts/count-tests.sh

# Run benchmarks
./scripts/benchmark.sh

# Run specific package tests
./scripts/test-runner.sh -p nites -v

# Generate coverage with HTML report
./scripts/test-runner.sh -c -p all
```

## Windows Usage

On Windows with bash.exe, all scripts work as intended:

```bash
# From project root
bash scripts/test-runner.sh -h
bash scripts/test-coverage.sh
bash scripts/count-tests.sh
bash scripts/benchmark.sh -t 3s
```
