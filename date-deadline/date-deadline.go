package date_deadline

import (
	"fmt"
	"math"
	"time"
)

func DeadlineLife() {
	fmt.Printf("Deadline life => day(%.0f|%.0f) <=> week(%.0f|%.0f)  <=> month(%.0f|%.0f) | %.3f%%",
		daysBetwayTodayBirth(), daysBetwayTodayDeadline(),
		weeksBetwayTodayBirth(), weeksBetwayTodayDeadline(),
		monthsBetwayTodayBirth(), monthsBetwayTodayDeadline(),
		percentageLived())
}

func deadline() time.Time {
	return time.Date(2058, 12, 31, 0, 0, 0, 0, time.UTC)
}

func birth() time.Time {
	return time.Date(1993, 8, 22, 0, 0, 0, 0, time.UTC)
}

func life() float64 {
	return math.RoundToEven(deadline().Sub(birth()).Hours() / 24)
}

func daysBetwayTodayDeadline() float64 {
	return math.RoundToEven(deadline().Sub(time.Now().UTC()).Hours() / 24)
}

func weeksBetwayTodayDeadline() float64 {
	return math.RoundToEven(daysBetwayTodayDeadline() / 7)
}

func monthsBetwayTodayDeadline() float64 {
	return math.RoundToEven(daysBetwayTodayDeadline() / 30)
}

func daysBetwayTodayBirth() float64 {
	return math.RoundToEven(time.Now().UTC().Sub(birth()).Hours() / 24)
}

func weeksBetwayTodayBirth() float64 {
	return math.RoundToEven(daysBetwayTodayBirth() / 7)
}

func monthsBetwayTodayBirth() float64 {
	return math.RoundToEven(daysBetwayTodayBirth() / 30)
}

func percentageLived() float64 {
	return daysBetwayTodayBirth() / life() * 100
}
