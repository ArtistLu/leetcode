package main

import (
	"fmt"
	"strconv"
)

func main() {
	matrix := []int{
		1, 5, 8,
	}

	for _, n := range matrix {
		fmt.Println(n, " -> ", fizzBuzz(n))
	}
}

func fizzBuzz(n int) []string {
	re := []string{}

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			re = append(re, "FizzBuzz")
		} else if i%5 == 0 {
			re = append(re, "Buzz")
		} else if i%3 == 0 {
			re = append(re, "Fizz")
		} else {
			re = append(re, strconv.Itoa(i))
		}
	}

	return re
}
