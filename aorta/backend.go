package aorta

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

func _gostack_back_UpdatePosData(_data ...interface{}) (data interface{}) {
	if len(_data) == 1 {
		data = _data[0] // just so there is only one optional param
	} else {
		data = nil
	}
	return
}

func _gostack_back_GetIdxFromData(stack *Stack, position POSITION, _data ...interface{}) (idx int) {
	return _gostack_back_GetIdxFromPosition(stack, position, _gostack_back_UpdatePosData(_data)).(int)
}

func _gostack_back_MakeSlice(x, y int) Slice {
	return Slice{x, y}
}

// returns index of searched item if valid
// else, returns -1
func _gostack_back_GetIdxFromPosition(stack *Stack, position POSITION, _data ...interface{}) (idx interface{}) {

	data := _gostack_back_UpdatePosData(_data...)

	switch position {

	case POSITION_First:
		idx = 0 // nil
	case POSITION_Last:
		idx = stack.size - 1 // nil
	case POSITION_Card:
		idx = -1
		for i, c := range stack.cards {
			if c == data { // key
				idx = i
				break
			}
		}
	case POSITION_Idx:
		idx = data // int
	case POSITION_Key:
		idx = -1
		for i, c := range stack.cards {
			if c.key == data { // key
				idx = i
				break
			}
		}
	case POSITION_Val:
		idx = -1
		for i, c := range stack.cards {
			if c.val == data { // card
				idx = i
				break
			}
		}
	case POSITION_Slice:
		idx = data // Slice
	case POSITION_All:
		idx = _gostack_back_MakeSlice(0, stack.size-1) // nil

	}

	return

}
