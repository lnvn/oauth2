package main

import (
	"errors"
	"fmt"
)

func divide(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("can not devide by zero")
	}

	result := numerator / denominator
	return result, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Result is:", result)
	}
	fmt.Println("-------------------------")
	result1, err1 := divide(10, 0)
	if err1 != nil {
		fmt.Println("error:", err1)
	} else {
		fmt.Println("Result is:", result1)
	}
}
