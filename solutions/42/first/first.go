package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 0, 2}
	fmt.Print(trap(height))
}

func trap(height []int) int {
	//1.去掉头部为0元素
	//2.去掉尾部为0元素
	//3.遍历去掉后数组，为0的项累加雨水值，不为0的项值减一
	//4.重复以上三步直到数组项小于3

	re := 0
	for len(height) > 2 {

		s := 0
		for i := 0; i < len(height)-1; i++ {
			if height[i] != 0 {
				break
			}
			s = i + 1
		}
		height = height[s:]

		e := len(height)
		for i := len(height) - 1; i >= 0; i-- {
			if height[i] != 0 {
				break
			}
			e = i
		}
		height = height[0:e]

		for i := 0; i < len(height); i++ {
			if height[i] > 0 {
				height[i]--
			} else {
				re++
			}
		}
	}

	return re
}
