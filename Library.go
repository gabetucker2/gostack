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
	card.Key = key
	card.Val = val

	// return
	return card

}

/** Creates a stack of cards with optional starting cards
 
 @param optional `input1` type{[]any, map[any]any} default nil
 @param optional `input2` type{[]any} default nil
 @param optional `repeats` type{int} default 1
 @returns type{*Stack} the newly-constructed stack of newly-constructed cards
 @constructs type{*Stack} a newly-constructed stack of newly-constructed type{*Card} cards
 @requires
  * `input1` is map and nil `input2`
      OR `input1` is an array and nil `input2`
	  OR `input1` is an array and `input2` is an array
	  OR `input1` is nil and `input2` is an array
  * IF `input1` and `input2` are both passed as arguments
      |`input1`| == |`input2`|
  * `MakeStackMatrix()` has been implemented
 @ensures
  * repeats the function filling `repeats` (or, if nil or under 0, 1) amount of times
  * IF `input1` is passed
      IF `input1` is a map
        unpack the map into new cards with corresponding keys and vals
      ELSEIF `input1` is an array and `input2` is not passed/nil
        unpack values from `input1` into new cards
      ELSEIF `input1` is an array and `input2` is an array
        unpack keys from `input1` and values from `input2` into new cards
      ELSEIF `input1` is nil and `input2` is an array
        unpack keys from `input2` into new cards
    ELSE
      the stack is empty
 */
func MakeStack(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var input1, input2, repeats any
	gogenerics.UnpackVariadic(variadic, &input1, &input2, &repeats)

	// set default
	if repeats == nil {
		repeats = 1
	}

	// new stack
	stack := new(Stack)

	// run MakeStackMatrix to 1D array and add to our stack `repeats` times
	for i := 0; i < repeats.(int); i++ {
		matrixShape := -1
		if input1 == nil {
			if input2 == nil {
				matrixShape = 0
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

		if !(input1 == nil && input2 == nil) {
			stack.Cards = append(stack.Cards, MakeStackMatrix(input1, input2, []int{matrixShape}).Cards...)
		}
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
 @returns type{*Stack} a new stack
 @constructs type{*Stack} a new stack with type{*Card} new cards
 @requires
  * `MakeCard()` has been implemented
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
	  treating `input1`/`input2` as 1D arrays:
	  IF `input1` is passed
        IF `input1` is a map
          unpack the map into matrix of shape `matrixShape` with corresponding keys and vals
        ELSEIF `input1` is an array and `input2` is not passed/nil
          unpack values from `input1` into matrix of shape `matrixShape`
        ELSEIF `input1` is an array and `input2` is an array
          unpack keys from `input1` and values from `input2` into matrix of shape `matrixShape`
        ELSEIF `input1` is nil and `input2` is an array
          unpack keys from `input2` into matrix of shape `matrixShape`
	  ELSEIF `input1` is not passed
	    create a StackMatrix of shape `matrixShape` whose deepest card vals are nil
 */
func MakeStackMatrix(variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var input1, input2, matrixShape any
	gogenerics.UnpackVariadic(variadic, &input1, &input2, &matrixShape)

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
					stack.makeStackMatrixFrom1D(matrixShape.([]int), keys, vals, new(int))
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
						// unpack values from `input1` into matrix of shape `matrixShape`
						stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, input1Array, new(int))
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
						stack.makeStackMatrixFrom1D(matrixShape.([]int), input1Array, input2Array, new(int))
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
				stack.makeStackMatrixFrom1D(matrixShape.([]int), input2Array, nil, new(int))
			}

		}

	// ELSEIF `input1` is not passed
	} else {

		// IF no `matrixShape` is passed
		if matrixShape == nil {
			// the stack is empty

		// ELSEIF `matrixShape` is passed
		} else {
			// create a StackMatrix of shape `matrixShape` whose deepest card vals are nil
			stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, nil, new(int))

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

	// iterate through each selection and add them to our new stack
	for _, idx := range selections {
		c := stack.Cards[idx]
		switch c.Val.(type) {
		case Stack:
			stripped := newStack.StripStackMatrix(variadic)
			for _, idx := range firstSelection.([]int) {
				newStack.Cards = append(newStack.Cards, stripped.Cards[idx])
			}
		case Card:
			newStack.Cards = append(newStack.Cards, c.Val.(*Card))
		}
	}

	// set properties
	newStack.setStackProperties()

	// return
	return newStack

}

/** Creates a new any array from values of `stack`

 @receiver `stack` type{*Stack}
 @returns type{[]any} new array
 @requires `stack.ToMatrix()` has been implemented
 @ensures new array values correspond to `stack` values
 */
func (stack *Stack) ToArray() (arr []any) {

	// return
	return stack.ToMatrix(1).([]any)

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
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{[]any, [][]any, ..., []...[]any}
 @ensures
  * -1 depth means it will go as deep as it can
  * new map keys and values correspond to `stack` keys and values
  * example: Stack{Stack{"Hi"}, Stack{"Hello", "Hola"}, "Hey"} =>
      []any{[]any{"Hi"}, []any{"Hola", "Hello"}, "Hey"}
 */
func (stack *Stack) ToMatrix(variadic ...any) (matrix any) {

	// unpack variadic into optional parameters
	var depth any
	gogenerics.UnpackVariadic(variadic, &depth)

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
				matrix = append(matrix.([]any), subStack.ToMatrix(depth.(int) - 1))
			} else {
				matrix = append(matrix.([]any), c.Val)
			}
		}
	}

	// return
	return

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
	clone.Key = gogenerics.IfElse(cloneKey == CLONE_True, gogenerics.CloneInterface(&card.Key), card.Key)
	clone.Val = gogenerics.IfElse(cloneVal == CLONE_True, gogenerics.CloneInterface(&card.Val), card.Val)

	// return
	return clone

}

/** Returns a clone of the given stack

 @receiver `stack` type{*Stack}
 @optional param `cloneCards` type{CLONE} default false
 @optional param `cloneKeys` type{CLONE} default false
 @optional param `cloneVals` type{CLONE} default false
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
	var cloneCards, cloneKeys, cloneVals any
	gogenerics.UnpackVariadic(variadic, &cloneCards, &cloneKeys, &cloneVals)
	// cast them to bools once so we don't have to do it every iteration
	_cloneKeys := cloneKeys.(bool)
	_cloneVals := cloneVals.(bool)

	// init
	clone := new(Stack)
	clone.Size = stack.Size
	clone.Depth = stack.Depth
	if cloneCards.(bool) {
		for i := range stack.Cards {
			clone.Cards = append(clone.Cards, stack.Cards[i].Clone(_cloneKeys, _cloneVals))
		}
	} else {
		clone.Cards = stack.Cards
	}

	// return
	return clone

}

/** Removes all cards from `stack` which share the same field value as another card in that stack and returns the new stack
 Assuming elements represent the values of cards in the pre-existing stack,
 Stack{"Hi", "Hey", "Hello", "Hi", "Hey", "Howdy"}.Unique(TYPE_Val) => Stack{"Hi", "Hey", "Hello", "Howdy"}

 @receiver `stack` type{*Stack}
 @param `typeType` type{TYPE}
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
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
	var matchByType, deepSearchType, depth, uniqueType any
	gogenerics.UnpackVariadic(variadic, &matchByType, &deepSearchType, &depth, &uniqueType)

	// allow deepSearchHandler to handle Unique
	*stack = *stack.deepSearchHandler("Unique", false, FIND_All, nil, nil, matchByType, deepSearchType, depth, typeType, uniqueType, nil, nil, nil, nil, nil, nil, nil, nil)

	// set properties
	stack.setStackProperties()

	return stack

}

/** Returns whether two cards equal one another
 
 @receiver `thisCard` type{*Card}
 @param `otherCard` type{*Card}
 @param optional `compareCards` type{bool} default false
	By default, does not compare the card structs, but rather their individual values; can be set true and adjusted with `matchByTypeCard`
 @param optional `matchByTypeCard` type{MATCHBY} default MATCHBY_Object
 @param optional `matchByTypeKey` type{MATCHBY} default MATCHBY_Object
 @param optional `matchByTypeVal` type{MATCHBY} default MATCHBY_Object
 @param optional `compareIdxs` type{bool} default false
 @returns type{bool}
 */
func (thisCard *Card) Equals(otherCard *Card, variadic ...any) bool {

	// unpack variadic into optional parameters
	var compareCards, matchByTypeCard, matchByTypeKey, matchByTypeVal, compareIdxs any
	gogenerics.UnpackVariadic(variadic, &compareCards, &matchByTypeCard, &matchByTypeKey, &matchByTypeVal, &compareIdxs)
	// set default vals
	if compareCards == nil {compareCards = true}
	setMATCHBYDefaultIfNil(matchByTypeCard)
	setMATCHBYDefaultIfNil(matchByTypeKey)
	setMATCHBYDefaultIfNil(matchByTypeVal)
	if compareIdxs == nil {compareIdxs = false}

	// return whether conditions yield true
	return 	(compareCards == false || ((matchByTypeCard == MATCHBY_Object && thisCard == otherCard) || (matchByTypeCard == MATCHBY_Reference && &thisCard == &otherCard))) &&
			((matchByTypeKey == MATCHBY_Object && thisCard.Key == otherCard.Key) || (matchByTypeKey == MATCHBY_Reference && &thisCard.Key == &otherCard.Key)) &&
			((matchByTypeVal == MATCHBY_Object && thisCard.Val == otherCard.Val) || (matchByTypeVal == MATCHBY_Reference && &thisCard.Val == &otherCard.Val)) &&
			(compareIdxs == false || thisCard.Idx == otherCard.Idx)

}

/** Returns whether two stacks equal one another
 
 @receiver `thisStack` type{*Stack}
 @param `otherStack` type{*Stack}
 @param `compareStacks` type{bool} default false
	By default, does not compare the stack structs, but rather their cards; can be set true and adjusted with `matchByTypeStack`
 @param `matchByTypeStack` type{MATCHBY} default MATCHBY_Object
 @param `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param `compareCards` type{bool} default true
	By default, does not compare the card structs, but rather their individual values; can be set true and adjusted with `matchByTypeCard`
 @param `matchByTypeCard` type{MATCHBY} default MATCHBY_Object
 @param `matchByTypeKey` type{MATCHBY} default MATCHBY_Object
 @param `matchByTypeVal` type{MATCHBY} default MATCHBY_Object
 @param `compareIdxs` type{bool} default false
 @returns type{bool}
 */
func (thisStack *Stack) Equals(otherStack *Stack, variadic ...any) bool {

	// unpack variadic into optional parameters
	var compareStacks, matchByTypeStack, deepSearchType, compareCards, matchByTypeCard, matchByTypeKey, matchByTypeVal, compareIdxs any
	gogenerics.UnpackVariadic(variadic, &compareStacks, &matchByTypeStack, &deepSearchType, &compareCards, &matchByTypeCard, &matchByTypeKey, &matchByTypeVal, &compareIdxs)
	// set default vals
	if compareStacks == nil {compareStacks = true}
	setMATCHBYDefaultIfNil(matchByTypeStack)
	setDEEPSEARCHDefaultIfNil(deepSearchType)

	matches := true

	if compareStacks.(bool) {
		// just test whether the stacks equal one another
		matches = (matchByTypeStack == MATCHBY_Object && thisStack == otherStack) || (matchByTypeStack == MATCHBY_Reference && &thisStack == &otherStack)
	} else {
		// test whether the properties of the cards within each stack equal one another (recursively self-call)
		for i := range thisStack.Cards {
			thisCard := thisStack.Cards[i]
			otherCard := gogenerics.IfElse(i < len(otherStack.Cards), thisStack.Cards[i], nil).(*Card)
			
			matches = thisCard.Equals(otherCard, compareCards, matchByTypeCard, matchByTypeKey, matchByTypeVal, compareIdxs)
			
			if matches && deepSearchType == DEEPSEARCH_True {
				switch thisCard.Val.(type) { // go deeper if possible, otherwise don't worry
				case *Stack:
					switch otherCard.Val.(type) { // check whether otherCard can go deeper since thisCard can; if not, it's not an equal stack
					case *Stack:
						matches = thisCard.Val.(*Stack).Equals(otherCard.Val.(*Stack), compareStacks, matchByTypeStack, deepSearchType, compareCards, matchByTypeCard, matchByTypeKey, matchByTypeVal, compareIdxs)
					default:
						matches = false
					}
				}
			}
		}

	}

	return matches

}

/** Shuffles the order of `stack` cards

 @receiver `stack` type{*Stack}
 @returns `stack`
 @updates
  * `stack` card ordering is randomized
  * rand.Seed is updated to time.Now().UnixNano()
 */
func (stack *Stack) Shuffle() *Stack {

	// pseudo-randomize seed
	rand.Seed(time.Now().UnixNano())

	// shuffle
	rand.Shuffle(stack.Size, func(i, j int) { stack.Cards[i], stack.Cards[j] = stack.Cards[j], stack.Cards[i] })
	
	// set indices
	setIndices(stack.Cards)

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
 @updates terminal logs
 */
func (card *Card) Print() {

	fmt.Println("gostack: PRINTING CARD")
	fmt.Printf("- card.Idx: %v\n", card.Idx)
	fmt.Printf("- card.Key: %v\n", card.Key)
	fmt.Printf("- card.Val: %v\n", card.Val)

}

/** Prints information regarding `stack` to the console
 
 @receiver `stack` type{*Stack}
 @param optional `depth` type{int} default 0
  the starting depth of `stack` within other stacks (usually there is no reason for the client to pass this)
 @updates terminal logs
 @requires `card.Print()` has been implemented
 */
func (stack *Stack) Print(depth ...int) {

	fmt.Println("gostack: PRINTING STACK")
	if len(depth) == 0 {
		depth[0] = 0
	}
	for i := 0; i < depth[0]; i++ {
		fmt.Print("-")
	}
	fmt.Printf("- stack.Size: %v\n", stack.Size)
	for i := range stack.Cards {
		c := stack.Cards[i]
		c.Print()

		switch c.Val.(type) {
		case *Stack:
			c.Val.(*Stack).Print(depth[0]+1)
		}
	}

}

/** Iterate through a stack calling your lambda function on each card
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(*Card, *Stack, ...any)}
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @ensures
  * Each card in `stack` is passed into your lambda function
  * `stack` is the first argument passed into your variadic parameter on the first call
 */
func (stack *Stack) Lambda(lambda any, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &deepSearchType, &depth)
	
	// main
	generalIterator(stack, lambda.(func(*Card, *Stack, ...any)), deepSearchType.(DEEPSEARCH), depth.(int))

	return stack

}

/** Adds to a stack of cards or a cards at (each) position(s) and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `insert` type{Card, Stack}
 @param optional `orderType` type{ORDER} default ORDER_Before
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack` if cards were added OR nil if no cards were added (due to invalid find)
 @updates `stack` to have new cards before/after each designated position
 @requires `stack.Clone()` has been implemented
 */
func (stack *Stack) Add(insert any, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var orderType, findType, findData, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &orderType, &findType, &findData, &matchByType, &deepSearchType, &depth)

	// allow deepSearchHandler to handle function
	*stack = *stack.deepSearchHandler("Add", true, findType, findData, nil, matchByType, deepSearchType, depth, nil, nil, insert, orderType, nil, nil, nil, nil, nil, nil)

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
 @param optional `matchByType_from` type{MATCHBY} default MATCHBY_Object
 @param optional `matchByType_to` type{MATCHBY} default MATCHBY_Object
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
	var findData_from, findData_to, matchByType_from, matchByType_to, deepSearchType_from, deepSearchType_to, depth_from, depth_to any
	gogenerics.UnpackVariadic(variadic, &findData_from, &findData_to, &matchByType_from, &matchByType_to, &deepSearchType_from, &deepSearchType_to, &depth_from, &depth_to)

	// 1) Get the cards to move
	from := stack.ExtractMany(findType_from, findData_from, matchByType_from, RETURN_Cards, deepSearchType_from, depth_from)
	// 2) Get the card to put froms before/after depending on whether before or after
	var toCard *Card
	toStack := stack.GetMany(findType_to, findData_to, RETURN_Cards, matchByType_to, CLONE_False, CLONE_False, CLONE_False, deepSearchType_to, depth_to)
	if orderType == ORDER_After {
		toCard = toStack.Get(FIND_Last)
	} else if orderType == ORDER_Before {
		toCard = toStack.Get(FIND_First)
	}
	// 3) Insert 2 to 1 (works since to.Idx is procedurally updated by ExtractMany())
	stack.Add(from, orderType, FIND_Idx, toCard.Idx, matchByType_to, deepSearchType_to, depth_to)

	// return
	return stack

}

/** Swaps one element or slice with the position of another element or slice
 
 @receiver `stack` type{*Stack}
 @param `findType_first` type{FIND}
 @param `findType_second` type{FIND}
 @param optional `findData_first` type{any} default nil
 @param optional `findData_second` type{any} default nil
 @param optional `matchByType_first` type{MATCHBY} default MATCHBY_Object
 @param optional `matchByType_second` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType_first` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `deepSearchType_second` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth_first` type{int} default -1 (deepest)
 @param optional `depth_second` type{int} default -1 (deepest)
 @returns `stack` if moved OR nil if no move occurred (due second bad find)
 @requires you are not swapping a stack with a location within that own stack
 @ensures a stack of cards, or individual cards, can be targeted
 */
func (stack *Stack) Swap(findType_first FIND, findType_second FIND, variadic ...any) *Stack {

	// unpack variadic insecond optional parameters
	var findData_first, findData_second, matchByType_first, matchByType_second, deepSearchType_first, deepSearchType_second, depth_first, depth_second any
	gogenerics.UnpackVariadic(variadic, &findData_first, &findData_second, &matchByType_first, &matchByType_second, &deepSearchType_first, &deepSearchType_second, &depth_first, &depth_second)

	// 1) Get the first and second cards being targeted
	firsts := stack.GetMany(findType_second, findData_second, RETURN_Cards, matchByType_second, CLONE_False, CLONE_False, CLONE_False, deepSearchType_second, depth_second)
	seconds := stack.GetMany(findType_first, findData_first, RETURN_Cards, matchByType_first, CLONE_False, CLONE_False, CLONE_False, deepSearchType_first, depth_first)
	// 2) Determine which card is before the other in the stack.  If the second is before the first, switch the first variable and the second variable and all corresponding variables.
	if seconds.Cards[0].Idx < firsts.Cards[0].Idx {
		_tempFirst := firsts//lint:ignore SA4006 Ignore warning
		firsts = seconds
		seconds = _tempFirst//lint:ignore SA4006 Ignore warning

		_tempFindData_first := findData_first
		findData_first = findData_second
		findData_second = _tempFindData_first

		_tempMatchByType_first := matchByType_first
		matchByType_first = matchByType_second
		matchByType_second = _tempMatchByType_first

		_tempDeepSearchType_first := deepSearchType_first
		deepSearchType_first = deepSearchType_second
		deepSearchType_second = _tempDeepSearchType_first

		_tempDepth_first := depth_first
		depth_first = depth_second
		depth_second = _tempDepth_first
	}
	// 3) Now, in order second preserve the integrity of indices should the client choose second use FIND_Idx(s)...
	//    * Insert a copy of firsts after seconds,
	stack.Add(firsts, ORDER_After, findData_second, matchByType_second, deepSearchType_second, depth_second)
	//    * move second after first,
	stack.Move(findType_second, ORDER_After, findType_first, findData_first, findData_second, matchByType_first, matchByType_second, deepSearchType_first, deepSearchType_second, depth_first, depth_second)
	//    * and remove the original copy of first
	stack.Remove(findType_first, findData_first, matchByType_first, deepSearchType_first, depth_first)

	// return
	return stack

}

/** Returns a boolean representing whether a search exists in the stack

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns true IF successful search, false IF unsuccessful search
 @requires `stack.Get()` has been implemented
 */
func (stack *Stack) Has(variadic ...any) bool {

	// unpack variadic into optional parameters
	var findType, findData, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findType, &findData, &matchByType, &deepSearchType, &depth)

	// return
	return stack.Get(findType, findData, matchByType, nil, nil, nil, deepSearchType, depth) != nil
}

/** Gets a card from specified parameters in a stack, or nil if does not exist

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
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
	var findType, findData, matchByType, clonesType_card, clonesType_key, clonesType_val, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findType, &findData, &matchByType, &clonesType_card, &clonesType_key, &clonesType_val, &deepSearchType, &depth)

	// allow deepSearchHandler to take care of function
	return stack.deepSearchHandler("Get", true, findType, findData, nil, matchByType, deepSearchType, depth, nil, nil, nil, nil, nil, nil, nil, clonesType_card, clonesType_key, clonesType_val).Cards[0]

}

/** Gets a stack from specified parameters in a stack
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `clonesType` type{CLONES} default CLONE_False
 @param optional `clonesType_keys` type{CLONES} default CLONE_False
 @param optional `clonesType_vals` type{CLONES} default CLONE_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} the new stack (if find fails, then an empty stack)
 @constructs type{*Stack} new stack of specified values from specified cards in `stack`
 @requires
  * `MakeStack()` has been implemented
  * `clonesType_keys` and `clonesType_vals` are only passed if `returnType` == RETURN_Cards
 @ensures
  * CLONE_True means the cards in the returned stack are clones
  * CLONE_True for `clonesType_keys` means the cards in the returned stack keys are clones
  * CLONE_True for `clonesType_vals` means the cards in the returned stack vals are clones
 */
func (stack *Stack) GetMany(findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, returnType, matchByType, clonesType, clonesType_keys, clonesType_vals, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &returnType, &matchByType, &clonesType, &clonesType_keys, &clonesType_vals, &deepSearchType, &depth)

	// allow deepSearchHandler to take care of function
	return stack.deepSearchHandler("Get", false, findType, findData, returnType, matchByType, deepSearchType, depth, nil, nil, nil, nil, nil, nil, nil, clonesType, clonesType_keys, clonesType_vals)

}

/** Returns a found card before its respective field is updated to `replaceData` (OR nil if not found)
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Card} a clone of extracted card OR nil if found no cards
 @updates first found card to `replaceData`
 @requires `stack.Get()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 */
func (stack *Stack) Replace(replaceType REPLACE, replaceData any, findType FIND, variadic ...any) (ret *Card) {

	// unpack variadic into optional parameters
	var findData, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &matchByType, &deepSearchType, &depth)
	
	// get deep copy of targeted card OR nil
	ret = stack.Get(findType, findData, matchByType, CLONE_True, CLONE_True, CLONE_True, deepSearchType, depth)
	// get target data
	_, targetCards, targetStacks := stack.getPositions(true, findType, findData, matchByType.(MATCHBY), deepSearchType.(DEEPSEARCH), depth.(int))
	
	//stack.Get(findType, findData, matchByType, CLONE_False, CLONE_False, CLONE_False, deepSearchType, depth)

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
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} a stack whose values are the extracted cards pre-update (if find fails, then an empty stack)
 @updates all found cards to `replaceData`
 @requires `stack.GetMany()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the cards found will be removed from `stack`
 */
func (stack *Stack) ReplaceMany(replaceType REPLACE, replaceData any, findType FIND, variadic ...any) (ret *Stack) {

	// unpack variadic into optional parameters
	var findData, returnType, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &returnType, &matchByType, &deepSearchType, &depth)

	// get deep copy of targeted cards to return
	ret = stack.GetMany(findType, findData, returnType, matchByType, CLONE_True, CLONE_True, CLONE_True)
	// get target data
	_, targetCards, targetStacks := stack.getPositions(false, findType, findData, matchByType.(MATCHBY), deepSearchType.(DEEPSEARCH), depth.(int))

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
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates the found card in `stack`
 @requires `stack.Replace()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 */
func (stack *Stack) Update(replaceType REPLACE, replaceData any, findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &matchByType, &deepSearchType, &depth)

	// update stack
	stack.Replace(replaceType, replaceData, findType, findData, matchByType, deepSearchType, depth)

	// return the original stack
	return stack

}

/** Updates cards in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates  the found cards in `stack`
 @requires `stack.ReplaceMany()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the cards found will be removed from `stack`
 */
func (stack *Stack) UpdateMany(replaceType REPLACE, replaceData any, findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &matchByType, &deepSearchType, &depth)

	// update stack
	stack.ReplaceMany(replaceType, replaceData, findType, findData, matchByType, nil, deepSearchType, depth)

	// return the original stack
	return stack

}

/** Gets and removes a card from `stack`, or returns nil if does not exist
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Card} the extracted card OR nil (if invalid find)
 @updates `stack` to no longer have found card
 @requires `stack.Replace()` has been implemented
 */
func (stack *Stack) Extract(findType FIND, variadic ...any) *Card {

	// unpack variadic into optional parameters
	var findData, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &matchByType, &deepSearchType, &depth)

	// return the original value
	return stack.Replace(REPLACE_Card, nil, findType, findData, matchByType, deepSearchType, depth)

}

/** Gets and removes a set of data from `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} the extracted card (if find fails, then an empty stack)
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 */
func (stack *Stack) ExtractMany(findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, returnType, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &returnType, &matchByType, &deepSearchType, &depth)

	// return the original value
	return stack.ReplaceMany(REPLACE_Card, nil, findType, findData, returnType, matchByType, deepSearchType, depth)

}

/** Removes a card from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates `stack` to no longer have found card
 @requires `stack.Replace()` has been implemented
 */
func (stack *Stack) Remove(findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &matchByType, &deepSearchType, &depth)

	// remove the card
	stack.Replace(REPLACE_Card, nil, findType, findData, matchByType, deepSearchType, depth)

	// return stack
	return stack

}

/** Removes a set of cards from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns `stack`
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 */
func (stack *Stack) RemoveMany(findType FIND, variadic ...any) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType, deepSearchType, depth any
	gogenerics.UnpackVariadic(variadic, &findData, &matchByType, &deepSearchType, &depth)

	// remove the cards
	stack.ReplaceMany(REPLACE_Card, nil, findType, findData, matchByType, nil, deepSearchType, depth)

	// return stack
	return stack

}