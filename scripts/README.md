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

### `release.sh`
One-pass release script: validates the version and its release notes, runs the
quality gates, then tags, pushes, and publishes the GitHub release.

The version is passed **without** the leading `v` (e.g. `2.0.4`); the script
derives the `v2.0.4` tag. It refuses to release unless `docs/releases/v<version>.md`
exists and the install commands in `README.md`/`docs/` already pin the new version.
Run it with no arguments for guided next-version suggestions.

**Usage:**
```bash
# Show guidance and suggested next versions (patch/minor/major)
./scripts/release.sh

# Validate everything without making changes
./scripts/release.sh 2.0.4 --dry-run

# Perform the release
./scripts/release.sh 2.0.4
```

**Options:**
- `--dry-run`: Run all checks and print the plan, but make no changes
- `--no-push`: Create the tag locally but do not push branch or tag
- `--no-github`: Do not create a GitHub release (via `gh`)
- `--race`: Include `-race` in the test gate
- `--skip-tests`: Skip the build/vet/test gates (not recommended)
- `--allow-dirty`: Proceed even with uncommitted changes
- `--allow-branch`: Proceed even if not on master/main
- `--remote NAME`: Git remote to push to (default: `origin`)
- `-y, --yes`: Skip the confirmation prompt
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
