package main

import (
	"fmt"

	"protobufApp/simple/example_simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int{1, 4, 7, 9},
	}
	fmt.Println(sm)
}
