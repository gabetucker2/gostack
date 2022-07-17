package casetests

import (
	"strconv"
	"fmt"

	. "github.com/gabetucker2/gostack"
)

// test variables (test with these variables only after MakeCard's test)

var testCardA = MakeCard("Card A") // in sample stack
var testCardB = MakeCard("Card B") // in sample stack
var testCardC = MakeCard("Card C") // in sample stack
var testCardD = MakeCard("Card D") // out of sample stack

// test functions

/** return whether len(cards) == cards.size */
func test_LenAndSize(stack *Stack, size int) bool {

	return len(stack.Cards) == size && stack.Size == size

}

/** Return whether the indices correspond to their position in a stack */
func test_IdxsAreGood(stack *Stack) bool {

	good := true
	for i := range stack.Cards {
		if stack.Cards[i].Idx == i {
			good = false
			break
		}
	}
	return good

}

func test_Start(funcName string, showTestText bool) {

	// print TESTING line only if showTestText var set to true
	if showTestText {
		fmt.Println("-   TESTING " + funcName + "()")
	}

}

func test_End(funcName string, conditions []bool) {

	// set test to -1 (true) by default
	test := -1

	// test each condition and update test flag to index of condition which failed
	for i, c := range conditions {
		if !c {
			test = i
			break
		}
	}

	// set SUCCESS/FAILURE based on which condition, if any, failed
	out := "-   "
	if test == -1 {
		out += "SUCCESS"
	} else {
		out += "FAILURE AT CONDITION IDX = " + strconv.Itoa(test) + " in"
	}

	// print all the data together
	fmt.Println(out + " " + funcName + "()")

}

func test_SampleStack() *Stack {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack := MakeStack()

	// create stack (don't use .Add() function, or else you'll have to case test)
	stack.Cards = append(stack.Cards, testCardA)
	stack.Cards = append(stack.Cards, testCardB)
	stack.Cards = append(stack.Cards, testCardC)
	
	return stack

}
