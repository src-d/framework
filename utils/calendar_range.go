package utils

import "time"

type CalendarRanges []CalendarRange

// Has reports whether a concrete moment is in any of the ranges.
// If two ranges overlap, the most concrete one is taken into account.
// (E. g. if one range is from 9 AM to 5 PM, and another range is from 9 AM to 2 PM
// and from 4 PM to 7 PM on April, Has("2015/07/13 3 PM") is true but
// Has("2015/04/13 3 PM") is false.)
func (crs CalendarRanges) Has(t time.Time) bool {
	// TODO
	return false
}

func (crs CalendarRanges) String() string {
	// TODO
	return ""
}

type CalendarRange struct {
	From CalendarTime
	To   CalendarTime
}

func (cr CalendarRange) Has(t time.Time) bool {
	// TODO
	return false
}

func (cr CalendarRange) String() string {
	// TODO
	return ""
}

type CalendarTime struct {
	HasYear    bool
	Year       int
	HasMonth   bool
	Month      time.Month
	HasWeekday bool
	Weekday    time.Weekday
	HasDay     bool
	Day        int
	HasHour    bool
	Hour       int
	HasMinute  bool
	Minute     int
}

func (ct CalendarTime) Has(t time.Time) bool {
	// TODO
	return false
}

func (ct CalendarTime) String() string {
	// TODO
	return ""
}
