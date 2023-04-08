# temporal (WIP)

A golang library for parsing and parsing, formatting and processing dates and times using simple human-friendly formats such as `yesterday`, `tomorrow`, `dd/mm/yyyy`, etc. The temporal does not aims to be a replacement for the standard time package, but rather addtional facilities to make regular date and time operations such as formatting, parsing, relative time, data range, etc. easier.

It does not depend on any third-party libraries and is fully compatible with Tinygo. It uses caching to improve performance and reduce allocations.

💯 **100% test coverage** 💯

✔️ Tinygo Compatible

## Installation

Installation is simple. Just run the following command in your terminal to install the temporal package in your project.

```sh
go get github.com/maniartech/temporal
```

## Usage

The following example shows how to use the temporal package.

```go

import "github.com/maniartech/temporal"

// Parse a date
t, err := temporal.Parse("tomorrow")

// Format a date
s := temporal.Format(t, "yyyy-mm-dd")

// Convert date string to different format
s, err := temporal.Convert("2012-01-01", "yyyy-mm-dd", "dd/mm/yyyy")

// Get the relative date range
start, end, err := temporal.RelativeRange("last-week")

temporal.Today()
temporal.EoD()
temporal.Yesterday()
temporal.Tomorrow()
temporal.LastWeek()
temporal.LastMonth()
temporal.LastYear()
temporal.NextWeek()
temporal.NextMonth()
temporal.NextYear()
```

## Date time formatting

Temporal support simple human friendly date time formatting. The following table shows the supported formats. Temporal internally uses the time.Time.Format() and converts the human friendly format to the time.Time format. For example,
it converts `yyyy-mm-dd` to `2006-01-02` and then uses `time.Time.Format()` to format the date.

### Date

| Format | Description | Example |
|--------|-------------|---------|
| yy     | Two digit year with leading zero | 06 |
| yyyy   | Four digit year | 2006 |
| m      | Month without leading zero | 1 |
| mm     | Month in two digits with leading zero | 01 |
| mmm    | Month in short name | Jan |
| mmmm   | Month in full name | January |
| d      | Day without leading zero | 2 |
| dd     | Day in two digits with leading zero | 02 |
| ddd    | Zero padded day of year | 002 |
| www    | Three letter weekday name | Mon |
| wwww   | Full weekday name | Monday |

### Time

| Format | Description | Example |
|--------|-------------|---------|
| h      | Hour in 12 hour format without leading zero | 3 |
| hh     | Hour in 12 hour format with leading zero | 03 |
| hhh    | Hour in 24 hour format with leading zero | 15 |
| a      | am/pm | pm |
| aa     | AM/PM | PM |
| i      | Minute without leading zero | 4 |
| ii     | Minute with leading zero | 04 |
| s      | Second without leading zero | 5 |
| ss     | Second with leading zero | 05 |
| u      | Microsecond | 000000 |

### Timezone

| Format | Description | Example |
|--------|-------------|---------|
| z      | UTC offset | ±07:00 |
| zz     | UTC offset with colon | ±07:00 |
| zh     | Numeric timezone hour with hours only | ±07 |
| zzz    | Timezone abbreviation | MST |
| zzzz   | Timezone in long format | GMT-07:00 |
