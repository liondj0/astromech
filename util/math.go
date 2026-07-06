package util

func Clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func ClampBuilder(min, max float64) func(value float64) float64 {
	return func(value float64) float64 {
		return Clamp(value, min, max)
	}
}
