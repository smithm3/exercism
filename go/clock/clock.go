package clock

import (
	"fmt"
)

const testVersion = 4

const minsInHour = 60
const hoursInDay = 24
const minsInDay = minsInHour * hoursInDay

type Clock struct {
	hour int
	minute  int
}

func New(hour, minute int) Clock {

	totalMins := determineTotalMins(hour, minute)
	hour, minute = determineTime(totalMins)

	return Clock{hour, minute}
}

func determineTotalMins(hour int, minute int) int {
	totalMins := hour * minsInHour
	totalMins += minute
	return totalMins
}

func determineTime(totalMins int) (hour int, minute int) {
	mod := totalMins % minsInDay

	if mod < 0 {
		mod += minsInDay
	}

	hour = mod / minsInHour
	minute = mod % minsInHour

	return hour, minute
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {

	totalMins := determineTotalMins(c.hour, c.minute) + minutes
	c.hour, c.minute = determineTime(totalMins)

	return c
}
