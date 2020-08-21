package main

import "fmt"

func sortedSquaresOfSortedArray(nums []int) []int {
	length := len(nums)
	arr := make([]int, length)
	var l, r int = 0, length-1

	for l<=r {
		left := nums[l]*nums[l]
		right := nums[r]*nums[r]
		if left>right {
			arr[r-l] = left
			l += 1
		}else{
			arr[r-l] = right
			r -= 1
		}
	}
	return arr
}

func main() {
	nums := []int{-7,-3,2,3,11}
	fmt.Println("Sorted Array: ",nums)
	fmt.Println("Sorted array of squares of Sorted Array:",sortedSquaresOfSortedArray(nums))
}
