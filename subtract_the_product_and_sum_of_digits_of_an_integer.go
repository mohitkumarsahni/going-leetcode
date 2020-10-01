package main

import "fmt"

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n>0 {
		temp := n%10
		product = product * temp
		sum = sum + temp

		n = n/10
	}

	return product-sum
}

func main() {
	n := 1

	fmt.Println("difference between the product of its digits and the sum of its digits: ", subtractProductAndSum(n))
}