package main

func main() {

}

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if j != i {
				nums[i] = 0
			}
			j++
		}
	}
}
