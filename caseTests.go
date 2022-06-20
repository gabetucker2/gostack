package main

import (
	"fmt"
	"strconv"
)

var enabled = false
var testCard1 = "Card A" // in sample stack
var testCard2 = "Card B" // in sample stack
var testCard3 = "Card C" // in sample stack
var testCard4 = "Card D" // out of sample stack

func _gostack_test_Start(funcName string, enabled bool) {

	// print TESTING line only if enabled var set to true (set to false for cleaner console)
	if enabled {
		fmt.Println("-   TESTING " + funcName + "()")
	}

}

func _gostack_test_End(conditions []bool, funcName string) {

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
	stack.Push(testCard1)
	stack.Push(testCard2)
	stack.Push(testCard3)

	return

}

func _gostack_case_MakeStack(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := MakeStack()

	conditions := []bool{
		len(stack.cards) == 0,
		stack.len == 0,
	}

	_gostack_test_End(conditions, funcName)

}

func _gostack_case_Push(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := MakeStack()
	stack.Push(testCard3)
	stack.Push(testCard2)
	stack.Push(testCard1)

	conditions := []bool{
		len(stack.cards) == 3,
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		stack.cards[2] == testCard3,
	}

	_gostack_test_End(conditions, funcName)

}

func _gostack_case_1_IndexOf(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := _gostack_test_SampleStack()
	correctIdx := 1

	conditions := []bool{
		stack.IndexOf(testCard2) == correctIdx,
	}

	_gostack_test_End(conditions, funcName)

}

func _gostack_case_2_IndexOf(funcName string) {

	_gostack_test_Start(funcName, enabled)

	stack := _gostack_test_SampleStack()
	correctIdx := -1

	conditions := []bool{
		stack.IndexOf(testCard4) == correctIdx,
	}

	_gostack_test_End(conditions, funcName)

}

// layer two is dependent on layer one, layer three dependent on layer two, etc
func main() {

	fmt.Println("- BEGINNING TESTS")

	// layer one
	_gostack_case_MakeStack("MakeStack")

	// layer two
	_gostack_case_Push("stack.Push")

	// layer three
	_gostack_case_1_IndexOf("T1:stack.IndexOf")
	_gostack_case_1_IndexOf("T2:stack.IndexOf")

}
