# GoTime 🕐🕑🕒

> **Intuitive time manipulation for Go** - Making date/time operations as simple as they should be.

GoTime extends Go's standard `time` package with human-friendly operations that developers actually need in real-world applications. It leverages Go's powerful time handling while providing an intuitive API that makes working with dates and times a breeze.

## Quick Example

```go
import "github.com/maniartech/gotime"

// Intuitive formatting - no more "2006-01-02"!
formatted := gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss")
// → "2025-07-07 14:30:45"

// Easy format conversion
converted, _ := gotime.Convert("07/07/2025", "mm/dd/yyyy", "mmmm dt, yyyy")
// → "July 7th, 2025"

// Human-readable relative time
timeAgo := gotime.TimeAgo(time.Now().Add(-5 * time.Minute))
// → "5 minutes ago"

// Simple date arithmetic
nextBusinessDay := gotime.WorkDay(1, time.Now())
tenDaysFromNow := gotime.Days(10, time.Now())
```

## Why GoTime?

**The Problem:** Go's time formatting is cryptic and error-prone
```go
// Standard Go - Who remembers this?
time.Now().Format("2006-01-02 15:04:05")  // 😵‍💫
```

**The Solution:** Human-readable format specifiers
```go
// GoTime - Intuitive and memorable
gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss")  // 😊
```

### Key Benefits

| What You Get | Standard Go | GoTime |
|-------------|-------------|---------|
| **Format Clarity** | `2006-01-02` | `yyyy-mm-dd` |
| **Case Sensitivity** | `MM` vs `mm` matters | Case-insensitive |
| **Relative Time** | 15+ lines of code | `TimeAgo()` |
| **Format Conversion** | Parse + Format | `Convert()` |
| **Date Arithmetic** | Complex calculations | `WorkDay()`, `Days()` |

### Perfect For

- **Web APIs** - Consistent date formatting across endpoints
- **Reports** - Human-readable timestamps and date ranges
- **Business Logic** - Working day calculations, relative dates
- **Data Processing** - Converting between date formats
- **User Interfaces** - "2 hours ago" style timestamps

## Installation

```bash
go get github.com/maniartech/gotime
```

**Requirements:** Go 1.13+ • Zero dependencies • TinyGo compatible

## Core Features

### Intuitive Formatting
```go
// Remember yyyy-mm-dd, not 2006-01-02
formatted := gotime.Format(time.Now(), "mmmm dt, yyyy")
// → "July 7th, 2025"
```

### Smart Date Parsing
```go
date, _ := gotime.Parse("2025-07-07", "yyyy-mm-dd")
date, _ := gotime.Parse("07/07/2025", "mm/dd/yyyy")
```

### One-Line Format Conversion
```go
iso, _ := gotime.Convert("07/07/2025", "mm/dd/yyyy", "yyyy-mm-dd")
// → "2025-07-07"
```

### Human-Friendly Relative Time
```go
gotime.TimeAgo(fiveMinutesAgo)  // → "5 minutes ago"
gotime.TimeAgo(nextWeek)        // → "Next week"
```

### Business Date Calculations
```go
nextBusinessDay := gotime.WorkDay(1, time.Now())
businessDaysCount := gotime.NetWorkDays(startDate, endDate)
```

## Documentation

**[Complete Documentation](docs/)**

### Quick Links
- **[5-Minute Quick Start](docs/getting-started/quick-start.md)** - Get productive immediately
- **[Why GoTime?](docs/core-concepts/why-gotime.md)** - Detailed comparison with standard library
- **[Format Specifiers](docs/core-concepts/idfs.md)** - Complete IDFS reference
- **[API Reference](docs/api-reference/)** - All functions with examples
- **[Real-World Examples](docs/examples/common-use-cases.md)** - Web APIs, databases, business logic

### By Use Case
- **[Web Development](docs/examples/common-use-cases.md#web-development)** - APIs, user timestamps
- **[Database Integration](docs/examples/common-use-cases.md#database-integration)** - Custom types, queries
- **[Business Applications](docs/examples/common-use-cases.md#business-applications)** - Invoices, schedules
- **[Analytics](docs/examples/common-use-cases.md#analytics-and-reporting)** - Time-series, monitoring

## Quality & Reliability

- ✓ **100% Test Coverage** - Every function thoroughly tested
- ✓ **Production Ready** - Used in real-world applications
- ✓ **Zero Dependencies** - Only uses Go standard library
- ✓ **TinyGo Compatible** - Works in embedded and WebAssembly
- ✓ **MIT Licensed** - Free for commercial use

## Contributing

Contributions are welcome! Please ensure your code includes tests and follows existing patterns.

**Found an issue?** [Open an issue](https://github.com/maniartech/gotime/issues)
**Want to contribute?** See our [contribution guidelines](CONTRIBUTING.md)

## License

MIT Licensed - see [LICENSE](LICENSE) for details.

---

**[Browse Documentation](docs/)** | **[Star on GitHub](https://github.com/maniartech/gotime)** | **[Report Issues](https://github.com/maniartech/gotime/issues)**
