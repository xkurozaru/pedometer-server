package common

import (
	"time"

	model_errors "github.com/xkurozaru/pedometer-server/domain/errors"
)

const (
	HyphenDateFormat = "2006-01-02"
)

type DateTime time.Time

func DateTimeNow() DateTime {
	return DateTime(time.Now())
}

func DateTimeFromString(s string, layout string) (DateTime, error) {
	t, err := time.Parse(layout, s)
	if err != nil {
		return DateTime{}, model_errors.NewInvalidError(err.Error())
	}
	return DateTime(t), nil
}

func (d DateTime) Time() time.Time {
	return time.Time(d)
}

func (d DateTime) StartOfDay() DateTime {
	t := d.Time()
	return DateTime(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()))
}

func (d DateTime) EndOfDay() DateTime {
	return d.StartOfDay().addDay(1).add(-1)
}

func (d DateTime) StartOfWeek() DateTime {
	return d.StartOfDay().addDay(-int(d.weekOfDay()))
}

func (d DateTime) EndOfWeek() DateTime {
	return d.StartOfWeek().addDay(7).add(-1)
}

func (d DateTime) StartOfMonth() DateTime {
	t := d.Time()
	return DateTime(time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()))
}

func (d DateTime) EndOfMonth() DateTime {
	return d.StartOfMonth().addMonth(1).add(-1)
}

func (d DateTime) add(dur time.Duration) DateTime {
	return DateTime(d.Time().Add(dur))
}

func (d DateTime) addDay(n int) DateTime {
	return DateTime(d.Time().AddDate(0, 0, n))
}

func (d DateTime) addMonth(n int) DateTime {
	return DateTime(d.Time().AddDate(0, n, 0))
}

func (d DateTime) Format(layout string) string {
	return d.Time().Format(layout)
}

func (d DateTime) weekOfDay() int {
	return int(d.Time().Weekday())
}
