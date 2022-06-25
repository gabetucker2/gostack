package gostack_casetests

import (
	"strconv"

	"github.com/gabetucker2/gostack/gostack_casetests"
	"github.com/gabetucker2/gostack/gostack_tutorials"
	"github.com/gabetucker2/gostack/gostack_aorta"
	"github.com/gabetucker2/gostack/gostack"
)

// test variables

var testCardA = gostack_aorta.GOSTACK_back_MakeCard("Card A") // in sample stack
var testCardB = GOSTACK_back_MakeCard("Card B") // in sample stack
var testCardC = GOSTACK_back_MakeCard("Card C") // in sample stack
var testCardD = GOSTACK_back_MakeCard("Card D") // out of sample stack

// test functions

func gostack_test_LenAndSize(stack *Stack, size int) bool {

	// return whether len(cards) == cards.size
	return len(stack.cards) == size && stack.size == size

}

func gostack_test_Start(funcName string, showTestText bool) {

	// print TESTING line only if showTestText var set to true
	if showTestText {
		println("-   TESTING " + funcName + "()")
	}

}

func gostack_test_End(funcName string, conditions []bool) {

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
	println(out + " " + funcName + "()")

}

func gostack_test_SampleStack() (stack *Stack) {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack = MakeStack()

	// create stack (don't use stack.Add because we can't assume it is working in case tests)
	gostack_back_AddCard(stack, testCardA, gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)
	gostack_back_AddCard(stack, testCardB, gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)
	gostack_back_AddCard(stack, testCardC, gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)

	return

}
