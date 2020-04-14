package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 4}
	re := removeDuplicates(nums)
	fmt.Print(re, nums)
}

func removeDuplicates(nums []int) int {
	var j int

	for i := 0; i < len(nums); i, j = i+1, j+1 {
		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
		}
		nums[j] = nums[i]
	}

	return j
}
