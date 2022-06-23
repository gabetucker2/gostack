package aorta

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

func (stack *Stack) Add(toAdd *interface{}, beforeNotAfter bool, position POSITION, data ...interface{}) *Stack {

	// get idx
	idx := _gostack_back_GetIdxFromData(stack, position, data)

	// add only if valid idx found
	if idx != -1 {
		if position == POSITION_Last {
			idx++ // since we're doing AddBefore, increment final's idx to size
		}

		// push card to front
		_gostack_back_AddCard(stack, card, idx, true)
	}

	// return
	return stack

}

func (stack *Stack) Extract(position POSITION, data ...interface{}) *Card {

	// get idx
	idx := _gostack_back_GetIdxFromData(stack, position, data)

	// extract card if valid idx
	var extract *Card = nil
	if idx != -1 {
		extract = _gostack_back_ExtractCard(stack, idx)
	}

	// return
	return extract

}

func (stack *Stack) Replace(toInsert *[]interface{}, position POSITION, data ...interface{}) (oldCards *Stack) {

	// get idx
	idx := _gostack_back_GetIdxFromData(stack, position, data)

	if idx != -1 {
		// extract card
		oldCard = _gostack_back_ExtractCard(stack, idx)

		// insert card at previous location if got out card
		if oldCard != nil {
			_gostack_back_AddCard(stack, toInsert, idx, true)
		}
	}

	// return
	return

}

func (stack *Stack) Has(lookFor interface{}, position POSITION, data ...interface{}) bool {

	// get idx
	idx := _gostack_back_GetIdxFromData(stack, position, data)

	// return
	return idx != -1

}

//TODO : Implement get Function
