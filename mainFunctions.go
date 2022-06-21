package main

func MakeStack() *Stack {

	// initialize new stack
	stack := new(Stack)

	// initialize stack properties
	stack.size = 0

	// return
	return stack

}

// not just call MakeStack since that would replace the object
func (stack *Stack) Empty() *Stack {

	stack.size = 0
	stack.cards = MakeStack().cards

	// return
	return stack

}

func (stack *Stack) Push(card *Card) *Stack {

	// push card to front
	_gostack_back_AddCardAfter(stack, card, 0)

	// return
	return stack

}

func (stack *Stack) Append(card *Card) *Stack {

	// append card to back
	_gostack_back_AddCardAfter(stack, card, stack.size-1)

	// return
	return stack

}

func (stack *Stack) Pop() *Card {

	// return
	return _gostack_back_RemoveCard(stack, stack.size-1)

}

func (stack *Stack) Behead() *Card {

	// return
	return _gostack_back_RemoveCard(stack, 0)

}

func (stack *Stack) Has(card interface{}) (has bool) {

	// sets the default has to false, the return value for a failed search
	has = false

	// searches through each card and, if match, sets has flag to true
	for _, c := range stack.cards {
		if c == card {
			has = true
			break
		}
	}

	// return
	return

}

func (stack *Stack) IndexOf(card interface{}) (idx int) {

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
