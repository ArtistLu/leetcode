package main

func main() {

}

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	//判断线段是否相交
	//不交返回空数组
	if overlap(start1, end1, start2, end2) == false {
		return []float64{}
	}

	if parallel(start1, end1, start2, end2) {

	}

	return overlapPoint(start1, end1, start2, end2)
	//相交判断是否平行，平行返回交点最小值
	//相交返回

	return []float64{1, 2}
}

func overlap(start1 []int, end1 []int, start2 []int, end2 []int) bool {
	xmin, xmax := start1[0], end1[0]
	if xmax < xmin {
		xmin, xmax = xmax, xmin
	}
	if start2[0]
}

func parallel(start1 []int, end1 []int, start2 []int, end2 []int) bool {
	return true
}

func overlapPoint(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	return []float64{}
}
