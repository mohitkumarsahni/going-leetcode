package main

import (
	"fmt"
	"sort"
)

func heightCheckerBruteForce(heights []int) int {
	length := len(heights)

	fP := 0
	aP := 1
	movements := 0
	min := heights[0]
	minI := 0

	for fP<length {
		min = heights[fP]
		minI = fP
		aP = fP + 1
		for aP<length {
			if heights[aP]<min {
				min = heights[aP]
				minI = aP
			}
			aP += 1
			continue
		}
		if minI!=fP {
			movements += 1
			temp := heights[fP]
			heights[fP] = heights[minI]
			heights[minI] = temp
		}
		fP += 1
	}

	if movements==0 {
		return 0
	}

	return movements+1
}

func heightChecker(heights []int) int {
	length := len(heights)

	arr := make([]int, length)
	copy(arr, heights)
	sort.Ints(arr)

	count := 0
	for i:=0;i<length;i++ {
		if arr[i]!=heights[i] {
			count += 1
		}
	}

	return count
}

func main() {
	slice := []int{2,1,2,1,1,2,2,1}

	fmt.Println(slice)
	fmt.Println("Number of Moved Students: ",heightChecker(slice))
	fmt.Println(slice)
}
