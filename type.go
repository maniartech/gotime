package dateutils

import (
	"time"
)

type DateTime interface {
	time.Time | string | int64 | uint64
}

func Format() {
	//
}