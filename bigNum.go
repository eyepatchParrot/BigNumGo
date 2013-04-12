package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello world\n")
	var testRNS [3]RNSNum
	testRNS[0].Assign(36)
	testRNS[1].Assign(45)
	fmt.Println(testRNS[0], testRNS[1])
	testRNS[2] = *testRNS[0].Add(&testRNS[1])
	fmt.Println(testRNS[0], testRNS[1], testRNS[2])
//	testRNS[0] = NewRNSNum(36)
//	testRNS[1] = NewRNSNum(45)
//	testRNS[2] = testRNS[0].add(testRNS[1])

//	var testDeque Deque
}
