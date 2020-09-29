package main

import "fmt"

func singleNumber(nums []int) int {
	xored := 0

	for _, val := range nums {
		xored = xored ^ val
	}

	return xored
}

func main() {
	slice := []int{2,2,1,3,3}

	fmt.Println("Single Number", singleNumber(slice))
}