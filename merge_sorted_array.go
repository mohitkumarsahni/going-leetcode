package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	l1,l2,end := m-1,n-1,m+n-1
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	for l1>=0 && l2>=0 {
		if nums2[l2] > nums1[l1] {
			nums1[end] = nums2[l2]
			l2 -= 1
		} else {
			nums1[end] = nums1[l1]
			l1 -= 1
		}
		end -= 1
	}
	for ; l2 >= 0; end-- {
		nums1[end] = nums2[l2]
		l2 -= 1
	}
	return
}

func main() {
	slice1 := []int{1,2,3,0,0,0}
	slice2 := []int{2,5,6}
	merge(slice1, 3, slice2, 3)
	fmt.Println("Merge: ",slice1)
}