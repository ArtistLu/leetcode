package main

import "fmt"

func main() {
	testMatix := [][]int{
		{1, 2, 4, 5, 6},
		{1, 2, 4, 5, 6, 9},
		{1, 2, 2, 3, 6, 4},
	}

	for _, nums := range testMatix {
		target := nums[len(nums)-1]
		nums := nums[0 : len(nums)-1]
		fmt.Print(nums, " -> ")
		indexs := twoSum(nums, target)
		fmt.Println(nums[indexs[0]], " + ", nums[indexs[1]], " = ", target)
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {

		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return []int{}
}
