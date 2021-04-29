package date_today

import (
	"fmt"
	"math"
	"time"
)

func Main() {
	fmt.Printf("dia(%.0f|%.0f) <=> semana(%.0f|%.0f) \n",
		daysBetwayTodayStartYear(), daysBetwayTodayFinishYear(),
		weeksBetwayTodayStartYear(), weeksBetwayTodayFinishYear(),
	)

	DeadlineFormated()
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
