package utils

import (
	"strconv"
	"time"
)

const (
	FullYearFormat = "2006.01.02"
	HalfYearFormat = "06.01.02"
	MonthDayFormat = "1.02"
	DayFormat      = ".02"
)

func ParseDate(date string) (time.Time, error) {
	t, err := ParseFormattedDate(date)
	if err == nil {
		return t, nil
	}

	t, err = ParseRelativeDate(date)
	if err == nil {
		return t, nil
	}

	return time.Time{}, err
}

func ParseFormattedDate(date string) (time.Time, error) {
	now := time.Now()
	t, err := time.Parse(DayFormat, date)
	if err == nil {
		return time.Date(now.Year(), now.Month(), t.Day(), 0, 0, 0, 0, now.Location()), nil
	}

	t, err = time.Parse(MonthDayFormat, date)
	if err == nil {
		return time.Date(now.Year(), t.Month(), t.Day(), 0, 0, 0, 0, now.Location()), nil
	}

	t, err = time.Parse(HalfYearFormat, date)
	if err == nil {
		return t, nil
	}

	t, err = time.Parse(FullYearFormat, date)
	if err == nil {
		return t, nil
	}

	return time.Time{}, err
}

func ParseRelativeDate(date string) (time.Time, error) {
	offset, err := strconv.Atoi(date)
	if err != nil {
		return time.Time{}, err
	}
	now := time.Now()
	t := now.AddDate(0, 0, offset)
	return t, nil
}
