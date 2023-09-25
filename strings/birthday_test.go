package strings

import (
	"testing"
	"time"
)

func TestBirthday_ToAge(t *testing.T) {
	b := Birthday("1990-01-01")
	today := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	expected := [2]int{33, 9}
	if result, err := b.ToAge(today); err != nil {
		t.Errorf("ToAge(%v) returned an error: %v", b, err)
	} else if result != expected {
		t.Errorf("ToAge(%v) = %v, expected %v", b, result, expected)
	}
}
