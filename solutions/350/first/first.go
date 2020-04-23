package main

import "fmt"

func main() {
	restMatrix := [][][]int{
		{{1, 2, 3}, {2, 5, 4, 3}},
		{{}, {2, 5, 4, 3}},
		{{1, 2, 3}, {}},
		{{1, 2, 3}, {2}},
		{{1, 1, 3, 7}, {3, 11, 7, 1}},
	}

	for _, arg := range restMatrix {
		fmt.Println(arg[0], "|", arg[1], " -> ", intersect(arg[0], arg[1]))
	}
}

func intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	m := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	res := []int{}
	for _, v := range nums2 {
		if c, ok := m[v]; ok && c > 0 {
			m[v]--
			res = append(res, v)
		}
	}
	return res
}
