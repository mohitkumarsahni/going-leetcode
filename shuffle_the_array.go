package main

import "fmt"

func shuffle(nums []int, n int) []int {
	i := 0
	j := n

	arr := make([]int, 2*n)
	for k:=0;k<2*n;k+=2 {
		arr[k] = nums[i]
		arr[k+1] = nums[j]

		i+=1
		j+=1
	}

	return arr
}

func main() {
	slice := []int{2,5,1,3,4,7}

	fmt.Println("Running Sum of Array:",shuffle(slice, len(slice)/2))
}