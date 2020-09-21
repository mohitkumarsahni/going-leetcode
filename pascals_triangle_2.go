package main

import "fmt"

func getRow(rowIndex int) []int {
	arr := make([][]int, rowIndex+1)

	for i:=0; i<=rowIndex; i++ {
		arr[i] = make([]int, i+1)
		for j:=0; j<=i; j++ {
			if j==0 || j==i {
				arr[i][j] = 1
				continue
			} else {
				arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
			}
		}
	}

	return arr[rowIndex]
}

func main() {
	row := 0;

	fmt.Println("Printing Rows of Pascals Triangle:",getRow(row));
}