package main

import "fmt"

func reverseInt(x int) int {
	var temp int = 0

	if x > 2147483647 || x < -2147483648 {
		return 0
	}
	for x!=0 {
		temp = (temp*10) + (x%10)
		if temp > 2147483647 || temp < -2147483648{
			return 0
		}
		x = x/10
	}
	return temp
}

func main()  {
	var x int = -214748367

	fmt.Println("Reversing the Number: ",x)
	fmt.Println(reverseInt(x))
}
