package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	rotate(nums, 6)
	fmt.Print(nums)
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	for i := 0; i < k; i++ {
		n := nums[l-1]
		for j := l - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = n
	}
}
