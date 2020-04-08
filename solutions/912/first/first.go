package main

import "fmt"

func main() {
	arr := []int{1, 7, 2, 4}
	arr = []int{4, 1}
	re := sortArray(arr)

	fmt.Print(re)

}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, s, e int) {
	if s >= e {
		return
	}

	m := parting(nums, s, e)

	quickSort(nums, s, m-1)
	quickSort(nums, m+1, e)
}

func parting(nums []int, s, e int) int {
	tail := nums[e]

	i := s
	for j := s; j < e; j++ {
		if nums[j] < tail {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[e] = nums[e], nums[i]

	return i
}
