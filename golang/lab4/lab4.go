package lab4

import "math"

func сalc(a, b, x float64) float64 {
	y := (math.Pow(a, x) - math.Pow(b, x)) / (math.Log10(a/b) * math.Cbrt(a*b))
	return y
}
