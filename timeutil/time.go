package timeutil

import "time"

const (
	DAY_IN_HOURS = 24

	DATE_FORMAT_SIMPLE = "2006-01-02"
	DATE_FORMAT_FULL   = "2006-01-02 15:04:05"
)

var (
	YEAR = time.Date(1, 0, 0, 0, 0, 0, 0, time.Local)
)

// AddDays Adds a number of days to a date returning a new time.Time
func AddDays(d time.Time, days int) time.Time {
	return d.AddDate(0, 0, days)
}

// AddHours Adds a number of hours to a date returning a new time.Time
func AddHours(d time.Time, hours int) time.Time {
	return d.Add(time.Duration(int(time.Hour) * hours))
}

// AddMinutes Adds a number of minutes to a date returning a new time.Time
func AddMinutes(d time.Time, minutes int) time.Time {
	return d.Add(time.Duration(int(time.Minute) * minutes))
}

// AddMonths Adds a number of months to a date returning a new time.Time
func AddMonths(d time.Time, months int) time.Time {
	return d.AddDate(0, months, 0)
}

// AddYears Adds a number of years to a date returning a new time.Time
func AddYears(d time.Time, years int) time.Time {
	return d.AddDate(years, 0, 0)
}

// Truncate this date, leaving the year field specified as the most significant field.
func TruncateYear(d time.Time) time.Time {
	return time.Date(d.Year(), 0, 0, 0, 0, 0, 0, d.Location())
}

// Truncate this date, leaving the month field specified as the most significant field.
func TruncateMonth(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), 0, 0, 0, 0, 0, d.Location())
}

// Truncate this date, leaving the day field specified as the most significant field.
func TruncateDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// Truncate this date, leaving the hour field specified as the most significant field.
func TruncateHour(d time.Time) time.Time {
	return d.Truncate(time.Hour)
}

// Truncate this date, leaving the minute field specified as the most significant field.
func TruncateMinute(d time.Time) time.Time {
	return d.Truncate(time.Minute)
}

// Truncate this date, leaving the second field specified as the most significant field.
func TruncateSecond(d time.Time) time.Time {
	return d.Truncate(time.Second)
}

//
func ParseSimpleFormat(s string) (time.Time, error) {
	return time.Parse(DATE_FORMAT_SIMPLE, s)
}

func ParseFormat(s string) (time.Time, error) {
	return time.Parse(DATE_FORMAT_FULL, s)
}

func FormatSimple(t time.Time) string {
	return t.Format(DATE_FORMAT_SIMPLE)
}

func Format(t time.Time) string {
	return t.Format(DATE_FORMAT_FULL)
}

//
func SetYears(t time.Time, years int) time.Time {
	return time.Date(years, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func SetMonths(t time.Time, months int) time.Time {
	return time.Date(t.Year(), time.Month(months), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func setDays(t time.Time, days int) time.Time {
	return time.Date(t.Year(), t.Month(), days, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func SetHours(t time.Time, hours int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hours, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func SetMinutes(t time.Time, minutes int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minutes, t.Second(), t.Nanosecond(), t.Location())
}

func setSeconds(t time.Time, seconds int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), seconds, t.Nanosecond(), t.Location())
}

func setMilliSeconds(t time.Time, milliseconds int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), milliseconds*1000, t.Location())
}
