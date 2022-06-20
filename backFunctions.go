package main

import (
	"fmt"
	"strconv"
)

// TEST FUNCTIONS (to facilitate cases)

func _gostack_test_Start(funcName string, enabled bool) {

	// print TESTING line only if enabled var set to true
	if enabled {
		fmt.Println("-   TESTING " + funcName + "()")
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
	fmt.Println(out + " " + funcName + "()")

}

func _gostack_test_SampleStack() (stack *Stack) {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack = MakeStack()
	stack.Append(testCard1).Append(testCard2).Append(testCard3)

	return

}

// only use this function for cases which make/update a Stack, not that solely get a value
func _gostack_test_LenAndSize(stack *Stack, size int) bool {

	return len(stack.cards) == size && stack.size == size

}
