package aorta

func MakeCard(idx int, val ...*interface{}, key ...*interface{}) (card *Card) {

	// initialize and set new Card
	card = new(Card)
	card.idx = idx
	card.key = key
	card.val = val

	// return
	return

}

func MakeCards(structureType STRUCTURE, input1 ...*interface{}, input2 ...*interface{}) (stack *Stack) {

	stack = MakeStack()

	if structureType == STRUCTURE_Map {

		switch input1.(type) {
		case []interface{}:

			for i := range len(input1) {
				k := input1[i] // avoid for-loop cloning
				v := input2[i] // avoid for-loop cloning
				MakeCard(i, v, k)
			}

		case map[interface{}]interface{}:

			for i := range len(input1) {
				k, v := input1[i] // avoid for-loop cloning
				MakeCard(i, v, k)
			}

		}

	} else if STRUCTURETYPE == STRUCTURE_Arr {

		for i := range len(input1) {
			v := input1[i] // avoid for-loop cloning
			MakeCard(i, v)
		}
		
	}

	return

}

func MakeStack(structureType ...STRUCTURE, input1 ...*interface{}, input2 ...*interface{}) (stack *Stack) {

	if structureType != nil {
		// if structureType passed in, get stack of cards from MakeCards
		if input2 != nil {
			stack = MakeCards(structureType, input1, input2)
		} else {
			stack = MakeCards(structureType, input1)
		}
	} else {
		// if no structureType passed in, just execute normally
		stack = new(Stack)
		stack.size = 0
	}

	// return
	return stack

}

func (stack *Stack) Empty() *Stack {

	stack.size = 0
	stack.cards = []*Card{} // avoid replacing stack object

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

func (stack *Stack) Get(returnType RETURN, positionType POSITION, POSITIONDATA, matchType ...MATCH) interface{} {



}

func (stack  *Stack) Has(lookFor interface{}, position POSITION, data ...interface{}) bool 

	// get id
	idx := _gostack_back_GetIdxFromData(stack, position, data)

	// retur
	return idx != -1

}
