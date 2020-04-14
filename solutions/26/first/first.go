package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 4}
	re := removeDuplicates(nums)
	fmt.Print(re, nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		r := false
		for k := 0; k <= i; k++ {
			if nums[j] == nums[k] {
				r = true
				break
			}
		}

		if r == false {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
