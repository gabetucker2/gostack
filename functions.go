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

func (stack *Stack) IndexOf(card interface{}) (idx int) {

	// sets the default index to -1, the return value for a failed search
	idx = -1

	// searches through each card and, if match, sets index to that target's index
	for i, c := range stack.cards {
		if c == card {
			idx = i
			break
		}
	}

	// return
	return

}
