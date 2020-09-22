package main

import "fmt"

func reverseString(s []byte) {
	pointer1 := 0
	pointer2 := len(s)-1

	for pointer1 < pointer2 {

		s[pointer1], s[pointer2] = s[pointer2], s[pointer1]

		pointer1++
		pointer2--
	}
	return
}

func main() {
	slice := []byte{'H','a','n','n','a','h'}

	fmt.Println(slice)
	fmt.Println("Reversing the Array: ")
	reverseString(slice)
	fmt.Println(slice)
}