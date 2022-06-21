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

// BACKEND FUNCTIONS

func _gostack_back_SampleStack() (stack *Stack) {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack = MakeStack()

	// create stack
	stack.Append(testCard1).Append(testCard2).Append(testCard3)

	return

}

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

func _gostack_back_AddCardAfter(stack *Stack, card *Card, idx int) *Stack {

	// insert card into new array slice to satisfy append function
	newCards := []*Card{}

	if stack.size == 0 { // add card to empty list

		newCards = append(newCards, card)

	} else { // append each card in stack.cards to card

		for i, _ := range stack.cards {
			c := stack.cards[i]
			if i != idx {
				newCards = append(newCards, c)
			} else if i == idx {
				newCards = append(newCards, c)
				newCards = append(newCards, card)
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

func _gostack_back_RemoveCard(stack *Stack, idx int) *Card {

	card := new(Card) // have to do this if we want to set as nil

	if stack.size == 0 { // if we can't pop it, return nil

		card = nil

	} else { // if we can pop it, return popped card

		// insert card into new array slice to satisfy append function
		newCards := []*Card{}

		// append each card in stack.cards to card
		for i, _ := range stack.cards {
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

	return card

}
