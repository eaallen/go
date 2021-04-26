package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	str string
}

func getVarType() {
	node := Node{str: "hello"}

	fmt.Println(reflect.ValueOf(node).Kind())
}

func main() {
	// get varuable type
	getVarType()
}
