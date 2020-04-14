package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 4}
	re := removeDuplicates(nums)
	fmt.Print(re, nums)
}

func removeDuplicates(nums []int) int {
	for index := 1; index < len(nums); index++ {
		if nums[index-1] == nums[index] {
			nums = append(nums[:index], nums[index+1:]...)
			index--
		}
	}
	return len(nums)
}
