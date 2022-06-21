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

func (stack *Stack) Add(card *Card, position Position, _idxData ...interface{}) *Stack {

	// get idxData
	idxData := _gostack_back_GetIdxData(_idxData)
	idx := _gostack_back_GetIdxFromPosition(stack, position, idxData).(int)

	if position == Position_Last {
		idx++ // since we're doing AddBefore, increment final's idx to size
	}

	// push card to front
	_gostack_back_AddCard(stack, card, idx, true)

	// return
	return stack

}

func (stack *Stack) Extract(position Position, _idxData ...interface{}) *Card {

	// get idxData
	var idxData = _gostack_back_GetIdxData(_idxData)

	// return
	return _gostack_back_ExtractCard(stack, _gostack_back_GetIdxFromPosition(stack, position, idxData))

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

func (stack *Stack) IndexCard(card *Card) int {

	return _gostack_back_IndexCard(stack, card)

}
