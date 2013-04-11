package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello world\n")
	var testDeque Deque
	for i := uint32(0); i < 25; i++ {
		testDeque.PushBack(i)
	}

	fmt.Printf("testDeque = %v\n", &testDeque)
	//	testDeque.GetAt(2)
	//	testDeque.SetAt(2, 10)
}
