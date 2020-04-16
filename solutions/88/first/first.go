package main

import "fmt"

func main() {
	testMatrix := [][2][]int{
		{{1, 2, 3}, {1, 2, 3}},
		{{4, 5, 6}, {1, 2, 3}},
		{{}, {1, 2, 3}},
	}
	for _, test := range testMatrix {
		nums1 := test[0]
		nums2 := test[1]
		m := len(nums1)
		n := len(nums2)
		for i := 0; i < len(nums2); i++ {
			nums1 = append(nums1, 0)
		}

		fmt.Print(nums1, " merge ", nums2)
		merge(nums1, m, nums2, n)
		fmt.Print("->", nums1)
		fmt.Println()
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
		for j := m + i; j > 0; j-- {
			if nums1[j] < nums1[j-1] {
				nums1[j], nums1[j-1] = nums1[j-1], nums1[j]
			} else {
				break
			}

		}
	}
}
