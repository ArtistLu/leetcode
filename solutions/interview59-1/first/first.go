package main

import "fmt"

func main() {
	testMatrix := [][]int{
		{1, 3, 2, 5, 5, 7},
		{9, 8, 2, 1, 9, 36, 8},
		{},
		{2},
	}

	testMatrixK := []int{
		3, 4, 5, 1,
	}

	for i, nums := range testMatrix {
		fmt.Println(nums, "|", testMatrixK[i], "->", maxSlidingWindow(nums, testMatrixK[i]))
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length == 1 || k <= 1 {
		return nums
	}

	wi, re := make([]int, 0), make([]int, 0)

	for i, n := range nums {
		if i > k-1 && wi[0] <= i-k {
			wi = wi[1:]
		}

		for len(wi) != 0 && n >= nums[wi[len(wi)-1]] {
			wi = wi[0 : len(wi)-1]
		}
		wi = append(wi, i)
		if i >= k-1 {
			re = append(re, nums[wi[0]])
		}
	}

	return re
}
