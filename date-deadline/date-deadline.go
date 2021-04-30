package date_deadline

import (
	"fmt"
	"math"
	"time"
)

func DeadlineLife() {
	fmt.Printf("Deadline life => dia(%.0f|%.0f) <=> semana(%.0f|%.0f) | %.4f",
		daysBetwayTodayBirth(), daysBetwayTodayDeadline(),
		weeksBetwayTodayBirth(), weeksBetwayTodayDeadline(),
		percentageLived())

}
func deadline() time.Time {
	return time.Date(2075, 10, 11, 0, 0, 0, 0, time.UTC)
}

func birth() time.Time {
	return time.Date(1993, 8, 22, 0, 0, 0, 0, time.UTC)
}

func daysBetwayTodayDeadline() float64 {
	return math.RoundToEven(deadline().Sub(time.Now().UTC()).Hours() / 24)
}

func weeksBetwayTodayDeadline() float64 {
	return math.RoundToEven(daysBetwayTodayDeadline() / 7)
}

func daysBetwayTodayBirth() float64 {
	return math.RoundToEven(time.Now().UTC().Sub(birth()).Hours() / 24)
}

func weeksBetwayTodayBirth() float64 {
	return math.RoundToEven(daysBetwayTodayBirth() / 7)
}

func percentageLived() float64 {
	return daysBetwayTodayBirth() / 300
}
