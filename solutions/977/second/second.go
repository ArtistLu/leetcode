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
	for j := len(A) - 1; j >= 0; j-- {
		f := A[0] * A[0]
		cur := A[j] * A[j]
		if f > cur {
			for i := 1; i <= j; i++ {
				A[i-1] = A[i]
			}
			A[j] = f
		} else {
			A[j] = cur
		}
	}

	return A
}
