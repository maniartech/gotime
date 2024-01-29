# temporal

A golang library for parsing and parsing, formatting and processing dates and
times using simple human-friendly formats such as `yesterday`, `tomorrow`,
`dd/mm/yyyy`, etc. The temporal does not aims to be a replacement for the
standard time package, but rather addtional facilities to make regular date
and time operations such as formatting, parsing, relative time, data range, etc.
easier.

## Why Temporal?

### ✔️ Designed for Practicality

It provides features are practical and useful in real-world applications. These
features are either missing or not easy to use in the standard time package.

- [x] Parse dates using simple human-friendly formats such as `dd/mm/yyyy`, etc.
- [x] Format dates using simple human-friendly formats such as `dd/mm/yyyy`, etc.
- [x] Convert dates from one format to another. For example, `dd/mm/yyyy` to
      `yyyy-mm-dd`.
- [x] Convert the datetime to relative time such as `1 hour ago`, `2 days ago`,
      etc.
- [x] Provides range of date finder functions such as `Yesterday()`,
      `Tomorrow()`, `SoD()`, `EoD()`, etc.
- [x] Provides range of date time utility functions such as `Latest()`,
      `Earliest()`, `IsBetween()`, `TruncateDate()`, etc.

### ✔️ Developer Friendly

It provides a comprehensive range of specifiers for all your date and time
formatting needs, making it an indispensable tool for Go developers.

- [x] 100% test coverage 👌🏼
- [x] TinyGo compatible 👌🏼
- [x] No external dependencies  👌🏼
- [x] Fully utilises the standard time package and does not reinvent the wheel 👌🏼
- [x] Simple, intuitive and hackable API 👌🏼
- [x] Fully documented 👌🏼
- [x] Performance focused 👌🏼

## Installation

Installation is simple. Just run the following command in your terminal to
install the temporal package in your project.

```sh
go get github.com/maniartech/temporal
```

## Usage

The following example shows how to use the temporal package.

## Date Parsing

Temporal supports parsing of dates using [Intuitive Date Format (IDF)
](#intuitive-date-format-idf) . The following example shows how to parse a
date in the `dd/mm/yyyy` format.

```go
dt := temporal.Parse("01/01/2020", "dd/mm/yyyy")

// You can also specify the timezone
dt2 := temporal.Parse("01/01/2020", "dd/mm/yyyy", time.UTC)
dt3 := temporal.Parse("01/01/2020", "dd/mm/yyyy", time.Local)
dt4 := temporal.Parse("01/01/2020", "dd/mm/yyyy", time.FixedZone("IST", 5.5*60*60))
```

```go

import "github.com/maniartech/temporal"

// Parse a date
t, err := temporal.Parse("2012-01-01", "yyyy-mm-dd")

range, err := temporal.ParseRange("yesterday", time.Now())

// Format a date
s := temporal.Format(t, "yyyy-mm-dd")

// Convert date string to different format
s, err := temporal.Convert("2012-01-01", "yyyy-mm-dd", "dd/mm/yyyy")

// Some handy date finders
temporal.Yesterday()  // Returns yesterday's date

temporal.Tomorrow()   // Returns tomorrow's date

// temporal.SoD()        // Returns start of day (00:00:00.000000000)
// temporal.EoD()        // Returns end of day (23:59:59.999999999)

// temporal.Date(-2)     // Returns date of 2 days ago
// temporal.Date(10)     // Returns date of 10 days from now

// temporal.Week(-1)   // Returns last week's date
// temporal.Week(2)    // Returns date of 2 weeks from now

// temporal.Month(-1)  // Returns last month's date from now

// temporal.MonthStart() // Returns start of current month
// temporal.MonthEnd()   // Returns end of current month

// temporal.MonthStart(temporal.Month(-1)) // Returns start of last month
// temporal.MonthEnd(time.Now())   // Returns end of last month

// temporal.Year()       // Returns last year's date
// temporal.YearStart()  // Returns start of current year
// temporal.YearEnd()    // Returns end of current year

// Combining date finders
// temporal.SoD(temporal.Yesterday()) // Returns start of yesterday
// temporal.EoD(temporal.Today())      // Returns end of today
```

## Intuitive Date Format (IDF)

We've developed the Intuitive Date Format (IDF) for Temporal. `IDF is a cAsE-insensitive format`, eliminating ambiguity often associated with dd-mm-yyyy formats. This intuitive format makes date and time formatting simple and hackable. entry by removing the need to remember upper and lower case attributes, a common issue with other similar formats. For example, %Y represents a four-digit year, while %y denotes a two-digit year in strftime. In contrast, IDF intuitively uses yyyy for a four-digit year and yy for a two-digit year. Typing dates is also more straightforward with IDF, as the format yyyy-mm-dd is easier to remember and input compared to the less intuitive 2006-01-02.

Temporal supports simple, human-friendly date-time formatting. The table below displays the supported formats. Internally, Temporal utilizes time.Time.Format() and converts human-friendly formats into the time.Time format. For instance, it transforms yyyy-mm-dd into 2006-01-02 before using time.Time.Format() to format the date.

### Date Formats

| Format | Output   | Description                                   |
| ------ | -------- | --------------------------------------------- |
| `yy`   | `06`     | Two-digit year with leading zero              |
| `yyyy` | `2006`   | Four-digit year                               |
| `m`    | `1`      | Month without leading zero                    |
| `mm`   | `01`     | Month in two digits with leading zero         |
| `mt`   | `1st`    | Month in ordinal format (not for parsing)     |
| `mmm`  | `Jan`    | Month in short name                           |
| `mmmm` | `January`| Month in full name                            |
| `d`    | `2`      | Day without leading zero                      |
| `dd`   | `02`     | Day in two digits with leading zero           |
| `db`   | ` 2`     | Day in blank-padded two digits                |
| `dt`   | `2nd`    | Day in ordinal format (not for parsing)       |
| `ddd`  | `002`    | Zero padded day of year                       |
| `www`  | `Mon`    | Three-letter weekday name                     |
| `wwww` | `Monday` | Full weekday name                             |

### Time Formats

| Format | Output | Description                                      |
| ------ | ------ | ------------------------------------------------ |
| `h`    | `3`    | Hour in 12-hour format without leading zero      |
| `hh`   | `03`   | Hour in 12-hour format with leading zero         |
| `hhh`  | `15`   | Hour in 24-hour format without leading zero      |
| `a`    | `pm`   | AM/PM in lowercase                               |
| `aa`   | `PM`   | AM/PM in uppercase                               |
| `ii`   | `04`   | Minute with leading zero                         |
| `i`    | `4`    | Minute without leading zero                      |
| `ss`   | `05`   | Second with leading zero                         |
| `s`    | `5`    | Second without leading zero                      |
| `.0`    | `.00`    | Microsecond with leading zero                    |
| `.9`    | `.99`    | Microsecond without leading zero                 |

### Timezone Formats

| Format | Output  | Description                                        |
| ------ | ------- | -------------------------------------------------- |
| `z`    | `Z`     | The Z literal represents UTC                       |
| `zz`   | `MST`   | Timezone abbreviation                              |
| `o`    | `±07`   | Timezone offset with leading zero (only hours)     |
| `oo`   | `±0700` | Timezone offset with leading zero without colon    |
| `ooo`  | `±07:00`| Timezone offset with leading zero with colon       |

### Built-in Formats

Temporal provides all the built-in formats supported by the standard time package.
Such as `time.Layout`, `time.ANSIC`, `time.UnixDate`, `time.RubyDate`, `time.RFC822`,
`time.RFC822Z`, `time.RFC850`, `time.RFC1123`, `time.RFC1123Z`, `time.RFC3339`,
`time.RFC3339Nano`, `time.Kitchen`, etc.

For more information, see the [time package documentation](https://golang.org/pkg/time/#pkg-constants).

### Code Clean up - WIP ⚠️

While the temporal is fully functional, and API has been finalized thoroughly
tested and documented, and can be used in production, there are few area that
needs to be cleaned up. Such as:

- [ ] Some of the source code and tests are not well organized
- [ ] In some cases, documentation needs to be improved
- [ ] Uncomented code in the source files etc

> We shall be working on these issues in the coming days. If you find any issues
or have any suggestions, please feel free to open an issue or submit a pull
request.

## Contributing

Contributions to `temporal` are welcome. Please ensure that your code adheres to the existing style and includes tests covering new features or bug fixes.

## License

`temporal` is [MIT licensed](./LICENSE).
