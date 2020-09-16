package main

import "fmt"

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	digits = append(digits, 0)
	copy(digits[1:], digits[0:])
	digits[0] = 1
	return digits
}

func main() {
	slice := []int{0}

	fmt.Println("Adding 1: ",plusOne(slice))
}
