package main

import (
	"fmt"
	"strconv"
)

type Group struct {
	size int
	space int
	arr []int
}

func groupThePeople(groupSizes []int) [][]int {
	groupMap := make(map[string]*Group)
	groupMapSize := 0

	for i:=0; i<len(groupSizes); i++ {
		grouped := false
		for _, val := range groupMap {
			if val.size == groupSizes[i] && val.space!=0 {
				val.arr = append(val.arr, i)
				val.space = val.space - 1
				grouped = true
				break
			}
		}
		if !grouped {
			groupMapSize++
			group := Group{space: groupSizes[i]-1, size: groupSizes[i], arr: []int{i}}
			groupMap[strconv.Itoa(groupMapSize)+strconv.Itoa(groupSizes[i])] = &group
		}
	}

	finalArray := make([][]int, groupMapSize)
	i := 0
	for _, val := range groupMap {
		finalArray[i] = make([]int, val.size)
		copy(finalArray[i], val.arr)
		i++
	}

	return finalArray
}

func main () {
	groupSizes := []int{2,1,3,3,3,2}

	fmt.Println("Grouping the People Given the Group Size They Belong To: ",groupThePeople(groupSizes))
}