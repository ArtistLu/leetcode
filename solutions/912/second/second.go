package main

import "fmt"

func main() {
	arr := []int{1, 7, 2, 4}
	// arr = []int{4, 1}
	re := sortArray(arr)

	fmt.Print(re)
}

func sortArray(nums []int) []int {
	mm := make([]int, 100001)
	for i := 0; i < len(nums); i++ {
		mm[nums[i]]++
	}
	var out []int
	for i := 0; i < len(mm); i++ {
		for mm[i] > 0 {
			out = append(out, i)
			mm[i]--
		}
	}
	return out
}
