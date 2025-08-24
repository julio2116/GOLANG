package main

import (
	"fmt"
	"strconv"
)

func convNumber(a string) (float64, error) {
	number, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("Cannot convert '%s' into a number: %w", a, err)
	}
	return number, nil
}

func calculator(a [2]float64, b string) (float64, error) {
	if len(a) > 2 {
		return 0, fmt.Errorf("Invalid values you must provid only two number, vlaues provideds: %.2f, %.2f", a[0], a[1])
	}

	switch b {
	case "*":
		return a[0] * a[1], nil
	case "/":
		if a[1] == 0 {
			return 0, fmt.Errorf("Error: division by zero")
		}
		return a[0] / a[1], nil
	case "-":
		return a[0] - a[1], nil
	case "+":
		return a[0] + a[1], nil
	default:
		return 0, fmt.Errorf("Provid of one the follow operators only: (+ - * /)")
	}
}

func main() {
	var values [3]string
	b := [...]string{"first number", "second number", "operation (+ - * /)"}

	for i := range len(values) {
		fmt.Printf("Enter %s:", b[i])
		fmt.Scan(&values[i])
		if values[i] == "" {
			fmt.Print("Error: must provid a value")
			return
		}
	}
	var numbers [2]float64

	for i := range 2 {
		number, err := convNumber(values[i])
		if err != nil {
			fmt.Printf("Error in conversion: %s", err)
			return
		}
		numbers[i] = number
	}

	result, err := calculator(numbers, values[2])

	if err != nil {
		fmt.Printf("Error calculating: %s", err)
		return
	}

	fmt.Printf("Result: %.2f", result)
}
