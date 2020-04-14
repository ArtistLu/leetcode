package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	rotate(nums, 2)
	fmt.Print(nums)
}

func rotate(nums []int, k int) {
	l := len(nums)
	temp := make([]int, l)
	k = k % l
	for i, v := range nums {
		temp[(i+k)%l] = v
	}

	copy(nums, temp)
}
