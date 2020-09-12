package main

import "fmt"

func runningSum(nums []int) []int {
	length := len(nums)

	for i:=1; i<length; i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}

func main() {
	slice := []int{3,1,2,10,1}

	fmt.Println("Running Sum of Array:",runningSum(slice))
}