package casetests

import "strconv"

// test variables

var testCardA = _gostack_back_NewCard("Card A") // in sample stack
var testCardB = _gostack_back_NewCard("Card B") // in sample stack
var testCardC = _gostack_back_NewCard("Card C") // in sample stack
var testCardD = _gostack_back_NewCard("Card D") // out of sample stack

// test functions

func _gostack_test_LenAndSize(stack *Stack, size int) bool {

	// return whether len(cards) == cards.size
	return len(stack.cards) == size && stack.size == size

}

func _gostack_test_Start(funcName string, showTestText bool) {

	// print TESTING line only if showTestText var set to true
	if showTestText {
		println("-   TESTING " + funcName + "()")
	}

}

func _gostack_test_End(funcName string, conditions []bool) {

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

func _gostack_test_SampleStack() (stack *Stack) {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack = MakeStack()

	// create stack (don't use stack.Add because we can't assume it is working in case tests)
	_gostack_back_AddCard(stack, testCardA, _gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)
	_gostack_back_AddCard(stack, testCardB, _gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)
	_gostack_back_AddCard(stack, testCardC, _gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)

	return

}
