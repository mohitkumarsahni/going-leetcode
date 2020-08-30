package main

import "fmt"

func strStrBruteForce(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)

	if needleLen==0 {
		return 0
	}else if haystackLen==0 || haystackLen<needleLen {
		return -1
	}

	isPresent := false
	for i:=0;i<haystackLen;i++ {
		j:=0
		k:=i
		for j=0;j<needleLen;j++ {
			if k<haystackLen && haystack[k]==needle[j] {
				k += 1
				isPresent = true
				continue
			}else {
				isPresent = false
				break
			}
		}
		if j==needleLen && isPresent {
			return k-j
		}
	}

	return -1
}

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if haystack == needle {
		return 0
	}
	for i := 0; i < haystackLen - needleLen + 1; i++ {
		if (haystack[i: i + needleLen] == needle) {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "mississippi"
	needle := "issipi"

	fmt.Println("Is the needle present in haystack: ",strStr(haystack, needle))
}