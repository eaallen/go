package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type Test struct {
	value string
	next  *Test
}

type Methods interface {
	test() string
}

type LinkedList struct {
	head Node
}

func (ll *LinkedList) add(node Node) {
	temp := ll.head
	ll.head = node
	ll.head.next = &temp
}

func NewTest() *Test {
	return &Test{value: "test1"}
}

func New() LinkedList {
	return LinkedList{head: Node{value: "head"}}
}

var LL = New()
var T = NewTest()

func main() {
	fmt.Println(LL)

	// add a new node
	n := Node{value: "new node"}
	LL.add(n)
	fmt.Println(LL.head.next)
}
