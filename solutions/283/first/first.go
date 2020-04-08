package main

import "fmt"

func main() {
	nums := []int{0, 1}
	moveZeroes(nums)
	fmt.Print(nums)
}

func moveZeroes(nums []int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] != 0 {
			continue
		}
		for j := i + 1; j < l; j++ {
			if nums[j] == 0 {
				continue
			}
			nums[j], nums[i] = nums[i], nums[j]
			break
		}
	}
}
