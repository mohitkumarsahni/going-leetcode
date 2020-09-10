package main

import "fmt"

type ListNode1 struct {
	Val int
	Next *ListNode1
}

func hasCycle(head *ListNode1) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == slow {
		return true
	}

	return false
}

func main() {
	list3 := &ListNode1{Val: 3, Next: nil}
	list2 := &ListNode1{Val: 2, Next: nil}
	list0 := &ListNode1{Val: 0, Next: nil}
	list4 := &ListNode1{Val: 4, Next: nil}

	list3.Next = list2
	list2.Next = list0
	list0.Next = list4
	list4.Next = list2

	fmt.Println("Is there a cycle: ",hasCycle(list3))
}