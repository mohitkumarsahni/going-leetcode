package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

func (receiver *ListNode) traverse() {
	if receiver==nil {
		fmt.Println("List is Empty.")
		return
	}
	temp := receiver
	for temp!=nil {
		fmt.Printf("%d->",temp.Val)
		temp = temp.Next
	}
	fmt.Printf("\n")
	return
}

func addAtBack(l *ListNode, val int) *ListNode{
	if l == nil {
		l = &ListNode{Val: val, Next: nil}
		return l
	}

	var temp *ListNode = l
	for temp.Next!=nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{Val: val, Next: nil}

	return l
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Printf("Merging The Lists...\n")
	var l3 *ListNode = nil

	for l1!=nil || l2!=nil {
		if l1==nil && l2!=nil {
			if l3 == nil {
				l3 = l2
				return l3
			}
			for l2!=nil {
				l3 = addAtBack(l3, l2.Val)
				l2 = l2.Next
			}
			return l3
		}else if l1!=nil && l2==nil {
			if l3 == nil {
				l3 = l1
				return l3
			}
			for l1!=nil {
				l3 = addAtBack(l3, l1.Val)
				l1 = l1.Next
			}
			return l3
		}

		if l1.Val == l2.Val {
			l3 = addAtBack(l3, l1.Val)
			l3 = addAtBack(l3, l2.Val)
			l1 = l1.Next
			l2 = l2.Next
			continue
		}else if l1.Val < l2.Val {
			l3 = addAtBack(l3, l1.Val)
			l1 = l1.Next
			continue
		}else {
			l3 = addAtBack(l3, l2.Val)
			l2 = l2.Next
			continue
		}
	}

	return l3
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	l1 = addAtBack(l1, 5)
	l1.traverse()
	l1 = addAtBack(l1, 6)
	l1 = addAtBack(l1, 7)
	l1.traverse()

	fmt.Println("List L1: ")
	l1.traverse()
	fmt.Println("List L2: ")
	l2.traverse()

	l3 := mergeTwoLists(l1, l2)
	fmt.Println("Merged List L3: ")
	l3.traverse()
}