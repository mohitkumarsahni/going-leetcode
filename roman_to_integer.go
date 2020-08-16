package main

import "fmt"

var romanSymbols = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInteger(s string) int {
	total := 0
	for i:=0; i<len(s); i++ {
		if i==(len(s)-1) {
			total = total + romanSymbols[string(s[i])]
			break
		}

		if romanSymbols[string(s[i])] >= romanSymbols[string(s[i+1])] {
			total = total + romanSymbols[string(s[i])]
		}else{
			total = total - romanSymbols[string(s[i])]
		}
	}

	return total
}

func main()  {
	roman := "DCCCLVI"

	fmt.Print("Converting the Roman To Integer:",romanToInteger(roman))
}
