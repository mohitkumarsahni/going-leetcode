package main

import (
	"sort"
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var prefix string = ""

	if len(strs)==0 {
		return prefix
	}else if len(strs)==1 {
		return strs[0]
	}

	sort.Strings(strs)
	fmt.Print(strs)
	x := strs[0]
	y := strs[len(strs)-1]

	for i:=0;i<len(x);i++ {
		if len(x)==0 {
			break
		}

		if x[i]==y[i] {
			prefix = prefix + string(x[i])
		}else{
			break
		}
	}

	return prefix
}

func main() {
	b := []string{"aa","a"}
	fmt.Println("Finding LCP: ", longestCommonPrefix(b))
}