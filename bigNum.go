package main

import (
	"fmt"
)

func main() {
	testMRSMod()

	var testRNS [3]RNSNum
	testRNS[0].AssignTo(5)
	testRNS[1].AssignTo(67901)
	testRNS[2] = *testRNS[0].Plus(&testRNS[1])
	fmt.Println("testRNS[0] : ", &testRNS[0], " ", testRNS[0], "testRNS[1] : ", &testRNS[1], " ", testRNS[1], "testRNS[2] : ", &testRNS[2], " ", testRNS[2], "\n")
	x, _ := testRNS[0].ToMRSNum()
	y, _ := testRNS[1].ToMRSNum()
	r, numGuesses := testRNS[2].ToMRSNum()
	fmt.Println(x.BigInt().String(), " + ", y.BigInt().String(), " = ", r.BigInt().String(), " ", numGuesses)
	fmt.Println(r)
	for i := 0; i < len(primes); i++ {
		fmt.Print(primes[i], " ")
	}
	fmt.Println()

//	testRNS[0] = NewRNSNum(36)
//	testRNS[1] = NewRNSNum(45)
//	testRNS[2] = testRNS[0].add(testRNS[1])

//	var testDeque Deque
}

func testMRSMod() {
	var n MRSNum
	n.Add(1, 7)
	n.Add(10, 3)
	if n.Mod(5) != 2 {
		panic("testMRSMod fail")
	}
}
