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

// BACKEND FUNCTIONS

func _gostack_back_SampleStack() (stack *Stack) {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack = MakeStack()
	stack.Append(testCard1).Append(testCard2).Append(testCard3)

	return

}

func _gostack_back_LenAndSize(stack *Stack, size int) bool {

	// return whether len(cards) == cards.size
	return len(stack.cards) == size && stack.size == size

}

func _gostack_back_NewCard() Card {

	// return newly-created card
	return Card{}

}

func (stack *Stack) _gostack_back_AddCard(card interface{}) *Stack {

	stack.size++
	stack.

	// indices.insert
	// etc

	return stack

}

func (stack *Stack) _gostack_back_RemoveCard(card interface{}) {

	// size++
	// indices.insert
	// etc

}
