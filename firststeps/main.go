package main

import "fmt"

func verifyValue(a float64) error {
	if a <= 0 {
		return fmt.Errorf("invalid number, zero or less")
	}
	return nil
}

func divide(a, b float64) float64 {
	return a / b
}

func sum(a, b float64) float64 {
	return a + b
}

func difference(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func main() {
	var fnOne float64
	var fnTwo float64
	fmt.Print("Provid a number: ")
	fmt.Scan(&fnOne)
	fmt.Print("Provid a second number: ")
	fmt.Scan(&fnTwo)

	err := verifyValue(fnTwo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The sum is %.2f\n", sum(fnOne, fnTwo))
	fmt.Printf("The difference is %.2f\n", difference(fnOne, fnTwo))
	fmt.Printf("The product is %.2f\n", multiply(fnOne, fnTwo))
	fmt.Printf("The quotient is %.2f\n", divide(fnOne, fnTwo))
}
