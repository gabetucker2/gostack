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
func (stack *Stack) Clear() *Stack {

	// clear stack
	stack.cards = MakeStack().cards
	stack.size = 0

	// return
	return stack

}

func (stack *Stack) Push(card Card) *Stack {

	// insert card into new interface slice to satisfy append function
	newCards := []Card{card}

	// append each card in stack.cards to card
	for _, c := range stack.cards {
		newCards = append(newCards, c)
	}

	// set stack.cards to our new interface
	stack.cards = newCards

	// update stack properties
	stack.size++

	// return
	return stack

}

func (stack *Stack) Append(card Card) *Stack {

	// create new interface slice
	newCards := []Card{}

	// append each card in stack.cards to card
	for _, c := range stack.cards {
		newCards = append(newCards, c)
	}

	// insert card into interface last
	newCards = append(newCards, card)

	// set stack.cards to our new interface
	stack.cards = newCards

	// update stack properties
	stack.size++

	// return
	return stack

}

func (stack *Stack) Pop() (card Card) {

	if stack.size == 0 { // if we can't pop it, return nil

		card = _gostack_back_NewCard()

	} else { // if we can pop it, return popped card

		// get card we're removing
		card = stack.cards[stack.size-1]

		// remove the last card from the stack
		stack.cards = stack.cards[:stack.size-1]

		// update stack properties
		stack.size--

	}

	// return
	return

}

func (stack *Stack) Behead() (card interface{}) {

	if stack.size == 0 { // if we can't behead it, return nil

		card = nil

	} else { // if we can pop it, return popped card

		// get card we're removing
		card = stack.cards[0]

		// remove the first card from the stack
		stack.cards = stack.cards[1:]

		// update stack properties
		stack.size--

	}

	// return
	return

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
