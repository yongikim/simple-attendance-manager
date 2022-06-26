package utility

import "time"

type SimpleDate struct {
	Year  int
	Month int
	Day   int
}

func (s SimpleDate) Time() time.Time {
	return time.Date(s.Year, time.Month(s.Month), s.Day, 0, 0, 0, 0, time.Local)
}

func SimpleDateFromTime(t time.Time) SimpleDate {
	return SimpleDate{
		Year:  t.Year(),
		Month: int(t.Month()),
		Day:   t.Day(),
	}
}
