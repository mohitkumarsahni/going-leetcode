package main

import "fmt"

func reverseInt(x int) int {
	var temp int = 0
	var ld int = 0
	var temp_n int = x

	if x > 2147483647 || x < -2147483648 {
		return 0
	}

	if x < 0 {
		temp_n = temp_n*-1
	}

	for temp_n>0 {
		ld = temp_n%10
		temp_n = temp_n/10
		temp = temp*10+ld
	}

	if x < 0 {
		temp = temp*-1
	}

	if temp > 2147483647 {
		return 0
	}else if temp < -2147483648 {
		return 0
	}
	return temp
}

func main()  {
	var x int = -214748364

	fmt.Println("Reversing the Number: ",x)
	fmt.Println(reverseInt(x))
}