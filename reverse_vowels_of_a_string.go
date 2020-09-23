package main

import "fmt"

func isVowel(c byte) bool {
	fmt.Println("For: ",c)
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func reverseVowels(s string) string {
	var newString string
	arr := make([]byte, len(s))

	for i:=0; i<len(s); i++ {
		arr[i] = s[i]
	}

	pointer1, pointer2 := 0, len(arr)-1

	for pointer1 < pointer2 {
		for pointer1 < pointer2 {
			if !isVowel(arr[pointer1]) {
				pointer1++
			} else {
				break
			}
		}

		for pointer1 < pointer2 {
			if !isVowel(arr[pointer2]) {
				pointer2--
			} else {
				break
			}
		}

		temp := arr[pointer1]
		arr[pointer1] = arr[pointer2]
		arr[pointer2] = temp
		pointer1++
		pointer2--
	}

	for _, val:=range arr {
		newString = newString + string(val)
	}

	return newString
}

func main() {
	slice := "a."

	fmt.Println("Reversing the Vowels: ",reverseVowels(slice))
}