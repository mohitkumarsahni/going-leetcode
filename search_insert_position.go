package main

import "fmt"

func searchInsert(nums []int, target int) int {
	first := 0
	last := len(nums)-1
	mid := 0
	for first<=last {
		mid = int(first+last)/2
		if nums[mid] == target {
			return mid
		}else {
			if nums[mid] < target {
				first = mid+1
			}else {
				last = mid-1
			}
		}
	}

	return first
}

func main() {
	arr := []int{1,3,5,6}
	target := 0
	fmt.Println("Element is present at or should be present at : ",searchInsert(arr, target))
}