package main

import "fmt"

func smallerNumbersThanCurrentBruteForce(nums []int) []int {
	arr := make([]int, len(nums))

	for i:=0; i<len(nums); i++ {
		count := 0
		for j:=0; j<len(nums); j++ {
			if nums[j]<nums[i] {
				count += 1
			}
		}
		arr[i] = count
	}

	return arr
}

func main() {
	slice := []int{8,1,2,2,3}

	fmt.Println("via brute force: ",smallerNumbersThanCurrentBruteForce(slice))

}