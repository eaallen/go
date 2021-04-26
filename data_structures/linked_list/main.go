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
	head Node
}

func (ll *LinkedList) add(node Node) {
	println(ll)
	self := *ll
	if self.head.value == nil {
		// head is empty, so just put the value there
		ll.head = node
	} else {
		temp := ll.head
		ll.head = node
		ll.head.next = &temp
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
	LL.add(n)
	fmt.Println(LL.head.next)
}
