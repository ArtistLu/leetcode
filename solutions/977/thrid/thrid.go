package main

import "fmt"

func main() {
	testMatrix := [][]int{
		{},
		{-1, 1, 2},
		{-5, -2, -1, 1, 2},
		{-9, -3, 0, 0, 6},
	}

	for _, test := range testMatrix {
		fmt.Print(test, "->")
		fmt.Println(sortedSquares(test))
	}
}

func sortedSquares(A []int) []int {
	length := len(A)
	re := make([]int, length)
	i, j, k := 0, length-1, length-1

	for i <= j {
		l := A[i] * A[i]
		r := A[j] * A[j]
		if l > r {
			re[k] = l
			i++
		} else {
			re[k] = r
			j--
		}
		k--
	}

	return re
}
