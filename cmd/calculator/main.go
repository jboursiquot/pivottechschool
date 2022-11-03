package main

import (
	"fmt"

	"github.com/jboursiquot/pivottechschool/calculator"
)

func main() {
	fmt.Printf("Add(1, 2) = %d\n", calculator.Add(1, 2))
	fmt.Printf("Subtract(1, 2) = %d\n", calculator.Subtract(1, 2))
	fmt.Printf("Multiply(2, 3) = %d\n", calculator.Multiply(2, 3))

	r, err := calculator.Divide(6, 3)
	fmt.Printf("Divide(6, 3) = %d, %v\n", r, err)

	r, err = calculator.Divide(6, 0)
	fmt.Printf("Divide(6, 3) = %d, %v\n", r, err)
}
