package main

import (
	"strconv"
)

// TEST FUNCTIONS (to facilitate cases)

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
	_gostack_back_AddCard(stack, testCardA, _gostack_back_GetIdxFromPosition(stack, Position_Last), false)
	_gostack_back_AddCard(stack, testCardB, _gostack_back_GetIdxFromPosition(stack, Position_Last), false)
	_gostack_back_AddCard(stack, testCardC, _gostack_back_GetIdxFromPosition(stack, Position_Last), false)

	return

}

// BACKEND FUNCTIONS

func _gostack_back_LenAndSize(stack *Stack, size int) bool {

	// return whether len(cards) == cards.size
	return len(stack.cards) == size && stack.size == size

}

func _gostack_back_NewCard(val interface{}) (card *Card) {

	// make newly-created card
	card = &Card{}
	card.key = nil
	card.value = val

	// return
	return

}

// TODO: implement for range
func _gostack_back_AddCard(stack *Stack, card *Card, idx interface{}, beforeNotAfter bool) *Stack {

	// insert card into new array slice to satisfy append function
	newCards := []*Card{}

	if stack.size == 0 { // add card to empty list

		newCards = append(newCards, card)

	} else { // append each card in stack.cards to card

		if beforeNotAfter {

			for i := range stack.cards {
				c := stack.cards[i]
				if i != idx {
					newCards = append(newCards, c)
				} else if i == idx {
					newCards = append(newCards, card)
					newCards = append(newCards, c)
				}
			}

			if idx == stack.size {
				newCards = append(newCards, card)
			}

		} else {

			for i := range stack.cards {
				c := stack.cards[i]
				if i != idx {
					newCards = append(newCards, c)
				} else if i == idx {
					newCards = append(newCards, c)
					newCards = append(newCards, card)
				}
			}

		}

	}

	// set stack.cards to our new array
	stack.cards = newCards

	// update stack properties
	stack.size++

	// return
	return stack

}

func _gostack_back_ExtractCard(stack *Stack, idx interface{}) (card *Card) {

	if stack.size == 0 { // if we can't pop it, return nil

		card = nil

	} else { // if we can pop it, return popped card

		// insert card into new array slice to satisfy append function
		newCards := []*Card{}

		// append each card in stack.cards to card
		for i := range stack.cards {
			c := stack.cards[i]
			if i != idx {
				newCards = append(newCards, c)
			} else if i == idx {
				card = c
			}
		}

		// set stack.cards to our new array
		stack.cards = newCards

		// update stack properties
		stack.size--

	}

	return

}

func _gostack_back_IndexKey(stack *Stack, key interface{}) (idx int) {

	// sets the default index to -1, the return value for a failed search
	idx = -1

	// searches through each card and, if match, sets idx to that target's index
	for i, c := range stack.cards {
		if c.key == key {
			idx = i
			break
		}
	}

	// return
	return

}

func _gostack_back_IndexCard(stack *Stack, card *Card) (idx int) {

	// sets the default index to -1, the return value for a failed search
	idx = -1

	// searches through each card and, if match, sets idx to that target's index
	for i, c := range stack.cards {
		if c == card {
			idx = i
			break
		}
	}

	// return
	return

}

func _gostack_back_GetIdxData(_idxData ...interface{}) (idxData interface{}) {
	if len(_idxData) == 1 {
		idxData = _idxData[0] // just so there is only one optional param
	} else {
		idxData = nil
	}
	return
}

func _gostack_back_GetIdxFromPosition(stack *Stack, position Position, _idxData ...interface{}) (idx interface{}) {

	var idxData = _gostack_back_GetIdxData(_idxData...)

	switch position {

	case Position_First:
		idx = 0 // nil
	case Position_Last:
		idx = stack.size - 1 // nil
	case Position_Idx:
		idx = idxData // int
	case Position_Key:
		idx = _gostack_back_IndexKey(stack, idxData) // key (interface)
	case Position_Val:
		idx = _gostack_back_IndexCard(stack, idxData.(*Card)) // card
	case Position_Slice:
		idx = idxData // {int, int}

	}

	return

}
