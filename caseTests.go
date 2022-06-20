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

	conditions := []bool{
		len(stack.cards) == 0,
		stack.len == 0,
	}

	gostack_SetTest(&test, conditions)
	gostack_PrintOut(test, thisFuncName)

	////////////////////////////////////////////////////// <>
	thisFuncName = "stack.Push"
	gostack_PrintStart(thisFuncName)

	testCard1 := "Card 1"
	testCard2 := "Card 2"
	testCard3 := "Card 3"
	stack.Push(testCard3)
	stack.Push(testCard2)
	stack.Push(testCard1)

	conditions = []bool{
		len(stack.cards) == 2,
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		stack.cards[2] == testCard3,
	}

	gostack_SetTest(&test, conditions)
	gostack_PrintOut(test, thisFuncName)

	////////////////////////////////////////////////////// <"Card 1", "Card 2", "Card 3">
	thisFuncName = "T1:stack.IndexOf"
	gostack_PrintStart(thisFuncName)

	correctIdx := 1

	conditions = []bool{
		stack.IndexOf(testCard2) == correctIdx,
	}

	gostack_SetTest(&test, conditions)
	gostack_PrintOut(test, thisFuncName)

	////////////////////////////////////////////////////// <"Card 1", "Card 2", "Card 3">
	thisFuncName = "T2:stack.IndexOf"
	gostack_PrintStart(thisFuncName)

	correctIdx = -1
	testCard4 := "Card 4"

	conditions = []bool{
		stack.IndexOf(testCard4) == correctIdx,
	}

	gostack_SetTest(&test, conditions)
	gostack_PrintOut(test, thisFuncName)

}
