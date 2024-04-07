# GoTime 🕧🕐🕜🕑🕝🕒

> ⚠️ Package name changed to `gotime` on community request.

Designed on the top of Golang's built-in `time` package, instead of reinventing
the wheel `gotime` uses internal `time` package to provide additional
day-to-day functionalities that are practical for real-world applications.

Before we understand why there is a need for yet another time paackage, let's see
some of the code that shocases what this library is capable of doing.

```go
import "github.com/maniartech/gotime"

// // Format the date using human-friendly case-insensitive specifiers like yyyy-mm-dd
dt = time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
s := gotime.Format(dt, "dt mmmm, yyyy")
fmt.Println(s) // 1st January, 2012

// Convert date string to different format
s, err := gotime.Convert("2012-01-01", "yyyy-mm-dd", "wwww, dt mmmm, yyyy")
fmt.Println(s) // Sunday, 1st January, 2012

// Use time ago to get relative time
s, err := gotime.TimeAgo(time.Now().Add(-5 * time.Minute))
fmt.Println(s) // 5 minutes ago

// Parse a date string using easy to remember specifiers
t, err := gotime.Parse("2012-01-01", "yyyy-mm-dd")
fmt.Println(t) // 2012-01-01 00:00:00 +0000 UTC

tz := time.FixedZone("IST", 5.5*60*60)
t, err := gotime.ParseInLocation("01/01/2020", "dd/mm/yyyy", tz)
fmt.Println(t) // 2020-01-01 00:00:00 +0530 IST

// Some handy date finders and other utility functions
gotime.Yesterday()  // Returns yesterday's date
gotime.NextWeek()   // Returns data exactly one week from now

// Other utility functions
d1 := gotime.Date(10) // Returns date of 10 days from now
d2 := gotime.Date(-2) // Returns date of 2 days ago
d3 := gotime.Date(10, d1) // Returns date of 10 days from t1
gotime.Earliest(d1, d2, d3) // Returns the earliest date from the given list of dates
gotime.Latest(d1, d2, d3) // Returns the latest date from the given list of dates
gotime.IsBetween(d1, d2, d3) // Returns true if d1 is between d2 and d3

gotime.IsLeapYear(2020) // Returns true if the given year is a leap year
weekDdays := bool{true, true, true, true, true, false, false}
gotime.WorkDay(time.Now(), 15, weekDays) // Returns the date after the specified
                                           // number of workdays, considering
                                           // holidays and weekends

// So on...
```

## Why GoTime?

### ✨ Key Features (Designed for Practicality)

It provides features are practical and useful in real-world applications. These
features are either missing or not easy to use in the standard time package.

- [x] **Easy Formatting:** Format dates using straightforward, human-readable and intuitive specifiers like dd/mm/yyyy.
- [x] **Human-Friendly Parsing:** Effortlessly parse dates using intuitive date specifiers
- [x] **Format Conversion:** Seamlessly convert dates between formats, e.g., from `dd/mm/yyyy` to `yyyy-mm-dd`.
- [x] **Relative Time Processing:** Translate datetime into relative terms like `a few minutes ago`, `yesterday`, `5 days ago`, or `3 years from now`.
- [x] **Finder Functions:** Utilize functions like `Yesterday()`, `Tomorrow()`, `SoD()` (Start of Day), and `EoD()` (End of Day) etc.
- [x] **Utility Functions:** Access a suite of utility functions including  `Latest()`,
      `Earliest()`, `IsBetween()`, `TruncateDate()`, and more.

### ✨ Developer Friendly

It provides a comprehensive range of specifiers for all your date and time
formatting needs, making it an indispensable tool for Go developers.

- [x] 💯% test coverage 👌🏼
- [x] TinyGo compatible 👌🏼
- [x] No external dependencies  👌🏼
- [x] Fully utilises the standard time package and does not reinvent the wheel 👌🏼
- [x] Simple, intuitive and hackable API 👌🏼
- [x] Fully documented 👌🏼
- [x] Performance focused 👌🏼

### ✨ Ideal Use Cases

- [x] Developers needing formatting and parsing dates using human-friendly specifiers like `yyyy-mm-dd`.
- [x] Applications requiring conversion between different date formats.
- [x] Systems that display relative time representations.
- [x] Software dealing with date range calculations and comparisons.
- [x] Projects where standard time package features are insufficient or cumbersome.

The `gotime` stands out by offering features that are either missing or not as user-friendly in the standard time package, making it an invaluable addition to any Go developer's toolkit.

## Installation

Installation is simple. Just run the following command in your terminal to
install the `gotime` package in your project.

```sh
go get github.com/maniartech/gotime
```

## Intuitive Date Specifiers (IDS)

We've developed the Intuitive Date Specifiers (IDS) for `gotime`. `IDS is a cAsE-insensitive format`, eliminating ambiguity often associated with dd-mm-yyyy formats. This intuitive specifiers makes date and time formatting simple and hackable. entry by removing the need to remember upper and lower case attributes, a common issue with other similar formats. For example, %Y represents a four-digit year, while %y denotes a two-digit year in strftime. In contrast, IDS intuitively uses yyyy for a four-digit year and yy for a two-digit year. Typing dates is also more straightforward with IDS, as the format yyyy-mm-dd is easier to remember and input compared to the less intuitive 2006-01-02.

It supports simple, human-friendly date-time formatting. The table below displays the supported formats. Internally, `gotime` utilizes time.Time.Format() and converts human-friendly formats into the time.Time format. For instance, it transforms yyyy-mm-dd into 2006-01-02 before using time.Time.Format() to format the date.

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

It provides all the built-in formats supported by the standard time package.
Such as `time.Layout`, `time.ANSIC`, `time.UnixDate`, `time.RubyDate`, `time.RFC822`,
`time.RFC822Z`, `time.RFC850`, `time.RFC1123`, `time.RFC1123Z`, `time.RFC3339`,
`time.RFC3339Nano`, `time.Kitchen`, etc.

## API Reference

### Date Parsing, Formatting and Conversion

| Function         | Description                                                                                                   |
|------------------|---------------------------------------------------------------------------------------------------------------|
| Parse            | Parses the given date string using [IDS](#intuitive-date-specifiers-ids), a human-friendly, hackable date format                                |
| ParseInLocation  | Parses the given date string in the specified location using [IDS](#intuitive-date-specifiers-ids), a human-friendly, hackable date format, and the given location |
| Format           | Formats the given date using [IDS](#intuitive-date-specifiers-ids) format specifiers                                                            |
| FormatTimestamp  | Formats the given timestamp in Unix format using [IDS](#intuitive-date-specifiers-ids) format specifiers                                        |
| Convert          | Converts the given date string from one format to another using [IDS](#intuitive-date-specifiers-ids) format specifiers                         |

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

#### DateRange Functions

| Function  | Description                                                                                   |
|-----------|-----------------------------------------------------------------------------------------------|
| IsInRange | Returns true if the given date is in the specified range                                      |
| IsInDateRange | Returns true if the given date is in the specified range. Before performing inclusive comparison, it sets the time to the start ~~of~~ the day for the start date and the end of the day for the end date. |

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
| NetWorkDays    | Returns the number of workdays between two dates, considering holidays and weekends           |
| WorkDay        | Returns the date after the specified number of workdays, considering holidays and weekends    |
| PrevWorkDay    | Returns the date before the specified number of working days, considering holidays and weekends   |

For more information, see the [time package documentation](https://golang.org/pkg/time/#pkg-constants).

### Code Clean up - WIP ⚠️

While the `gotime` is fully functional, and API has been finalized thoroughly
tested and documented, and can be used in production, there are few area that
needs to be cleaned up. Such as:

- [ ] Some of the source code and tests are not well organized
- [ ] In some cases, documentation needs to be improved
- [ ] Uncomented code in the source files etc

> We shall be working on these issues in the coming days. If you find any issues
or have any suggestions, please feel free to open an issue or submit a pull
request.

## Contributing

Contributions to `gotime` are welcome. Please ensure that your code adheres to the existing style and includes tests covering new features or bug fixes.

## License

`gotime` is [MIT licensed](./LICENSE).
