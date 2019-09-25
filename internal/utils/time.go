package utils

import (
	"errors"
	"time"
)

func NextTime(hour, minute, second int) time.Time {
	now := time.Now()
	next := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		hour,
		minute,
		second,
		0,
		now.Location(),
	)

	if now.After(next) {
		return next.AddDate(0, 0, 1)
	}
	return next
}

func TimeTill(t time.Time) (*time.Duration, error) {
	n := time.Now()
	d := t.Sub(n)
	if n.After(t) {
		return nil, errors.New("Time has to be in the future!")
	}
	return &d, nil
}
