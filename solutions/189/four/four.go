package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	rotate(nums, 2)
	fmt.Print(nums)
}

func rotate(nums []int, k int) {
	numsLen := len(nums)
	if numsLen != 1 {
		clone := nums[numsLen-(k%numsLen):]
		latter := nums[:numsLen-(k%numsLen)]
		clone = append(clone, latter...)
		copy(nums, clone)
	}
}
