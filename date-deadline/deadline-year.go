package date_deadline

import (
	"fmt"
	"math"
	"time"
)

func DeadlineYear() {
	fmt.Printf("Deadline year => dia(%.0f|%.0f) <=> semana(%.0f|%.0f) | %.4f \n",
		daysBetwayTodayStartYear(), daysBetwayTodayFinishYear(),
		weeksBetwayTodayStartYear(), weeksBetwayTodayFinishYear(),
		percentageLivedYear())
}
func finishYearCurrent() time.Time {
	return time.Date(time.Now().Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)
}

func startYearCurrent() time.Time {
	return time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC)
}

func daysBetwayTodayFinishYear() float64 {
	return math.RoundToEven(finishYearCurrent().Sub(time.Now().UTC()).Hours() / 24)
}

func weeksBetwayTodayFinishYear() float64 {
	return math.RoundToEven(daysBetwayTodayFinishYear() / 7)
}

func daysBetwayTodayStartYear() float64 {
	return math.RoundToEven(time.Now().UTC().Sub(startYearCurrent()).Hours() / 24)
}

func weeksBetwayTodayStartYear() float64 {
	return math.RoundToEven(daysBetwayTodayStartYear() / 7)
}

func percentageLivedYear() float64 {
	return daysBetwayTodayStartYear() / 365.25 * 100
}
