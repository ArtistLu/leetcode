package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 4}
	re := removeDuplicates(nums)
	fmt.Print(re, nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
			return 0
	}
	
	prev := nums[0]
	slot := 1
	for i := 1; i < len(nums); i++ {
			n := nums[i]
			if n == prev {
					continue
			}
			
			nums[slot] = n
			slot++
			prev = n
	}
	
	return slot
}