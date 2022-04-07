package validation

import (
	"errors"
	"time"
)

//vars errors

var ErrEmptyField = errors.New("пустое поле")
var ErrDateInvalid = errors.New("неверный формат даты")

func IsDatePattern(layout, date string) (time.Time, error) {
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, ErrDateInvalid
	}

	return t, nil
}

func IsEmpty(s string) error {
	if s == "" {
		return ErrEmptyField
	}

	return nil
}
