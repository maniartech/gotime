package gotime_test

import (
	"testing"
	"time"

	"github.com/maniartech/gotime/v2"
	"github.com/maniartech/gotime/v2/internal/utils"
)

func TestParse(t *testing.T) {
	utils.AssertEqual(t, "2022-12-31 00:00:00 +0000 UTC", testParse("yyyy-mm-dd", "2022-12-31").String())
	utils.AssertEqual(t, "2022-12-31 00:00:00 +0000 UTC", testParse("2006/01/02", "2022/12/31").String())
	utils.AssertEqual(t, "2010-05-15 00:00:00 +0000 UTC", testParse("01/02/2006", "05/15/2010").String())
	utils.AssertEqual(t, "2022-12-31 12:34:56 +0000 UTC", testParse("2006-01-02 hh:ii:ss", "2022-12-31 12:34:56").String())
	utils.AssertEqual(t, "2022-12-31 12:34:00 +0000 UTC", testParse("2006-01-02 15:04", "2022-12-31 12:34").String())
	utils.AssertEqual(t, "2001-01-15 23:45:00 +0000 UTC", testParse("01/02/2006 15:04", "01/15/2001 23:45").String())
	utils.AssertEqual(t, "2022-12-31 03:00:00 +0300 UTC+3", testParse("yyyy-mm-dd", "2022-12-31").In(time.FixedZone("UTC+3", 3*60*60)).String())
	utils.AssertEqual(t, "2022-12-31 12:34:56 +0000 UTC", testParse("yyyy-mm-dd hh:ii:ss", "2022-12-31 12:34:56").String())
	utils.AssertEqual(t, "2022-12-31 12:34:56.789 +0000 UTC", testParse("2006-01-02 15:04:05.999", "2022-12-31 12:34:56.789").String())
	utils.AssertEqual(t, "2022-12-31 12:34:56.789 +0000 UTC", testParse("01/02/2006 15:04:05.999", "12/31/2022 12:34:56.789").String())
}

func TestParseInLocation(t *testing.T) {
	layout := "yyyy-mm-dd"
	value := "2022-12-31"
	loc := time.FixedZone("UTC+5.5", int(5.5*60*60))
	expectedTime := time.Date(2022, time.December, 31, 0, 0, 0, 0, loc)

	parsedDate, err := gotime.ParseInLocation(layout, value, loc)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedTime, parsedDate)

	layout = "2006/01/02"
	value = "2022/12/31"
	loc = time.FixedZone("UTC+2", int(2*60*60))
	expectedTime = time.Date(2022, time.December, 31, 0, 0, 0, 0, loc)

	parsedDate, err = gotime.ParseInLocation(layout, value, loc)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedTime, parsedDate)

	layout = "01/02/2006"
	value = "05/15/2010"
	loc = time.FixedZone("UTC-3", int(-3*60*60))
	expectedTime = time.Date(2010, time.May, 15, 0, 0, 0, 0, loc)

	parsedDate, err = gotime.ParseInLocation(layout, value, loc)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedTime, parsedDate)

	layout = "2006-01-02 15:04"
	value = "2022-12-31 12:34"
	loc = time.FixedZone("UTC+8", int(8*60*60))
	expectedTime = time.Date(2022, time.December, 31, 12, 34, 0, 0, loc)

	parsedDate, err = gotime.ParseInLocation(layout, value, loc)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedTime, parsedDate)

	layout = "01/02/2006 15:04"
	value = "01/15/2001 23:45"
	loc = time.FixedZone("UTC-2", int(-2*60*60))
	expectedTime = time.Date(2001, time.January, 15, 23, 45, 0, 0, loc)

	parsedDate, err = gotime.ParseInLocation(layout, value, loc)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedTime, parsedDate)

	layout = "2006-01-02 15:04:05.999"
	value = "2022-12-31 12:34:56.789"
	loc = time.FixedZone("UTC-5", int(-5*60*60))
	expectedTime = time.Date(2022, time.December, 31, 12, 34, 56, 789000000, loc)

	parsedDate, err = gotime.ParseInLocation(layout, value, loc)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedTime, parsedDate)

	layout = "01/02/2006 15:04:05.999"
	value = "12/31/2022 12:34:56.789"
	loc = time.FixedZone("UTC+2", int(2*60*60))
	expectedTime = time.Date(2022, time.December, 31, 12, 34, 56, 789000000, loc)

	parsedDate, err = gotime.ParseInLocation(layout, value, loc)
	utils.AssertNoError(t, err)
	utils.AssertEqual(t, expectedTime, parsedDate)

}

func testParse(layout, value string) time.Time {
	dt, err := gotime.Parse(layout, value)
	if err != nil {
		// Use t.Fatal in real tests, but for this helper, return zero time
		return time.Time{}
	}

	return dt
}
