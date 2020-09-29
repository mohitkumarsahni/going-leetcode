package main

import "fmt"

func numberOfSteps (num int) int {
	if num==0 {
		return 0
	}

	steps := 0
	for num!=0 {
		if num%2 == 0 {
			num = num/2
			steps = steps + 1
		}else if (num%2!=0) || (num == 1) {
			num = num - 1
			steps = steps + 1
		}
	}

	return steps
}

func main() {
	num := 3

	fmt.Println("Number of Steps to reduce the given number to 0: ",numberOfSteps(num))
}