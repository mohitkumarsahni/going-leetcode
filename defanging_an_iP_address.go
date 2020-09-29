package main

import "fmt"

func defangIPaddr(address string) string {
	newStr := ""

	for _, val := range address {
		if val=='.' {
			newStr = newStr + "[.]"
		}else {
			newStr = newStr + string(val)
		}
	}

	return newStr
}

func main() {
	address := "1.1.1.1"

	fmt.Println("Defanging the IP: ",defangIPaddr(address))
}