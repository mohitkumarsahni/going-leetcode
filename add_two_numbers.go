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


func addToBack(l *ListNode, value int) *ListNode {
	if l==nil {
		l = &ListNode{Val: value, Next: nil}
		return l
	}

	var temp *ListNode = l
	for temp.Next!=nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{Val: value, Next: nil}

	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode = nil
	carry := 0
	if l1==nil && l2==nil {
		return nil
	}

	for l1!=nil || l2!=nil {
		if l1==nil && l2!=nil {
			sum := l2.Val + carry
			if sum>9 {
				l3 = addToBack(l3, sum%10)
				carry = sum/10
			}else {
				l3 = addToBack(l3, sum)
				carry = 0
			}
			l2 = l2.Next
			continue
		}else if l1!=nil && l2==nil {
			sum := l1.Val + carry
			if sum>9 {
				l3 = addToBack(l3, sum%10)
				carry = sum/10
			}else {
				l3 = addToBack(l3, sum)
				carry = 0
			}
			l1 = l1.Next
			continue
		}

		sum := l1.Val + l2.Val + carry
		if sum>9 {
			l3 = addToBack(l3, sum%10)
			carry = sum/10
		}else {
			l3 = addToBack(l3, sum)
			carry = 0
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry!=0 {
		l3 = addToBack(l3, carry)
	}

	return l3
}

func main() {
	var l1 *ListNode = nil
	var l2 *ListNode = nil
	l1 = addToBack(l1, 2)
	l1 = addToBack(l1, 4)
	l1 = addToBack(l1, 3)

	l2 = addToBack(l2, 5)
	l2 = addToBack(l2, 6)
	//l2 = addToBack(l2, 9)

	l3 := addTwoNumbers(l1, l2)
	l3.traverse()
}
