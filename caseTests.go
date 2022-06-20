package main

import (
	"fmt"
	"strconv"
)

func gostack_PrintStart(thisFuncName string) {

	fmt.Println("-   TESTING " + thisFuncName + "()")

}

func gostack_PrintOut(test int, thisFuncName string) {

	var out = "-   "
	if test == -1 {
		out += "SUCCESS"
	} else {
		out += "FAILURE AT CONDITION IDX = " + strconv.Itoa(test) + " in"
	}

	fmt.Println(out + " " + thisFuncName + "()")

}

func gostack_SetTest(test *int, conditions []bool) {

	// set test to -1 (true) nu default
	*test = -1

	// test each condition and update test flag to index of condition which failed
	for i, c := range conditions {
		if !c {
			*test = i
			break
		}
	}

}

func main() {

	fmt.Println("- BEGINNING TESTS")
	test := -1

	//////////////////////////////////////////////////////
	thisFuncName := "MakeStack"
	gostack_PrintStart(thisFuncName)

	stack := MakeStack()
	conditions := []bool{len(stack.cards) == 0, stack.len == 0}

	gostack_SetTest(&test, conditions)
	gostack_PrintOut(test, thisFuncName)

	//////////////////////////////////////////////////////
	thisFuncName = "stack.Push"
	gostack_PrintStart(thisFuncName)

	testStr1 := "Card 1"
	testStr2 := "Card 2"
	stack.Push(testStr2)
	stack.Push(testStr1)
	conditions = []bool{len(stack.cards) == 2, stack.cards[0] == testStr1, stack.cards[1] == testStr2}

	gostack_SetTest(&test, conditions)
	gostack_PrintOut(test, thisFuncName)

	//////////////////////////////////////////////////////

}
