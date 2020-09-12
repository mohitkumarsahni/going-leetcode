package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	length := len(candies)
	max := candies[0]

	boolArr := make([]bool, length)

	for _, val := range candies {
		if max < val {
			max = val
		}
	}

	for i:=0; i<length; i++ {
		if candies[i]+extraCandies >= max {
			boolArr[i] = true
		}else {
			boolArr[i] = false
		}
	}

	return boolArr
}

func main() {
	slice := []int{12,1,12}
	extraCandies := 10

	fmt.Println("Running Sum of Array:",kidsWithCandies(slice, extraCandies))
}