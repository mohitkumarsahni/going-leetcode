package main

import "fmt"

func countDigitsInNumber(num int) int {
	count := 0
	for num>0 {
		count += 1
		num = num/10
	}

	return count
}

func findNumbers(nums []int) int {
	count := 0
	for _, i := range nums {
		x := countDigitsInNumber(i)
		if x%2==0{
			count += 1
		}
	}
	return count
}

func main() {
	nums := []int{555,901,482,12}
	fmt.Println("Number of even number of digits numbers are : ",findNumbers(nums))
}