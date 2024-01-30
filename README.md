# Temporal 🕧🕐🕜🕑🕝🕒

> /ˈtɛmp(ə)rəl/ - Relating to Time. The spatial and temporal dimensions of human interference in complex ecosystems

Temporal is a Go language library designed to simplify the parsing, formatting,
and processing of dates and times. While it does not intend to replace Go's
standard time package, Temporal enhances it by providing additional, user-friendly
functionalities that are practical for real-world applications. This library
focuses on making common date and time operations, such as formatting, parsing,
and working with relative times and date ranges, more accessible and efficient.

## Why Temporal?

### ✨ Key Features (Designed for Practicality)

It provides features are practical and useful in real-world applications. These
features are either missing or not easy to use in the standard time package.

- [x] **Human-Friendly Parsing:** Effortlessly parse dates in intuitive formats
      like dd/mm/yyyy.
- [x] **Easy Formatting:** Format dates using straightforward, human-readable formats.
- [x] **Format Conversion:** Seamlessly convert dates between formats, e.g., from `dd/mm/yyyy` to `yyyy-mm-dd`.
- [x] **Relative Time Processing:** Translate datetime into relative terms like `a few minutes ago`, `yesterday`, `5 days ago`, or `3 years from now`.
- [x] **Finder Functions:** Utilize functions like `Yesterday()`, `Tomorrow()`, `SoD()` (Start of Day), and `EoD()` (End of Day) etc.
- [x] **Utility Functions:** Access a suite of utility functions including  `Latest()`,
      `Earliest()`, `IsBetween()`, `TruncateDate()`, and more.

### ✨ Developer Friendly

It provides a comprehensive range of specifiers for all your date and time
formatting needs, making it an indispensable tool for Go developers.

- [x] 100% test coverage 👌🏼
- [x] TinyGo compatible 👌🏼
- [x] No external dependencies  👌🏼
- [x] Fully utilises the standard time package and does not reinvent the wheel 👌🏼
- [x] Simple, intuitive and hackable API 👌🏼
- [x] Fully documented 👌🏼
- [x] Performance focused 👌🏼

### ✨ Ideal Use Cases

- [x] Developers needing intuitive date parsing and formatting.
- [x] Applications requiring conversion between different date formats.
- [x] Systems that display relative time representations.
- [x] Software dealing with date range calculations and comparisons.
- [x] Projects where standard time package features are insufficient or cumbersome.

Temporal stands out by offering features that are either missing or not as user-friendly in the standard time package, making it an invaluable addition to any Go developer's toolkit.

## Installation

Installation is simple. Just run the following command in your terminal to
install the temporal package in your project.

```sh
go get github.com/maniartech/temporal
```

## Usage

The following example shows how to use the temporal package.

```go
import "github.com/maniartech/temporal"

// Parse a date
t, err := temporal.Parse("2012-01-01", "yyyy-mm-dd")
fmt.Println(t) // 2012-01-01 00:00:00 +0000 UTC

tz := time.FixedZone("IST", 5.5*60*60)
t, err := temporal.ParseInLocation("01/01/2020", "dd/mm/yyyy", tz)
fmt.Println(t) // 2020-01-01 00:00:00 +0530 IST

// Format a date
s := temporal.Format(t, "dt mmmm, yyyy")
fmt.Println(s) // 1st January, 2012

// Convert date string to different format
s, err := temporal.Convert("2012-01-01", "yyyy-mm-dd", "wwww, dt mmmm, yyyy")
fmt.Println(s) // Sunday, 1st January, 2012

// Time ago
s, err := temporal.TimeAgo(time.Now().Add(-5 * time.Minute))
fmt.Println(s) // 5 minutes ago

// Some handy date finders and other utility functions
temporal.Yesterday()  // Returns yesterday's date
temporal.NextWeek()   // Returns data exactly one week from now

// Other utility functions
d1 := temporal.Date(10) // Returns date of 10 days from now
d2 := temporal.Date(-2) // Returns date of 2 days ago
d3 := temporal.Date(10, d1) // Returns date of 10 days from t1
temporal.Earliest(d1, d2, d3) // Returns the earliest date from the given list of dates
temporal.Latest(d1, d2, d3) // Returns the latest date from the given list of dates
temporal.IsBetween(d1, d2, d3) // Returns true if d1 is between d2 and d3

temporal.IsLeapYear(2020) // Returns true if the given year is a leap year
weekDdays := bool{true, true, true, true, true, false, false}
temporal.WorkDay(time.Now(), 15, weekDays) // Returns the date after the specified
                                           // number of workdays, considering
                                           // holidays and weekends
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

## API Reference

### Date Parsing, Formatting and Conversion

| Function         | Description                                                                                                   |
|------------------|---------------------------------------------------------------------------------------------------------------|
| Parse            | Parses the given date string using [IDF](#intuitive-date-format-idf), a human-friendly, hackable date format                                |
| ParseInLocation  | Parses the given date string in the specified location using [IDF](#intuitive-date-format-idf), a human-friendly, hackable date format, and the given location |
| Format           | Formats the given date using [IDF](#intuitive-date-format-idf) format specifiers                                                            |
| FormatTimestamp  | Formats the given timestamp in Unix format using [IDF](#intuitive-date-format-idf) format specifiers                                        |
| Convert          | Converts the given date string from one format to another using [IDF](#intuitive-date-format-idf) format specifiers                         |

### Time Ago

| Function | Description                                        |
|----------|----------------------------------------------------|
| TimeAgo  | Returns the humanized relative time from the given date |


### Relative Date Finders

#### Years

| Function   | Description                                            |
|------------|--------------------------------------------------------|
| YearStart  | Returns the start date of the given year               |
| YearEnd    | Returns the end date of the given year                 |
| LastYear   | Returns the date exactly one year before today         |
| NextYear   | Returns the date exactly one year after today          |
| Years      | Returns the date of the given number of years from now or the given date |

#### Months

| Function   | Description                                             |
|------------|---------------------------------------------------------|
| MonthStart | Returns the start date of the given month               |
| MonthEnd   | Returns the end date of the given month                 |
| LastMonth  | Returns the date exactly one month before today         |
| NextMonth  | Returns the date exactly one month after today          |
| Months     | Returns the date of the given number of months from now or the given date |

#### Weeks

| Function     | Description                                                                          |
|--------------|--------------------------------------------------------------------------------------|
| WeekStart    | Returns the start date of the given week                                             |
| WeekStartOn  | Considers the given day as the start of the week and returns the start date of the week |
| WeekEnd      | Returns the end date of the given week                                               |
| WeekEndOn    | Considers the given day as the end of the week and returns the end date of the week   |
| LastWeek     | Returns the date exactly one week before today                                       |
| NextWeek     | Returns the date exactly one week after today                                        |
| Weeks        | Returns the date of the given number of weeks from now or the given date             |

#### Days

| Function  | Description                                                 |
|-----------|-------------------------------------------------------------|
| SoD       | Returns the start of the given day                           |
| EoD       | Returns the end of the given day                             |
| Yesterday | Returns the date exactly one day before today                |
| Tomorrow  | Returns the date exactly one day after today                 |
| Days      | Returns the date of the given number of days from now or the given date |

### DateTime Utility Functions

| Function       | Description                                                                                   |
|----------------|-----------------------------------------------------------------------------------------------|
| IsLeapYear     | Returns true if the given year is a leap year                                                 |
| DaysInMonth    | Returns the number of days in the given month                                                 |
| DaysInYear     | Returns the number of days in the given year                                                  |
| DaysInQuarter  | Returns the number of days in the given quarter                                               |
| NewDate        | Returns the date from the given year, month, and day                                          |
| NewTime        | Returns the time from the given hour, minute, and second                                      |
| DateValue      | Returns the serial number of the given date from 1900-01-01                                   |
| Diff           | Returns the difference between two durations. Can also return a rounded result.               |
| Latest         | Returns the latest time from the specified list of times                                      |
| Earliest       | Returns the earliest time from the specified list of times                                    |
| IsBetween      | Returns true if the given date is between the specified range                                  |
| TruncateTime   | Truncates the time part from the given date                                                   |
| TruncateDate   | Truncates the date part from the given date                                                   |
| WorkDay        | Returns the date after the specified number of workdays, considering holidays and weekends    |
| NetWorkDays    | Returns the number of workdays between two dates, considering holidays and weekends           |

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
