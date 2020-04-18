package main

import "fmt"

func main() {
	testMatrix := [][]int{
		{},
		{-1, 1, 2},
		{-5, -1, 1, 2},
		{-9, -3, 0, 0, 6},
	}

	for _, test := range testMatrix {
		fmt.Print(test, "->")
		fmt.Println(sortedSquares(test))
	}
}

func sortedSquares(A []int) []int {
	l := len(A)
	if l == 0 {
		return A
	}

	A[l-1] *= A[l-1]
	for i := l - 2; i >= 0; i-- {
		A[i] *= A[i]
		j := i
		for j < l-1 && A[j] > A[j+1] {
			A[j], A[j+1] = A[j+1], A[j]
			j++
		}
	}

	return A
}
