package common

import "time"

type DateTime time.Time

func DateTimeNow() DateTime {
	return DateTime(time.Now())
}

func (d DateTime) Time() time.Time {
	return time.Time(d)
}

func (d DateTime) StartOfDay() DateTime {
	t := d.Time()
	return DateTime(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()))
}
