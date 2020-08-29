package main

import "fmt"

func validMountainArraySinglePointer(A []int) bool {
	length := len(A)
	if A==nil || length<3 {
		return false
	}

	i := 0
	for (i+1 < length)&&(A[i]<A[i+1]) {
		i += 1
	}

	if (i==0)||(i==length-1) {
		return false
	}

	for (i+1 < length)&&(A[i]>A[i+1]) {
		i += 1
	}

	return i==length-1
}

func main() {
	slice := []int{2,0,2}
	fmt.Println("Is Valid Maountain Array: ",validMountainArraySinglePointer(slice))
}
