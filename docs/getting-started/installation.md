[Home](../README.md) > [Getting Started](README.md) > Installation

# Installation

GoTime is designed to be easy to install and integrate into your Go projects.

## Requirements

- **Go Version**: Go 1.13 or later
- **TinyGo**: Fully compatible with TinyGo
- **Dependencies**: Zero external dependencies

## Installation Methods

### 1. Using Go Modules (Recommended)

For projects using Go modules (Go 1.11+):

```bash
go get github.com/maniartech/gotime
```

This will automatically download and install the latest version of GoTime and make it available in your project.

### 2. Using GOPATH

For older Go versions or GOPATH-based projects:

```bash
go get -u github.com/maniartech/gotime
```

### 3. Specific Version

To install a specific version:

```bash
go get github.com/maniartech/gotime@v1.0.0
```

## Verify Installation

Create a simple test file to verify the installation:

```go
// test_installation.go
package main

import (
    "fmt"
    "time"
    "github.com/maniartech/gotime"
)

func main() {
    // Test basic functionality
    now := time.Now()
    formatted := gotime.Format(now, "yyyy-mm-dd hh:ii:ss")
    fmt.Printf("GoTime is working! Current time: %s\n", formatted)

    // Test parsing
    parsed, err := gotime.Parse("2024-01-01", "yyyy-mm-dd")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Parsed date: %s\n", parsed)
}
```

Run the test:

```bash
go run test_installation.go
```

Expected output:
```
GoTime is working! Current time: 2025-07-07 12:34:56
Parsed date: 2024-01-01 00:00:00 +0000 UTC
```

## Import Statement

Once installed, import GoTime in your Go files:

```go
import "github.com/maniartech/gotime"
```

### Alternative Import with Alias

If you prefer a shorter alias:

```go
import gt "github.com/maniartech/gotime"

// Usage
formatted := gt.Format(time.Now(), "yyyy-mm-dd")
```

## Project Setup

### 1. New Project

For a new project, initialize Go modules first:

```bash
mkdir my-time-app
cd my-time-app
go mod init my-time-app
go get github.com/maniartech/gotime
```

### 2. Existing Project

For an existing project with Go modules:

```bash
cd your-project
go get github.com/maniartech/gotime
```

The package will be added to your `go.mod` file automatically.

## IDE Setup

### VS Code

If you're using VS Code with the Go extension, it will automatically:
- Download dependencies when you save files
- Provide code completion for GoTime functions
- Show function documentation on hover

### Other IDEs

Most modern Go IDEs will automatically handle Go modules and provide:
- Auto-completion
- Import suggestions
- Documentation tooltips

## Troubleshooting

### Common Issues

#### 1. Module Not Found

**Error**: `module github.com/maniartech/gotime: not found`

**Solutions**:
- Ensure you have internet connectivity
- Check if you're behind a corporate firewall/proxy
- Try setting GOPROXY: `export GOPROXY=https://proxy.golang.org,direct`

#### 2. Version Conflicts

**Error**: Version conflicts with other dependencies

**Solutions**:
- Update to the latest version: `go get -u github.com/maniartech/gotime`
- Check `go.mod` for version constraints
- Use `go mod tidy` to clean up dependencies

#### 3. Import Issues

**Error**: Package not found in imports

**Solutions**:
- Ensure the import path is correct: `github.com/maniartech/gotime`
- Run `go mod download` to download dependencies
- Check that your Go version supports modules

### Getting Help

If you encounter issues:

1. Check the [GitHub Issues](https://github.com/maniartech/gotime/issues)
2. Create a new issue with:
   - Your Go version (`go version`)
   - Your OS and architecture
   - The error message
   - A minimal reproduction case

## What's Next?

Now that GoTime is installed:

1. **Start Learning**: Check out the [Quick Start Guide](quick-start.md)
2. **Explore Features**: Read about [Basic Usage](basic-usage.md)
3. **Understand NITES**: Learn about [Natural and Intuitive Time Expression Syntax](../core-concepts/nites.md)
4. **See Examples**: Browse [Common Use Cases](../examples/common-use-cases.md)

## Package Information

- **Repository**: https://github.com/maniartech/gotime
- **License**: MIT
- **Go Modules**: ✅ Supported
- **TinyGo**: ✅ Compatible
- **External Dependencies**: ❌ None
- **Test Coverage**: 💯% Coverage

---

Ready to start using GoTime? Head to the [Quick Start Guide](quick-start.md)!
