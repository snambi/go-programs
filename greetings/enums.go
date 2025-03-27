package greetings

import "fmt"

type DAYS int

const (
	SUNDAY DAYS = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

var dayNames = map[DAYS]string{
	SUNDAY:    "Sunday",
	MONDAY:    "Monday",
	TUESDAY:   "Tuesday",
	WEDNESDAY: "Wednesday",
	THURSDAY:  "Thursday",
	FRIDAY:    "Friday",
	SATURDAY:  "Saturday",
}

func (d DAYS) String() string {
	return dayNames[d]
}

func Call(d DAYS) {
	fmt.Println("Today is ", d)
}
