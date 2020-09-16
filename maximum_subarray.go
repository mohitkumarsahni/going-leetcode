package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	length := len(nums)
	maxCurrent, maxGlobal := nums[0], nums[0]

	for i := 1; i < length ; i++ {
		maxCurrent = int(math.Max(float64(nums[i]), float64(maxCurrent+nums[i])))
		if maxCurrent > maxGlobal {
			maxGlobal = maxCurrent;
		}
	}

	return maxGlobal;
}

func main() {
	slice := []int{-2,1,-3,4,-1,2,1,-5,4}

	fmt.Println("Maximum Sum of Contiguous sub-array: ",maxSubArray(slice))
}