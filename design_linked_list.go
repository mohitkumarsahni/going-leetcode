package main

import (
	"fmt"
)

type Node struct {
	val int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, size: 0}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.head == nil || index < 0 || index > this.size {
		return -1
	}

	temp := this.head
	for i:=0; temp!=nil; i++ {
		if i==index {
			return temp.val
		}
		temp = temp.next
	}

	return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.head == nil {
		this.head = &Node{val: val, next: nil}
		this.size += 1
		return
	}

	this.head = &Node{val: val, next: this.head}
	this.size += 1
	return
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.head = &Node{val: val, next: nil}
		this.size += 1
		return
	}

	temp := this.head
	for temp.next!=nil {
		temp = temp.next
	}

	temp.next = &Node{val: val, next: nil}
	this.size += 1
	return
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size+1 {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}else if index == this.size+1 {
		this.AddAtTail(val)
		return
	}

	temp := this.head
	for i:=0; temp!=nil; i++ {
		if i==index-1 {
			temp.next = &Node{val: val, next: temp.next}
			this.size += 1
			break
		}
		temp = temp.next
	}
	return
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size-1 {
		return
	}

	if index == 0 {
		this.head = this.head.next
		this.size -= 1
		return
	}

	temp := this.head
	for i:=0; temp!=nil; i++ {
		if i==index-1 {
			temp.next = temp.next.next
			this.size -= 1
			break
		}
		temp = temp.next
	}
	return
}

func (receiver *MyLinkedList) traverse() {
	if receiver.head==nil {
		fmt.Println("List is Empty.")
		return
	}
	temp := receiver.head
	for temp!=nil {
		fmt.Printf("%d->",temp.val)
		temp = temp.next
	}
	fmt.Printf("\n")
	return
}

func main() {
	obj := Constructor();
	obj.traverse()
	obj.AddAtHead(2)
	obj.traverse()
	println("Size: ",obj.size)
	obj.DeleteAtIndex(1)
	obj.traverse()
	obj.AddAtHead(2)
	obj.traverse()
	obj.AddAtHead(7)
	obj.traverse()
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(5)
	obj.AddAtTail(5)
	obj.traverse()
	obj.DeleteAtIndex(6)
	obj.DeleteAtIndex(4)

	obj.traverse()
}

//["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
//[[],[1],[3],[1,2],[1],[0],[0]]

//["MyLinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
//[[]				,[2]		,[1]			,[2]		,[7]		,[3]		,[2]		,[5]		,[5]	,[5]		,[6]		,[4]]