package main

import "fmt"

func restoreString(s string, indices []int) string {
	newStr := make([]byte, len(indices))

	for i:=0; i<len(indices); i++ {
		newStr[indices[i]] = s[i]
	}

	newString := ""
	for _, val := range newStr {
		newString = newString + string(val)
	}

	return newString
}

func main() {
	s := "codeleet"
	indices := []int{4,5,6,7,0,2,1,3}

	fmt.Println("Shuffled String: ", restoreString(s, indices))
}