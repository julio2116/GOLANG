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

func calcSum(a int) (int, error){
	if a > 1000000 || a < -1000000 {
		return 0, fmt.Errorf("Error: too large number limit from -1000000 to 1000000")
	}

	sum := 0

	if a < 0 {
		for i := a; i <= 0; i++ {
			fmt.Println(i)
			sum += i
		}
		return sum, nil
	}

	for i := 0; i <= a; i++ {
		fmt.Println(i)
		sum += i
	}
	return sum, nil
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

	sum, calcErr := calcSum(number)
	if calcErr != nil {
		fmt.Println(calcErr)
		return
	}

	fmt.Printf("The total sum is %d", sum)

}
