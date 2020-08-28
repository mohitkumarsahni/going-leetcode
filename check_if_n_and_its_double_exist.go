package main

import "fmt"

func checkIfExistBruteForce(arr []int) bool {
	length := len(arr)
	if length==0 || arr==nil {
		return false
	}
	for i:=0;i<length;i++ {
		for j:=i+1;j<length;j++ {
			if arr[i]==(2*arr[j]) || arr[j]==(2*arr[i]) {
				return true
			}
		}
	}
	return false
}

func checkIfExist(arr []int) bool {
	length := len(arr)
	if length==0 || arr==nil {
		return false
	}

	maintain := make(map[int]int, length)

	for i:=0;i<length;i++ {
		if _, ok := maintain[2*arr[i]]; ok{
			return true
		}else if arr[i]%2==0 {
			if _, ok := maintain[arr[i]/2]; ok{
				return true
			}else {
				maintain[arr[i]] = i
			}
		}else {
			maintain[arr[i]] = i
		}
	}

	return false
}

func main() {
	slice1 := []int{10,2,5,3}
	fmt.Println("Does N & 2N Exist: ",checkIfExistBruteForce(slice1))
	fmt.Println("Does N & 2N Exist: ",checkIfExist(slice1))
}