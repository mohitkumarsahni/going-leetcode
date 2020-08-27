package main

import "fmt"

func duplicateZerosBruteForce(arr []int)  {
	length := len(arr)
	if arr==nil || length==0 {
		return
	}

	for i:=0;i<length;i++ {
		if arr[i]==0 {
			for j:=length-1;j>i;j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
	return
}

func duplicateZerosBigOn(arr []int) {
	coz := 0
	length := len(arr)
	if arr==nil || length==0 {
		return
	}

	for i := 0 ; i < length; i++ {
		if arr[i] == 0 {
			coz = coz + 1
		}
	}

	for i, j := length-1, length+coz-1; i>=0; i, j = i-1, j-1 {
		//System.out.println(i + "  " + j);
		if arr[i] == 0 {
			if j < length {
				arr[j] = arr[i]
			}
			j--
			if j < length {
				arr[j] = arr[i]
			}
		}else{
			if j < length {
				arr[j] = arr[i]
			}
		}
	}
	return
}

func main() {
	slice1 := []int{1,2,3,0,4,5}
	duplicateZerosBigOn(slice1)
	fmt.Println("New Slice1: ",slice1)
}
