package main

import (
	"fmt"
)

func main() {
	var result float64

	var a float64 = 0.4
	var b float64 = 0.8
	var x [5]float64 = [5]float64{4.48, 3.56, 2.78, 5.28, 3.21}

	fmt.Println("\nзадача А:")
	for i := 3.2; i < 6.2; i += 0.6 {
		result = lab4.colc(a, b, i)
		fmt.Println("При x = %.2f, y = %.2f \n", i, result)
	}

	fmt.Println("\nЗадача В:")
	for i := 0; i < len(x); i++ {
		result = lab4.colc(a, b, x[i])
		fmt.Println("При x =", x[i], "y =", result)
	}
}
