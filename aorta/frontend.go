package aorta

// (val ...any, key ...any, idx ...int)
func MakeCard(variadic ...interface{}) (card *Card) {

	// unpack variadic into optional parameters
	var val, key, idx *interface{}
	GOSTACK_back_UnpackVariadic(variadic, val, key, idx)

	// return
	return GOSTACK_back_MakeCard(val, key, idx)

}

/**
 Makes a stack of cards with inputted vals and keys
 @param optional `input1` type{any, []any, map[any]any}
 @param optional `input2` type{any, []any}
 @returns type{*Stack} the newly-constructed stack of cards
 @constructs type{*Stack} a newly-constructed stack of cards
 @requires
  * `input1` is map and no `input2` passed
      OR `input1` is an array and no `input2` passed
	  OR `input1` is an array and `input2` is an array
  * IF `input1` and `input2` are both passed as arguments
      |`input1`| == |`input2`|
  * MakeCard() has been implemented
 @ensures
  * IF `input1` is passed
	    IF `input1` is a map
          unpack the map into new cards with corresponding keys and vals
        ELSEIF `input1` is an array and `input2` is not passed
          unpack values from `input1` into new cards
        ELSEIF `input1` is an array and `input2` is an array
		  unpack keys from `input1` and values from `input2` into new cards
	ELSE
	  the stack is empty
 */
func MakeCards(variadic ...interface{}) (stack *Stack) {

	// unpack variadic into optional parameters
	var input1, input2 interface{}
	GOSTACK_back_UnpackVariadic(variadic, &input1, &input2)

	stack = MakeStack()

	switch structureType {

	case STRUCTURE_Arr:

		vals := input1.([]interface{})
		for i := range vals {
			MakeCard(&vals[i], nil, i)
		}

	case STRUCTURE_Map:

		switch reflect.TypeOf(input1).Key() {

		case []interface{}:

			keys := input1.([]interface{})
			vals := input2.([]interface{})
			for i := range keys {
				MakeCard(&vals[i], &keys[i], i)
			}

		case map[interface{}]interface{}:

			m := input1.(map[interface{}]interface{})
			i := 0
			for k, v := range m {
				MakeCard(v, k, i)
				i++
			}

		}

	}

	return

}

// dependent upon MakeCards
// (input1 ...interface{}, input2 ...interface{})
func MakeStack(variadic ...interface{}) (stack *Stack) {

	var input1, input2 *interface{}
	GOSTACK_back_UnpackVariadic(variadic, input1, input2)

	if input1 != nil {
		// if input is passed in, pass input values to MakeCards
		stack = MakeCards(input1, input2)
	} else {
		// if no input is passed in, just make an empty stack
		stack = new(Stack)
		stack.size = 0
	}

	// return
	return stack

}

// ()
func (stack *Stack) Empty() *Stack {

	stack.size = 0
	stack.cards = []*Card{} // avoid replacing stack object

	// return
	return stack

}

/*
func (stack *Stack) Add(toAdd *interface{}, beforeNotAfter bool, position POSITION, data ...interface{}) *Stack {

	// get idx
	idx := gostack_back_GetIdxFromData(stack, position, data)

	// add only if valid idx found
	if idx != -1 {
		if position == POSITION_Last {
			idx++ // since we're doing AddBefore, increment final's idx to size
		}

		// push card to front
		gostack_back_AddCard(stack, card, idx, true)
	}

	// return
	return stack

}

func (stack *Stack) Extract(position POSITION, data ...interface{}) *Card {

	// get idx
	idx := gostack_back_GetIdxFromData(stack, position, data)

	// extract card if valid idx
	var extract *Card = nil
	if idx != -1 {
		extract = gostack_back_ExtractCard(stack, idx)
	}

	// return
	return extract

}

func (stack *Stack) Replace(toInsert *[]interface{}, position POSITION, data ...interface{}) (oldCards *Stack) {

	// get idx
	idx := gostack_back_GetIdxFromData(stack, position, data)

	if idx != -1 {
		// extract card
		oldCard = gostack_back_ExtractCard(stack, idx)

		// insert card at previous location if got out card
		if oldCard != nil {
			gostack_back_AddCard(stack, toInsert, idx, true)
		}
	}

	// return
	return

}

func (stack *Stack) Get(returnType RETURN, positionType POSITION, POSITIONDATA, matchType ...MATCH) interface{} {



}

func (stack  *Stack) Has(lookFor interface{}, position POSITION, data ...interface{}) bool

	// get id
	idx := gostack_back_GetIdxFromData(stack, position, data)

	// retur
	return idx != -1

}
*/
