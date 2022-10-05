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
 @param optional `idx` type{int} default -1
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
	card.Val = val
	card.Key = key

	// return
	return card

}

/** Creates a stack of cards with optional starting cards
 
 @param optional `input1` type{[]any, map[any]any} default nil
 @param optional `input2` type{[]any} default nil
 @param optional `repeats` type{int} default 1
 @param optional `overrideCards` type{OVERRIDE} default OVERRIDE_False
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
  * assuming all mentions of array are interchangeable with *Stack:
    IF `input1` is passed
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
				if reflect.ValueOf(input2).Kind() == reflect.Ptr {
					matrixShape = input2.(*Stack).Size
				} else {
					matrixShape = len(gogenerics.UnpackArray(input2))
				}
			}
		} else {
			switch reflect.ValueOf(input1).Kind() {
			case reflect.Map:
				matrixShape = len(gogenerics.UnpackMap(input1))
			case reflect.Ptr:
				matrixShape = input1.(*Stack).Size
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
 @param optional `overrideCards` type{OVERRIDE} default OVERRIDE_False
   By default, if you do MakeStackMatrix([]*Card {cardA}), stack.Cards = []*Card {cardA}.  If you would like your cards to have vals pointing to other cards, where stack.Cards = []*Card { card {Idx = 0, Key = nil, Val = cardA} }, set this variable to true.
   This only has an effect when `matrixShape` is passed.
 @returns type{*Stack} a new stack
 @constructs type{*Stack} a new stack with type{*Card} new cards
 @requires
  * If no `matrixShape` is passed, keys dimension must match the vals dimension
  * IF `input1` and `input2` are both passed as arguments
      |`input1`| == |`input2`|
 @ensures
  * assuming all mentions of array are interchangeable with *Stack:
    IF no `matrixShape` is passed
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
	setOVERRIDEDefaultIfNil(&overrideCards)

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
	
					// unpack the map into matrix of shape `inputx` with corresponding keys and vals
					stack.makeStackMatrixFromNDMap(input1)

				// ELSEIF `matrixShape` is passed
				} else {

					keys, vals = gogenerics.GetKeysValsFromMap(input1)
					// unpack the map into matrix of shape `matrixShape` with corresponding keys and vals
					stack.makeStackMatrixFrom1D(matrixShape.([]int), keys, vals, new(int), overrideCards.(OVERRIDE))
				}
			
			// ELSEIF `input1` is an array (or slice)...
			default:

				var input1Array []any
				if reflect.ValueOf(input1).Kind() == reflect.Ptr {
					for _, c := range input1.(*Stack).Cards {
						input1Array = append(input1Array, c.Val)
					}
				} else {
					input1Array = gogenerics.UnpackArray(input1)
				}

				// ...and `input2` is not passed
				if input2 == nil {

					// IF no `matrixShape` is passed
					if matrixShape == nil {
						// unpack values from `input1` into matrix of shape `inputx`
						stack.makeStackMatrixFromND(nil, input1Array)
					
					// ELSEIF `matrixShape` is passed
					} else {

						// set `stack.Cards` to cards in `input1` in matrix of shape `matrixShape`
						stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, input1Array, new(int), overrideCards.(OVERRIDE))

					}

				// ...and `input2` is an array
				} else {

					var input2Array []any
					if reflect.ValueOf(input2).Kind() == reflect.Ptr {
						for _, c := range input2.(*Stack).Cards {
							input2Array = append(input2Array, c.Val)
						}
					} else {
						input2Array = gogenerics.UnpackArray(input2)
					}
					
					// IF no `matrixShape` is passed
					if matrixShape == nil {
						// unpack keys from `input1` and values from `input2` into matrix of shape `inputx`
						stack.makeStackMatrixFromND(input1Array, input2Array)
						
					// ELSEIF `matrixShape` is passed
					} else {
						// unpack keys from `input1` and values from `input2` into matrix of shape `matrixShape`
						stack.makeStackMatrixFrom1D(matrixShape.([]int), input1Array, input2Array, new(int), overrideCards.(OVERRIDE))
					}

				}

			}

		// ELSEIF `input1` is nil and `input2` is an array
		} else {

			var input2Array []any
			if reflect.ValueOf(input2).Kind() == reflect.Ptr {
				for _, c := range input2.(*Stack).Cards {
					input2Array = append(input2Array, c.Val)
				}
			} else {
				input2Array = gogenerics.UnpackArray(input2)
			}
			
			// IF no `matrixShape` is passed
			if matrixShape == nil {
				// unpack keys from `input2` into matrix of shape `inputx`
				stack.makeStackMatrixFromND(input2Array, nil)

			// ELSEIF `matrixShape` is passed
			} else {
				// unpack keys from `input2` into matrix of shape `matrixShape`
				stack.makeStackMatrixFrom1D(matrixShape.([]int), input2Array, nil, new(int), overrideCards.(OVERRIDE))
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
			stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, nil, new(int), OVERRIDE_False)

		}

	}

	// set properties
	stack.setStackProperties()

	// return
	return stack
	
}

/** Returns a stack representing a selection within a stack matrix
 
 @receiver `stack` type{*Stack}
 @param optional `selections` type{int variadic, []int} default []int {0, ..., stack.Size - 1}
	a set of args representing the indices being selected within an array
 @returns type{*Stack} a new Stack representing the selection
 @constructs type{*Stack} a new Stack representing the selection
 @requires `idx` arguments get valid index positions from the stack
 */
func (stack *Stack) StripStackMatrix(variadic ...any) *Stack {

	// TODO: update to just use the Get() function
	// TODO: add support for inserting a variadic of ints instead of []int

	// unpack variadic into optional parameters
	var firstSelection any
	gogenerics.UnpackVariadic(variadic, &firstSelection)

	/*
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
	*/
	return stack

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
	return stack.ToMatrix(returnType, 1)

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

/** Creates a new matrix from a stack

 @receiver `stack` type{*Stack}
 @param optional `returnType` type{RETURN} default RETURN_Vals
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{ []any {., ..., .}, []any {[]any {., ..., .}, []any {., ..., .}, ...}, ..., []any { []any {...{ []any{., ..., .}, ... }, ... }, ... } }
 @ensures
  * example: Stack{Stack{"Hi"}, Stack{"Hello", "Hola"}, "Hey"} | deepsearchtrue =>
      []any{[]any{"Hi"}, []any{"Hola", "Hello"}, "Hey"}
  * example: Stack{Stack{"Hi"}, Stack{"Hello", "Hola"}, "Hey"} | deepsearchfalse =>
      []any{[]any{}, []any{}, "Hey"}
 */
func (stack *Stack) ToMatrix(variadic ...any) (matrix []any) {

	// unpack variadic into optional parameters
	var returnType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &returnType, &deepSearchType, &depth)
	// set defaults
	setRETURNDefaultIfNil(&returnType)
	setDepthDefaultIfNil(&depth)
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	if depth == -1 || depth.(int) > stack.Depth { depth = stack.Depth }
	if deepSearchType == DEEPSEARCH_False { depth = 1 }

	// break recursion at depth == 0
	if depth.(int) != 0 {
		// add to return
		for i := range stack.Cards {
			c := stack.Cards[i]
			// if this Card's val is a Stack
			subStack, isStack := c.Val.(*Stack)
			
			if isStack {
				if depth.(int) > 1 {
					matrix = append(matrix, subStack.ToMatrix(returnType, deepSearchType, depth.(int) - 1))
				} else {
					matrix = append(matrix, []any {})
				}
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

/** Returns the shape of this stackMatrix, or nil if irregular shape

 @receiver `stack` type{*Stack}
 @returns type{[]int}
 */
 func (stack *Stack) Shape() (shape []int) {

	// body
	if stack.IsRegular() {

		shape = append(shape, stack.Size)

		if stack.Size > 0 {
			_, hasSubstack := stack.Cards[0].Val.(*Stack)
			if hasSubstack {
				shape = append(shape, stack.Cards[0].Val.(*Stack).Shape()...)
			}
		}

	} else {
		shape = nil
	}

	// return
	return shape

}

/** Returns whether the matrix is of a regular shape

 @receiver `stack` type{*Stack}
 @returns type{bool}
 @ensures
   * example:
       {{1, 2}, 3} == irregular/false
       {{1, 2}, {3}} == irregular/false
       {{1, 2}, {3, 4}} == regular/true
	   {1, 3} == regular/true
	   {} == regular/true
 */
 func (stack *Stack) IsRegular() bool {

	// init
	test := true
	normDepth := -1
	normSize := -1
	normSubstack := -1

	// body
	for _, c := range stack.Cards {
		substack, hasSubstack := c.Val.(*Stack)

		if hasSubstack {
			if normDepth == -1 {
				normDepth = substack.Depth
			} else if normDepth != substack.Depth {
				test = false
			}
			if normSize == -1 {
				normSize = substack.Size
			} else if normSize != substack.Size {
				test = false
			}
			if normSubstack == -1 {
				normSubstack = 0
			} else if normSubstack != 0 {
				test = false
			}
			test = test && substack.IsRegular()
		} else {
			if normSubstack == -1 {
				normSubstack = 1
			} else if normSubstack != 1 {
				test = false
			}
		}
	}

	return test

}

/** Adds the cards in `stack` to itself `n` - 1 times
  (duplicate 4 means 3 duplicates made; duplicate 1 means don't duplicate; duplicate 0 means empty)
 
 @receiver `stack` type{*Stack}
 @param optional `n` type{int} default 2
 @updates `stack`
 @returns `stack`
 */
 func (stack *Stack) Duplicate(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var n any
	gogenerics.UnpackVariadic(variadic, &n)
	if n == nil {
		n = 2
	}

	var cardsSave []*Card
	cardsSave = append(cardsSave, stack.Cards...)
	stack.Cards = []*Card {}

	for i := 0; i < n.(int); i++ {
		stack.Cards = append(stack.Cards, cardsSave...)
	}

	stack.setStackProperties()

	return stack

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

 @receiver `card` type{*Card}
 @returns type{*Card} card clone
 @constructs type{*Card} clone of `card`
*/
func (card *Card) Clone() *Card {

	// init
	clone := gogenerics.CloneObject(card).(*Card)
	clone.Idx = card.Idx
	if gogenerics.IsPointer(card.Key) {
		clone.Key = gogenerics.CloneObject(card.Key)
	} else {
		clone.Key = card.Key
	}
	if gogenerics.IsPointer(card.Val) {
		clone.Val = gogenerics.CloneObject(card.Val)
	} else {
		clone.Val = card.Val
	}
	
	// return
	return clone

}

/** Returns a clone of the given stack

 @receiver `stack` type{*Stack}
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int} default -1 (deepest)
 @param optional `cloneCardKeys` type{CLONE} default CLONE_True
 @param optional `cloneCardVals` type{CLONE} default CLONE_True
 @param optional `cloneSubstackKeys` type{CLONE} default CLONE_True
 @param optional `cloneSubstackVals` type{CLONE} default CLONE_True
	only set to false if you want all cards containing substacks to erase those substacks, replacing them with nils, preserving their space in the stack as cards with nil values
      (will only work on the first layer)
 @returns type{*Stack} stack clone
 @constructs type{*Stack} clone of `stack`
 @ensures
   * if you set clone to false, then that property will be cloned as nil
   * if you shallow clone a deep stack and cloneSubstackVals is true, then the original substacks will be held in the clone
*/
func (stack *Stack) Clone(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var deepSearchType, depth, cloneCardKeys, cloneCardVals, cloneSubstackKeys, cloneSubstackVals any
	gogenerics.UnpackVariadic(variadic, &deepSearchType, &depth, &cloneCardKeys, &cloneCardVals, &cloneSubstackKeys, &cloneSubstackVals)
	// set defaults
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setDepthDefaultIfNil(&depth)
	setCLONEDefaultIfNil(&cloneCardKeys)
	setCLONEDefaultIfNil(&cloneCardVals)
	setCLONEDefaultIfNil(&cloneSubstackKeys)
	setCLONEDefaultIfNil(&cloneSubstackVals)
	if depth == -1 || depth.(int) > stack.Depth { depth = stack.Depth }
	if deepSearchType == DEEPSEARCH_False { depth = 1 }

	// init
	clone := new(Stack)
	// recursive loop
	for _, originalCard := range stack.Cards {

		// set up new card
		newCard := MakeCard()
		clone.Cards = append(clone.Cards, newCard)

		// set card properties depending on whether it's a substack
		substack, isSubstack := originalCard.Val.(*Stack)
		if isSubstack {

			if cloneSubstackKeys == CLONE_True {
				newCard.Key = originalCard.Key
			}
			if cloneSubstackVals == CLONE_True {

				// forwardpropagate if depth > 1 and we're deepsearching
				if deepSearchType == DEEPSEARCH_True && depth.(int) > 1 {
					newCard.Val = substack.Clone(deepSearchType, depth.(int) - 1, cloneCardKeys, cloneCardVals, cloneSubstackKeys, cloneSubstackVals)
				} else {
					newCard.Val = substack
				}

			}

		} else { // is non-substack card

			if cloneCardKeys == CLONE_True {
				newCard.Key = originalCard.Key
			}
			if cloneCardVals == CLONE_True {
				newCard.Val = originalCard.Val
			}

		}


	}

	// set depth and size
	clone.setStackProperties()

	// return
	return clone

}

/** Removes all cards from `stack` which share the same field value as another card in that stack and returns the new stack
 Assuming elements represent the values of cards in the pre-existing stack,
 Stack{"Hi", "Hey", "Hello", "Hi", "Hey", "Howdy"}.Unique(TYPE_Val) => Stack{"Hi", "Hey", "Hello", "Howdy"}

 @receiver `stack` type{*Stack}
 @param `unique optionalType` type{TYPE} default TYPE_Val
 @returns `stack`
 @updates `stack` to have no repeating values between field `uniqueType`
 */
func (stack *Stack) Unique(variadic ...any) *Stack {
	
	// unpack variadic into optional parameters
	var uniqueType any
	gogenerics.UnpackVariadic(variadic, &uniqueType)
	if uniqueType == nil { uniqueType = TYPE_Val }

	// main
	return stack.GetMany(FIND_Lambda, func(card *Card, _ *Stack, _ bool, workingStack *Stack, wm ...any) (bool) {
		if workingStack.Size == 0 {
			return true
		} else {
			switch wm[0].(TYPE) {
			case TYPE_Key:
				return !workingStack.Has(FIND_Key, card.Key)
			case TYPE_Val:
				return !workingStack.Has(FIND_Val, card.Val)
			}
			return false // just so it compiles
		}
	}, nil, nil, nil, nil, nil, nil, nil, []any {uniqueType})

}

/** Returns whether two cards equal one another
 
 @receiver `thisCard` type{*Card}
 @param `otherCard` type{*Card}
 @param optional `pointerTypeKey` type{POINTER} default POINTER_False
 @param optional `pointerTypeVal` type{POINTER} default POINTER_False
 @param optional `compareIdxs` type{COMPARE} default COMPARE_False
 @param optional `compareKeys` type{COMPARE} default COMPARE_True
 @param optional `compareVals` type{COMPARE} default COMPARE_True
 @returns type{bool}
 */
func (thisCard *Card) Equals(otherCard *Card, variadic ...any) bool {

	// unpack variadic into optional parameters
	var pointerTypeKey, pointerTypeVal, compareIdxs, compareKeys, compareVals any
	gogenerics.UnpackVariadic(variadic, &pointerTypeKey, &pointerTypeVal, &compareIdxs, &compareKeys, &compareVals)
	// set default vals
	setPOINTERDefaultIfNil(&pointerTypeKey)
	setPOINTERDefaultIfNil(&pointerTypeVal)
	if compareIdxs == nil {compareIdxs = COMPARE_False}
	setCOMPAREDefaultIfNil(&compareKeys)
	setCOMPAREDefaultIfNil(&compareVals)
	/*setPRINTDefaultIfNil(&printType)

	print := func(printType any, stringToPrint string) {
		if printType.(PRINT) == PRINT_True {
			fmt.Printf("-     DETAIL: CONDITION: %v\n", stringToPrint)
		}
	}*/

	condition := thisCard != nil && otherCard != nil
	
	condition = condition && 
		(compareKeys == COMPARE_False ||
		(compareKeys == COMPARE_True &&
			(
				(pointerTypeKey == POINTER_False && thisCard.Key == otherCard.Key) ||
				(pointerTypeKey == POINTER_True && gogenerics.PointersEqual(thisCard.Key, otherCard.Key) ) ) ) )
	//print(printType, fmt.Sprintf("KEY PASSES EQUALITY CHECK: %v: (compareKeys == COMPARE_False [%v] || (compareKeys == COMPARE_True [%v] && ( (pointerTypeKey == POINTER_False [%v] && thisCard.Key == otherCard.Key [%v]) || (pointerTypeKey == POINTER_True [%v] && gogenerics.PointersEqual(thisCard.Key, otherCard.Key) [%v] ) ) ) )", (compareKeys == COMPARE_False || (compareKeys == COMPARE_True && ( (pointerTypeKey == POINTER_False && thisCard.Key == otherCard.Key) || (pointerTypeKey == POINTER_True && gogenerics.PointersEqual(thisCard.Key, otherCard.Key) ) ) ) ), compareKeys == COMPARE_False, compareKeys == COMPARE_True, pointerTypeKey == POINTER_False, thisCard.Key == otherCard.Key, pointerTypeKey == POINTER_True, gogenerics.PointersEqual(thisCard.Key, otherCard.Key) ))
	
	condition = condition && 
		(compareVals == COMPARE_False ||
		(compareVals == COMPARE_True &&
			(
				(pointerTypeVal == POINTER_False && thisCard.Val == otherCard.Val) ||
				(pointerTypeVal == POINTER_True && gogenerics.PointersEqual(thisCard.Val, otherCard.Val) ) ) ) )
	//print(printType, fmt.Sprintf("Val PASSES EQUALITY CHECK: %v: (compareVals == COMPARE_False [%v] || (compareVals == COMPARE_True [%v] && ( (pointerTypeVal == POINTER_False [%v] && thisCard.Val == otherCard.Val [%v]) || (pointerTypeVal == POINTER_True [%v] && gogenerics.PointersEqual(thisCard.Val, otherCard.Val) [%v] ) ) ) )", (compareVals == COMPARE_False || (compareVals == COMPARE_True && ( (pointerTypeVal == POINTER_False && thisCard.Val == otherCard.Val) || (pointerTypeVal == POINTER_True && gogenerics.PointersEqual(thisCard.Val, otherCard.Val) ) ) ) ), compareVals == COMPARE_False, compareVals == COMPARE_True, pointerTypeVal == POINTER_False, thisCard.Val == otherCard.Val, pointerTypeVal == POINTER_True, gogenerics.PointersEqual(thisCard.Val, otherCard.Val) ))

	condition = condition && (compareIdxs == COMPARE_False || (compareIdxs == COMPARE_True && thisCard.Idx == otherCard.Idx))
	//print(printType, fmt.Sprintf("IDX PASSES EQUALITY CHECK: %v: (compareIdxs == COMPARE_False [%v] || (compareIdxs == COMPARE_True [%v] && thisCard.Idx == otherCard.Idx [%v]))", (compareIdxs == COMPARE_False || (compareIdxs == COMPARE_True && thisCard.Idx == otherCard.Idx)), compareIdxs == COMPARE_False, compareIdxs == COMPARE_True, thisCard.Idx == otherCard.Idx))

	// return whether conditions yield true
	return condition

}

/** Returns whether two stacks equal one another
 
 @receiver `thisStack` type{*Stack}
 @param `otherStack` type{*Stack}
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `compareCardKeys` type{COMPARE} default COMPARE_True
 @param optional `compareCardVals` type{COMPARE} default COMPARE_True
 @param optional `compareSubstackKeys` type{SUBSTACKKEYS} default COMPARE_True
 @param optional `compareSubstackVals` type{SUBSTACKKEYS} default COMPARE_False
   this being set to true will compare the substacks themselves
 @param optional `pointerCardKeys` type{POINTER} default POINTER_False
 @param optional `pointerCardVals` type{POINTER} default POINTER_False
 @param optional `pointerSubstackKeys` type{POINTER} default POINTER_False
 @param optional `pointerSubstackVals` type{POINTER} default POINTER_True
   this being set to true will compare the addresses of the substacks
 @ensures
   * `stack`.Size == `otherStack`.Size is tested on the deepest layer
   * if `stack`.Depth != `otherStack`.Depth and the N-deep comparison finds that they're equal, then return that they're Equal 
 @returns type{bool}
 */
func (stack *Stack) Equals(otherStack *Stack, variadic ...any) (test bool) {

	/*
	PSEUDOCODE OUTLINE:

	stack.Equals(otherStack, ...variadic) bool

		set up variadic stuff

		if depth is stack
			depth = stack to array
		testLayer = true
		if deepSearchType is false
			if depth == -1 // first input
				depth = 1
			else
				depth = 0
				testLayer = false
		else
			if depth is an int && (depth == -1 || depth > stack.Depth)
				depth = stack.Depth
			else if depth is an []int
				if []depth does not have an element == 1
					testLayer = false
			if depth == 0
				testLayer = false
		
		test = neither stack is nil && (depth == 0 || stack.Size == otherStack.Size) // backpropagate if not considering this layer; or, if you are considering this layer, ensure each stack has the same size

		for each cardA in this stack
			for each cardB in other stack
				if cardA corresponds to cardB and test and depth != 0
					
					if cardA.Val and cardB.Val is a substack

						// compare substack properties
						if testLayer
							if compareSubstackKeys
								if pointerSubstackKeys
									test = test && compare substack keys as pointers
								else
									test = test && compare substack keys regularly
							if compareSubstackVals
								if pointerSubstackVals
									test = test && compare substack vals as pointers
								else
									test = test && compare substack vals regularly
						
						// forwardpropagate
						test = test && substackA.Equals(substackB, ..., depth = depth - 1 OR depth[] = depth[i - 1, ..., n - 1])
						
					else if one holds a substack and the other doesn't
						test = false
					
					else if neither are substacks and they are both just cards
						
						// compare card properties
						if testLayer
							test = test && cardA.Equals(cardB, [pass in pointer and compare stuff for key and val])

		// backpropagate
		return test

	*/
	
	// unpack variadic into optional parameters
	var deepSearchType, depth, compareCardKeys, compareCardVals, compareSubstackKeys, compareSubstackVals, pointerCardKeys, pointerCardVals, pointerSubstackKeys, pointerSubstackVals any
	gogenerics.UnpackVariadic(variadic, &deepSearchType, &depth, &compareCardKeys, &compareCardVals, &compareSubstackKeys, &compareSubstackVals, &pointerCardKeys, &pointerCardVals, &pointerSubstackKeys, &pointerSubstackVals)
	// set default vals
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setDepthDefaultIfNil(&depth)

	setCOMPAREDefaultIfNil(&compareCardKeys)
	setCOMPAREDefaultIfNil(&compareCardVals)
	setCOMPAREDefaultIfNil(&compareSubstackKeys)
	if compareSubstackVals == nil { compareSubstackVals = COMPARE_False }
	
	setPOINTERDefaultIfNil(&pointerCardKeys)
	setPOINTERDefaultIfNil(&pointerCardVals)
	setPOINTERDefaultIfNil(&pointerSubstackKeys)
	if pointerSubstackVals == nil { pointerSubstackVals = POINTER_True }

	depthStack, depthIsStack := depth.(*Stack)
	if depthIsStack {
		depth = []int {}
		for _, d := range depthStack.Cards {
			depth = append(depth.([]int), d.Val.(int))
		}
	}
	testLayer := true
	_, depthIsInt := depth.(int)
	if deepSearchType == DEEPSEARCH_False {
		if depth == -1 {// first input
			depth = 1
		} else {
			depth = 0
			testLayer = false
		}
	} else {
		if depthIsInt && (depth == -1 || depth.(int) > stack.Depth) {
			depth = stack.Depth
		} else if !depthIsInt { // depth is an []int
			has1 := false
			for _, d := range depth.([]int) {
				if d == 1 {
					has1 = true
					break
				}
			}
			if !has1 {
				testLayer = false
			}
		}
		if depth == 0 {
			testLayer = false
		}
	}
	
	test = stack != nil && otherStack != nil && (depth == 0 || stack.Size == otherStack.Size)

	for _, cardA := range stack.Cards {
		for _, cardB := range otherStack.Cards {
			if cardA.Idx == cardB.Idx && test && depth != 0 {

				substackA, cardAIsSubstack := cardA.Val.(*Stack)
				substackB, cardBIsSubstack := cardB.Val.(*Stack)
				
				if cardAIsSubstack && cardBIsSubstack {

					// compare substack properties
					if testLayer {
						if compareSubstackKeys == COMPARE_True {
							if pointerSubstackKeys == POINTER_True {
								test = test && gogenerics.PointersEqual(cardA.Key, cardB.Key)
							} else {
								test = test && cardA.Key == cardB.Key
							}
						}
						if compareSubstackVals == COMPARE_True {
							if pointerSubstackVals == POINTER_True {
								test = test && gogenerics.PointersEqual(cardA.Val, cardB.Val)
							} else {
								test = test && cardA.Val == cardB.Val
							}
						}
					}
					
					// forwardpropagate
					var transformedDepth any
					if depthIsInt {
						transformedDepth = depth.(int) - 1
					} else {
						transformedDepth = []int {}
						for i := range depth.([]int) {
							transformedDepth = append(transformedDepth.([]int), depth.([]int)[i] - 1)
						}
					}
					test = test && substackA.Equals(substackB, deepSearchType, transformedDepth, compareCardKeys, compareCardVals, compareSubstackKeys, compareSubstackVals, pointerCardKeys, pointerCardVals, pointerSubstackKeys, pointerSubstackVals)

				} else if (cardAIsSubstack && !cardBIsSubstack) || (!cardAIsSubstack && cardBIsSubstack) { // one holds a substack and the other doesnt

					test = false

				} else { // neither are substacks and they are both just cards

					// compare card properties
					if testLayer {
						test = test && cardA.Equals(cardB, pointerCardKeys, pointerCardVals, COMPARE_False, compareCardKeys, compareCardVals)
					}

				}
			}
		}
	}

	// backpropagate
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

	for ok := true; ok; ok = (newOrder.(bool) && stack.Size > 1 && initClone.Equals(stack)) { // emulate a do-while loop
		
		// pseudo-randomize seed
		rand.Seed(time.Now().UnixNano())

		// shuffle
		rand.Shuffle(stack.Size, func(i, j int) { stack.Cards[i], stack.Cards[j] = stack.Cards[j], stack.Cards[i] })
		
		// set indices
		stack.setStackProperties()

	}

	// return
	return stack

}

/** Transposes the ordering of `stack.Cards`
  Conceptually, this performs the transverse function assuming maximum depth.
 
 @receiver `stack` type{*Stack}
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates `stack` to have its ordering reversed
 */
func (stack *Stack) Transpose(variadic ...any) *Stack {

	// TODO: implement substackKeysType
	// TODO: ensure depth works

	// unpack variadic into optional parameters
	var deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &deepSearchType, &depth)
	// set defaults
	setDepthDefaultIfNil(&depth)
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	if depth == -1 || depth.(int) > stack.Depth { depth = stack.Depth }
	if deepSearchType == DEEPSEARCH_False { depth = 1 }

	// body
	/*
	stack.Lambda(func(card *Card, parentStack *Stack, _ ...any) {
		stack.Move(FIND_Card, ORDER_Before, FIND_Idx, card, 0)
		switch card.Val.(type) {
		case *Stack:
			if depth.(int) > 1 { // forwardpropagate
				card.Val.(*Stack).Transpose(deepSearchType, depth)
			}
		}
	})
	*/

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
	if depth == nil { depth = 0 }

	// prints
	fmt.Printf("%v|%vCARD\n", depthPrinter(depth.(int)), gogenerics.IfElse(depth == 0, "gostack: PRINTING ", ""))
	if card == nil {
		fmt.Printf("%v- card:          %v\n", depthPrinter(depth.(int)), nil)
	} else {
		fmt.Printf("%v- &card:         %v\n", depthPrinter(depth.(int)), fmt.Sprintf("%p", card))
		fmt.Printf("%v- card.Idx:      %v\n", depthPrinter(depth.(int)), card.Idx)
		if gogenerics.IsPointer(card.Key) {
			fmt.Printf("%v- &card.Key:     %v\n", depthPrinter(depth.(int)), fmt.Sprintf("%p", card.Key))
			fmt.Printf("%v- card.Key:      %v\n", depthPrinter(depth.(int)), reflect.ValueOf(card.Key).Elem())
			fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(depth.(int)), reflect.TypeOf(reflect.ValueOf(card.Key).Elem().Interface()))
		} else {
			fmt.Printf("%v- card.Key:      %v\n", depthPrinter(depth.(int)), card.Key)
			if card.Key != nil {
				fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(depth.(int)), reflect.TypeOf(card.Key))
			}
		}
		if gogenerics.IsPointer(card.Val) {
			fmt.Printf("%v- &card.Val:     %v\n", depthPrinter(depth.(int)), fmt.Sprintf("%p", card.Val))
			fmt.Printf("%v- card.Val:      %v\n", depthPrinter(depth.(int)), reflect.ValueOf(card.Val).Elem())
			fmt.Printf("%v- card.Val.Type: (%v)\n", depthPrinter(depth.(int)), reflect.TypeOf(reflect.ValueOf(card.Val).Elem().Interface()))
		} else {
			fmt.Printf("%v- card.Val:      %v\n", depthPrinter(depth.(int)), card.Val)
			if card.Val != nil {
				fmt.Printf("%v- card.Val.Type: (%v)\n", depthPrinter(depth.(int)), reflect.TypeOf(card.Val))
			}
		}
	}

}

/** Prints information regarding `stack` to the console
 
 @receiver `stack` type{*Stack}
 @param optional `depth` type{int} default 0
   This variable only exists for text-indenting purposes to make your terminal output look a bit cleaner.  1 depth => 4 "-" added before the print.
 @updates terminal logs
 */
func (stack *Stack) Print(variadic ...any) {
	
	// unpack variadic into optional parameters
	var depth, idx, key any
	gogenerics.UnpackVariadic(variadic, &depth, &idx, &key)
	if depth == nil { depth = 0 }

	fmt.Printf("%v|%vSTACK\n", depthPrinter(depth.(int)), gogenerics.IfElse(idx == nil, "gostack: PRINTING ", "SUB"))
	if stack == nil {
		fmt.Printf("%v- stack:         %v\n", depthPrinter(depth.(int)), nil)
	} else {
		fmt.Printf("%v- &stack:        %v\n", depthPrinter(depth.(int)), fmt.Sprintf("%p", stack))
		if idx != nil {
			fmt.Printf("%v- card.Idx:      %v\n", depthPrinter(depth.(int)), idx)
		}
		if key != nil {
			if gogenerics.IsPointer(key) {
				fmt.Printf("%v- &card.Key:     %v\n", depthPrinter(depth.(int)), key)
				fmt.Printf("%v- card.Key:      %v\n", depthPrinter(depth.(int)), reflect.ValueOf(key).Elem())
				fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(depth.(int)), reflect.TypeOf(reflect.ValueOf(key).Elem().Interface()))
			} else {
				fmt.Printf("%v- card.Key:      %v\n", depthPrinter(depth.(int)), key)
				if key != nil {
					fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(depth.(int)), reflect.TypeOf(key))
				}
			}
		}
		fmt.Printf("%v- stack.Size:    %v\n", depthPrinter(depth.(int)), stack.Size)
		fmt.Printf("%v- stack.Depth:   %v\n", depthPrinter(depth.(int)), stack.Depth)
		for i := range stack.Cards {
			c := stack.Cards[i]
	
			switch c.Val.(type) {
			case *Stack:
				c.Val.(*Stack).Print(depth.(int)+4, i, c.Key)
			default:
				c.Print(depth.(int)+4)
			}
		}
	}
	
}

/** Iterate through a stack calling your lambda function on each card
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @returns `stack` type{*Stack}
 @returns `retStack` type{*Stack}
 @returns `retCard` type{*Card}
 @returns `retVarAdr` type{any}
 */
func (stack *Stack) Lambda(lambda func(*Card, *Stack, bool, *Stack, *Card, any, ...any), variadic ...any) (*Stack, *Stack, *Card, any) {

	/**
	PSEUDOCODE OUTLINE:

	(stack) Lambda(lambda, ...variadic) (retAdr any)

		set up variadic stuff

		if depth is Stack
			depth = stack to array
		passLayer = true
		if deepSearchType is false
			if depth == -1 // first input
				depth = 1
			else
				depth = 0
				passLayer = false
		else
			if depth is an int && (depth == -1 || depth > stack.Depth)
				depth = stack.Depth
			else if depth is an []int
				if []depth does not have an element == 1
					passLayer = false
			if depth == 0
				passLayer = false
		
		for each card in this stack
			
			if card is substack

				if passSubstacks and passLayer
					pass substack into lambda(isSubstack = true)

				if depth > 1 or depth[] has an element > 1 // forwardpropagate
					substack.Lambda(..., depth = depth - 1 OR depth[] = depth[i - 1, ..., n - 1])

			else if card is not substack

				if passCards and passLayer
					pass card into lambda(isSubstack = false)

		return *retAdr.(*any)

	*/

	// unpack variadic into optional parameters
	var retStack, retCard, retVarAdr, workingMem, deepSearchType, depth, passSubstacks, passCards any
	gogenerics.UnpackVariadic(variadic, &retStack, &retCard, &retVarAdr, &workingMem, &deepSearchType, &depth, &passSubstacks, &passCards)
	if retStack == nil {retStack = MakeStack()}
	if retCard == nil {retCard = MakeCard()}
	if retVarAdr == nil {var o any; retVarAdr = &o;}
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setDepthDefaultIfNil(&depth)
	if passSubstacks == nil {passSubstacks = PASS_False}
	if passCards == nil {passCards = PASS_True}
	
	// main
	depthStack, depthIsStack := depth.(*Stack)
	if depthIsStack {
		depth = []int {}
		for _, d := range depthStack.Cards {
			depth = append(depth.([]int), d.Val.(int))
		}
	}
	passLayer := true
	has1 := false
	hasOver1 := false
	_, depthIsInt := depth.(int)
	if deepSearchType == DEEPSEARCH_False {
		if depth == -1 {// first input
			depth = 1
		} else {
			depth = 0
			passLayer = false
		}
	} else {
		if depthIsInt && (depth == -1 || depth.(int) > stack.Depth) {
			depth = stack.Depth
		} else if !depthIsInt { // depth is an []int
			for _, d := range depth.([]int) {
				if d == 1 {
					has1 = true
				} else if d > 1 {
					hasOver1 = true
				}
			}
			if !has1 {
				passLayer = false
			}
		}
		if depth == 0 {
			passLayer = false
		}
	}
	
	for _, card := range stack.Cards {
		
		substack, isSubstack := card.Val.(*Stack)
		if isSubstack {

			if passSubstacks == PASS_True && passLayer {
				lambda(card, stack, true, toTypeStack(retStack), toTypeCard(retCard), &retVarAdr, workingMem.([]any)...)

				// update properties
				stack.setStackProperties()
				if retStack != nil {
					retStack.(*Stack).setStackProperties()
				}
			}

			// forwardpropagate
			if (depthIsInt && depth.(int) > 1) || hasOver1 {
				var transformedDepth any
				if depthIsInt {
					transformedDepth = depth.(int) - 1
				} else {
					transformedDepth = []int {}
					for i := range depth.([]int) {
						transformedDepth = append(transformedDepth.([]int), depth.([]int)[i] - 1)
					}
				}

				substack.Lambda(lambda, retStack, retCard, retVarAdr, workingMem, deepSearchType, transformedDepth, passSubstacks, passCards)
			}

		} else { // card is not substack

			if passCards == PASS_True && passLayer {
				lambda(card, stack, false, toTypeStack(retStack), toTypeCard(retCard), &retVarAdr, workingMem.([]any)...)

				// update properties
				stack.setStackProperties()
				if retStack != nil {
					retStack.(*Stack).setStackProperties()
				}
			}

		}

	}

	return stack, toTypeStack(retStack), toTypeCard(retCard), retVarAdr

}

/** Iterate through a stack calling your lambda function on each card, returning only `stack`
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @returns `stack` type{*Stack}
 */
func (stack *Stack) LambdaThis(lambda func(*Card, *Stack, bool, *Stack, *Card, any, ...any), variadic ...any) *Stack {
	stack.Lambda(lambda, variadic...)
	return stack
}

/** Iterate through a stack calling your lambda function on each card, returning only `retStack`
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @returns `retStack` type{*Stack}
 */
func (stack *Stack) LambdaStack(lambda func(*Card, *Stack, bool, *Stack, *Card, any, ...any), variadic ...any) *Stack {
	_, thisStack, _, _ := stack.Lambda(lambda, variadic...)
	return thisStack
}

/** Iterate through a stack calling your lambda function on each card, returning only `retCard`
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @returns `retCard` type{*Card}
 */
func (stack *Stack) LambdaCard(lambda func(*Card, *Stack, bool, *Stack, *Card, any, ...any), variadic ...any) *Card {
	_, _, thisCard, _ := stack.Lambda(lambda, variadic...)
	return thisCard
}

/** Iterate through a stack calling your lambda function on each card, returning only `retVarAdr`
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @returns `retVarAdr` type{any}
 */
func (stack *Stack) LambdaVarAdr(lambda func(*Card, *Stack, bool, *Stack, *Card, any, ...any), variadic ...any) any {
	_, _, _, retVarAdr := stack.Lambda(lambda, variadic...)
	return retVarAdr
}

/** Adds to a stack of cards or a cards at (each) position(s) and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `insert` type{any, []any, Stack}
 @param optional `orderType` type{ORDER} default ORDER_After
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `actionType` type{ACTION} default ACTION_All
 @param optional `overrideCards` type{OVERRIDE} default OVERRIDE_False
   By default, if you do stack.Add(cardA), stack = {cardA}.  If you instead desire stack = {Card {val = cardA}}, do true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack` if valid find, or nil if invalid find
 */
func (stack *Stack) Add(insert any, variadic ...any) *Stack {
	
	// unpack variadic into optional parameters
	var orderType, findType, findData, findCompareRaw, actionType, overrideCards, deepSearchType, depth, pointerType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(variadic, &orderType, &findType, &findData, &findCompareRaw, &actionType, &overrideCards, &deepSearchType, &depth, &pointerType, &passSubstacks, &passCards, &workingMem)
	setACTIONDefaultIfNil(&actionType)
	setOVERRIDEDefaultIfNil(&overrideCards)
	setORDERDefaultIfNil(&orderType)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// add card
	cardAddedAdr := stack.LambdaVarAdr(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, wmadrs ...any) {
		
		actionType := wmadrs[4].(ACTION)
		// only do add to the first match if ACTION_First, otherwise do for every match
		if (selectCard(findType, findData, pointerType, wmadrs[0].(COMPARE), "card", card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs[5:]...)) && actionType == ACTION_All || (actionType == ACTION_First && !gogenerics.GetPointer(retVarAdr).(bool)) {

			// update variable to show that we have successfully found a card to whom's stack we will add before or after
			gogenerics.SetPointer(retVarAdr, true)

			// initialize variables
			insert := wmadrs[1]
			orderType := wmadrs[2].(ORDER)
			overrideCards := wmadrs[3].(OVERRIDE)
			insertArr := []any {}
			insertCards := []*Card {}

			// set up insertArr
			switch getType(insert, false) {
			case "element":
				insertArr = append(insertArr, insert)
			case "slice":
				insertArr = gogenerics.UnpackArray(insert)
			case "stack":
				insertArr = insert.(*Stack).ToArray()
			}

			// set up insertCards
			for _, ins := range insertArr {
				insCard, isCard := ins.(*Card)
				if isCard && overrideCards == OVERRIDE_False {
					insertCards = append(insertCards, insCard.Clone()) // insert a clone of this card
				} else {
					insertCards = append(insertCards, MakeCard(ins)) // insert a card whose val is ins
				}
			}

			// update the stack to have insert at its respective location
			targetIdx := card.Idx
			switch orderType {
			case ORDER_Before:
				targetIdx += 0
			case ORDER_After:
				targetIdx += 1
			}

			beginningSegment := parentStack.Cards[:targetIdx]
			endSegment := parentStack.Cards[targetIdx:]

			parentStack.Cards = []*Card {}
			parentStack.Cards = append(parentStack.Cards, beginningSegment...)
			parentStack.Cards = append(parentStack.Cards, insertCards...)
			parentStack.Cards = append(parentStack.Cards, endSegment...)
			
		}

	}, nil, nil, false, append([]any{findCompareRaw, insert, orderType, overrideCards, actionType}, workingMem.([]any)...), deepSearchType, depth, passSubstacks, passCards)

	// return nil if no add was made, else return card
	if !gogenerics.GetPointer(cardAddedAdr).(bool) {
		return nil
	} else {
		return stack
	}

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
 @param optional `deepSearchType_from` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `deepSearchType_to` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `substackKeysType_from` type{SUBSTACKKEYS} default SUBSTACKKEYS_False
 @param optional `substackKeysType_to` type{SUBSTACKKEYS} default SUBSTACKKEYS_False
 @param optional `depth_from` type{int} default -1 (deepest)
 @param optional `depth_to` type{int} default -1 (deepest)
 @updates `stack` if move was performed
 @returns `stack` if moved OR nil if no move occurred (due to bad find)
 @requires you are not moving a stack to a location within that own stack
 @ensures a stack of cards, or individual cards, can be targeted
 */
func (stack *Stack) Move(findType_from FIND, orderType ORDER, findType_to FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData_from, findData_to, pointerType_from, pointerType_to, deepSearchType_from, deepSearchType_to, substackKeysType_from, substackKeysType_to, depth_from, depth_to any
	gogenerics.UnpackVariadic(variadic, &findData_from, &findData_to, &pointerType_from, &pointerType_to, &deepSearchType_from, &deepSearchType_to, &substackKeysType_from, &substackKeysType_to, &depth_from, &depth_to)

	// 1) Get the cards to move
	from := stack.Clone().ExtractMany(findType_from, findData_from, pointerType_from, RETURN_Cards, deepSearchType_from, depth_from)

	// 2) Get the card to put froms before/after depending on whether before or after
	var toCard *Card
	toStack := stack.GetMany(findType_to, findData_to, RETURN_Cards, pointerType_to, CLONE_False, CLONE_False, CLONE_False, deepSearchType_to, depth_to)
	toCard = gogenerics.IfElse(orderType == ORDER_After, toStack.Get(FIND_Last), toStack.Get(FIND_First)).(*Card)
	
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
 @param optional `deepSearchType_first` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `deepSearchType_second` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `substackKeysType_first` type{SUBSTACKKEYS} default SUBSTACKKEYS_False
 @param optional `substackKeysType_second` type{SUBSTACKKEYS} default SUBSTACKKEYS_False
 @param optional `depth_first` type{int} default -1 (deepest)
 @param optional `depth_second` type{int} default -1 (deepest)
 @returns `stack` if moved OR nil if no move occurred (due second bad find)
 @updates `stack`
 @requires you are not swapping a stack with a location within that own stack
 @ensures a stack of cards, or individual cards, can be targeted
 */
func (stack *Stack) Swap(findType_first FIND, findType_second FIND, variadic ...any) *Stack {

	// unpack variadic insecond optional parameters
	var findData_first, findData_second, pointerType_first, pointerType_second, deepSearchType_first, deepSearchType_second, depth_first, substackKeysType_first, substackKeysType_second, depth_second any
	gogenerics.UnpackVariadic(variadic, &findData_first, &findData_second, &pointerType_first, &pointerType_second, &deepSearchType_first, &deepSearchType_second, &substackKeysType_first, &substackKeysType_second, &depth_first, &depth_second)

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

/** Returns a bool for whether a card can be found in a stack

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
 @returns type{bool}
 */
func (stack *Stack) Has(variadic ...any) bool {

	// return
	return stack.Get(variadic...) != nil

}

/** Gets the first card from specified parameters in a stack, or nil if does not exist

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns type{*Card} the found card OR nil (if invalid find)
    IF `find` is FIND_Lambda, `findData` is of type{ func(card *Card, parentStack *Stack, isSubstack bool, workingMem ...any) (bool) }
*/
 func (stack *Stack) Get(variadic ...any) *Card {
	
	// unpack variadic into optional parameters
	var findType, findData, findCompareRaw, deepSearchType, depth, pointerType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(variadic, &findType, &findData, &findCompareRaw, &deepSearchType, &depth, &pointerType, &passSubstacks, &passCards, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// get card
	out := stack.LambdaCard(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, wmadrs ...any) {
		
		if selectCard(findType, findData, pointerType, wmadrs[0].(COMPARE), "card", card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs[1:]...) && retCard.Idx == -1 {
			*retCard = *card
		}

	}, nil, nil, nil, append([]any{findCompareRaw}, workingMem.([]any)...), deepSearchType, depth, passSubstacks, passCards)

	// return nil if no card found, else return card
	if out.Idx == -1 {
		return nil
	} else {
		return out
	}

}

/** Gets a stack of cards from specified parameters in a stack (whose cards are clones of the found cards)

 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
 @constructs a new stack
 @returns type{*Stack} the new stack
 @requires
   IF `find` is FIND_Lambda, `findData` is of type{ func(card *Card, parentStack *Stack, isSubstack bool, workingStack *Stack, workingMem ...any) (bool) }
 */
func (stack *Stack) GetMany(findType FIND, variadic ...any) *Stack {
	
	// unpack variadic into optional parameters
	var findData, findCompareRaw, returnType, deepSearchType, depth, pointerType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(variadic, &findData, &findCompareRaw, &returnType, &deepSearchType, &depth, &pointerType, &passSubstacks, &passCards, &workingMem)
	if returnType == nil {returnType = RETURN_Cards}
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// make new stack and return
	return stack.LambdaStack(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, wmadrs ...any) {
		
		if selectCard(findType, findData, pointerType, wmadrs[0].(COMPARE), "stack", card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs[1:]...) {

			var outCard *Card
			switch returnType {
			case RETURN_Keys:
				outCard = MakeCard(card.Key)
			case RETURN_Vals:
				outCard = MakeCard(card.Val)
			case RETURN_Idxs:
				outCard = MakeCard(card.Idx)
			case RETURN_Cards:
				outCard = card.Clone()
			}
			retStack.Cards = append(retStack.Cards, outCard)
			
		}

	}, nil, nil, nil, append([]any{findCompareRaw}, workingMem.([]any)...), deepSearchType, depth, passSubstacks, passCards)

}

/** Returns a clone of a found card before its respective field is updated to `replaceWith` (OR nil if not found)
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceWith` type{any, []any, *Stack}
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `overrideCards` type{OVERRIDE} default OVERRIDE_False
   By default, if you do stack.Add(cardA), stack = {cardA}.  If you instead desire stack = {Card {val = cardA}}, do true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int} default -1 (deepest)
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns type{*Card}
 */
func (stack *Stack) Replace(replaceType REPLACE, replaceWith any, variadic ...any) *Card {

	// unpack variadic into optional parameters
	var findType, findData, findCompareRaw, overrideCards, deepSearchType, depth, pointerType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(variadic, &findType, &findData, &findCompareRaw, &overrideCards, &deepSearchType, &depth, &pointerType, &passSubstacks, &passCards, &workingMem)
	setOVERRIDEDefaultIfNil(&overrideCards)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// main
	return stack.LambdaCard(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, wmadrs ...any) {
		
		if selectCard(findType, findData, pointerType, wmadrs[0].(COMPARE), "card", card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs[4:]...) && retCard.Idx == -1 {

			*retCard = *card.Clone() // return the original card

			// initialize variables
			replaceType := wmadrs[1].(REPLACE)
			replaceWith := wmadrs[2]

			// replace mechanism
			switch replaceType {
			case REPLACE_Key:

				card.Key = replaceWith

			case REPLACE_Val:

				card.Val = replaceWith

			case REPLACE_Card:

				// initialize variables
				overrideCards := wmadrs[3].(OVERRIDE)
				insertArr := []any {}
				insertCards := []*Card {}

				// set up insertArr
				switch getType(replaceWith, false) {
				case "element":
					insertArr = append(insertArr, replaceWith)
				case "slice":
					insertArr = gogenerics.UnpackArray(replaceWith)
				case "stack":
					insertArr = replaceWith.(*Stack).ToArray()
				}
	
				// set up insertCards
				for _, ins := range insertArr {
					insCard, isCard := ins.(*Card)
					if isCard && overrideCards == OVERRIDE_False {
						insertCards = append(insertCards, insCard.Clone()) // insert a clone of this card
					} else {
						insertCards = append(insertCards, MakeCard(ins)) // insert a card whose val is ins
					}
				}
	
				// update the stack to have replaceWith at its respective location
				targetIdx := card.Idx
				beginningSegment := parentStack.Cards[:targetIdx]
				endSegment := parentStack.Cards[targetIdx+1:]
	
				parentStack.Cards = []*Card {}
				parentStack.Cards = append(parentStack.Cards, beginningSegment...)
				parentStack.Cards = append(parentStack.Cards, insertCards...)
				parentStack.Cards = append(parentStack.Cards, endSegment...)

			case REPLACE_Lambda:

				/*switch returnType {
				case "card":
					return findData.(func(*Card, *Stack, bool, ...any) (bool)) (card, stack, isSubstack, wmadrs...)
				case "stack":
					return findData.(func(*Card, *Stack, bool, *Stack, ...any) (bool)) (card, stack, isSubstack, retStack, wmadrs...)
				}*/

			}

		}

	}, nil, nil, false, append([]any{findCompareRaw, replaceType, replaceWith, overrideCards}, workingMem.([]any)...), deepSearchType, depth, passSubstacks, passCards)

}

/** Returns a stack whose values are the original fields updated to `replaceWith`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceWith` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} a stack whose values are the extracted cards pre-update (if find fails, then an empty stack)
 @updates all found cards to `replaceWith` in `stack`
 @ensures IF `replaceWith` is nil and `replaceType is REPLACE_Card`, the cards found will be removed from `stack`
 */
func (stack *Stack) ReplaceMany(replaceType REPLACE, replaceWith any, findType FIND, variadic ...any) (ret *Stack) {

	// unpack variadic into optional parameters
	var findData, returnType, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &returnType, &pointerType, &deepSearchType, &depth)
/*
	// get deep copy of targeted cards to return
	ret = stack.GetMany(findType, findData, returnType, pointerType, CLONE_True, CLONE_True, CLONE_True)
	// get target data
	_, targetCards, targetStacks := stack.getPositions(false, findType, findData, pointerType.(POINTER), deepSearchType.(DEEPSEARCH), depth.(int))

	// set targeted cards' fields to replaceWith if was found (updateRespectiveField fulfills our ensures clause)
	if len(targetCards) != 0 {
		for i := range targetCards {
			targetStacks[i].updateRespectiveField(replaceType, replaceWith, targetCards[i])
		}
	}

	// update properties
	stack.setStackProperties()
*/
	// return
	return

}

/** Updates a card in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceWith` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates the found card in `stack`
 @requires `stack.Replace()` has been implemented
 @ensures IF `replaceWith` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 */
func (stack *Stack) Update(replaceType REPLACE, replaceWith any, findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, pointerType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &pointerType, &deepSearchType, &depth)

	// update stack
	stack.Replace(replaceType, replaceWith, findType, findData, pointerType, deepSearchType, depth)

	// return the original stack
	return stack

}

/** Gets and removes a card from `stack`, or returns nil if does not exist
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
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
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
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
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
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
