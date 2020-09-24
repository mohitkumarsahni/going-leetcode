package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	ransomCharCounter := make([]int, 26)
	for _, ch := range ransomNote {
		ransomCharCounter[ch-'a']++
	}

	for _, ch := range magazine {
		ransomCharCounter[ch-'a']--
	}

	for _, val := range ransomCharCounter {
		if val > 0 {
			return false
		}
	}

	return true
}

func main() {
	ransomNote := ""
	magazine := ""

	fmt.Println("Output: ",canConstruct(ransomNote, magazine))
}
