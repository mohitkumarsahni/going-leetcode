package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s)==0 {
		return 0
	}

	strArray := strings.Split(s, " ")

	for i:=len(strArray)-1; i>=0; i-- {
		if strArray[i]=="" {
			continue
		}else {
			return len(strArray[i])
		}
	}

	return 0
}

func main() {
	slice := "s"

	fmt.Println("Length of last word: ",lengthOfLastWord(slice))
}