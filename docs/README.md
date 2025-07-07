# GoTime Documentation

Welcome to the comprehensive documentation for GoTime - a powerful, intuitive time manipulation library for Go that makes working with dates and times as simple as it should be.

## Quick Start

New to GoTime? Start here:

```go
import "github.com/maniartech/gotime"

// Format dates intuitively
formatted := gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss")

// Parse dates easily
date, err := gotime.Parse("2025-07-07", "yyyy-mm-dd")

// Convert between formats
converted, err := gotime.Convert("07/07/2025", "mm/dd/yyyy", "yyyy-mm-dd")

// Get relative time
timeAgo := gotime.TimeAgo(time.Now().Add(-5 * time.Minute)) // "5 minutes ago"

// Date calculations
tomorrow := gotime.Tomorrow()
nextWeek := gotime.NextWeek()
tenDaysFromNow := gotime.Days(10)
```

**[→ Full Quick Start Guide](getting-started/quick-start.md)**

## Documentation Sections

### Getting Started
Perfect for newcomers to GoTime:
- **[Installation](getting-started/installation.md)** - Get GoTime installed and verified
- **[Quick Start Guide](getting-started/quick-start.md)** - Essential operations in 5 minutes
- **[Basic Usage](getting-started/basic-usage.md)** - Comprehensive usage patterns

### Core Concepts
Understand what makes GoTime powerful:
- **[Why GoTime?](core-concepts/why-gotime.md)** - Problems solved and design decisions
- **[Intuitive Date Format Specifiers (IDFS)](core-concepts/idfs.md)** - The heart of GoTime's simplicity

### API Reference
Complete function documentation:
- **[Date Parsing and Formatting](api-reference/parsing-formatting.md)** - Parse, Format, FormatTimestamp
- **[Date Conversion](api-reference/conversion.md)** - Convert between any date formats
- **[Relative Time Functions](api-reference/relative-time.md)** - TimeAgo, Days, Weeks, Months, Years
- **[Time Calculations](api-reference/time-calculations.md)** - WorkDay, Latest, Earliest, Diff
- **[Date Range Operations](api-reference/date-ranges.md)** - IsBetween, range validation
- **[Utility Functions](api-reference/utilities.md)** - IsLeapYear, DaysInMonth, date construction

### Examples
Real-world applications and patterns:
- **[Common Use Cases](examples/common-use-cases.md)** - Web APIs, databases, business logic
- **[Real-world Examples](examples/real-world.md)** - Complete applications
- **[Migration Guide](examples/migration.md)** - Moving from standard library

### Advanced Topics
Optimization and best practices:
- **[Performance Considerations](advanced/performance.md)** - Optimization techniques
- **[Best Practices](advanced/best-practices.md)** - Production-ready patterns
- **[Working with Timezones](advanced/timezones.md)** - Global application support

## Key Features

### Intuitive Format Specifiers (IDFS)
Replace Go's cryptic `2006-01-02` with readable `yyyy-mm-dd`:

```go
// Standard Go - hard to remember
time.Now().Format("2006-01-02 15:04:05")

// GoTime - intuitive and memorable
gotime.Format(time.Now(), "yyyy-mm-dd hh:ii:ss")
```

### Seamless Format Conversion
Transform dates between any formats effortlessly:

```go
// One-line format conversion
iso, _ := gotime.Convert("07/07/2025", "mm/dd/yyyy", "yyyy-mm-dd")
readable, _ := gotime.Convert("2025-07-07", "yyyy-mm-dd", "mmmm dt, yyyy")
```

### Smart Relative Time
Human-friendly time descriptions:

```go
gotime.TimeAgo(fiveMinutesAgo)  // "5 minutes ago"
gotime.TimeAgo(tomorrow)        // "Tomorrow"
gotime.TimeAgo(lastWeek)        // "Last week"
```

### Powerful Date Arithmetic
Intuitive date calculations:

```go
// Simple and readable
nextBusinessDay := gotime.WorkDay(time.Now(), 1, workingDays)
quarterStart := gotime.MonthStart(gotime.Months(-2))
weekEnd := gotime.WeekEnd()
```

## Why Choose GoTime?

| Challenge | Standard Go | GoTime |
|-----------|-------------|---------|
| **Format Clarity** | `2006-01-02` | `yyyy-mm-dd` |
| **Case Sensitivity** | `MM` vs `mm` matters | Case-insensitive |
| **Common Operations** | 5-10 lines | 1 line |
| **Relative Time** | Manual implementation | Built-in |
| **Learning Curve** | High | Low |
| **Error Prone** | Yes | No |

## Format Specifier Quick Reference

| Format | Output | Description |
|--------|--------|-------------|
| `yyyy` | `2025` | 4-digit year |
| `mm` | `07` | 2-digit month |
| `dd` | `07` | 2-digit day |
| `hh` | `14` | 24-hour format |
| `ii` | `30` | Minutes |
| `ss` | `45` | Seconds |
| `mmmm` | `July` | Full month name |
| `wwww` | `Monday` | Full weekday name |
| `dt` | `7th` | Day with ordinal suffix |

**[→ Complete IDFS Reference](core-concepts/idfs.md)**

## Quick Links by Use Case

### Web Development
- [API Date Handling](examples/common-use-cases.md#web-development)
- [User-Friendly Timestamps](examples/common-use-cases.md#user-friendly-timestamps)
- [Form Validation](examples/common-use-cases.md#form-date-validation)

### Database Integration
- [Custom Database Types](examples/common-use-cases.md#database-integration)
- [Query Builders](examples/common-use-cases.md#query-builder-with-date-filters)

### Business Applications
- [Invoice Management](examples/common-use-cases.md#business-applications)
- [Employee Scheduling](examples/common-use-cases.md#employee-schedule-management)
- [Project Timelines](examples/common-use-cases.md#project-timeline-management)

### Analytics & Reporting
- [Time-Series Analysis](examples/common-use-cases.md#analytics-and-reporting)
- [Performance Monitoring](examples/common-use-cases.md#performance-monitoring)

## Installation

```bash
go get github.com/maniartech/gotime
```

**Requirements:**
- ✓ Go 1.13+
- ✓ Zero external dependencies
- ✓ TinyGo compatible

## Need Help?

- **New to GoTime?** → [Quick Start Guide](getting-started/quick-start.md)
- **Looking for examples?** → [Common Use Cases](examples/common-use-cases.md)
- **Need specific functions?** → [API Reference](api-reference/)
- **Performance questions?** → [Best Practices](advanced/best-practices.md)
- **Migration help?** → [Migration Guide](examples/migration.md)

## Quality Assurance

- ✓ **100% Test Coverage** - Every function thoroughly tested
- ✓ **Production Ready** - Used in real-world applications
- ✓ **Zero Dependencies** - No external package dependencies
- ✓ **TinyGo Compatible** - Works in embedded and WebAssembly
- ✓ **MIT Licensed** - Open source and commercial-friendly
- ✓ **Comprehensive Docs** - Every feature documented with examples

## Success Stories

GoTime is successfully used in:
- **Web APIs** for consistent date formatting
- **Financial systems** for invoice and payment processing
- **Analytics platforms** for time-series data processing
- **Content management** for publishing schedules
- **Project management** for timeline calculations

---

## Get Started Now

Ready to make date/time handling in Go intuitive and powerful?

1. **[Install GoTime](getting-started/installation.md)** - Get up and running
2. **[Quick Start](getting-started/quick-start.md)** - Learn the basics in 5 minutes
3. **[Explore Examples](examples/common-use-cases.md)** - See real-world applications
4. **[Master IDFS](core-concepts/idfs.md)** - Understand the format system

---

*GoTime - Making time manipulation in Go as intuitive as it should be.*

**[Repository](https://github.com/maniartech/gotime)** | **[Issues](https://github.com/maniartech/gotime/issues)** | **[License](https://github.com/maniartech/gotime/blob/main/LICENSE)**
