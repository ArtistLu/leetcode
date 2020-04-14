package main

import "fmt"

func main() {
	digits := []int{9, 9, 9}
	fmt.Print(plusOne(digits))
}

func plusOne(digits []int) []int {
	e := len(digits) - 1
	digits[e]++
	for i := e - 1; i >= 0; i-- {
		if digits[i+1] == 10 {
			digits[i+1] = 0
			digits[i]++
		} else {
			break
		}
	}

	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}
