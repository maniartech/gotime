#!/bin/bash

# Script to fix godoc examples to use gotime. prefix for better user-friendliness

# List of function patterns to fix
FUNCTIONS=(
    "MonthsBetween"
    "DaysBetween"
    "WeeksBetween"
    "DurationInWords"
    "IsValidAge"
    "IsLeapYear"
    "DaysInMonth"
    "DaysInYear"
    "DaysInQuarter"
    "NewDate"
    "NewTime"
    "ReplaceDate"
    "ReplaceTime"
    "QuarterStart"
    "QuarterEnd"
    "LastQuarter"
    "NextQuarter"
    "Quarters"
    "QuarterOfYear"
    "IsBetween"
    "IsBetweenDates"
    "IsWeekdayPresentInRange"
    "CountWeekdaysInRange"
    "Hours"
    "Minutes"
    "Seconds"
    "Earliest"
    "TruncateTime"
    "WorkDay"
    "PrevWorkDay"
    "NetWorkDays"
)

# Find and replace function calls in godoc examples
for func in "${FUNCTIONS[@]}"; do
    echo "Fixing $func examples..."
    # Use sed to replace naked function calls with gotime. prefix in example comments
    sed -i "s|//\([[:space:]]*[a-zA-Z_][a-zA-Z0-9_]*[[:space:]]*:=\)[[:space:]]*$func(|//\1 gotime.$func(|g" *.go
done

echo "Fixed godoc examples to use gotime. prefix"
