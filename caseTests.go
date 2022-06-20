package main

import (
	"fmt"
	"strconv"
)

// VARIABLES

var enabled = false      // false for cleaner console, true for debugging
var testCard1 = "Card A" // in sample stack
var testCard2 = "Card B" // in sample stack
var testCard3 = "Card C" // in sample stack
var testCard4 = "Card D" // out of sample stack

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

func _gostack_test_SampleStack() (stack Stack) {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack = MakeStack()
	stack.Push(testCard3)
	stack.Push(testCard2)
	stack.Push(testCard1)

	return

}

// only use this function for cases which make/update a Stack, not that solely get a value
func _gostack_test_LenAndSize(stack Stack, size int) bool {

	return len(stack.cards) == size && stack.size == size

}

// CASE FUNCTIONS

func _gostack_case_MakeStack(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := MakeStack()

	conditions := []bool{
		_gostack_test_LenAndSize(stack, 0),
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Push(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := MakeStack()
	stack.Push(testCard3)
	stack.Push(testCard2)
	stack.Push(testCard1)

	conditions := []bool{
		_gostack_test_LenAndSize(stack, 3),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		stack.cards[2] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_IndexOf(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := _gostack_test_SampleStack()

	conditions := []bool{
		stack.IndexOf(testCard1) == 0,
		stack.IndexOf(testCard2) == 1,
		stack.IndexOf(testCard3) == 2,
		stack.IndexOf(testCard4) == -1,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_1_Pop(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := _gostack_test_SampleStack()
	pop := stack.Pop()

	conditions := []bool{
		_gostack_test_LenAndSize(stack, 2),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		pop == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_2_Pop(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := MakeStack()
	pop := stack.Pop()

	conditions := []bool{
		_gostack_test_LenAndSize(stack, 0),
		pop == nil,
	}

	_gostack_test_End(funcName, conditions)

}

// MAIN FUNCTION

func main() {

	// layer two is dependent on layer one, layer three dependent on layer two, etc
	fmt.Println("- BEGINNING TESTS")

	// layer one
	_gostack_case_MakeStack("MakeStack") // regular case

	// layer two
	_gostack_case_Push("stack.Push")       // regular case
	_gostack_case_1_Pop("T1:stack.Pop")    // regular case
	_gostack_case_2_Pop("T2:stack.Pop")    // edge case
	_gostack_case_IndexOf("stack.IndexOf") // regular and edge cases

}
