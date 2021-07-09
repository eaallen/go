package main

import (
	"fmt"
)

type Node struct {
	value string
	next  *Node
}

type Methods interface {
	test() string
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) add(node *Node) {
	// nodes are passed to the function, but we only use the pointers
	self := *ll
	if self.head == nil {
		self.head = node
	} else {
		temp := self.head
		self.head = node
		self.head.next = temp
	}
}

func New() LinkedList {
	return LinkedList{}
}

var LL = New()

func main() {
	fmt.Println(LL)

	// add a new node
	n := Node{value: "new node"}
	LL.add(&n)
	fmt.Println(LL.head.next)
}
