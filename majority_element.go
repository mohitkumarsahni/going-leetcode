package main

import "fmt"

func majorityElement(nums []int) int {
	freqMap := make(map[int]int)

	if len(nums) == 1 {
		return nums[0]
	}

	for _, val := range nums {
		_, ok := freqMap[val]
		if ok {
			freqMap[val] += 1
			if freqMap[val] > len(nums)/2 {
				return val
			}
		}else {
			freqMap[val] = 1
		}
	}

	return 0
}

func main() {
	slice := []int{3,2,3}

	fmt.Println("Majority Element: ",majorityElement(slice))
}