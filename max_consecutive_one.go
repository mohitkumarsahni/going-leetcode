package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	length := len(nums)
	count := 0
	max := 0
	for i:=0; i<length; i++ {
		if nums[i] == 1{
			count += 1
		}else{
			if max < count {
				max = count
			}
			count = 0
		}
	}
	if max < count {
		max = count
	}

	return max
}

func main() {
	nums := []int{1,0,1}
	fmt.Println("Maximum Consicutive Ones: ",findMaxConsecutiveOnes(nums))
}
