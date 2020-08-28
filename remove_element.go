package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length==0 || nums==nil {
		return 0
	}

	countOfVal := 0
	firstP := 0
	secondP := length-1

	for ;firstP<length && secondP>=firstP; {
		if nums[firstP]==val {
			countOfVal += 1
			if nums[secondP]==val {
				secondP -= 1
				continue
			}else{
				nums[firstP]=nums[secondP]
				secondP -= 1
			}
		}
		firstP += 1
	}

	return length-countOfVal
}


//Need to be coded.
func main() {
	slice1 := []int{0,1,2,2,3,0,4,2}
	val := 2
	fmt.Println("Removing the Val:",val,". | Number of values removed aka New Length: ",removeElement(slice1, val))
	fmt.Println(slice1)
}