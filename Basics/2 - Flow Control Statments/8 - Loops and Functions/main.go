package main

import (
	"fmt"
	"math"
)

// Sqrt calculates the square root of x using Newton's method.
func Sqrt(x float64) float64 {
	z := 1.0 // initial guess
	for i := 0; i < 10; i++ {
		fmt.Printf("Iteration %d: z = %f\n", i+1, z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	// Test the Sqrt function with various values of x
	values := []float64{1, 2, 3, 4, 9, 16, 25}
	for _, x := range values {
		fmt.Printf("Calculating square root of %f\n", x)
		result := Sqrt(x)
		fmt.Printf("Approximated square root: %f\n", result)
		fmt.Printf("Math library sqrt: %f\n", math.Sqrt(x))
		fmt.Println()
	}
}
