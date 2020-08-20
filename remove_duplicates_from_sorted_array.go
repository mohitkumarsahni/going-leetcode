package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length==0 {
		return 0
	}

	currentPointer := 0
	uniquePointer := 1

	for uniquePointer<length {
		if nums[currentPointer] == nums[uniquePointer] {
			uniquePointer += 1
			continue
		}
		currentPointer += 1
		nums[currentPointer] = nums[uniquePointer]
	}

	return currentPointer+1
}

func main() {
	nums := []int{-7,-2,1,8,8,9,9,9}
	fmt.Println("Initial Nums: ",nums)
	fmt.Println("Performing Duplicate Removal...")
	len := removeDuplicates(nums)
	fmt.Println("Nums after: ",nums," | Length till Unique Element: ",len)
}
