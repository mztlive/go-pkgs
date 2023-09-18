package currency

import "math"

type Money int64

// Cent 分
type Cent Money

// Dollar 元
type Dollar float64

// Dime 角
type Dime float64

func (a Cent) ToDollar() Dollar {
	tmp := float64(a) / 100.
	out := math.Round(tmp*100) / 100
	return Dollar(out)
}

func (a Cent) ToDime() Dime {
	tmp := float64(a) / 10.
	out := math.Round(tmp*10) / 10
	return Dime(out)
}

func (a Dollar) ToCent() Cent {
	return Cent(math.Round(float64(a) * 100))
}

func (a Dollar) ToDime() Dime {
	return Dime(math.Round(float64(a) * 10))
}
