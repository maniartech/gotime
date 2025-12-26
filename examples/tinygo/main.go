package main

import (
	"fmt"
	"time"

	"github.com/maniartech/gotime/v2"
)

func main() {
	now := time.Date(2025, 7, 7, 12, 0, 0, 0, time.UTC)
	fmt.Println("Now:", now)

	// Add 5 days
	future := gotime.Days(5, now)
	fmt.Println("5 days from now:", future)

	// Start of day
	startOfDay := gotime.SoD(now)
	fmt.Println("Start of day:", startOfDay)

	// Is business day
	weekends := []time.Weekday{time.Saturday, time.Sunday}
	holiday := time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC)
	isBiz := gotime.IsBusinessDay(now, weekends, holiday)
	fmt.Println("Is business day:", isBiz)
}
