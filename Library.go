package gostack

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/gabetucker2/gogenerics"
)

/** Creates a card with inputted val, key, and idx

 @param optional `val` type{any} default nil
 @param optional `key` type{any} default nil
 @param optional `idx` type{int} default -1 no pass-by-reference
 @returns type{*Card} the newly-constructed card
 @constructs type{*Card} a newly-constructed card
 @ensures the new card will have val `val`, key `key`, and idx `idx`
*/
func MakeCard(variadic ...any) *Card {

	// unpack variadic into optional parameters
	var val, key, idx any
	gogenerics.UnpackVariadic(variadic, &val, &key, &idx)

	// initialize and set new Card
	card := new(Card)
	if idx == nil { card.Idx = -1 } else { card.Idx = idx.(int) }
	setCardProps(card, val, true) // TODO: add warning for *string, *int, *etc passed instead of *any
	setCardProps(card, key, false)

	// return
	return card

}

/** Creates a stack of cards with optional starting cards
 
 @param optional `input1` type{[]any, map[any]any} default nil
 @param optional `input2` type{[]any} default nil
 @param optional `repeats` type{int} default 1
 @param optional `overrideCards` type{bool} default false
   By default, if you do MakeStack([]*Card {cardA}), stack.Cards = []*Card {cardA}.  If you would like your cards to have vals pointing to other cards, where stack.Cards = []*Card { card {Idx = 0, Key = nil, Val = cardA} }, set this variable to true.
 @returns type{*Stack} the newly-constructed stack of newly-constructed cards
 @constructs type{*Stack} a newly-constructed stack of newly-constructed type{*Card} cards
 @requires
  * `input1` is map and nil `input2`
      OR `input1` is an array and nil `input2`
	  OR `input1` is an array and `input2` is an array
	  OR `input1` is nil and `input2` is an array
  * IF `input1` and `input2` are both passed as arguments
      |`input1`| == |`input2`|
 @ensures
  * repeats the function filling `repeats` (or, if nil or under 0, 1) amount of times
  * IF `input1` is passed
      IF `input1` is a map
        unpack the map into new cards with corresponding keys and vals
      ELSEIF `input1` is an array and `input2` is not passed/nil
	    IF `input1` is an array of cards:
		  set `stack.Cards` to `input1`
		ELSE:
          unpack values from `input1` into new cards
      ELSEIF `input1` is an array and `input2` is an array
        unpack keys from `input1` and values from `input2` into new cards
      ELSEIF `input1` is nil and `input2` is an array
        unpack keys from `input2` into new cards
	ELSEIF `input1` is nil and `input2` is nil and `repeats` is passed
		make `repeats` cards with nil value and nil key
    ELSE
      the stack is empty
 */
func MakeStack(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var input1, input2, repeats, overrideCards any
	gogenerics.UnpackVariadic(variadic, &input1, &input2, &repeats, &overrideCards)
	// set default
	if repeats == nil {
		if input1 == nil && input2 == nil {
			repeats = 0
		} else {
			repeats = 1
		}
	}

	// new stack
	stack := new(Stack)

	// run MakeStackMatrix to 1D array and add to our stack `repeats` times
	for i := 0; i < repeats.(int); i++ {
		matrixShape := -1
		if input1 == nil {
			if input2 == nil {
				matrixShape = 1
			} else {
				matrixShape = len(gogenerics.UnpackArray(input2))
			}
		} else {
			switch reflect.ValueOf(input1).Kind() {
			case reflect.Map:
				matrixShape = len(gogenerics.UnpackMap(input1))
			default: // reflect.Array OR reflect.Slice
				matrixShape = len(gogenerics.UnpackArray(input1))
			}
		}
		
		stack.Cards = append(stack.Cards, MakeStackMatrix(input1, input2, []int{matrixShape}, overrideCards).Cards...)
	}

	// property sets
	stack.setStackProperties()

	// return
	return stack

}

/** Returns a new stack-within-stack-structured stack
 
 @param optional `input1` type{any} default nil
 @param optional `input2` type{any} default nil
 @param optional `matrixShape` type{[]int} default nil
  * an int array representing the shape of the matrix
  * the first int is the largest container
  * the last int is the container directly containing the inputted cards
 @param optional `overrideCards` type{bool} default false
   By default, if you do MakeStackMatrix([]*Card {cardA}), stack.Cards = []*Card {cardA}.  If you would like your cards to have vals pointing to other cards, where stack.Cards = []*Card { card {Idx = 0, Key = nil, Val = cardA} }, set this variable to true.
 @returns type{*Stack} a new stack
 @constructs type{*Stack} a new stack with type{*Card} new cards
 @requires
  * If no `matrixShape` is passed, keys dimension must match the vals dimension
  * IF `input1` and `input2` are both passed as arguments
      |`input1`| == |`input2`|
 @ensures
  * IF no `matrixShape` is passed
      treating `input1`/`input2` as matrices/a map of matrices:
      IF `input1` is passed
        IF `input1` is a map
          unpack the map into matrix of shape `inputx` with corresponding keys and vals
        ELSEIF `input1` is an array and `input2` is not passed/nil
          unpack values from `input1` into matrix of shape `inputx`
        ELSEIF `input1` is an array and `input2` is an array
          unpack keys from `input1` and values from `input2` into matrix of shape `inputx`
        ELSEIF `input1` is nil and `input2` is an array
          unpack keys from `input2` into matrix of shape `inputx` 
      ELSEIF `input1` is not passed
        the stack is empty
	ELSEIF `matrixShape` is passed
	  treating `input1`/`input2` as 1D structures:
	  IF `input1` is passed
        IF `input1` is a map
          unpack the map into matrix of shape `matrixShape` with corresponding keys and vals
        ELSEIF `input1` is an array and `input2` is not passed/nil
	      IF `input1` is an array of cards:
		    set `stack.Cards` to cards in `input1` in matrix of shape `matrixShape`
		  ELSE:
            unpack values from `input1` into new cards in matrix of shape `matrixShape`
        ELSEIF `input1` is an array and `input2` is an array
          unpack keys from `input1` and values from `input2` into matrix of shape `matrixShape`
        ELSEIF `input1` is nil and `input2` is an array
          unpack keys from `input2` into matrix of shape `matrixShape`
	  ELSEIF `input1` is not passed AND `input2` is not passed
	    create a StackMatrix of shape `matrixShape` whose deepest card keys/vals are nil
 */
func MakeStackMatrix(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var input1, input2, matrixShape, overrideCards any
	gogenerics.UnpackVariadic(variadic, &input1, &input2, &matrixShape, &overrideCards)

	stack := new(Stack)

	// IF `input1` is passed
	if !(input1 == nil && input2 == nil) {

		if input1 != nil {

			input1Type := reflect.ValueOf(input1).Kind()
			switch input1Type {
			
			// IF `input1` is a map
			case reflect.Map:
				
				// get keys and vals from the input1 map
				var keys, vals []any
				
				// IF no `matrixShape` is passed
				if matrixShape == nil {
	
					// TODO: FIX LATER
					// unpackDeepMapToKeysVals(input1, keys, vals)

					// unpack the map into matrix of shape `inputx` with corresponding keys and vals
					stack.makeStackMatrixFromND(keys, vals)

				// ELSEIF `matrixShape` is passed
				} else {

					for k, v := range gogenerics.UnpackMap(input1) {
						keys = append(keys, k)
						vals = append(vals, v)
					}
					// unpack the map into matrix of shape `matrixShape` with corresponding keys and vals
					stack.makeStackMatrixFrom1D(matrixShape.([]int), keys, vals, new(int), false)
				}
			
			// ELSEIF `input1` is an array (or slice)...
			default:

				input1Array := gogenerics.UnpackArray(input1)

				// ...and `input2` is not passed
				if input2 == nil {

					// IF no `matrixShape` is passed
					if matrixShape == nil {
						// unpack values from `input1` into matrix of shape `inputx`
						stack.makeStackMatrixFromND(nil, input1)
					
					// ELSEIF `matrixShape` is passed
					} else {

						// IF `input1` is an array of cards:
						if true {
							// set `stack.Cards` to cards in `input1` in matrix of shape `matrixShape`
							stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, input1Array, new(int), overrideCards)
						// ELSE:
						} else {
							// unpack values from `input1` into new cards in matrix of shape `matrixShape`
							stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, input1Array, new(int), overrideCards)
						}
					}

				// ...and `input2` is an array
				} else {

					input2Array := gogenerics.UnpackArray(input2)
					
					// IF no `matrixShape` is passed
					if matrixShape == nil {
						// unpack keys from `input1` and values from `input2` into matrix of shape `inputx`
						stack.makeStackMatrixFromND(input1, input2)
						
					// ELSEIF `matrixShape` is passed
					} else {
						// unpack keys from `input1` and values from `input2` into matrix of shape `matrixShape`
						stack.makeStackMatrixFrom1D(matrixShape.([]int), input1Array, input2Array, new(int), false)
					}

				}

			}

		// ELSEIF `input1` is nil and `input2` is an array
		} else {

			input2Array := gogenerics.UnpackArray(input2)
			
			// IF no `matrixShape` is passed
			if matrixShape == nil {
				// unpack keys from `input2` into matrix of shape `inputx`
				stack.makeStackMatrixFromND(input2, nil)

			// ELSEIF `matrixShape` is passed
			} else {
				// unpack keys from `input2` into matrix of shape `matrixShape`
				stack.makeStackMatrixFrom1D(matrixShape.([]int), input2Array, nil, new(int), false)
			}

		}

	// ELSEIF `input1` is nil AND `input2` is nil
	} else {

		// IF no `matrixShape` is passed
		if matrixShape == nil {
			// the stack is empty

		// ELSEIF `matrixShape` is passed
		} else {
			// create a StackMatrix of shape `matrixShape` whose deepest card vals are nil
			stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, nil, new(int), false)

		}

	}

	// set properties
	stack.setStackProperties()

	// return
	return stack
	
}

/** Returns a stack representing a selection within a stack matrix
 
 @receiver `stack` type{*Stack}
 @param variadic `selections` type{int, []int} a set of args representing the indices being selected within an array
 @returns type{*Stack} a new Stack representing the selection
 @constructs type{*Stack} a new Stack representing the selection
 @requires `idx` arguments get valid index positions from the stack
 */
func (stack *Stack) StripStackMatrix(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var firstSelection any
	gogenerics.UnpackVariadic(variadic, &firstSelection)

	// init
	newStack := MakeStack()
	var selections []int

	// put firstSelection type{int, []int} into array selections type{[]int}
	switch fs := firstSelection.(type) {
	case int:
		selections = append(selections, fs)
	case []int:
		selections = append(selections, fs...)
	}

	// fmt.Println(len(selections))
	// iterate through each selection and add them to our new stack
	for _, idx := range selections {
		c := stack.Cards[idx]
		switch c.Val.(type) {
		case Stack:
			stripped := newStack.StripStackMatrix(variadic[1:])
			for _, idx := range firstSelection.([]int) {
				newStack.Cards = append(newStack.Cards, stripped.Cards[idx])
			}
		default:
			newStack.Cards = append(newStack.Cards, c)
		}
	}

	// set properties
	newStack.setStackProperties()

	// return
	return newStack

}

/** Creates a new any array from values of `stack`

 @receiver `stack` type{*Stack}
 @parameter optional `returnType` type{RETURN} default RETURN_Vals
 @returns type{[]any} new array
 @requires `stack.ToMatrix()` has been implemented
 @ensures new array values correspond to `stack` values
 */
func (stack *Stack) ToArray(variadic ...any) (arr []any) {

	// unpack variadic into optional parameters
	var returnType any
	gogenerics.UnpackVariadic(variadic, &returnType)

	// return
	return stack.ToMatrix(returnType, 1).([]any)

}

/** Creates a new any-any map from values of `stack`

 @receiver `stack` type{*Stack}
 @returns type{map[any]any} new map
 @ensures new map keys and values correspond to `stack` keys and values
 */
func (stack *Stack) ToMap() (m map[any]any) {

	// add all card keys and values in stack to m
	m = make(map[any]any)
	for i := range stack.Cards {
		c := stack.Cards[i]
		m[c.Key] = c.Val
	}

	// return
	return

}

/** Creates a new matrix from a stack by depth.  For instance, if depth = 2, then returns the stacks inside stack as an [][]any

 @receiver `stack` type{*Stack}
 @parameter optional `returnType` type{RETURN} default RETURN_Vals
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{[]any, [][]any, ..., []...[]any}
 @ensures
  * -1 depth means it will go as deep as it can
  * new map keys and values correspond to `stack` keys and values
  * example: Stack{Stack{"Hi"}, Stack{"Hello", "Hola"}, "Hey"} =>
      []any{[]any{"Hi"}, []any{"Hola", "Hello"}, "Hey"}
 */
func (stack *Stack) ToMatrix(variadic ...any) any {

	// unpack variadic into optional parameters
	var returnType, depth any
	gogenerics.UnpackVariadic(variadic, &returnType, &depth)
	setRETURNDefaultIfNil(&returnType)
	
	var matrix []any

	// update depth
	if depth == nil {
		depth = -1
	}

	// break recursion at depth == 0
	if depth.(int) != 0 {
		// add to return
		for i := range stack.Cards {
			c := stack.Cards[i]
			// if this Card's val is a Stack
			subStack, isStack := c.Val.(*Stack)
			if isStack {
				matrix = append(matrix, subStack.ToMatrix(depth.(int) - 1))
			} else {
				switch returnType {
				case RETURN_Vals:
					matrix = append(matrix, c.Val)
				case RETURN_Keys:
					matrix = append(matrix, c.Key)
				case RETURN_Idxs:
					matrix = append(matrix, c.Idx)
				case RETURN_Cards:
					matrix = append(matrix, c)
				}
			}
		}
	}

	// return
	return matrix

}

/** Makes a card with inputted vals and keys

 @receiver `stack` type{*Stack}
 @returns `stack`
 @updates `stack` to be empty
*/
func (stack *Stack) Empty() *Stack {

	// clear stack.Cards
	stack.Size = 0
	stack.Depth = 1
	stack.Cards = []*Card{} // avoid replacing stack object

	// return
	return stack

}

/** Returns a clone of the given card
 NOTE: Memory address of Key/Val in new card will be different even if CLONE_False, but updating one will update the other

 @receiver `card` type{*Card}
 @param optional `cloneKey` type{CLONE} default CLONE_False
 @param optional `cloneVal` type{CLONE} default CLONE_False
 @returns type{*Card} card clone
 @constructs type{*Card} clone of `card`
*/
func (card *Card) Clone(variadic ...any) *Card {

	// unpack variadic into optional parameters
	var cloneKey, cloneVal any
	gogenerics.UnpackVariadic(variadic, &cloneKey, &cloneVal)
	// set default vals
	setCLONEDefaultIfNil(&cloneKey)
	setCLONEDefaultIfNil(&cloneVal)

	// init
	clone := new(Card)
	clone.Idx = card.Idx
	if cloneKey == CLONE_True {
		clone.Key = gogenerics.CloneInterface(&card.Key)
	} else {
		clone.Key = reflect.ValueOf(&card.Key).Elem().Interface()
		card.Key = reflect.ValueOf(&clone.Key).Elem().Interface()
	}
	if cloneVal == CLONE_True {
		clone.Val = gogenerics.CloneInterface(&card.Val)
	} else {
		clone.Val = reflect.ValueOf(&card.Val).Elem().Interface()
		card.Val = reflect.ValueOf(&clone.Val).Elem().Interface()
	}

	// return
	return clone

}

/** Returns a clone of the given stack

 @receiver `stack` type{*Stack}
 @optional param `cloneKeys` type{CLONE} default CLONE_True
 @optional param `cloneVals` type{CLONE} default CLONE_True
 @returns type{*Stack} stack clone
 @constructs type{*Stack} clone of `stack`
 @ensures
  * the stack clone has the same card pointers as `stack`
  * `cloneCards` => each Card in the stack clone is cloned
  * `cloneKeys` => each Card in the stack's Key is cloned
  * `cloneVals` => each Card in the stack's Val is cloned
*/
func (stack *Stack) Clone(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var cloneKeys, cloneVals any
	gogenerics.UnpackVariadic(variadic, &cloneKeys, &cloneVals)

	// init
	clone := new(Stack)
	clone.Size = stack.Size
	clone.Depth = stack.Depth
	for i := range stack.Cards {
		clone.Cards = append(clone.Cards, stack.Cards[i].Clone(cloneKeys, cloneVals))
	}

	// return
	return clone

}

/** Removes all cards from `stack` which share the same field value as another card in that stack and returns the new stack
 Assuming elements represent the values of cards in the pre-existing stack,
 Stack{"Hi", "Hey", "Hello", "Hi", "Hey", "Howdy"}.Unique(TYPE_Val) => Stack{"Hi", "Hey", "Hello", "Howdy"}

 @receiver `stack` type{*Stack}
 @param `typeType` type{TYPE}
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates `stack` to have no repeating values between field `typeType`
 @requires `stack.Clone()` has been implemented
 @ensures
  * IF no deepSearch
      removes repeat cards from stack
    ELSE
	  removes cards matching other cards within the scope of each substack
	    For instance, Stack{Stack{1, 2, 1}, Stack{1, 2}}.Unique => Stack{Stack{1, 2}, Stack{1, 2}}
 */
func (stack *Stack) Unique(typeType TYPE, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var pointerType, deepSearchType, depth, uniqueType any
	gogenerics.UnpackVariadic(variadic, &pointerType, &deepSearchType, &depth, &uniqueType)

	// allow deepSearchHandler to handle Unique
	*stack = *stack.deepSearchHandler("Unique", false, FIND_All, nil, nil, pointerType, deepSearchType, depth, typeType, uniqueType, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	// set properties
	stack.setStackProperties()

	return stack

}

/** Returns whether two cards equal one another
 
 @receiver `thisCard` type{*Card}
 @param `otherCard` type{*Card}
 @param optional `pointerTypeKey` type{POINTER} default POINTER_False
 @param optional `pointerTypeVal` type{POINTER} default POINTER_False
 @param optional `compareIdxs` type{COMPARE} default COMPARE_False
 @param optional `compareKeys` type{COMPARE} default COMPARE_True
 @param optional `compareVals` type{COMPARE} default COMPARE_True
 @param optional `printType` type{PRINT} default PRINT_False
 @returns type{bool}
 */
func (thisCard *Card) Equals(otherCard *Card, variadic ...any) bool {

	// unpack variadic into optional parameters
	var pointerTypeKey, pointerTypeVal, compareIdxs, compareKeys, compareVals, printType any
	gogenerics.UnpackVariadic(variadic, &pointerTypeKey, &pointerTypeVal, &compareIdxs, &compareKeys, &compareVals, &printType)
	// set default vals
	setPOINTERDefaultIfNil(&pointerTypeKey)
	setPOINTERDefaultIfNil(&pointerTypeVal)
	setCOMPAREDefaultIfNil(&compareIdxs)
	if compareKeys == nil {compareKeys = COMPARE_True}
	if compareVals == nil {compareVals = COMPARE_True}
	setPRINTDefaultIfNil(&printType)

	print := func(printType any, stringToPrint string) {
		if printType.(PRINT) == PRINT_True {
			fmt.Printf("-     DETAIL: CONDITION: %v\n", stringToPrint)
		}
	}

	condition := true
	
	condition = condition && 
		(compareKeys == COMPARE_False ||
		(compareKeys == COMPARE_True &&
			(
				(pointerTypeKey == POINTER_False && thisCard.Key == otherCard.Key) ||
				(pointerTypeKey == POINTER_True && gogenerics.PointersEqual(thisCard.Key, otherCard.Key) ) ) ) )
	print(printType, fmt.Sprintf("KEY PASSES EQUALITY CHECK: %v: (compareKeys == COMPARE_False [%v] || (compareKeys == COMPARE_True [%v] && ( (pointerTypeKey == POINTER_False [%v] && thisCard.Key == otherCard.Key [%v]) || (pointerTypeKey == POINTER_True [%v] && gogenerics.PointersEqual(thisCard.Key, otherCard.Key) [%v] ) ) ) )", (compareKeys == COMPARE_False || (compareKeys == COMPARE_True && ( (pointerTypeKey == POINTER_False && thisCard.Key == otherCard.Key) || (pointerTypeKey == POINTER_True && gogenerics.PointersEqual(thisCard.Key, otherCard.Key) ) ) ) ), compareKeys == COMPARE_False, compareKeys == COMPARE_True, pointerTypeKey == POINTER_False, thisCard.Key == otherCard.Key, pointerTypeKey == POINTER_True, gogenerics.PointersEqual(thisCard.Key, otherCard.Key) ))
	
	condition = condition && 
		(compareVals == COMPARE_False ||
		(compareVals == COMPARE_True &&
			(
				(pointerTypeVal == POINTER_False && thisCard.Val == otherCard.Val) ||
				(pointerTypeVal == POINTER_True && gogenerics.PointersEqual(thisCard.Val, otherCard.Val) ) ) ) )
	print(printType, fmt.Sprintf("Val PASSES EQUALITY CHECK: %v: (compareVals == COMPARE_False [%v] || (compareVals == COMPARE_True [%v] && ( (pointerTypeVal == POINTER_False [%v] && thisCard.Val == otherCard.Val [%v]) || (pointerTypeVal == POINTER_True [%v] && gogenerics.PointersEqual(thisCard.Val, otherCard.Val) [%v] ) ) ) )", (compareVals == COMPARE_False || (compareVals == COMPARE_True && ( (pointerTypeVal == POINTER_False && thisCard.Val == otherCard.Val) || (pointerTypeVal == POINTER_True && gogenerics.PointersEqual(thisCard.Val, otherCard.Val) ) ) ) ), compareVals == COMPARE_False, compareVals == COMPARE_True, pointerTypeVal == POINTER_False, thisCard.Val == otherCard.Val, pointerTypeVal == POINTER_True, gogenerics.PointersEqual(thisCard.Val, otherCard.Val) ))

	condition = condition && (compareIdxs == COMPARE_False || (compareIdxs == COMPARE_True && thisCard.Idx == otherCard.Idx))
	print(printType, fmt.Sprintf("IDX PASSES EQUALITY CHECK: %v: (compareIdxs == COMPARE_False [%v] || (compareIdxs == COMPARE_True [%v] && thisCard.Idx == otherCard.Idx [%v]))", (compareIdxs == COMPARE_False || (compareIdxs == COMPARE_True && thisCard.Idx == otherCard.Idx)), compareIdxs == COMPARE_False, compareIdxs == COMPARE_True, thisCard.Idx == otherCard.Idx))

	// return whether conditions yield true
	return condition

}

/** Returns whether two stacks equal one another
 
 @receiver `thisStack` type{*Stack}
 @param `otherStack` type{*Stack}
 @param optional `compareKeys` type{COMPARE} default COMPARE_True
 @param optional `compareVals` type{COMPARE} default COMPARE_True
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int} default -1 (deepest)
 @param optional `pointerTypeKey` type{POINTER} default POINTER_False
 @param optional `pointerTypeVal` type{POINTER} default POINTER_False
 @param optional `pointerTypeStack` type{POINTER} default POINTER_False
	Checks for if the vals of hidden cards (cards whose vals are substacks) are pointers
		Set true if and only if updating substack 1 in `stack` should update the same substack 1 in `otherStack`
 @ensures
	If `stack`.Depth != `otherStack`.Depth and the N-deep comparison finds that they're equal, then return that they're Equal 
 @returns type{bool}
 */
func (stack *Stack) Equals(otherStack *Stack, variadic ...any) (test bool) {

	/*

	PSEUDOCODE OUTLINE:

	stack.Equals(otherStack) bool

		set up variadic stuff

		if depth == -1 || depth > stack.Depth
			depth = stack.Depth
		if deepSearchType is false
			depth = 1

		test = (depth != 0 and stack and otherStack have the same Size) or (depth == 0)

		for each cardA in this stack
			for each cardB in other stack
				if cardA corresponds to cardB and test is true and depth != 0
						
					if compareKeys is true
						test = test && cardA.Equals(cardB, [pass in pointer and compare stuff for key])
						
						
					if cards' vals both hold substacks
						test = test && substackA.Equals(substackB, ..., depth: depth - 1)
						if pointerTypeStack is true
							test = test && substackA is same pointer as substackB and both are pointers

					else if one holds a substack and the other doesn't
						test = false

					else if compareKeys is true
						test = test && cardA.Equals(cardB, [pass in pointer and compare stuff for val])


		return test

	*/
	
	// unpack variadic into optional parameters
	var compareKeys, compareVals, deepSearchType, depth, pointerTypeKey, pointerTypeVal, pointerTypeStack any
	gogenerics.UnpackVariadic(variadic, &compareKeys, &compareVals, &deepSearchType, &depth, &pointerTypeKey, &pointerTypeVal, &pointerTypeStack)
	// set default vals
	if compareKeys == nil {compareKeys = COMPARE_True}
	if compareVals == nil {compareVals = COMPARE_True}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_True}
	setDepthDefaultIfNil(&depth)
	setPOINTERDefaultIfNil(&pointerTypeStack)
	setPOINTERDefaultIfNil(&pointerTypeKey)
	setPOINTERDefaultIfNil(&pointerTypeVal)
	
	if depth == -1 || depth.(int) > stack.Depth { depth = stack.Depth }
	if deepSearchType == DEEPSEARCH_False { depth = 1 }
	
	test = (depth != 0 && stack.Size == otherStack.Size) || depth == 0
	
	for _, cardA := range stack.Cards {
		for _, cardB := range otherStack.Cards {
			if cardA.Idx == cardB.Idx && test && depth != 0 {
				
				if compareKeys == COMPARE_True {
					test = test && cardA.Equals(cardB, pointerTypeKey, POINTER_False, COMPARE_False, compareKeys, COMPARE_False)
				}

				oneHoldsSubstack := false
				bothHoldSubstacks := false
				switch cardA.Val.(type) {
				case *Stack:
					oneHoldsSubstack = true
				}
				switch cardB.Val.(type) {
				case *Stack:
					if oneHoldsSubstack {
						bothHoldSubstacks = true
						oneHoldsSubstack = false
						test = test && cardA.Val.(*Stack).Equals(cardB.Val.(*Stack), compareKeys, compareVals, deepSearchType, depth.(int) - 1, pointerTypeKey, pointerTypeVal, pointerTypeStack)
						if pointerTypeStack == POINTER_True {
							test = test && gogenerics.PointersEqual(cardA.Val, cardB.Val)
						}
					} else {
						oneHoldsSubstack = true
					}
				}
				if oneHoldsSubstack {
					test = false
				}
				if !bothHoldSubstacks && !oneHoldsSubstack && compareVals == COMPARE_True {
					test = test && cardA.Equals(cardB, POINTER_False, pointerTypeVal, COMPARE_False, COMPARE_False, compareVals)
				}

			}
		}
	}

	return test

}

/** Shuffles the order of `stack` cards

 @receiver `stack` type{*Stack}
 @param optional `newOrder` type{bool} default true
 @returns `stack`
 @updates
  * `stack` card ordering is randomized
  * rand.Seed is updated to time.Now().UnixNano()
 @ensures if stack.Size > 1 and newOrder == true, then new order will be different than previous
 */
func (stack *Stack) Shuffle(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var newOrder any
	gogenerics.UnpackVariadic(variadic, &newOrder)
	if newOrder == nil {newOrder = false}

	// body
	initClone := stack.Clone()

	for ok := true; ok; ok = (newOrder.(bool) && stack.Size > 1 && initClone.Equals(stack, COMPARE_False, nil, DEEPSEARCH_True)) { // emulate a do-while loop
		
		// pseudo-randomize seed
		rand.Seed(time.Now().UnixNano())

		////////////////////////////////////// NEED A DEEP IMPLEMENTATION
		// shuffle
		rand.Shuffle(stack.Size, func(i, j int) { stack.Cards[i], stack.Cards[j] = stack.Cards[j], stack.Cards[i] })
		
		// set indices
		setIndices(stack.Cards)

	}

	// return
	return stack

}

/** Flips the ordering of `stack.Cards`
 
 @receiver `stack` type{*Stack}
 @returns `stack`
 @updates `stack` to have its ordering reversed
 */
func (stack *Stack) Flip() *Stack {

	stack.Lambda(func(card *Card, stack *Stack, _ ...any) {
		stack.Move(FIND_Card, ORDER_Before, FIND_Idx, card, 0)
	})

	// return
	return stack

}

/** Prints information regarding `card` to the console
 
 @receiver `card` type{*Card}
 @param optional `depth` type{int} default 0
 	This variable only exists for text-indenting purposes to make your terminal output look a bit cleaner.  1 depth => 4 "-" added before the print.
 @updates terminal logs
 */
func (card *Card) Print(variadic ...any) {

	// unpack variadic into optional parameters
	var depth any
	gogenerics.UnpackVariadic(variadic, &depth)

	if depth == nil {
		depth = 0
	}
	depth = depth.(int)

	// prints
	fmt.Printf("%v|%vCARD\n", depthPrinter(depth.(int)), gogenerics.IfElse(depth == 0, "gostack: PRINTING ", ""))
	fmt.Printf("%v- &card:    %v\n", depthPrinter(depth.(int)), &card)
	fmt.Printf("%v- card.Idx: %v\n", depthPrinter(depth.(int)), card.Idx)
	fmt.Printf("%v- card.Key: %v\n", depthPrinter(depth.(int)), card.Key)
	fmt.Printf("%v- card.Val: %v\n", depthPrinter(depth.(int)), card.Val)

}

/** Prints information regarding `stack` to the console
 
 @receiver `stack` type{*Stack}
 @param optional `depth` type{int} default 0
   This variable only exists for text-indenting purposes to make your terminal output look a bit cleaner.  1 depth => 4 "-" added before the print.
 @updates terminal logs
 */
func (stack *Stack) Print(depth ...int) {
	
	if depth == nil {
		depth = []int {0}
	}
	fmt.Printf("%v|%vSTACK\n", depthPrinter(depth[0]), gogenerics.IfElse(len(depth) != 2, "gostack: PRINTING ", "SUB"))
	fmt.Printf("%v- &stack:      %v\n", depthPrinter(depth[0]), &stack)
	if len(depth) == 2 {
		fmt.Printf("%v- card.Idx:    %v\n", depthPrinter(depth[0]), depth[1])
	}
	fmt.Printf("%v- stack.Size:  %v\n", depthPrinter(depth[0]), stack.Size)
	fmt.Printf("%v- stack.Depth: %v\n", depthPrinter(depth[0]), stack.Depth)
	for i := range stack.Cards {
		c := stack.Cards[i]

		switch c.Val.(type) {
		case *Stack:
			c.Val.(*Stack).Print(depth[0]+4, i)
		default:
			c.Print(depth[0]+4)
		}
	}
	
}

/** Iterate through a stack calling your lambda function on each card
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(*Card, *Stack, (returnVal) any, ...any)}
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns (returnVal) type{any}
 @ensures
  * Each card in `stack` is passed into your lambda function
  * `stack` is the first argument passed into your variadic parameter on the first call
 */
func (stack *Stack) Lambda(lambda any, variadic ...any) (ret any) {

	// unpack variadic into optional parameters
	var deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &deepSearchType, &depth)
	
	// main
	generalIterator(stack, lambda.(func(*Card, *Stack, any, ...any)), deepSearchType.(DEEPSEARCH), depth.(int), ret, nil) // TODO: replace nil final value

	return ret

}

/** Adds to a stack of cards or a cards at (each) position(s) and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `insert` type{Card, Stack}
 @param optional `orderType` type{ORDER} default ORDER_Before
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @param optional `overrideStackConversion` type{bool} default false
	if `insert` is of type{Stack}:
		if not `overrideStackConversion`:
			add to `stack` from `insert.Cards`
		else if `overrideStackConversion`:
			add the `insert` stack to `stack` as the val of a card
 @returns `stack` if cards were added OR nil if no cards were added (due to invalid find)
 @updates `stack` to have new cards before/after each designated position
 @requires `stack.Clone()` has been implemented
 */
func (stack *Stack) Add(insert any, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var orderType, findType, findData, pointerType, deepSearchType, depth, overrideStackConversion any
	gogenerics.UnpackVariadic(variadic, &orderType, &findType, &findData, &pointerType, &deepSearchType, &depth, &overrideStackConversion)

	// allow deepSearchHandler to handle function
	*stack = *stack.deepSearchHandler("Add", true, findType, findData, nil, pointerType, deepSearchType, depth, nil, nil, insert, orderType, nil, nil, nil, nil, nil, nil, overrideStackConversion)

	// allow deepSearchHandler to take care of function
	return stack

}

/** Moves one element or slice of cards to before or after another element or slice of cards
 
 @receiver `stack` type{*Stack}
 @param `findType_from` type{FIND}
 @param `orderType` type{ORDER}
 @param `findType_to` type{FIND}
 @param optional `findData_from` type{any} default nil
 @param optional `findData_to` type{any} default nil
 @param optional `pointerType_from` type{POINTER} default POINTER_False
 @param optional `pointerType_to` type{POINTER} default POINTER_False
 @param optional `deepSearchType_from` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `deepSearchType_to` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth_from` type{int} default -1 (deepest)
 @param optional `depth_to` type{int} default -1 (deepest)
 @returns `stack` if moved OR nil if no move occurred (due to bad find)
 @requires you are not moving a stack to a location within that own stack
 @ensures a stack of cards, or individual cards, can be targeted
 */
func (stack *Stack) Move(findType_from FIND, orderType ORDER, findType_to FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData_from, findData_to, pointerType_from, pointerType_to, deepSearchType_from, deepSearchType_to, depth_from, depth_to any
	gogenerics.UnpackVariadic(variadic, &findData_from, &findData_to, &pointerType_from, &pointerType_to, &deepSearchType_from, &deepSearchType_to, &depth_from, &depth_to)

	// 1) Get the cards to move
	from := stack.ExtractMany(findType_from, findData_from, pointerType_from, RETURN_Cards, deepSearchType_from, depth_from)
	// 2) Get the card to put froms before/after depending on whether before or after
	var toCard *Card
	toStack := stack.GetMany(findType_to, findData_to, RETURN_Cards, pointerType_to, CLONE_False, CLONE_False, CLONE_False, deepSearchType_to, depth_to)
	if orderType == ORDER_After {
		toCard = toStack.Get(FIND_Last)
	} else if orderType == ORDER_Before {
		toCard = toStack.Get(FIND_First)
	}
	// 3) Insert 2 to 1 (works since to.Idx is procedurally updated by ExtractMany())
	stack.Add(from, orderType, FIND_Idx, toCard.Idx, pointerType_to, deepSearchType_to, depth_to)

	// return
	return stack

}

/** Swaps one element or slice with the position of another element or slice
 
 @receiver `stack` type{*Stack}
 @param `findType_first` type{FIND}
 @param `findType_second` type{FIND}
 @param optional `findData_first` type{any} default nil
 @param optional `findData_second` type{any} default nil
 @param optional `pointerType_first` type{POINTER} default POINTER_False
 @param optional `pointerType_second` type{POINTER} default POINTER_False
 @param optional `deepSearchType_first` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `deepSearchType_second` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth_first` type{int} default -1 (deepest)
 @param optional `depth_second` type{int} default -1 (deepest)
 @updates `stack`
 @returns `stack` if moved OR nil if no move occurred (due second bad find)
 @requires you are not swapping a stack with a location within that own stack
 @ensures a stack of cards, or individual cards, can be targeted
 */
func (stack *Stack) Swap(findType_first FIND, findType_second FIND, variadic ...any) *Stack {

	// unpack variadic insecond optional parameters
	var findData_first, findData_second, pointerType_first, pointerType_second, deepSearchType_first, deepSearchType_second, depth_first, depth_second any
	gogenerics.UnpackVariadic(variadic, &findData_first, &findData_second, &pointerType_first, &pointerType_second, &deepSearchType_first, &deepSearchType_second, &depth_first, &depth_second)

	// 1) Get the first and second cards being targeted
	firsts := stack.GetMany(findType_second, findData_second, RETURN_Cards, pointerType_second, CLONE_False, CLONE_False, CLONE_False, deepSearchType_second, depth_second)
	seconds := stack.GetMany(findType_first, findData_first, RETURN_Cards, pointerType_first, CLONE_False, CLONE_False, CLONE_False, deepSearchType_first, depth_first)
	// 2) Determine which card is before the other in the stack.  If the second is before the first, switch the first variable and the second variable and all corresponding variables.
	if seconds.Cards[0].Idx < firsts.Cards[0].Idx {
		_tempFirst := firsts//lint:ignore SA4006 Ignore warning
		firsts = seconds
		seconds = _tempFirst//lint:ignore SA4006 Ignore warning

		_tempFindData_first := findData_first
		findData_first = findData_second
		findData_second = _tempFindData_first

		_temppointerType_first := pointerType_first
		pointerType_first = pointerType_second
		pointerType_second = _temppointerType_first

		_tempDeepSearchType_first := deepSearchType_first
		deepSearchType_first = deepSearchType_second
		deepSearchType_second = _tempDeepSearchType_first

		_tempDepth_first := depth_first
		depth_first = depth_second
		depth_second = _tempDepth_first
	}
	// 3) Now, in order second preserve the integrity of indices should the client choose second use FIND_Idx(s)...
	//    * Insert a copy of firsts after seconds,
	stack.Add(firsts, ORDER_After, findData_second, pointerType_second, deepSearchType_second, depth_second)
	//    * move second after first,
	stack.Move(findType_second, ORDER_After, findType_first, findData_first, findData_second, pointerType_first, pointerType_second, deepSearchType_first, deepSearchType_second, depth_first, depth_second)
	//    * and remove the original copy of first
	stack.Remove(findType_first, findData_first, pointerType_first, deepSearchType_first, depth_first)

	// return
	return stack

}

/** Returns a boolean representing whether a search exists in the stack

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns true IF successful search, false IF unsuccessful search
 @requires `stack.Get()` has been implemented
 */
func (stack *Stack) Has(variadic ...any) bool {

	// unpack variadic into optional parameters
	var findType, findData, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findType, &findData, &pointerType, &deepSearchType, &depth)

	// return
	return stack.Get(findType, findData, pointerType, nil, nil, nil, deepSearchType, depth) != nil
}

/** Gets a card from specified parameters in a stack, or nil if does not exist

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `clonesType_card` type{CLONES} default CLONE_False
 @param optional `clonesType_keys` type{CLONES} default CLONE_False
 @param optional `clonesType_vals` type{CLONES} default CLONE_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Card} the found card OR nil (if invalid find)
 @ensures
  * CLONE_True for `clonesType_card` means the returned card object itself is a clone
  * CLONE_True for `clonesType_key` means the returned card key is a clone
  * CLONE_True for `clonesType_val` means the returned card val is a clone
 */
func (stack *Stack) Get(variadic ...any) (ret *Card) {

	// unpack variadic into optional parameters
	var findType, findData, pointerType, clonesType_card, clonesType_key, clonesType_val, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findType, &findData, &pointerType, &clonesType_card, &clonesType_key, &clonesType_val, &deepSearchType, &depth)

	// allow deepSearchHandler to take care of function
	return stack.deepSearchHandler("Get", true, findType, findData, nil, pointerType, deepSearchType, depth, nil, nil, nil, nil, nil, nil, nil, clonesType_card, clonesType_key, clonesType_val, nil).Cards[0]

}

/** Gets a stack from specified parameters in a stack
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `clonesType` type{CLONES} default CLONE_False
 @param optional `clonesType_keys` type{CLONES} default CLONE_False
 @param optional `clonesType_vals` type{CLONES} default CLONE_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} the new stack (if find fails, then an empty stack)
 @constructs type{*Stack} new stack of specified values from specified cards in `stack`
 @requires
  * `clonesType_keys` and `clonesType_vals` are only passed if `returnType` == RETURN_Cards
 @ensures
  * CLONE_True means the cards in the returned stack are clones
  * CLONE_True for `clonesType_keys` means the cards in the returned stack keys are clones
  * CLONE_True for `clonesType_vals` means the cards in the returned stack vals are clones
 */
func (stack *Stack) GetMany(findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, returnType, pointerType, clonesType, clonesType_keys, clonesType_vals, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &returnType, &pointerType, &clonesType, &clonesType_keys, &clonesType_vals, &deepSearchType, &depth)

	// allow deepSearchHandler to take care of function
	return stack.deepSearchHandler("Get", false, findType, findData, returnType, pointerType, deepSearchType, depth, nil, nil, nil, nil, nil, nil, nil, clonesType, clonesType_keys, clonesType_vals, nil)

}

/** Returns a found card before its respective field is updated to `replaceData` (OR nil if not found)
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Card} a clone of extracted card OR nil if found no cards
 @updates first found card to `replaceData`
 @requires `stack.Get()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 */
func (stack *Stack) Replace(replaceType REPLACE, replaceData any, findType FIND, variadic ...any) (ret *Card) {

	// unpack variadic into optional parameters
	var findData, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &pointerType, &deepSearchType, &depth)
	
	// get deep copy of targeted card OR nil
	ret = stack.Get(findType, findData, pointerType, CLONE_True, CLONE_True, CLONE_True, deepSearchType, depth)
	// get target data
	_, targetCards, targetStacks := stack.getPositions(true, findType, findData, pointerType.(POINTER), deepSearchType.(DEEPSEARCH), depth.(int))
	
	//stack.Get(findType, findData, pointerType, CLONE_False, CLONE_False, CLONE_False, deepSearchType, depth)

	// set targeted card field to replaceData if was found (updateRespectiveField fulfills our ensures clause)
	if len(targetCards) != 0 {
		targetStacks[0].updateRespectiveField(replaceType, replaceData, targetCards[0])
	}

	// update properties
	stack.setStackProperties()

	// return
	return

}

/** Returns a stack whose values are the original fields updated to `replaceData`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} a stack whose values are the extracted cards pre-update (if find fails, then an empty stack)
 @updates all found cards to `replaceData`
 @requires `stack.GetMany()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the cards found will be removed from `stack`
 */
func (stack *Stack) ReplaceMany(replaceType REPLACE, replaceData any, findType FIND, variadic ...any) (ret *Stack) {

	// unpack variadic into optional parameters
	var findData, returnType, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &returnType, &pointerType, &deepSearchType, &depth)

	// get deep copy of targeted cards to return
	ret = stack.GetMany(findType, findData, returnType, pointerType, CLONE_True, CLONE_True, CLONE_True)
	// get target data
	_, targetCards, targetStacks := stack.getPositions(false, findType, findData, pointerType.(POINTER), deepSearchType.(DEEPSEARCH), depth.(int))

	// set targeted cards' fields to replaceData if was found (updateRespectiveField fulfills our ensures clause)
	if len(targetCards) != 0 {
		for i := range targetCards {
			targetStacks[i].updateRespectiveField(replaceType, replaceData, targetCards[i])
		}
	}

	// update properties
	stack.setStackProperties()

	// return
	return

}

/** Updates a card in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates the found card in `stack`
 @requires `stack.Replace()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 */
func (stack *Stack) Update(replaceType REPLACE, replaceData any, findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &pointerType, &deepSearchType, &depth)

	// update stack
	stack.Replace(replaceType, replaceData, findType, findData, pointerType, deepSearchType, depth)

	// return the original stack
	return stack

}

/** Updates cards in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates  the found cards in `stack`
 @requires `stack.ReplaceMany()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the cards found will be removed from `stack`
 */
func (stack *Stack) UpdateMany(replaceType REPLACE, replaceData any, findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &pointerType, &deepSearchType, &depth)

	// update stack
	stack.ReplaceMany(replaceType, replaceData, findType, findData, pointerType, nil, deepSearchType, depth)

	// return the original stack
	return stack

}

/** Gets and removes a card from `stack`, or returns nil if does not exist
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Card} the extracted card OR nil (if invalid find)
 @updates `stack` to no longer have found card
 @requires `stack.Replace()` has been implemented
 */
func (stack *Stack) Extract(findType FIND, variadic ...any) *Card {

	// unpack variadic into optional parameters
	var findData, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &pointerType, &deepSearchType, &depth)

	// return the original value
	return stack.Replace(REPLACE_Card, nil, findType, findData, pointerType, deepSearchType, depth)

}

/** Gets and removes a set of data from `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} the extracted card (if find fails, then an empty stack)
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 */
func (stack *Stack) ExtractMany(findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, returnType, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &returnType, &pointerType, &deepSearchType, &depth)

	// return the original value
	return stack.ReplaceMany(REPLACE_Card, nil, findType, findData, returnType, pointerType, deepSearchType, depth)

}

/** Removes a card from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates `stack` to no longer have found card
 @requires `stack.Replace()` has been implemented
 */
func (stack *Stack) Remove(findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &pointerType, &deepSearchType, &depth)

	// remove the card
	stack.Replace(REPLACE_Card, nil, findType, findData, pointerType, deepSearchType, depth)

	// return stack
	return stack

}

/** Removes a set of cards from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 */
func (stack *Stack) RemoveMany(findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &pointerType, &deepSearchType, &depth)

	// remove the cards
	stack.ReplaceMany(REPLACE_Card, nil, findType, findData, pointerType, nil, deepSearchType, depth)

	// return stack
	return stack

}