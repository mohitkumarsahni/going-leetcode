package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	jMap := make(map[rune]int, len(J))

	for _, val := range J {
		jMap[val] = 0
	}

	for _, val := range S {
		_, ok := jMap[val]
		if ok {
			jMap[val] = jMap[val] +1
		}
	}

	num := 0
	for _, val := range jMap {
		num = num + val
	}

	return num
}

func main() {
	J := "z"
	S := "ZZ"

	fmt.Println("how many of the stones you have are also jewels: ",numJewelsInStones(J, S))
}