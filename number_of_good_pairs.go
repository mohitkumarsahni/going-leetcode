package main

import "fmt"

func numIdenticalPairsBruteForce(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}

	pairs := 0
	for i:=0; i < length; i++ {
		for j:=i+1; j < length; j++ {
			if nums[i] == nums[j] {
				pairs += 1
			}
		}
	}

	return pairs
}

func numIdenticalPairs(nums []int) int {
	freqmap := make(map[int]int)

	for _, val := range nums {
		_, value := freqmap[val]
		if value!=false {
			freqmap[val] += 1
		}else {
			freqmap[val] = 1
		}
	}

	pairs := 0

	for _, val := range freqmap {
		temp := val*(val-1)/2

		pairs = pairs + temp
	}

	return pairs
}

func main() {
	slice := []int{1,2,3}

	fmt.Println("Number of Good Pairs via brute force: ",numIdenticalPairsBruteForce(slice))
	fmt.Println("Number of Good Pairs: ",numIdenticalPairs(slice))
}