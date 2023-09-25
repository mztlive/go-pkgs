package strings

import (
	"time"
)

type Birthday string

// ToAge returns the age of the birthday
// Example:
//
//	input: "2020-01-01", today: 2023-10-01 result: [3, 9]
func (b Birthday) ToAge(todayTime time.Time) ([2]int, error) {
	var (
		layout       = "2006-01-02"
		birthdayTime time.Time
		years        int
		months       int
		err          error
	)

	if birthdayTime, err = time.Parse(layout, string(b)); err != nil {
		return [2]int{}, err
	}

	// Calculate the difference in years
	years = todayTime.Year() - birthdayTime.Year()

	// If today's month is less than birthday's month, subtract a year
	// Or if they are on the same month, but today's day is less than birthday's day, also subtract a year
	if todayTime.Month() < birthdayTime.Month() ||
		(todayTime.Month() == birthdayTime.Month() && todayTime.Day() < birthdayTime.Day()) {
		years--
	}

	// Calculate the difference in months
	if todayTime.Month() >= birthdayTime.Month() {
		months = int(todayTime.Month() - birthdayTime.Month())
	} else {
		months = int(todayTime.Month() + 12 - birthdayTime.Month())
	}

	// If today's day is less than birthday's day, subtract a month
	if todayTime.Day() < birthdayTime.Day() {
		months--
	}

	return [2]int{years, months}, nil
}
