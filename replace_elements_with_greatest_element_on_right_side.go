package main

import "fmt"

func replaceElements(arr []int) []int {
	length := len(arr)
	if arr==nil || length==0 {
		return nil
	}

	max := 0
	arr1 := make([]int,length)
	for i:=0;i<length;i++ {
		max = 0
		for j:=i+1;j<length;j++ {
			if max<arr[j] {
				max = arr[j]
			}
		}
		arr1[i] = max
	}

	arr1[length-1] = -1

	return arr1
}

func replaceElementsOneLoop(arr []int) []int {
	length := len(arr)
	if arr==nil || length==0 {
		return nil
	}

	max := -1
	for i:=length-1; i>=0; i-- {
		temp := arr[i]
		arr[i] = max
		if temp > max {
			max = temp
		}
	}

	return arr
}

func main() {
	slice := []int{17,18,5,4,6,1}

	fmt.Println("Replace Elements with Greatest Element on Right Side: ",replaceElementsOneLoop(slice))
}
