package main

import "fmt"

func sortArrayByParity(A []int) []int {
	length := len(A)
	if A==nil || length==0 {
		return nil
	}

	fP := 0
	lP := length-1
	arr := make([]int, length)

	for _,val := range A {
		if val%2==0 {
			arr[fP] = val
			fP += 1
		}else {
			arr[lP] = val
			lP -= 1
		}
	}

	return arr
}

func main() {
	slice := []int{3,1,2,4,0}
	fmt.Println("Re-Arranging the array",sortArrayByParity(slice))
}