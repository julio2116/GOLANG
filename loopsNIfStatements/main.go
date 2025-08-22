package main

import (
	"fmt"
	"strconv"
)

func convertIntoNumber(a string) (int, error) {
	number, err := strconv.Atoi(a)
	if err != nil {
		return 0, fmt.Errorf("Cannot convert ' %s ', into a number: %w", a, err)
	}
	return number, nil
}

func main() {
	var numberString string
	fmt.Print("Enter a number: ")
	fmt.Scan(&numberString)

	number, err := convertIntoNumber(numberString)

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for i := 1; i <= number; i++ {
		fmt.Println(i)
		sum += i
	}

	fmt.Printf("The total sum is %d", sum)

}
