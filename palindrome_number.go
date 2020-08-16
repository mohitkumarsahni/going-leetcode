package main

import "fmt"

func palindromeCheck(n int) bool {
	var reverseN int
	var originalN int = n

	if originalN < 0 {
		return false
	}

	for n!=0 {
		reverseN = (reverseN*10) + (n%10)
		n = n/10
	}

	return reverseN == originalN
}

func main()  {
	n := 1421

	fmt.Println("Result of Palindrome Check for Number: ",n)
	fmt.Print(palindromeCheck(n))
}
