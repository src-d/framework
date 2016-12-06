package utils

import "time"

type Timezone string

type TimeWithTimezone struct {
	time.Time
	Timezone Timezone
}

func NewTimeWithTimezone(t time.Time) TimeWithTimezone {
	return TimeWithTimezone{
		Time:     t,
		Timezone: Timezone(t.Format("-0700")),
	}
}
