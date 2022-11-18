# datetime (WIP)

A golang library for parsing and parsing, formatting and processing dates and times using simple human-friendly formats such as "yesterday", "tomorrow", "dd/mm/yyyy", etc.

## Installation

    go get github.com/maniartech/datetime

## Usage

      import "github.com/lestrrat/datetime"

      // Parse a date
      t, err := datetime.Parse("tomorrow")

      // Format a date
      s := datetime.Format(t, "yyyy-mm-dd")

      // Convert date string to different format
      s, err := dateutils.Convert("2012-01-01", "yyyy-mm-dd", "dd/mm/yyyy")

      // Get the relative date range
      start, end, err := dateutils.RelativeRange("last-week")

## Date Format Table

| Format | Description | Example |
|--------|-------------|---------|
| **Date**  |
| yy     | Two digit year with leading zero | 06 |
| yyyy   | Four digit year | 2006 |
| m      | Month without leading zero | 1 |
| mm     | Month in two digits with leading zero | 01 |
| mmm    | Month in short name | Jan |
| mmmm   | Month in full name | January |
| d      | Day without leading zero | 2 |
| dd     | Day in two digits with leading zero | 02 |
| ddd    | Zero padded day of year | 002 |
| w      | Three letter weekday name | Mon |
| ww     | Full weekday name | Monday |
| **Time**  |
| h      | Hour in 12 hour format without leading zero | 3 |
| hh     | Hour in 12 hour format with leading zero | 03 |
| hhh   | Hour in 24 hour format with leading zero | 15 |
| a      | am/pm | pm |
| aa     | AM/PM | PM |
| i      | Minute without leading zero | 4 |
| ii     | Minute with leading zero | 04 |
| s      | Second without leading zero | 5 |
| ss     | Second with leading zero | 05 |
| u      | Microsecond | 000000 |
| **Timezone** |
| z      | UTC offset | ±0700 |
| zh     | Numeric timezone hour with hours only | ±07 |
| zz     | UTC offset with colon | ±07:00 |
| zzz    | Timezone abbreviation | MST |
| zzzz   | Timezone in long format | GMT-07:00 |


Refs:

https://docs.djangoproject.com/en/4.1/ref/templates/builtins/#date
https://programming.guide/go/format-parse-string-time-date-example.html