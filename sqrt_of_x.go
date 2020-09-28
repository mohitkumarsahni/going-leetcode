package main

import "fmt"

func mySqrt(x int) int {
	if x==0 || x==1 {
		return x
	}

	start, end := 0, x
	sqrt := 0

	for start<=end {
		mid := (start+end)/2

		if (mid*mid)==x {
			return mid
		}

		if (mid*mid)<x {
			start = mid+1
			sqrt = mid
		}else {
			end = mid-1
		}
	}

	return sqrt
}

func main() {
	x := 30

	fmt.Println("Sqrt of X: ",mySqrt(x))
}