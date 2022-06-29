package date_deadline

import (
	"fmt"
	"math"
	"time"
)

func DeadlineMonth() {
	fmt.Printf("Deadline month => day(%.0f|%.0f) | %.1f%% \n",
		daysBetwayTodayStartMonth(), daysBetwayTodayFinishMonth(), percentageLivedMouth())
}

func finishMonthCurrent() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month()+1, 1, 0, 0, 0, 0, time.UTC)
}

func startMonthCurrent() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
}


func daysBetwayTodayStartMonth() float64 {
	return math.RoundToEven(time.Now().UTC().Sub(startMonthCurrent()).Hours() / 24)
}


func daysBetwayTodayFinishMonth() float64 {
	return math.RoundToEven(finishMonthCurrent().Sub(time.Now().UTC()).Hours() / 24)
}

func days() float64 {
	return math.RoundToEven(finishMonthCurrent().Sub(startMonthCurrent()).Hours() / 24) 
}

func percentageLivedMouth() float64 {
	return daysBetwayTodayStartMonth()/days()*100
}
