package dateutils

import "time"

type DateTime interface {
	time.Time | string | int64
}

func Format() {
	//
}