package temporal

import "time"

// Add adds the given duration to the DateTime
func (d DateTime) Add(duration time.Duration) DateTime {
	return DateTime(d.Time().Add(duration))
}

// AddDate adds the given number of years, months, and days to the DateTime
func (d DateTime) AddDate(years, months, days int) DateTime {
	return DateTime(d.Time().AddDate(years, months, days))
}

func (d DateTime) UTC() DateTime {
	// return DateTime((*time.Time)(&d).UTC())
	return DateTime(time.Time(d).UTC())
}

// In returns the DateTime in the given location
func (d DateTime) In(loc *time.Location) DateTime {
	return DateTime(d.Time().In(loc))
}

// Round returns the DateTime rounded to the nearest unit
func (d DateTime) Round(unit time.Duration) DateTime {
	return DateTime(d.Time().Round(unit))
}

// Truncate returns the DateTime truncated to the given unit
func (d DateTime) Truncate(unit time.Duration) DateTime {
	return DateTime(d.Time().Truncate(unit))
}
