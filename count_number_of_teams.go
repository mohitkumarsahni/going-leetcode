package main

import "fmt"

func numTeams(rating []int) int {
	teams := 0
	length := len(rating)
	for i := 0; i < length; i++ {
		leftSmaller, rightLarger := 0, 0
		leftLarger, rightSmaller := 0, 0
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				leftSmaller++
			} else if rating[j] > rating[i] {
				leftLarger++
			}
		}
		for k := i + 1; k < length; k++ {
			if rating[i] < rating[k] {
				rightLarger++;
			} else if rating[i] > rating[k] {
				rightSmaller++;
			}
		}
		teams += leftSmaller * rightLarger + leftLarger * rightSmaller
	}

	return teams
}

func main() {
	rating := []int{1,2,3,4}

	fmt.Println("Total number of teams: ",numTeams(rating))
}