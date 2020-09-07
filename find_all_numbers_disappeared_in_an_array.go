package main

import "fmt"

func absInt(x int) int {
	if x<0 {
		return -x
	}

	return x
}

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	missed := []int{}

	for i:=0;i<length;i++ {
		val := absInt(nums[i])-1
		if nums[val]>0 {
			nums[val] = -nums[val]
		}
	}

	for i:=0;i<length;i++ {
		if nums[i]>0{
			missed = append(missed, i+1)
		}
	}

	return missed
}

func main() {
	slice := []int{4,3,2,7,8,2,3,1}

	fmt.Println("Missing Elements: ",findDisappearedNumbers(slice))
}