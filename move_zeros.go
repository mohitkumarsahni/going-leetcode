package main

import "fmt"

func moveZeroes(nums []int) {
	length := len(nums)
	if nums==nil || length==0 {
		return
	}

	zP := 0
	nzP := 1

	for zP<length && nzP<length {
		if nums[zP]==0 {
			if nums[nzP]==0 {
				nzP += 1
				continue
			}else{
				temp := nums[zP]
				nums[zP] = nums[nzP]
				nums[nzP] = temp
				zP += 1
				continue
			}
		}else {
			zP += 1
			nzP = zP+1
		}
	}
}

func moveZeroesTwoLoop(nums []int) {
	idx := 0

	for _, v := range nums {
		if v != 0 {
			nums[idx] = v
			idx++
		}
	}

	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	slice := []int{0,1,2}

	moveZeroesTwoLoop(slice)
	fmt.Println("Moving Zeros to last: ")
	fmt.Println(slice)
}