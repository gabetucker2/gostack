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
	_gostack_back_AddCard(stack, testCardA, _gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)
	_gostack_back_AddCard(stack, testCardB, _gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)
	_gostack_back_AddCard(stack, testCardC, _gostack_back_GetIdxFromPosition(stack, POSITION_Last), false)

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
	card.val = val

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

func _gostack_back_UpdatePosData(_posData ...interface{}) (posData interface{}) {
	if len(_posData) == 1 {
		posData = _posData[0] // just so there is only one optional param
	} else {
		posData = nil
	}
	return
}

func _gostack_back_GetIdxFromData(stack *Stack, position POSITION, _posData ...interface{}) (idx int) {
	return _gostack_back_GetIdxFromPosition(stack, position, _gostack_back_UpdatePosData(_posData)).(int)
}

func _gostack_back_MakeSlice(x, y int) Slice {
	return Slice{x, y}
}

// returns index of searched item if valid
// else, returns -1
func _gostack_back_GetIdxFromPosition(stack *Stack, position POSITION, _posData ...interface{}) (idx interface{}) {

	posData := _gostack_back_UpdatePosData(_posData...)

	switch position {

	case POSITION_First:
		idx = 0 // nil
	case POSITION_Last:
		idx = stack.size - 1 // nil
	case POSITION_Card:
		idx = -1
		for i, c := range stack.cards {
			if c == posData { // key
				idx = i
				break
			}
		}
	case POSITION_Idx:
		idx = posData // int
	case POSITION_Key:
		idx = -1
		for i, c := range stack.cards {
			if c.key == posData { // key
				idx = i
				break
			}
		}
	case POSITION_Val:
		idx = -1
		for i, c := range stack.cards {
			if c.val == posData { // card
				idx = i
				break
			}
		}
	case POSITION_Slice:
		idx = posData // Slice
	case POSITION_All:
		idx = _gostack_back_MakeSlice(0, stack.size-1) // nil

	}

	return

}
