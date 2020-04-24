package main

import "fmt"

func main() {
	test := [][]int{
		{1, 2, 3, 0, 5, 0, 7},
		{0, 0, 1, 5, 0, 9},
		{},
		{0, 0},
	}

	for _, nums := range test {
		fmt.Print(nums, "|")
		moveZeroes(nums)
		fmt.Println(nums)
	}
}

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
