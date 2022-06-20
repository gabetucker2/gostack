package main

func MakeStack() Stack {

	// creates new stack
	var stack Stack

	// initialized new stack's variables
	stack.len = 0

	// return
	return stack

}

func (stack *Stack) Push(card interface{}) *Stack {

	// insert card into new interface slice to satisfy append function
	newInterface := []interface{}{card}

	// append each card in stack.cards to card
	for _, e := range stack.cards {
		newInterface = append(newInterface, e)
	}

	// set stack.cards to our new interface
	stack.cards = newInterface

	// return
	return stack

}

func (stack *Stack) IndexOf(target interface{}) (idx int) {

	// sets
	idx = -1
	for i, c := range stack.cards {
		if c == target {
			idx = i
			break
		}
	}

	return

}
