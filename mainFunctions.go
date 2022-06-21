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
	stack.cards = []*Card{}

	// return
	return stack

}

func (stack *Stack) Add(card *Card, position Position, _idxData ...interface{}) *Stack {

	// get idx
	idx := _gostack_back_GetIdxFromData(stack, position, _idxData)

	// add only if valid idx found
	if idx != -1 {
		if position == Position_Last {
			idx++ // since we're doing AddBefore, increment final's idx to size
		}

		// push card to front
		_gostack_back_AddCard(stack, card, idx, true)
	}

	// return
	return stack

}

func (stack *Stack) Extract(position Position, _idxData ...interface{}) *Card {

	// get idx
	idx := _gostack_back_GetIdxFromData(stack, position, _idxData)

	// extract card if valid idx
	var extract *Card = nil
	if idx != -1 {
		extract = _gostack_back_ExtractCard(stack, idx)
	}

	// return
	return extract

}

func (stack *Stack) Replace(newCard *Card, position Position, _idxData ...interface{}) (oldCard *Card) {

	// get idx
	idx := _gostack_back_GetIdxFromData(stack, position, _idxData)

	if idx != -1 {
		// extract card
		oldCard = _gostack_back_ExtractCard(stack, idx)

		// insert card at previous location if got out card
		if oldCard != nil {
			_gostack_back_AddCard(stack, newCard, idx, true)
		}
	}

	// return
	return

}

func (stack *Stack) Has(position Position, _idxData ...interface{}) bool {

	// get idx
	idx := _gostack_back_GetIdxFromData(stack, position, _idxData)

	// return
	return idx != -1

}

func (stack *Stack) Index(position Position, _idxData ...interface{}) int {

	// return index
	return _gostack_back_GetIdxFromData(stack, position, _idxData)

}
