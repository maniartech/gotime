package dateutils

import (
	"testing"
)

func TestToday(t *testing.T) {
	today := *Today()
	if today != todayMidnight {
		t.Error("")
	}
}

func TestYesterday(t *testing.T) {
	yesterday := *Yesterday()
	if yesterday != todayMidnight.AddDate(0, 0, -1) {
		t.Error("")
	}
}

func TestTomorrow(t *testing.T) {
	tomorrow := *Tomorrow()
	if tomorrow != todayMidnight.AddDate(0, 0, 1) {
		t.Error("")
	}
}

/*func TestToday(t *testing.T) {
	today := Today()
	if today != todayMidnight {
		t.Error("bruh")
	}
}*/