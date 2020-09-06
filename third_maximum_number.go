package main

import "math"

func thirdMax(nums []int) int {
	max1 := math.MinInt32
	max2 := math.MinInt32
	max3 := math.MinInt32

	for _,v := range nums {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v < max1 && v > max2 {
			max3 = max2
			max2 = v
		} else if v < max2 && v > max3 {
			max3 = v
		}
	}

	if max3 == math.MinInt32 {
		return max1
	}
	return max3
}

func main() {
	slice := []int{1, 1, 2}

	println("Third Maximum Number/Maximum Number: ",thirdMax(slice))
}