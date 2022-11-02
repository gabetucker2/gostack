package gostack

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/gabetucker2/gogenerics"
)

/** Creates a card with given properties

 MakeCard(input1 any [nil], input2 any [nil], idx int [-1]) (*Card)
 
 @ensures
 | IF `input1` OR `input2` are nil:
 |     MakeCard := func(`val`, `key`, `idx`)
 | ELSE:
 |     MakeCard := func(`key`, `val`, `idx`)
 @examples
 | MakeCard("Hello") => Card{Val: "Hello"}
 | MakeCard(nil, "Hello") => Card{Key: "Hello"}
 | MakeCard(1, 2) => Card{Key: 1, Val: 2}
*/
 func MakeCard(arguments ...any) *Card {

	// unpack arguments into optional parameters
	var input1, input2, idx any
	gogenerics.UnpackVariadic(arguments, &input1, &input2, &idx)

	// initialize `key` and `val`
	var key, val any
	if input1 == nil || input2 == nil {
		key = input2
		val = input1
	} else {
		key = input1
		val = input2
	}

	// initialize and set new Card
	card := new(Card)
	if idx == nil { card.Idx = -1 } else { card.Idx = idx.(int) }
	card.Val = val
	card.Key = key

	// return
	return card

}

/** Creates a stack initialized with starting cards
 
 MakeStack(input1 []any|map[any]any|*Stack [nil], input2 []any|*Stack [nil], repeats int [1], overrideCards OVERRIDE [OVERRIDE_False]) (newStack *Stack)
 
 Where all mentions of array are interchangeable with Stack:
 @notes
 | Makes `repeats` repeats of `input1`/`input2`
 @requires
 | `input1` is a map and `input2` is nil
 |     OR `input1` is an array and `input2` is nil
 |     OR `input1` is an array and `input2` is an array
 |     OR `input1` is nil and `input2` is an array
 |
 | IF `input1` AND `input2` are both passed as arguments
 |      |`input1`| == |`input2`|
 @ensures
 |     IF `input1` is passed
 |       IF `input1` is a map
 |         unpack the map into new cards with corresponding keys and vals
 |       ELSEIF `input1` is an array and `input2` is not passed/nil
 |  	   IF `input1` is an array of cards:
 |           `overrideCards` == OVERRIDE_True:
 |               MakeStack([]*Card {cardA}) => stack.Cards = []*Card { card {Idx = 0, Key = nil, Val = cardA} }
 |           `overrideCards` == OVERRIDE_False:
 |               MakeStack([]*Card {cardA}) => stack.Cards = []*Card {cardA}
 |  	   ELSE:
 |           unpack values from `input1` into new cards
 |       ELSEIF `input1` is an array and `input2` is an array
 |         unpack keys from `input1` and values from `input2` into new cards
 |       ELSEIF `input1` is nil and `input2` is an array
 |         unpack keys from `input2` into new cards
 |  		make `repeats` cards with nil value and nil key
 |  		ELSEIF `input1` is nil and `input2` is nil and `repeats` is passed
 |     ELSE
 |       the stack is empty
 @examples
 | MakeStack([]int {1, 2, 3}) => Stack{Vals: {1, 2, 3}}
 | MakeStack(nil, []int {1, 2, 3}) => Stack{Keys: {1, 2, 3}}
 | MakeStack([]string {"a", "b", "c"}, []int {1, 2, 3}) => Stack{Keys: {"a", "b", "c"}, Vals: {1, 2, 3}}
 | MakeStack(map[string]int {"a":1, "b":2, "c":3}) => Stack{Keys: {"a", "b", "c"}, Vals: {1, 2, 3}} // but not necessarily in this order
 | MakeStack(nil, nil, 5) => Stack{nil, nil, nil, nil, nil}
 */
 func MakeStack(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var input1, input2, repeats, overrideCards any
	gogenerics.UnpackVariadic(arguments, &input1, &input2, &repeats, &overrideCards)
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

/** An identical implementation to `MakeStack()`

 MakeSubstack(input1 []any|map[any]any|*Stack [nil], input2 any|*Stack [nil], repeats int [1], overrideCards OVERRIDE [OVERRIDE_False]) (newSubstack *Stack)
 
 @examples
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}) => Stack{Stack{1, 2}, Stack{3, 4}}
 */
 func MakeSubstack(arguments ...any) *Stack {
	return MakeStack(arguments...)
}

/** Creates a stack matrix initialized with starting cards

 MakeStackMatrix(input1 []any (deep/shallow)|map[any]any (deep/shallow)|*Stack [nil], input2 []any (deep/shallow)|*Stack [nil], matrixShape []int [[]int {1}], overrideCards OVERRIDE [OVERRIDE_False]) (newStackMatrix *Stack)
 
 Where all mentions of array are interchangeable with Stack:
 @requires
 | `input1` is a map and `input2` is nil
 |     OR `input1` is an array and `input2` is nil
 |     OR `input1` is an array and `input2` is an array
 |     OR `input1` is nil and `input2` is an array
 |
 | IF `input1` AND `input2` are both passed as arguments:
 |      |`input1`| == |`input2`|
 |
 | `matrixShape` must be an int array representing the shape of a regularly-shaped matrix where:
 | * the first int defines `newStackMatrix.Size`
 | * the last int defines the size of each final stack
 | * the product of `matrixShape` is equal to the amount of elements in your input(s)
 @ensures
 | Using the same logic as MakeStack() in deciding which of the first two inputs is a key/val:
 |
 |  IF no `matrixShape` is passed:
 |    treating `input1`/`input2` as matrices ([]any {[]any {...}, []any {...}, ..., []any {...}})/a map of matrices (map[any]map[any]...map[any]any)/a StackMatrix:
 |    IF `input1` is passed:
 |      IF `input1` is a map:
 |        unpack the map into matrix of shape `inputx` with corresponding keys and vals
 |      ELSEIF `input1` is an array and `input2` is nil:
 |        unpack values from `input1` into matrix of shape `inputx`
 |      ELSEIF `input1` is an array and `input2` is an array:
 |        unpack keys from `input1` and values from `input2` into matrix of shape `inputx`
 |      ELSEIF `input1` is nil and `input2` is an array:
 |        unpack keys from `input2` into matrix of shape `inputx` 
 |    ELSEIF `input1` and `input2` are nil:
 |      the stack is empty
 |    ELSEIF `matrixShape` is passed:
 |      treating `input1`/`input2` as 1D structures ([]any, map[any]any, Stack):
 |      IF `input1` is a map:
 |        unpack the map into matrix of shape `matrixShape` with corresponding keys and vals
 |      ELSEIF `input1` is an array and `input2` is nil:
 |        IF `input1` is an array of cards:
 |          IF `overrideCards` == OVERRIDE_True:
 |            MakeStackMatrix([]*Card {cardA}) => stack.Cards = []*Card { card {Idx = 0, Key = nil, Val = cardA} }
 |          ELSEIF `overrideCards` == OVERRIDE_False:
 |            MakeStackMatrix([]*Card {cardA}) => stack.Cards = []*Card {cardA}
 |        ELSE:
 |           unpack values from `input1` into new cards
 |      ELSEIF `input1` is an array and `input2` is an array:
 |        unpack keys from `input1` and values from `input2` into matrix of shape `matrixShape`
 |      ELSEIF `input1` is nil and `input2` is an array:
 |        unpack keys from `input2` into matrix of shape `matrixShape`
 |      ELSEIF `input1` is nil AND `input2` is nil:
 |        create a StackMatrix of shape `matrixShape` whose heightest card keys/vals are nil
 @examples
 | MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2}) => Stack{Stack{1, 2}, Stack{3, 4}}
 | MakeStackMatrix([]int {1, 2, 3, 4, 5, 6}, nil, []int {2, 3}) => Stack{Stack{1, 2, 3}, Stack{4, 5, 6}}
 | MakeStackMatrix([]int {1, 2, 3, 4, 5, 6}, nil, []int {3, 2}) => Stack{Stack{1, 2}, Stack{3, 4}, Stack{5, 6}}
 | MakeStackMatrix([]any {[]any {1, 2}, []any {3, 4}}} =>  Stack{Stack{1, 2}, Stack{3, 4}}
 */
 func MakeStackMatrix(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var input1, input2, matrixShape, overrideCards any
	gogenerics.UnpackVariadic(arguments, &input1, &input2, &matrixShape, &overrideCards)
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
			// create a StackMatrix of shape `matrixShape` whose heightest card vals are nil
			stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, nil, new(int), OVERRIDE_False)

		}

	}

	// set properties
	stack.setStackProperties()

	// return
	return stack
	
}

/** Updates a stack to represent a selection within that stack matrix
 
 stack.StripStackMatrix(idx ...int|[]int|*Stack [[]int {0, 1, ..., stack.Size - 1}]) (stack)

 @requires
 | `idx` refers to valid index positions from the stack
 @examples
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).StripStackMatrix() => Stack {1, 2, 3, 4}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).StripStackMatrix([]int {0, 1}) => Stack {1, 2, 3, 4}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).StripStackMatrix(0) => Stack {1, 2}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).StripStackMatrix(1) => Stack {3, 4}
 */
func (stack *Stack) StripStackMatrix(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var idx any
	gogenerics.UnpackVariadic(arguments, &idx)

	// main
	if idx == nil {
		return stack.Filter(FIND_All, nil, nil, RETURN_Stacks)
	} else {
		idxInt, isInt := idx.(int)
		if isInt {
			return stack.Filter(FIND_Idx, idxInt, nil, RETURN_Stacks)
		} else {
			return stack.Filter(FIND_Slice, idx, nil, RETURN_Stacks)
		}
	}

}

/** Creates a new any array whose elements are the values of the cards in `stack`
 
 stack.ToArray(returnType RETURN [RETURN_Vals]) (newArray []any)

 @examples
 | MakeStack([]int {1, 2, 3}, []string {"a", "b", "c"}).ToArray() => []any {1, 2, 3}
 | MakeStack([]int {1, 2, 3}, []string {"a", "b", "c"}).ToArray(RETURN_Keys) => []any {"a", "b", "c"}
 | MakeStack([]int {1, 2, 3}, []string {"a", "b", "c"}).ToArray(RETURN_Idxs) => []any {0, 1, 2}
 | MakeStack([]*Card {cardA, cardB, cardC}).ToArray(RETURN_Cards) => []any {cardA, cardB, cardC}
 | MakeStack([]*Stack {substackA, substackB}).ToArray(RETURN_Cards) => []any {Card{Val:substackA}, Card{Val:substackA}}
 | MakeStack([]*Stack {substackA, substackB}).ToArray(RETURN_Stacks) => []any {substackA, substackB}
 */
func (stack *Stack) ToArray(arguments ...any) []any {

	// unpack arguments into optional parameters
	var returnType any
	gogenerics.UnpackVariadic(arguments, &returnType)

	// return
	return stack.ToMatrix(returnType, 1)

}

/** Creates a new map whose keys and values correspond to the cards in `stack`

 stack.ToMap() (newMap map[any]any)

 @examples
 | MakeStack([]int {1, 2, 3}, []string {"a", "b", "c"}).ToMap() => map[any]any {1:"a", 2:"b", 3:"c"} // in any order
 */
func (stack *Stack) ToMap() map[any]any {

	// add all card keys and values in stack to m
	newMap := make(map[any]any)
	for i := range stack.Cards {
		c := stack.Cards[i]
		newMap[c.Key] = c.Val
	}

	// return
	return newMap

}

/** Creates a new matrix structure from `stack`

 stack.ToMatrix(returnType RETURN [RETURN_Vals], depth int [-1]) (newMatrix []any {elem/[]any{}})

 @examples
 | MakeStack([]int {1, 2, 3, 4}).ToMatrix() => []any {1, 2, 3, 4}
 | MakeStack(*Stack{MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).ToMatrix() => []any {[]any {1, 2}, []any {3, 4}}
 */
func (stack *Stack) ToMatrix(arguments ...any) []any {

	// unpack arguments into optional parameters
	var returnType, depth any
	gogenerics.UnpackVariadic(arguments, &returnType, &depth)
	// set defaults
	setRETURNDefaultIfNil(&returnType)
	setHeightDefaultIfNil(&depth)
	if depth == -1 || depth.(int) > stack.Height { depth = stack.Height }
	newMatrix := []any {}

	// break recursion at depth == 0
	if depth.(int) != 0 {
		// add to return
		for i := range stack.Cards {
			c := stack.Cards[i]
			// if this Card's val is a Stack
			subStack, isStack := c.Val.(*Stack)
			
			if isStack {
				if depth.(int) > 1 {
					newMatrix = append(newMatrix, subStack.ToMatrix(returnType, depth.(int) - 1))
				} else {
					newMatrix = append(newMatrix, []any {})
				}
			} else {
				switch returnType {
				case RETURN_Vals:
					newMatrix = append(newMatrix, c.Val)
				case RETURN_Keys:
					newMatrix = append(newMatrix, c.Key)
				case RETURN_Idxs:
					newMatrix = append(newMatrix, c.Idx)
				case RETURN_Cards:
					newMatrix = append(newMatrix, c)
				}
			}
		}
	}
	
	// return
	return newMatrix

}

/** Returns an array representing the shape of `stack`

 stack.Shape() (stackShape []int)

 @ensures
 | returns nil if it's not regular and thus doesn't have a shape
 @examples
 | MakeStack([]*Stack{MakeSubstack([]int {1, 2, 3}), MakeSubstack([]int {4, 5, 6})}).Shape() => []int {2, 3}
 | MakeStack([]*Stack{MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4}), MakeSubstack([]int {5, 6})}).Shape() => []int {3, 2}
 | MakeStack([]*Stack{MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4, 5}), MakeSubstack([]int {6, 7})}) => nil
 */
 func (stack *Stack) Shape() []int {

	// init
	stackShape := []int {}

	// body
	if stack.IsRegular() {

		stackShape = append(stackShape, stack.Size)

		if stack.Size > 0 {
			_, hasSubstack := stack.Cards[0].Val.(*Stack)
			if hasSubstack {
				stackShape = append(stackShape, stack.Cards[0].Val.(*Stack).Shape()...)
			}
		}

	} else {
		stackShape = nil
	}

	// return
	return stackShape

}

/** Returns whether the matrix is of a regular shape

 stack.IsRegular() (stackIsRegular bool)

 @examples
 | MakeStack([]*Stack{MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4}), MakeSubstack([]int {5, 6})}) => true
 | MakeStack([]*Stack{MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4, 5}), MakeSubstack([]int {6, 7})}) => false
 */
 func (stack *Stack) IsRegular() bool {

	// init
	test := true
	normHeight := -1
	normSize := -1
	normSubstack := -1

	// body
	for _, c := range stack.Cards {
		substack, hasSubstack := c.Val.(*Stack)

		if hasSubstack {
			if normHeight == -1 {
				normHeight = substack.Height
			} else if normHeight != substack.Height {
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
 func (stack *Stack) Duplicate(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var n any
	gogenerics.UnpackVariadic(arguments, &n)
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
	stack.Height = 1
	stack.Cards = []*Card{} // avoid replacing stack object

	// return
	return stack

}

/** Returns a clone of `card`
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
 @param optional `depth` type{int} default -1 (heightest)
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
   * if you shallow clone a depth stack and cloneSubstackVals is true, then the original substacks will be held in the clone
*/
func (stack *Stack) Clone(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var deepSearchType, depth, cloneCardKeys, cloneCardVals, cloneSubstackKeys, cloneSubstackVals any
	gogenerics.UnpackVariadic(arguments, &deepSearchType, &depth, &cloneCardKeys, &cloneCardVals, &cloneSubstackKeys, &cloneSubstackVals)
	// set defaults
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setHeightDefaultIfNil(&depth)
	setCLONEDefaultIfNil(&cloneCardKeys)
	setCLONEDefaultIfNil(&cloneCardVals)
	setCLONEDefaultIfNil(&cloneSubstackKeys)
	setCLONEDefaultIfNil(&cloneSubstackVals)
	if depth == -1 || depth.(int) > stack.Height { depth = stack.Height }
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

				// forwardpropagate if depth > 1 and we're heightsearching
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
func (stack *Stack) Unique(arguments ...any) *Stack {
	
	// unpack arguments into optional parameters
	var uniqueType any
	gogenerics.UnpackVariadic(arguments, &uniqueType)
	if uniqueType == nil { uniqueType = TYPE_Val }

	// main
	return stack.GetMany(FIND_Lambda, func(card *Card, _ *Stack, _ bool, workingStack *Stack, wm ...any) (bool) {
		if workingStack.Size == 0 {
			return true
		} else {
			switch uniqueType.(TYPE) {
			case TYPE_Key:
				return !workingStack.Has(FIND_Key, card.Key)
			case TYPE_Val:
				return !workingStack.Has(FIND_Val, card.Val)
			}
			return false // just so it compiles
		}
	})

}

/** Returns whether two cards equal one another
 
 @receiver `thisCard` type{*Card}
 @param `otherCard` type{*Card}
 @param optional `dereferenceTypeKey` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `dereferenceTypeVal` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `compareIdxs` type{COMPARE} default COMPARE_False
 @param optional `compareKeys` type{COMPARE} default COMPARE_True
 @param optional `compareVals` type{COMPARE} default COMPARE_True
 @param optional `compareObjectAdr` type{COMPARE} default COMPARE_False
   if true, ensures the cards are the same object
 @returns type{bool}
 */
func (thisCard *Card) Equals(otherCard *Card, arguments ...any) bool {

	// unpack arguments into optional parameters
	var dereferenceTypeKey, dereferenceTypeVal, compareIdxs, compareKeys, compareVals, compareObjectAdr any
	gogenerics.UnpackVariadic(arguments, &dereferenceTypeKey, &dereferenceTypeVal, &compareIdxs, &compareKeys, &compareVals, &compareObjectAdr)
	// set default vals
	setDEREFERENCEDefaultIfNil(&dereferenceTypeKey)
	setDEREFERENCEDefaultIfNil(&dereferenceTypeVal)
	if compareIdxs == nil {compareIdxs = COMPARE_False}
	setCOMPAREDefaultIfNil(&compareKeys)
	setCOMPAREDefaultIfNil(&compareVals)
	if compareObjectAdr == nil {compareObjectAdr = COMPARE_False}

	condition := thisCard != nil && otherCard != nil
	
	condition = condition && 
		(compareKeys == COMPARE_False ||
		(compareKeys == COMPARE_True &&
			(
				(dereferenceTypeKey == DEREFERENCE_None && thisCard.Key == otherCard.Key) ||
				(dereferenceTypeKey == DEREFERENCE_Both && gogenerics.PointersEqual(thisCard.Key, otherCard.Key) ) ) ) )
	
	condition = condition && 
		(compareVals == COMPARE_False ||
		(compareVals == COMPARE_True &&
			(
				(dereferenceTypeVal == DEREFERENCE_None && thisCard.Val == otherCard.Val) ||
				(dereferenceTypeVal == DEREFERENCE_Both && gogenerics.PointersEqual(thisCard.Val, otherCard.Val) ) ) ) )

	condition = condition && (compareIdxs == COMPARE_False || (compareIdxs == COMPARE_True && thisCard.Idx == otherCard.Idx))

	condition = condition && (compareObjectAdr == COMPARE_False || (compareObjectAdr == COMPARE_True && fmt.Sprintf("%p", thisCard) == fmt.Sprintf("%p", otherCard)))
	
	// return whether conditions yield true
	return condition

}

/** Returns whether two stacks equal one another
 
 @receiver `thisStack` type{*Stack}
 @param `otherStack` type{*Stack}
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `compareCardKeys` type{COMPARE} default COMPARE_True
 @param optional `compareCardVals` type{COMPARE} default COMPARE_True
 @param optional `compareSubstackKeys` type{SUBSTACKKEYS} default COMPARE_True
 @param optional `compareSubstackVals` type{SUBSTACKKEYS} default COMPARE_False
   this being set to true will compare the substacks themselves
 @param optional `pointerCardKeys` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `pointerCardVals` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `pointerSubstackKeys` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `pointerSubstackVals` type{DEREFERENCE} default DEREFERENCE_Both
   this being set to true will compare the addresses of the substacks
 @ensures
   * `stack`.Size == `otherStack`.Size is tested on the heightest layer
   * if `stack`.Height != `otherStack`.Height and the N-depth comparison finds that they're equal, then return that they're Equal 
 @returns type{bool}
 */
func (stack *Stack) Equals(otherStack *Stack, arguments ...any) (test bool) {

	/*
	PSEUDOCODE OUTLINE:

	stack.Equals(otherStack, ...arguments) bool

		set up arguments stuff

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
			if depth is an int && (depth == -1 || depth > stack.Height)
				depth = stack.Height
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

	// unpack arguments into optional parameters
	var deepSearchType, depth, compareCardKeys, compareCardVals, compareSubstackKeys, compareSubstackVals, pointerCardKeys, pointerCardVals, pointerSubstackKeys, pointerSubstackVals any
	gogenerics.UnpackVariadic(arguments, &deepSearchType, &depth, &compareCardKeys, &compareCardVals, &compareSubstackKeys, &compareSubstackVals, &pointerCardKeys, &pointerCardVals, &pointerSubstackKeys, &pointerSubstackVals)
	// set default vals
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setHeightDefaultIfNil(&depth)

	setCOMPAREDefaultIfNil(&compareCardKeys)
	setCOMPAREDefaultIfNil(&compareCardVals)
	setCOMPAREDefaultIfNil(&compareSubstackKeys)
	if compareSubstackVals == nil { compareSubstackVals = COMPARE_False }
	
	setDEREFERENCEDefaultIfNil(&pointerCardKeys)
	setDEREFERENCEDefaultIfNil(&pointerCardVals)
	setDEREFERENCEDefaultIfNil(&pointerSubstackKeys)
	if pointerSubstackVals == nil { pointerSubstackVals = DEREFERENCE_Both }

	heightStack, heightIsStack := depth.(*Stack)
	if heightIsStack {
		depth = []int {}
		for _, d := range heightStack.Cards {
			depth = append(depth.([]int), d.Val.(int))
		}
	}
	testLayer := true
	_, heightIsInt := depth.(int)
	if deepSearchType == DEEPSEARCH_False {
		if depth == -1 {// first input
			depth = 1
		} else {
			depth = 0
			testLayer = false
		}
	} else {
		if heightIsInt && (depth == -1 || depth.(int) > stack.Height) {
			depth = stack.Height
		} else if !heightIsInt { // depth is an []int
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
							if pointerSubstackKeys == DEREFERENCE_Both {
								test = test && gogenerics.PointersEqual(cardA.Key, cardB.Key)
							} else {
								test = test && cardA.Key == cardB.Key
							}
						}
						if compareSubstackVals == COMPARE_True {
							if pointerSubstackVals == DEREFERENCE_Both {
								test = test && gogenerics.PointersEqual(cardA.Val, cardB.Val)
							} else {
								test = test && cardA.Val == cardB.Val
							}
						}
					}
					
					// forwardpropagate
					var transformedHeight any
					if heightIsInt {
						transformedHeight = depth.(int) - 1
					} else {
						transformedHeight = []int {}
						for i := range depth.([]int) {
							transformedHeight = append(transformedHeight.([]int), depth.([]int)[i] - 1)
						}
					}
					test = test && substackA.Equals(substackB, deepSearchType, transformedHeight, compareCardKeys, compareCardVals, compareSubstackKeys, compareSubstackVals, pointerCardKeys, pointerCardVals, pointerSubstackKeys, pointerSubstackVals)

				} else if (cardAIsSubstack && !cardBIsSubstack) || (!cardAIsSubstack && cardBIsSubstack) { // one holds a substack and the other doesnt

					test = false

				} else { // neither are substacks and they are both just cards

					// compare card properties
					if testLayer {
						test = test && cardA.Equals(cardB, pointerCardKeys, pointerCardVals, COMPARE_True, compareCardKeys, compareCardVals)
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
func (stack *Stack) Shuffle(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var newOrder any
	gogenerics.UnpackVariadic(arguments, &newOrder)
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

/** Reverses the order of all cards and stacks in `stack`
 */
func (stack *Stack) Transpose() *Stack {

	flipper := func(thisStack *Stack) {
		newCardsArr := []*Card {}
		for i := range thisStack.Cards {
			newCardsArr = append(newCardsArr, thisStack.Cards[thisStack.Size - 1 - i])
		}
		thisStack.Cards = newCardsArr
		thisStack.setStackProperties()
	}

	// flip the root stack
	flipper(stack)

	// main
	return stack.LambdaThis(func(card *Card, _ *Stack, _ bool, _ *Stack, _ *Card, _ any, _ []any, _ ...any) {

		flipper(card.Val.(*Stack))

	}, nil, nil, nil, nil, nil, nil, PASS_True, PASS_False)

}

/** Switches the Key and the Val of `card`
 */
 func (card *Card) SwitchKeyVal() *Card {

	// main
	tempVal := card.Val
	card.Val = card.Key
	card.Key = tempVal

	// return
	return card

 }

 /** Switches the Key and the Val of each found card

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any, []any, *Stack any, func(*Card, *Stack, bool, *Stack, ...any) (bool)} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
 @constructs a new stack
 @returns type{*Stack} the new stack
 @requires
   IF `find` is FIND_Lambda, `findData` is of type{ func(card *Card, parentStack *Stack, isSubstack bool, workingStack *Stack, workingMem ...any) (bool) }
  */
  func (stack *Stack) SwitchKeysVals(arguments ...any) *Stack {

	// set up arguments
	var findType, findData, findCompareRaw, deepSearchType, depth, dereferenceType, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &findCompareRaw, &deepSearchType, &depth, &dereferenceType, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
 
	 // main
	// get card
	stack.Lambda(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any, wmadrs ...any) {
		
		if selectCard(findType, findData, dereferenceType, findCompareRaw.(COMPARE), card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs...) && retCard.Idx == -1 {
			
			card.SwitchKeyVal()

		}

	}, nil, nil, nil, workingMem.([]any), deepSearchType, depth, PASS_False, PASS_True)
 
	 // return
	 return stack
 
  }

/** Prints information surrounding `card` to the terminal
 
 card.Print(indent int [0]) (card)

 @ensures
 | prints "-" `indent` * 4 times before each line to indicate depth in a stackMatrix
 */
func (card *Card) Print(arguments ...any) *Card {

	// unpack arguments into optional parameters
	var indents any
	gogenerics.UnpackVariadic(arguments, &indents)
	if indents == nil { indents = 0 }

	// prints
	fmt.Printf("%v|%vCARD\n", depthPrinter(indents.(int)), gogenerics.IfElse(indents == 0, "gostack: PRINTING ", ""))
	if card == nil {
		fmt.Printf("%v- card:          %v\n", depthPrinter(indents.(int)), nil)
	} else {
		fmt.Printf("%v- &card:         %v\n", depthPrinter(indents.(int)), fmt.Sprintf("%p", card))
		fmt.Printf("%v- card.Idx:      %v\n", depthPrinter(indents.(int)), card.Idx)
		if gogenerics.IsPointer(card.Key) {
			fmt.Printf("%v- &card.Key:     %v\n", depthPrinter(indents.(int)), fmt.Sprintf("%p", card.Key))
			fmt.Printf("%v- card.Key:      %v\n", depthPrinter(indents.(int)), reflect.ValueOf(card.Key).Elem())
			fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(indents.(int)), reflect.TypeOf(reflect.ValueOf(card.Key).Elem().Interface()))
		} else {
			fmt.Printf("%v- card.Key:      %v\n", depthPrinter(indents.(int)), card.Key)
			if card.Key != nil {
				fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(indents.(int)), reflect.TypeOf(card.Key))
			}
		}
		if gogenerics.IsPointer(card.Val) {
			fmt.Printf("%v- &card.Val:     %v\n", depthPrinter(indents.(int)), fmt.Sprintf("%p", card.Val))
			fmt.Printf("%v- card.Val:      %v\n", depthPrinter(indents.(int)), reflect.ValueOf(card.Val).Elem())
			fmt.Printf("%v- card.Val.Type: (%v)\n", depthPrinter(indents.(int)), reflect.TypeOf(reflect.ValueOf(card.Val).Elem().Interface()))
		} else {
			fmt.Printf("%v- card.Val:      %v\n", depthPrinter(indents.(int)), card.Val)
			if card.Val != nil {
				fmt.Printf("%v- card.Val.Type: (%v)\n", depthPrinter(indents.(int)), reflect.TypeOf(card.Val))
			}
		}
	}

	return card

}

/** Prints information surrounding `stack` to the terminal

 stack.Print(indent int [0]) (stack)
 
 @ensures
 | prints "-" `indent` * 4 times before each line to indicate depth in a stackMatrix
 */
func (stack *Stack) Print(arguments ...any) *Stack {
	
	// unpack arguments into optional parameters
	var indent, idx, key any
	gogenerics.UnpackVariadic(arguments, &indent, &idx, &key)
	if indent == nil { indent = 0 }

	fmt.Printf("%v|%vSTACK\n", depthPrinter(indent.(int)), gogenerics.IfElse(idx == nil, "gostack: PRINTING ", "SUB"))
	if stack == nil {
		fmt.Printf("%v- stack:         %v\n", depthPrinter(indent.(int)), nil)
	} else {
		fmt.Printf("%v- &stack:        %v\n", depthPrinter(indent.(int)), fmt.Sprintf("%p", stack))
		if idx != nil {
			fmt.Printf("%v- card.Idx:      %v\n", depthPrinter(indent.(int)), idx)
		}
		if key != nil {
			if gogenerics.IsPointer(key) {
				fmt.Printf("%v- &card.Key:     %v\n", depthPrinter(indent.(int)), key)
				fmt.Printf("%v- card.Key:      %v\n", depthPrinter(indent.(int)), reflect.ValueOf(key).Elem())
				fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(indent.(int)), reflect.TypeOf(reflect.ValueOf(key).Elem().Interface()))
			} else {
				fmt.Printf("%v- card.Key:      %v\n", depthPrinter(indent.(int)), key)
				if key != nil {
					fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(indent.(int)), reflect.TypeOf(key))
				}
			}
		}
		fmt.Printf("%v- stack.Size:    %v\n", depthPrinter(indent.(int)), stack.Size)
		fmt.Printf("%v- stack.Height:   %v\n", depthPrinter(indent.(int)), stack.Height)
		for i := range stack.Cards {
			c := stack.Cards[i]
	
			switch c.Val.(type) {
			case *Stack:
				c.Val.(*Stack).Print(indent.(int)+4, i, c.Key)
			default:
				c.Print(indent.(int)+4)
			}
		}
	}

	return stack
	
}

/** Iterate through a stack calling your lambda function on each card
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any {cardAdr, parentStackAdr, retStackAdr, retCardAdr}, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `otherInfo` type{[]any {retStackAdr, retCardAdr}} default []any {nil, nil}
   in order to reference `retStack` and `retCard` to `card`, `parentStack`, or some other struct from within your custom lambda function, the addresses of these pointers must be passed into `otherInfo`
 @returns `stack` type{*Stack}
 @returns `retStack` type{*Stack}
 @returns `retCard` type{*Card}
 @returns `retVarAdr` type{any}
 */
func (stack *Stack) Lambda(lambda func(*Card, *Stack, bool, *Stack, *Card, any, []any, ...any), arguments ...any) (*Stack, *Stack, *Card, any) {

	/**
	PSEUDOCODE OUTLINE:

	(stack) Lambda(lambda, ...arguments) (retAdr any)

		set up arguments stuff

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
			if depth is an int && (depth == -1 || depth > stack.Height)
				depth = stack.Height
			else if depth is an []int
				if []depth does not have an element == 1
					passLayer = false
			if depth == 0
				passLayer = false
		
		for each card in this stack
			
			if card is substack

				if passSubstacks and passLayer
					pass substack into lambda(isSubstack = true)
					update pointers to reflect possible pointer pointer reassignments
					// https://stackoverflow.com/questions/74090485/why-is-my-interface-containing-a-pointer-not-updating-after-the-pointer-is-updat/74090525#74090525

				if depth > 1 or depth[] has an element > 1 // forwardpropagate
					substack.Lambda(..., depth = depth - 1 OR depth[] = depth[i - 1, ..., n - 1])
					update pointers to reflect possible pointer pointer reassignments

			else if card is not substack

				if passCards and passLayer
					pass card into lambda(isSubstack = false)
					update pointers to reflect possible pointer pointer reassignments

		return *retAdr.(*any)

	*/

	// unpack arguments into optional parameters
	var retStack, retCard, retVarAdr, workingMem, deepSearchType, depth, passSubstacks, passCards, otherInfo any
	gogenerics.UnpackVariadic(arguments, &retStack, &retCard, &retVarAdr, &workingMem, &deepSearchType, &depth, &passSubstacks, &passCards, &otherInfo)
	if retVarAdr == nil {var o any; retVarAdr = &o;}
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setHeightDefaultIfNil(&depth)
	if passSubstacks == nil {passSubstacks = PASS_False}
	if passCards == nil {passCards = PASS_True}
	if otherInfo == nil {otherInfo = []any {nil, nil}}
	retStackAdr := otherInfo.([]any)[0]
	retCardAdr := otherInfo.([]any)[1]
	if retStack == nil {
		newStack := MakeStack()
		retStack = newStack
		retStackAdr = &newStack
		otherInfo.([]any)[0] = retStackAdr
	}
	if retCard == nil {
		newCard := MakeCard()
		retCard = newCard
		retCardAdr = &newCard
		otherInfo.([]any)[1] = retCardAdr
	}
	
	// main
	heightStack, heightIsStack := depth.(*Stack)
	if heightIsStack {
		depth = []int {}
		for _, d := range heightStack.Cards {
			depth = append(depth.([]int), d.Val.(int))
		}
	}
	passLayer := true
	has1 := false
	hasOver1 := false
	_, heightIsInt := depth.(int)
	if deepSearchType == DEEPSEARCH_False {
		if depth == -1 {// first input
			depth = 1
		} else {
			depth = 0
			passLayer = false
		}
	} else {
		if heightIsInt && (depth == -1 || depth.(int) > stack.Height) {
			depth = stack.Height
		} else if !heightIsInt { // depth is an []int
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

				lambda(card, stack, true, toTypeStack(retStack), toTypeCard(retCard), &retVarAdr, []any {&card, &stack, retStackAdr, retCardAdr}, workingMem.([]any)...)
				if retCardAdr != nil && *retCardAdr.(**Card) != nil {
					retCard = *retCardAdr.(**Card)
				}
				if retStackAdr != nil && *retStackAdr.(**Stack) != nil {
					retStack = *retStackAdr.(**Stack)
				}

				// update properties
				stack.setStackProperties()
				if retStack != nil {
					retStack.(*Stack).setStackProperties()
				}
			}

			// forwardpropagate
			if (heightIsInt && depth.(int) > 1) || hasOver1 {
				var transformedHeight any
				if heightIsInt {
					transformedHeight = depth.(int) - 1
				} else {
					transformedHeight = []int {}
					for i := range depth.([]int) {
						transformedHeight = append(transformedHeight.([]int), depth.([]int)[i] - 1)
					}
				}

				substack.Lambda(lambda, retStack, retCard, retVarAdr, workingMem, deepSearchType, transformedHeight, passSubstacks, passCards, otherInfo)
				if retCardAdr != nil && *retCardAdr.(**Card) != nil {
					retCard = *retCardAdr.(**Card)
				}
				if retStackAdr != nil && *retStackAdr.(**Stack) != nil {
					retStack = *retStackAdr.(**Stack)
				}
			}

		} else { // card is not substack

			if passCards == PASS_True && passLayer {

				lambda(card, stack, false, toTypeStack(retStack), toTypeCard(retCard), &retVarAdr, []any {&card, &stack, retStackAdr, retCardAdr}, workingMem.([]any)...)
				if retCardAdr != nil && *retCardAdr.(**Card) != nil {
					retCard = *retCardAdr.(**Card)
				}
				if retStackAdr != nil && *retStackAdr.(**Stack) != nil {
					retStack = *retStackAdr.(**Stack)
				}

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
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any {cardAdr, parentStackAdr, retStackAdr, retCardAdr}, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `otherInfo` type{[]any {retStackAdr, retCardAdr}} default []any {nil, nil}
 @returns `stack` type{*Stack}
 */
func (stack *Stack) LambdaThis(lambda func(*Card, *Stack, bool, *Stack, *Card, any, []any, ...any), arguments ...any) *Stack {
	stack.Lambda(lambda, arguments...)
	return stack
}

/** Iterate through a stack calling your lambda function on each card, returning only `retStack`
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any {cardAdr, parentStackAdr, retStackAdr, retCardAdr}, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `otherInfo` type{[]any {retStackAdr, retCardAdr}} default []any {nil, nil}
 @returns `retStack` type{*Stack}
 */
func (stack *Stack) LambdaStack(lambda func(*Card, *Stack, bool, *Stack, *Card, any, []any, ...any), arguments ...any) *Stack {
	_, thisStack, _, _ := stack.Lambda(lambda, arguments...)
	return thisStack
}

/** Iterate through a stack calling your lambda function on each card, returning only `retCard`
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any {cardAdr, parentStackAdr, retStackAdr, retCardAdr}, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `otherInfo` type{[]any {retStackAdr, retCardAdr}} default []any {nil, nil}
 @returns `retCard` type{*Card}
 */
func (stack *Stack) LambdaCard(lambda func(*Card, *Stack, bool, *Stack, *Card, any, []any, ...any), arguments ...any) *Card {
	_, _, thisCard, _ := stack.Lambda(lambda, arguments...)
	return thisCard
}

/** Iterate through a stack calling your lambda function on each card, returning only `retVarAdr`
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any {cardAdr, parentStackAdr, retStackAdr, retCardAdr}, workingMem ...any)}
 @param optional `retStack` type{*Stack} default nil
 @param optional `retCard` type{*Card} default nil
 @param optional `retVarAdr` type{any} default nil
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_True
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `passSubstacks` type{PASS} default PASS_False
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `otherInfo` type{[]any {retStackAdr, retCardAdr}} default []any {nil, nil}
 @returns `retVarAdr` type{any}
 */
func (stack *Stack) LambdaVarAdr(lambda func(*Card, *Stack, bool, *Stack, *Card, any, []any, ...any), arguments ...any) any {
	_, _, _, retVarAdr := stack.Lambda(lambda, arguments...)
	return retVarAdr
}

/** The exact same as GetMany, except it replaces `stack` with the new stack */
func (stack *Stack) Filter(findType FIND, arguments ...any) *Stack {
	*stack = *stack.GetMany(findType, arguments...)
	return stack
}

/** Adds to a stack of cards or a cards the first time it is able and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `insert` type{any, []any, Stack}
 @param optional `orderType` type{ORDER} default ORDER_After
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any, []any, *Stack any, func(*Card, *Stack, bool, ...any) (bool)} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `overrideCards` type{OVERRIDE} default OVERRIDE_False
   By default, if you do stack.Add(cardA), stack = {cardA}.  If you instead desire stack = {Card {val = cardA}}, do true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack` if valid find, or nil if invalid find
 */
func (stack *Stack) Add(insert any, arguments ...any) *Stack {
	
	return stack.addHandler(false, insert, arguments...)

}

/** Adds to a stack of cards or a cards at each position and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `insert` type{any, []any, Stack}
 @param optional `orderType` type{ORDER} default ORDER_After
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any, []any, *Stack any, func(*Card, *Stack, bool, ...any) (bool)} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `overrideCards` type{OVERRIDE} default OVERRIDE_False
   By default, if you do stack.Add(cardA), stack = {cardA}.  If you instead desire stack = {Card {val = cardA}}, do true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack` if valid find, or nil if invalid find
 */
func (stack *Stack) AddMany(insert any, arguments ...any) *Stack {
	
	return stack.addHandler(true, insert, arguments...)

}

/** Moves one card to before or after another card
 
 @receiver `stack` type{*Stack}
 @param optional `orderType` type{ORDER} default ORDER_After
 @param optional `findTypeFrom` type{FIND} default FIND_First
 @param optional `findTypeTo` type{FIND} default FIND_Last
 @param optional `findDataFrom` type{any, []any, *Stack any, func(*Card, *Stack, bool, ...any) (bool)} default nil
 @param optional `findDataTo` type{any, []any, *Stack any, func(*Card, *Stack, bool, ...any) (bool)} default nil
 @param optional `findCompareRawFrom` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `findCompareRawTo` type{COMPARE} default COMPARE_False
 @param optional `deepSearchTypeFrom` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `deepSearchTypeTo` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `heightFrom` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `heightTo` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceTypeFrom` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `dereferenceTypeTo` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacksFrom` type{PASS} default PASS_True
 @param optional `passSubstacksTo` type{PASS} default PASS_True
 @param optional `passCardsFrom` type{PASS} default PASS_True
 @param optional `passCardsTo` type{PASS} default PASS_True
 @param optional `workingMemFrom` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
 @param optional `workingMemTo` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack`
 */
func (stack *Stack) Move(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var orderType, findTypeFrom, findTypeTo, findDataFrom, findDataTo, findCompareRawFrom, findCompareRawTo, deepSearchTypeFrom, deepSearchTypeTo, heightFrom, heightTo, dereferenceTypeFrom, dereferenceTypeTo, passSubstacksFrom, passSubstacksTo, passCardsFrom, passCardsTo, workingMemFrom, workingMemTo any
	gogenerics.UnpackVariadic(arguments, &orderType, &findTypeFrom, &findTypeTo, &findDataFrom, &findDataTo, &findCompareRawFrom, &findCompareRawTo, &deepSearchTypeFrom, &deepSearchTypeTo, &heightFrom, &heightTo, &dereferenceTypeFrom, &dereferenceTypeTo, &passSubstacksFrom, &passSubstacksTo, &passCardsFrom, &passCardsTo, &workingMemFrom, &workingMemTo)
	if findTypeFrom == nil {findTypeFrom = FIND_First}
	setORDERDefaultIfNil(&orderType)

	// main
	cardTo := stack.Get(findTypeTo, findDataTo, findCompareRawTo, deepSearchTypeTo, heightTo, dereferenceTypeTo, passSubstacksTo, passCardsTo, workingMemTo)
	cardFrom := stack.Extract(findTypeFrom, findDataFrom, findCompareRawFrom, deepSearchTypeFrom, heightFrom, dereferenceTypeFrom, passSubstacksFrom, passCardsFrom, workingMemFrom)
	
	// return
	return stack.Add(cardFrom, orderType, FIND_Card, cardTo, nil, nil, DEEPSEARCH_True)

}

/** Swaps one card with the position of another card
 
 @receiver `stack` type{*Stack}
 @param optional `findType1` type{FIND} default FIND_First
 @param optional `findType2` type{FIND} default FIND_Last
 @param optional `findData1` type{any, []any, *Stack any, func(*Card, *Stack, bool, ...any) (bool)} default nil
 @param optional `findData2` type{any, []any, *Stack any, func(*Card, *Stack, bool, ...any) (bool)} default nil
 @param optional `findCompareRaw1` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `findCompareRaw2` type{COMPARE} default COMPARE_False
 @param optional `deepSearchType1` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `deepSearchType2` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `height1` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `height2` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType1` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `dereferenceType2` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks1` type{PASS} default PASS_True
 @param optional `passSubstacks2` type{PASS} default PASS_True
 @param optional `passCards1` type{PASS} default PASS_True
 @param optional `passCards2` type{PASS} default PASS_True
 @param optional `workingMem1` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
 @param optional `workingMem2` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack`
 */
func (stack *Stack) Swap(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var findType1, findType2, findData1, findData2, findCompareRaw1, findCompareRaw2, deepSearchType1, deepSearchType2, height1, height2, dereferenceType1, dereferenceType2, passSubstacks1, passSubstacks2, passCards1, passCards2, workingMem1, workingMem2 any
	gogenerics.UnpackVariadic(arguments, &findType1, &findType2, &findData1, &findData2, &findCompareRaw1, &findCompareRaw2, &deepSearchType1, &deepSearchType2, &height1, &height2, &dereferenceType1, &dereferenceType2, &passSubstacks1, &passSubstacks2, &passCards1, &passCards2, &workingMem1, &workingMem2)
	if findType1 == nil {findType1 = FIND_First}

	// get card1, add a placeholder right after card1, then remove card1
	card1Placeholder := MakeCard()
	card1 := stack.Get(findType1, findData1, findCompareRaw1, deepSearchType1, height1, dereferenceType1, passSubstacks1, passCards1, workingMem1)
	stack.Add(card1Placeholder, ORDER_After, FIND_Card, card1, nil, nil, DEEPSEARCH_True)
	stack.Remove(FIND_Card, card1, nil, DEEPSEARCH_True)
	
	// get card2, insert card1 after card2, then remove card2
	card2 := stack.Get(findType2, findData2, findCompareRaw2, deepSearchType2, height2, dereferenceType2, passSubstacks2, passCards2, workingMem2)
	stack.Add(card1, ORDER_After, FIND_Card, card2, nil, nil, DEEPSEARCH_True)
	stack.Remove(FIND_Card, card2, nil, DEEPSEARCH_True)

	// insert card2 after the placeholder, then remove the placeholder
	stack.Add(card2, ORDER_After, FIND_Card, card1Placeholder, nil, nil, DEEPSEARCH_True)
	stack.Remove(FIND_Card, card1Placeholder, nil, DEEPSEARCH_True)
	
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
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
 @returns type{bool}
 */
func (stack *Stack) Has(arguments ...any) bool {

	// return
	return stack.Get(arguments...) != nil

}

/** Gets the first card from specified parameters in a stack, or nil if does not exist

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any, []any, *Stack any, func(*Card, *Stack, bool, ...any) (bool)} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns type{*Card} the found card OR nil (if invalid find)
    IF `find` is FIND_Lambda, `findData` is of type{ func(card *Card, parentStack *Stack, isSubstack bool, workingMem ...any) (bool) }
*/
 func (stack *Stack) Get(arguments ...any) *Card {
	
	// unpack arguments into optional parameters
	var findType, findData, findCompareRaw, deepSearchType, depth, dereferenceType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &findCompareRaw, &deepSearchType, &depth, &dereferenceType, &passSubstacks, &passCards, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// get card
	out := stack.LambdaCard(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any, wmadrs ...any) {

		if selectCard(findType, findData, dereferenceType, findCompareRaw.(COMPARE), card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs...) && retCard.Idx == -1 {

			*otherInfo[3].(**Card) = *otherInfo[0].(**Card)

		}

	}, nil, nil, nil, workingMem.([]any), deepSearchType, depth, passSubstacks, passCards)

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
 @param optional `findData` type{any, []any, *Stack any, func(*Card, *Stack, bool, *Stack, ...any) (bool)} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
 @constructs a new stack
 @returns type{*Stack} the new stack
 @requires
   IF `find` is FIND_Lambda, `findData` is of type{ func(card *Card, parentStack *Stack, isSubstack bool, workingStack *Stack, workingMem ...any) (bool) }
 */
func (stack *Stack) GetMany(findType FIND, arguments ...any) *Stack {
	
	// unpack arguments into optional parameters
	var findData, findCompareRaw, returnType, deepSearchType, depth, dereferenceType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(arguments, &findData, &findCompareRaw, &returnType, &deepSearchType, &depth, &dereferenceType, &passSubstacks, &passCards, &workingMem)
	if returnType == nil {returnType = RETURN_Cards}
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// make new stack and return
	return stack.LambdaStack(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any, wmadrs ...any) {
		
		if selectCard(findType, findData, dereferenceType, findCompareRaw.(COMPARE), card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs...) {

			switch returnType {
			case RETURN_Keys:
				retStack.Cards = append(retStack.Cards, MakeCard(card.Key))
			case RETURN_Vals:
				retStack.Cards = append(retStack.Cards, MakeCard(card.Val))
			case RETURN_Idxs:
				retStack.Cards = append(retStack.Cards, MakeCard(card.Idx))
			case RETURN_Cards:
				retStack.Cards = append(retStack.Cards, card.Clone())
			case RETURN_Stacks:
				retStack.Cards = append(retStack.Cards, card.Val.(*Stack).Cards...)
			}
			
		}

	}, nil, nil, nil, workingMem.([]any), deepSearchType, depth, passSubstacks, passCards)

}

/** Returns a clone of a found card before its respective field is updated to `replaceWith` (OR nil if not found)
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceWith` type{any, []any, *Stack, func(card *Card, _ *Stack, _ bool, _ ...any)}
   only set to []any or *Stack through which to iterate if `replaceType` == REPLACE_Cards
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns type{*Card}
 @updates `stack`
 @ensures
   * REPLACE_Card with nil as input ensures the card is removed
 */
func (stack *Stack) Replace(replaceType REPLACE, replaceWith any, arguments ...any) *Card {

	// unpack arguments into optional parameters
	var findType, findData, findCompareRaw, deepSearchType, depth, dereferenceType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &findCompareRaw, &deepSearchType, &depth, &dereferenceType, &passSubstacks, &passCards, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// main
	return stack.LambdaCard(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any, wmadrs ...any) {
		
		if selectCard(findType, findData, dereferenceType, findCompareRaw.(COMPARE), card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs...) && retCard.Idx == -1 {

			*retCard = *card.Clone() // return the original card

			// replace mechanism
			switch replaceType {
			case REPLACE_Key:

				card.Key = replaceWith

			case REPLACE_Val:

				card.Val = replaceWith

			case REPLACE_Card:

				// initialize variables
				insertArr := []any {}
				insertCards := []*Card {}

				// set up insertArr
				switch getType(replaceWith, false) {
				case "element":
					insertArr = append(insertArr, replaceWith)
				case "slice":
					insertArr = gogenerics.UnpackArray(replaceWith)
				case "stack":
					insertArr = replaceWith.(*Stack).ToArray(RETURN_Cards)
				}
	
				// set up insertCards
				for _, ins := range insertArr {
					if ins != nil {
						insertCards = append(insertCards, ins.(*Card)) // insert this card
					} else {
						insertCards = append(insertCards, nil)
					}
				}
	
				// update the stack to have replaceWith at its respective location
				targetIdx := card.Idx
				beginningSegment := parentStack.Cards[:targetIdx]
				endSegment := parentStack.Cards[targetIdx+1:]
	
				parentStack.Cards = []*Card {}
				parentStack.Cards = append(parentStack.Cards, beginningSegment...)
				if !(len(insertArr) == 1 && insertArr[0] == nil) {
					parentStack.Cards = append(parentStack.Cards, insertCards...)
				}
				parentStack.Cards = append(parentStack.Cards, endSegment...)

			case REPLACE_Lambda:
				
				conversion1, success := replaceWith.(func())
				if success {conversion1()}
				conversion2, success := replaceWith.(func(*Card))
				if success {conversion2(card)}
				conversion3, success := replaceWith.(func(*Card, *Stack))
				if success {conversion3(card, parentStack)}
				conversion4, success := replaceWith.(func(*Card, *Stack, bool))
				if success {conversion4(card, parentStack, isSubstack)}
				conversion5, success := replaceWith.(func(*Card, *Stack, bool, ...any))
				if success {conversion5(card, parentStack, isSubstack, wmadrs...)}
				
			}

		}

	}, nil, nil, false, workingMem.([]any), deepSearchType, depth, passSubstacks, passCards)

}

/** Returns a stack of clones of found cards before its respective field is updated to `replaceWith`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceWith` type{any, []any, *Stack, func(card *Card, _ *Stack, _ bool, _ ...any)}
   only set to []any or *Stack through which to iterate if `replaceType` == REPLACE_Cards
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns type{*Stack}
 @updates `stack`
 @ensures
   * REPLACE_Card with nil as input ensures the card is removed
 */
 func (stack *Stack) ReplaceMany(replaceType REPLACE, replaceWith any, arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var findType, findData, findCompareRaw, returnType, deepSearchType, depth, dereferenceType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &findCompareRaw, &returnType, &deepSearchType, &depth, &dereferenceType, &passSubstacks, &passCards, &workingMem)
	if returnType == nil {returnType = RETURN_Cards}
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// main
	return stack.LambdaStack(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any, wmadrs ...any) {
		
		if selectCard(findType, findData, dereferenceType, findCompareRaw.(COMPARE), card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs...) {

			retStack.Cards = append(retStack.Cards, card.Clone())

			// replace mechanism
			switch replaceType {
			case REPLACE_Key:

				card.Key = replaceWith

			case REPLACE_Val:

				card.Val = replaceWith

			case REPLACE_Card:

				// initialize variables
				insertArr := []any {}
				insertCards := []*Card {}

				// set up insertArr
				switch getType(replaceWith, false) {
				case "element":
					insertArr = append(insertArr, replaceWith)
				case "slice":
					insertArr = gogenerics.UnpackArray(replaceWith)
				case "stack":
					insertArr = replaceWith.(*Stack).ToArray(RETURN_Cards)
				}
	
				// set up insertCards
				for _, ins := range insertArr {
					if ins != nil {
						insertCards = append(insertCards, ins.(*Card).Clone()) // insert a clone of this card
					} else {
						insertCards = append(insertCards, nil)
					}
				}
	
				// update the stack to have replaceWith at its respective location
				targetIdx := card.Idx
				beginningSegment := parentStack.Cards[:targetIdx]
				endSegment := parentStack.Cards[targetIdx+1:]
	
				parentStack.Cards = []*Card {}
				parentStack.Cards = append(parentStack.Cards, beginningSegment...)
				if !(len(insertArr) == 1 && insertArr[0] == nil) {
					parentStack.Cards = append(parentStack.Cards, insertCards...)
				}
				parentStack.Cards = append(parentStack.Cards, endSegment...)

			case REPLACE_Lambda:

				replaceWith.(func(*Card, *Stack, bool, ...any)) (card, parentStack, isSubstack, wmadrs...)

			}

		}

	}, nil, nil, false, workingMem.([]any), deepSearchType, depth, passSubstacks, passCards).GetMany(FIND_All, nil, nil, returnType)

}

/** Returns a clone of a found card before it was removed (OR nil if not found)
 
 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns type{*Card}
 @updates `stack`
 @ensures
   * REPLACE_Card with nil as input ensures the card is removed
 */
func (stack *Stack) Extract(arguments ...any) *Card {

	// return the original value
	return stack.Replace(REPLACE_Card, nil, arguments...)

}

/** Returns a stack of clones of found cards before they were removed
 
 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns type{*Stack}
 @updates `stack`
 @ensures
   * REPLACE_Card with nil as input ensures the card is removed
 */
func (stack *Stack) ExtractMany(arguments ...any) *Stack {

	// return the original value
	return stack.ReplaceMany(REPLACE_Card, nil, arguments...)

}

/** Removes a card from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack`
 @updates `stack`
 @ensures
   * REPLACE_Card with nil as input ensures the card is removed
 */
func (stack *Stack) Remove(arguments ...any) *Stack {

	// remove the card
	stack.Replace(REPLACE_Card, nil, arguments...)

	// return stack
	return stack

}

/** Removes a card from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack`
 @updates `stack`
 @ensures
   * REPLACE_Card with nil as input ensures the card is removed
 */
func (stack *Stack) RemoveMany(arguments ...any) *Stack {

	// remove the card
	stack.ReplaceMany(REPLACE_Card, nil, arguments...)

	// return stack
	return stack

}

/** Updates a card in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceWith` type{any, []any, *Stack, func(card *Card, _ *Stack, _ bool, _ ...any)}
   only set to []any or *Stack through which to iterate if `replaceType` == REPLACE_Cards
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack`
 @updates `stack`
 @ensures
   * REPLACE_Card with nil as input ensures the card is removed
 */
 func (stack *Stack) Update(replaceType REPLACE, replaceWith any, arguments ...any) *Stack {

	// update stack
	stack.Replace(replaceType, replaceWith, arguments...)

	// return the original stack
	return stack

}

/** Updates all matched cards in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceWith` type{any, []any, *Stack, func(card *Card, _ *Stack, _ bool, _ ...any)}
   only set to []any or *Stack through which to iterate if `replaceType` == REPLACE_Cards
 @param optional `findType` type{FIND} default FIND_Last
 @param optional `findData` type{any} default nil
 @param optional `findCompareRaw` type{COMPARE} default COMPARE_False
   By default, if an array or Stack is passed into findData, it will iterate through each of its elements in its search.  If you would like to find an array or Stack itself without iterating through their elements, set this to true
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int, []int, *Stack ints} default -1 (heightest)
 @param optional `dereferenceType` type{DEREFERENCE} default DEREFERENCE_None
 @param optional `passSubstacks` type{PASS} default PASS_True
 @param optional `passCards` type{PASS} default PASS_True
 @param optional `workingMem` type{[]any} default []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
   to add more than 10 (n) working memory variables, you must initialize workingMem with an []any argument with n variables
 @returns `stack`
 @updates `stack`
 @ensures
   * REPLACE_Card with nil as input ensures the card is removed
 */
 func (stack *Stack) UpdateMany(replaceType REPLACE, replaceWith any, arguments ...any) *Stack {

	// update stack
	stack.ReplaceMany(replaceType, replaceWith, arguments...)

	// return the original stack
	return stack

}
