package casetests

import (
	"fmt"
	"strconv"

	. "github.com/gabetucker2/gostack"
)

// test variables (test with these variables only after MakeCard's test)

var testCardA = MakeCard("Card A") // in sample stack
var testCardB = MakeCard("Card B") // in sample stack
var testCardC = MakeCard("Card C") // in sample stack
var testCardD = MakeCard("Card D") // out of sample stack

var map1 map[string]int
var arrKeys []string
var arrVals []int

/** I hate Go */
func test_RemoveCompileError(...interface{}) {}

/** Initialize test variables */
func test_Setup() {
	map1 := map[string]int {"Alexander":111, "Breton":222, "Charles":333}
	arrKeys := []string {"Alex", "Bre", "Charlie"}
	arrVals := []int {11, 22, 33}
	test_RemoveCompileError(map1, arrKeys, arrVals)
}

/** Test whether stack equals array */
func test_StackEqualArray(stack *Stack, _vals, _keys, _ma interface{}) bool {
	vals := _vals.([]interface{})
	keys := _keys.([]interface{})
	ma := _ma.(map[interface{}]interface{})
	maKeys := make([]interface{}, len(ma))
	for i := range stack.Cards {
		c := stack.Cards[i]
		if (_vals != nil && vals[i] != c.Val) || (_keys != nil && keys[i] != c.Key) || (_ma != nil && (maKeys[i] != c.Key || ma[maKeys[i]] != c.Val)) {
			return false
		}
	}
	return true
}

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

/** Make a sample stack of cards */
func test_SampleStack() *Stack {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack := MakeStack()

	// create stack (don't use .Add() function, or else you'll have to case test)
	stack.Cards = append(stack.Cards, testCardA)
	stack.Cards = append(stack.Cards, testCardB)
	stack.Cards = append(stack.Cards, testCardC)
	
	return stack

}
