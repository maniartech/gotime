package gotime

import "time"

// IsBetween returns true if the given time.Time is between the given time.Time range
//
// # Arguments
//
// t1: (time.Time) The date to be checked
//
// t2: (time.Time) The first date of the range
//
// t3: (time.Time) The second date of the range
func IsBetween(t1, t2, t3 time.Time) bool {
	t1Unix := t1.UnixMilli()
	t2Unix := t2.UnixMilli()
	t3Unix := t3.UnixMilli()

	// Swapping the values if t2Unix is greater than t3Unix
	if t2Unix > t3Unix {
		t2Unix, t3Unix = t3Unix, t2Unix
	}

	return t1Unix >= t2Unix && t1Unix <= t3Unix
}

// IsBetweenDates checks if the given time is in the range of the start and end
// date. Before performing inclusive comparison, it sets the time to the start
// of the day for the start date and the end of the day for the end date.
func IsBetweenDates(t1, t2, t3 time.Time) bool {
	sec1 := SoD(t1).Unix()
	sec2 := SoD(t2).Unix()
	sec3 := EoD(t3).Unix()

	// Swapping the values if sec2 is greater than sec3
	if sec2 > sec3 {
		sec2, sec3 = sec3, sec2
	}

	return sec1 >= sec2 && sec1 <= sec3
}
