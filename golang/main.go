package main

import (
	"fmt"
)

func main() {
	fmt.Println("kozin_danila_sergeevich")

}
func SolveTaskA(a, b, xStart, xEnd, step float64) []float64 {
	results := []float64{}
	for x := xStart; x <= xEnd; x += step {
		y := CalculateY(a, b, x)
		results = append(results, y)
	}
	return results
}

func SolveTaskB(a, b float64, xValues []float64) []float64 {
	results := []float64{}
	for _, x := range xValues {
		y := CalculateY(a, b, x)
		results = append(results, y)
	}
	return results
}
func StartLab4() {
	a := 0.4
	b := 0.8

	fmt.Println("Задача A:")
	xStartA := 3.2
	xEndA := 6.2
	stepA := 0.6
	resultsA := SolveTaskA(a, b, xStartA, xEndA, stepA)
	fmt.Println(resultsA)

	fmt.Println("Задача B:")
	xValuesB := []float64{4.48, 3.56, 2.78, 5.28, 3.21}
	resultsB := SolveTaskB(a, b, xValuesB)
	fmt.Println(resultsB)
=======
}
