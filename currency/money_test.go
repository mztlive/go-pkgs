package currency

import (
	"math"
	"testing"
)

func TestCentToDollar(t *testing.T) {
	cases := []struct {
		in   Cent
		want Dollar
	}{
		{Cent(100), Dollar(1)},
		{Cent(50), Dollar(0.5)},
		{Cent(1234), Dollar(12.34)},
		{Cent(1999), Dollar(19.99)},
	}

	for _, c := range cases {
		got := c.in.ToDollar()
		if got != c.want {
			t.Errorf("ToDollar(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCentToDime(t *testing.T) {
	cases := []struct {
		in   Cent
		want Dime
	}{
		{Cent(10), Dime(1)},
		{Cent(50), Dime(5)},
		{Cent(1234), Dime(123.4)},
	}

	for _, c := range cases {
		got := c.in.ToDime()
		if got != c.want {
			t.Errorf("ToDime(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestDollarToCent(t *testing.T) {
	cases := []struct {
		in   Dollar
		want Cent
	}{
		{Dollar(1), Cent(100)},
		{Dollar(0.5), Cent(50)},
		{Dollar(12.34), Cent(1234)},
		{Dollar(19.99), Cent(1999)},
	}

	for _, c := range cases {
		got := c.in.ToCent()
		if got != c.want {
			t.Errorf("ToCent(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestDollarToDime(t *testing.T) {
	cases := []struct {
		in   Dollar
		want Dime
	}{
		{Dollar(1), Dime(10)},
		{Dollar(0.5), Dime(5)},
	}

	for _, c := range cases {
		got := c.in.ToDime()
		if math.Abs(float64(got-c.want)) > 0.0001 {
			t.Errorf("ToDime(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
