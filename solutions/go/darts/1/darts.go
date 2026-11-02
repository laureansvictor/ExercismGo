package darts

import "math"

func Score(x, y float64) int {
	d := math.Sqrt(math.Pow(float64(x), 2.0) + math.Pow(float64(y), 2.0))
	switch {
	case d <= 1:
		return 10
	case d <= 5:
		return 5
	case d <= 10:
		return 1
	}

	return 0
}
