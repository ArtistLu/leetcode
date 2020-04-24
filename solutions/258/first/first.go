package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 0, 5, 0, 7, 10, 11, 12, 100, 123, 19}

	for _, num := range test {
		fmt.Println(num, "|", addDigits(num))
	}
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	re := 0
	for num > 0 {
		re += num % 10
		num /= 10
	}

	return addDigits(re)
}
