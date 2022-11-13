package gostack

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
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
 
 MakeStack(input1 []any|map[any]any|*Stack [nil], input2 []any|*Stack [nil], repeats int [1]) (newStack *Stack)
 
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
 |           unpack cards in `input1` into `stack`
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
	var input1, input2, repeats any
	gogenerics.UnpackVariadic(arguments, &input1, &input2, &repeats)
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
		
		stack.Cards = append(stack.Cards, MakeStackMatrix(input1, input2, []int{matrixShape}).Cards...)
	}

	// property sets
	stack.setStackProperties()

	// return
	return stack

}

/** An identical implementation to `MakeStack()`

 MakeSubstack(input1 []any|map[any]any|*Stack [nil], input2 any|*Stack [nil], repeats int [1], overrideInsert OVERRIDE [OVERRIDE_False]) (newSubstack *Stack)
 
 @examples
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}) => Stack{Stack{1, 2}, Stack{3, 4}}
 */
 func MakeSubstack(arguments ...any) *Stack {
	return MakeStack(arguments...)
}

/** Creates a stack matrix initialized with starting cards

 MakeStackMatrix(input1 []any (deep/shallow)|map[any]any (deep/shallow)|*Stack [nil], input2 []any (deep/shallow)|*Stack [nil], matrixShape []int [[]int {1}]) (newStackMatrix *Stack)
 
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
 |          unpack cards from `input1` into `stack`
 |        ELSE:
 |          unpack values from `input1` into new cards
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
	var input1, input2, matrixShape any
	gogenerics.UnpackVariadic(arguments, &input1, &input2, &matrixShape)

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
					stack.makeStackMatrixFrom1D(matrixShape.([]int), keys, vals, new(int))
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
						stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, input1Array, new(int))

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
						stack.makeStackMatrixFrom1D(matrixShape.([]int), input1Array, input2Array, new(int))
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
				stack.makeStackMatrixFrom1D(matrixShape.([]int), input2Array, nil, new(int))
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
			stack.makeStackMatrixFrom1D(matrixShape.([]int), nil, nil, new(int))

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
		return stack.Filter(FIND_All, nil, RETURN_Stacks)
	} else {
		idxInt, isInt := idx.(int)
		if isInt {
			return stack.Filter(FIND_Idx, idxInt, RETURN_Stacks)
		} else {
			return stack.Filter(FIND_Slice, idx, RETURN_Stacks)
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

/** Appends the cards in `stack` to itself `n` - 1 times

 stack.Duplicate(n int [2]) (stack)

 @examples
 | MakeStack([]int {1}).Duplicate(0) // Stack{}
 | MakeStack([]int {1}).Duplicate(1) // Stack{1}
 | MakeStack([]int {1}).Duplicate(2) // Stack{1, 1}
 | MakeStack([]int {1, 2}).Duplicate(3) // Stack{1, 2, 1, 2, 1, 2}
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
		for _, c := range cardsSave {
			stack.Cards = append(stack.Cards, c.Clone())
		}
	}

	stack.setStackProperties()

	return stack

}

/** Removes all cards from `stack`

 stack.Empty() (stack)
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
	clone.Key = card.Key
	clone.Val = card.Val
	
	// return
	return clone

}

/**  Returns a clone of `card`

 stack.Clone(deepSearchType DEEPSEARCH [DEEPSEARCH_True], depth int [-1], cloneCardKeys CLONE [CLONE_True], cloneCardVals CLONE [CLONE_True], cloneSubstackKeys CLONE [CLONE_True], cloneSubstackVals CLONE [CLONE_True]) (stack)

 @ensures
 | If `cloneSubstackVals` == CLONE_False, then each card holding a substack as its Val will have its Val updated to nil
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

/** Removes all cards from `stack` which share a given property as another card in that stack

 stack.Unique(uniqueType TYPE [TYPE_Val], dereferenceType DEREFERENCE [DEREFERENCE_None]) (stack)

 @examples
 | MakeStack([]int {1, 2, 3, 1, 2, 4}).Unique() // Stack{1, 2, 3, 4}
 | MakeStack([]int {0, 1, 0, 0, 1, 0}, []int {1, 2, 3, 1, 2, 4}).Unique(TYPE_Key) // Stack{1, 2}
 */
func (stack *Stack) Unique(arguments ...any) *Stack {
	
	// unpack arguments into optional parameters
	var uniqueType, dereferenceType any
	gogenerics.UnpackVariadic(arguments, &uniqueType, &dereferenceType)
	if uniqueType == nil { uniqueType = TYPE_Val }
	if dereferenceType == nil { dereferenceType = DEREFERENCE_None }
	if dereferenceType != DEREFERENCE_None { dereferenceType = DEREFERENCE_This }

	// main
	return stack.GetMany(FIND_Lambda, func(card *Card, _ *Stack, _ bool, _ *Stack, workingStack *Stack) (bool) {
		if workingStack.Size == 0 {
			return true
		} else {
			switch uniqueType.(TYPE) {
			case TYPE_Key:
				return !workingStack.Has(FIND_Key, card.Key, nil, nil, nil, dereferenceType)
			case TYPE_Val:
				test := !workingStack.Has(FIND_Val, card.Val, nil, nil, nil, dereferenceType)
				return test
				// return !workingStack.Has(FIND_Val, card.Val, nil, nil, nil, dereferenceType)
			}
			return false // just so it compiles
		}
	})

}

/** Returns whether one card equals another

 card.Equals(
    otherCard *Card,
    compareIdxs COMPARE [COMPARE_False],
    compareKeys COMPARE [COMPARE_True],
    compareVals COMPARE [COMPARE_True],
	comparecardPtrs COMPARE [COMPARE_False],
    pointerKeys DEREFERENCE [DEREFERENCE_None],
    pointerVals DEREFERENCE [DEREFERENCE_None]
 ) (cardEqualsOtherCard bool)

 @examples
 | card1 := MakeCard("Hey")
 | card2 := MakeCard("Hey")
 | myStr := "Hey"
 | card1.Equals(card2, nil, nil, nil, COMPARE_False) // True
 | card1.Equals(card2, nil, nil, nil, COMPARE_True) // False
 | card1.Equals(MakeCard(&myStr), nil, nil, nil, nil, nil, DEREFERENCE_This) // True
 */
func (thisCard *Card) Equals(otherCard *Card, arguments ...any) bool {

	// unpack arguments into optional parameters
	var compareIdxs, compareKeys, compareVals, comparecardPtrs, pointerKeys, pointerVals any
	gogenerics.UnpackVariadic(arguments, &compareIdxs, &compareKeys, &compareVals, &comparecardPtrs, &pointerKeys, &pointerVals)
	// set default vals
	setDEREFERENCEDefaultIfNil(&pointerKeys)
	setDEREFERENCEDefaultIfNil(&pointerVals)
	if compareIdxs == nil {compareIdxs = COMPARE_False}
	setCOMPAREDefaultIfNil(&compareKeys)
	setCOMPAREDefaultIfNil(&compareVals)
	if comparecardPtrs == nil {comparecardPtrs = COMPARE_False}

	condition := thisCard != nil && otherCard != nil
	
	condition = condition && 
		(compareKeys == COMPARE_False ||
		(compareKeys == COMPARE_True && ( compareDereference(pointerKeys.(DEREFERENCE), thisCard.Key, otherCard.Key) ) ) )
	
	condition = condition && 
		(compareVals == COMPARE_False ||
		(compareVals == COMPARE_True && ( compareDereference(pointerVals.(DEREFERENCE), thisCard.Val, otherCard.Val) ) ) )

	condition = condition && (compareIdxs == COMPARE_False || (compareIdxs == COMPARE_True && thisCard.Idx == otherCard.Idx))

	condition = condition && (comparecardPtrs == COMPARE_False || (comparecardPtrs == COMPARE_True && fmt.Sprintf("%p", thisCard) == fmt.Sprintf("%p", otherCard)))
	
	// return whether conditions yield true
	return condition

}

/** Returns whether one stack equals another

 stack.Equals(
    otherStack *Stack,
    deepSearchType *DEEPSEARCH [DEEPSEARCH_True],
    depth int|[]int|*Stack,
    compareCardKeys COMPARE [COMPARE_True],
    compareCardVals COMPARE [COMPARE_True],
    compareSubstackKeys COMPARE [COMPARE_True],
    pointerCardKeys DEREFERENCE [DEREFERENCE_None],
    pointerCardVals DEREFERENCE [DEREFERENCE_None],
    pointerSubstackKeys DEREFERENCE [DEREFERENCE_None],
    compareSubstackAdrs COMPARE [COMPARE_False]
 ) (stackEqualsOtherStack bool)
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
	var deepSearchType, depth, compareCardKeys, compareCardVals, compareSubstackKeys, pointerCardKeys, pointerCardVals, pointerSubstackKeys, compareSubstackAdrs any
	gogenerics.UnpackVariadic(arguments, &deepSearchType, &depth, &compareCardKeys, &compareCardVals, &compareSubstackKeys, &pointerCardKeys, &pointerCardVals, &pointerSubstackKeys, &compareSubstackAdrs)
	// set default vals
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setHeightDefaultIfNil(&depth)

	setCOMPAREDefaultIfNil(&compareCardKeys)
	setCOMPAREDefaultIfNil(&compareCardVals)
	setCOMPAREDefaultIfNil(&compareSubstackKeys)
	
	setDEREFERENCEDefaultIfNil(&pointerCardKeys)
	setDEREFERENCEDefaultIfNil(&pointerCardVals)
	setDEREFERENCEDefaultIfNil(&pointerSubstackKeys)

	if compareSubstackAdrs == nil { compareSubstackAdrs = COMPARE_False }

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
						if compareSubstackAdrs == COMPARE_True {
							test = test && fmt.Sprintf("%p", substackA) == fmt.Sprintf("%p", substackB)
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
					test = test && substackA.Equals(substackB, deepSearchType, transformedHeight, compareCardKeys, compareCardVals, compareSubstackKeys, pointerCardKeys, pointerCardVals, pointerSubstackKeys, compareSubstackAdrs)

				} else if (cardAIsSubstack && !cardBIsSubstack) || (!cardAIsSubstack && cardBIsSubstack) { // one holds a substack and the other doesnt

					test = false

				} else { // neither are substacks and they are both just cards

					// compare card properties
					if testLayer {
						test = test && cardA.Equals(cardB, COMPARE_True, compareCardKeys, compareCardVals, COMPARE_False, pointerCardKeys, pointerCardVals)
					}

				}
			}
		}
	}

	// backpropagate
	return test

}

/** Shuffles the order of `stack` cards

 stack.Shuffle(repeatType REPEAT [REPEAT_False]) (stack)

 @ensures
 | IF `repeatType` == true AND stack.Size > 1:
 |   shuffles `stack` until it is no longer in its previous order
 | rand.Seed is updated to time.Now().UnixNano()
 */
func (stack *Stack) Shuffle(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var repeatType any
	gogenerics.UnpackVariadic(arguments, &repeatType)
	if repeatType == nil {repeatType = REPEAT_False}

	// body
	initClone := stack.Clone()

	for ok := true; ok; ok = (repeatType.(REPEAT) == REPEAT_True && stack.Size > 1 && initClone.Equals(stack)) { // emulate a do-while loop
		
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

/** Reverses the order of the first immediate cards/substacks in `stack`

 stack.Flip() (stack)
 */
func (stack *Stack) Flip() *Stack {

	// main
	*stack = *stack.LambdaStack(func(card *Card, _ *Stack, _ bool, _ *Stack, retStack *Stack) {
		retStack.Add(card, ORDER_Before, FIND_First)
	}, nil, nil, nil, nil, DEEPSEARCH_False)

	// return
	return stack

}

/** Switches the Key and the Val of `card`

 card.SwitchKeyVal() (card)
 */
 func (card *Card) SwitchKeyVal() *Card {

	// main
	tempVal := card.Val
	card.Val = card.Key
	card.Key = tempVal

	// return
	return card

 }

 /** Switches the Key and the Val of each found card and returns `stack`

  stack.SwitchKeysVals(
    findType FIND [FIND_All],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
  */
  func (stack *Stack) SwitchKeysVals(arguments ...any) *Stack {

	// set up arguments
	var findType, findData, deepSearchType, depth, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &deepSearchType, &depth, &dereferenceType, &overrideFindData, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if overrideFindData == nil {overrideFindData = OVERRIDE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if findType == nil {findType = FIND_All}
 
	// main
	stack.Lambda(func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarPtr any, otherInfo []any, workingMem ...any) {
		
		if selectCard(findType, findData, dereferenceType, overrideFindData.(OVERRIDE), card, parentStack, isSubstack, coords, retStack, retCard, retVarPtr, workingMem...) {
			
			card.SwitchKeyVal()

		}

	}, nil, nil, nil, workingMem.([]any), deepSearchType, depth, PASS_Cards)
 
	 // return
	 return stack
 
  }

/** Prints information surrounding `card` to the terminal and returns `card`
 
 card.Print(name string [""], indent int [0]) (card)

 @ensures
 | prints "-" `indent` * 4 times before each line to indicate depth in a stackMatrix
 */
func (card *Card) Print(arguments ...any) *Card {

	// unpack arguments into optional parameters
	var name, indents any
	gogenerics.UnpackVariadic(arguments, &name, &indents)
	if name == nil { name = "" }
	if indents == nil { indents = 0 }

	// prints
	fmt.Printf("%v|%vCARD%v\n", depthPrinter(indents.(int)), gogenerics.IfElse(indents == 0, "gostack: PRINTING ", ""), gogenerics.IfElse(name == "", "", ": \"" + name.(string) + "\""))
	if card == nil {
		fmt.Printf("%v- card:          %v\n", depthPrinter(indents.(int)), nil)
	} else {
		fmt.Printf("%v- &card:         %v\n", depthPrinter(indents.(int)), fmt.Sprintf("%p", card))
		fmt.Printf("%v- card.Idx:      %v\n", depthPrinter(indents.(int)), card.Idx)
		if gogenerics.IsPointer(card.Key) {
			fmt.Printf("%v- card.Key:      %v\n", depthPrinter(indents.(int)), fmt.Sprintf("%p", card.Key))
			fmt.Printf("%v- *card.Key:     %v\n", depthPrinter(indents.(int)), reflect.ValueOf(card.Key).Elem())
			fmt.Printf("%v- card.Key.Type: (*%v)\n", depthPrinter(indents.(int)), reflect.TypeOf(reflect.ValueOf(card.Key).Elem().Interface()))
		} else {
			fmt.Printf("%v- card.Key:      %v\n", depthPrinter(indents.(int)), card.Key)
			if card.Key != nil {
				fmt.Printf("%v- card.Key.Type: (%v)\n", depthPrinter(indents.(int)), reflect.TypeOf(card.Key))
			}
		}
		if gogenerics.IsPointer(card.Val) {
			fmt.Printf("%v- card.Val:      %v\n", depthPrinter(indents.(int)), fmt.Sprintf("%p", card.Val))
			fmt.Printf("%v- *card.Val:     %v\n", depthPrinter(indents.(int)), reflect.ValueOf(card.Val).Elem())
			fmt.Printf("%v- card.Val.Type: (*%v)\n", depthPrinter(indents.(int)), reflect.TypeOf(reflect.ValueOf(card.Val).Elem().Interface()))
		} else {
			fmt.Printf("%v- card.Val:      %v\n", depthPrinter(indents.(int)), card.Val)
			if card.Val != nil {
				fmt.Printf("%v- card.Val.Type: (%v)\n", depthPrinter(indents.(int)), reflect.TypeOf(card.Val))
			}
		}
	}

	return card

}

/** Prints information surrounding `stack` to the terminal and returns `stack`

 stack.Print(indent int [0]) (stack)
 
 @ensures
 | prints "-" `indent` * 4 times before each line to indicate depth in a stackMatrix
 @examples
 | MakeStack([]string {"Hey", "Hi"}).Print().Remove(FIND_Last).Print() // prints the stack before and after performing the remove function
 */
func (stack *Stack) Print(arguments ...any) *Stack {
	
	// unpack arguments into optional parameters
	var name, indent, idx, key any
	gogenerics.UnpackVariadic(arguments, &name, &indent, &idx, &key)
	if indent == nil { indent = 0 }
	if name == nil { name = "" }

	fmt.Printf("%v|%vSTACK%v\n", depthPrinter(indent.(int)), gogenerics.IfElse(idx == nil, "gostack: PRINTING ", "SUB"), gogenerics.IfElse(name == "", "", ": \"" + name.(string) + "\""))
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
				c.Val.(*Stack).Print("", indent.(int)+4, i, c.Key)
			default:
				c.Print("", indent.(int)+4)
			}
		}
	}

	return stack
	
}

/** Iterates through `stack` calling your lambda function on each card, returning `stack`, `retStack`, `retCard`, and `retVarPtr`

 stack.Lambda(
    lambda func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        retStack *Stack,
        retCard *Card,
        retVarPtr any,
        otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
		},
        workingMem ...any),
    retStack *Stack [nil],
    retCard *Card [nil],
    retVarPtr any [nil],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}],
    deepSearchType DEEPSEARCH [DEEPSEARCH_True],
    depth int [-1],
    passType PASS [PASS_Both],
    otherInfo []any {
        retStackPtr,
        retCardPtr
    } []any [[]any {nil, nil}],
 ) (stack, retStack, retCard, retVarPtr)

 @ensures
 | IF a version for `lambda` is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 | IF you would like to manage more than 10 variables via `workingMem`:
 |   you must pass an []any array into `workingMem` when you call this function
 | IF you would like to reference the object address of `retStack` or `retCard`:
 |   pass the addresses of `retStack` or `retCard` into `otherInfo`
 @examples
 | myStack := MakeStackMatrix([]int {1, 3, 2, 4}, nil, []int {2, 2}).LambdaThis(func(card *Card) {
 |   if card.Idx == 0 && card.Val.(int) % 2 == 0 {
 |     card.Key = "Marker"	
 |   }
 | }) // Stack{nil:1, nil:3, "Marker":2, nil:4}
 */
 func (stack *Stack) Lambda(lambda any, arguments ...any) (*Stack, *Stack, *Card, any) {

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

			define this coords
			
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
	var retStack, retCard, retVarPtr, workingMem, deepSearchType, depth, passType, otherInfo, _coords any
	gogenerics.UnpackVariadic(arguments, &retStack, &retCard, &retVarPtr, &workingMem, &deepSearchType, &depth, &passType, &otherInfo, &_coords)
	if retVarPtr == nil {var o any; retVarPtr = &o;}
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setHeightDefaultIfNil(&depth)
	if passType == nil {passType = PASS_Both}
	if otherInfo == nil {otherInfo = []any {nil, nil}}
	retStackPtr := otherInfo.([]any)[0]
	retCardPtr := otherInfo.([]any)[1]
	if retStack == nil {
		newStack := MakeStack()
		retStack = newStack
		retStackPtr = &newStack
		otherInfo.([]any)[0] = retStackPtr
	}
	if retCard == nil {
		newCard := MakeCard()
		retCard = newCard
		retCardPtr = &newCard
		otherInfo.([]any)[1] = retCardPtr
	}
	if _coords == nil { _coords = MakeStack() }
	
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
	
	for i, card := range stack.Cards {

		coords := _coords.(*Stack).Clone().Add(i)
		
		substack, isSubstack := card.Val.(*Stack)
		if isSubstack {

			if (passType == PASS_Both || passType == PASS_Substacks) && passLayer {

				callLambda(lambda, card, stack, true, coords, toTypeStack(retStack), toTypeCard(retCard), &retVarPtr, []any {&card, &stack, retStackPtr, retCardPtr}, workingMem.([]any))

				if retCardPtr != nil && *retCardPtr.(**Card) != nil {
					retCard = *retCardPtr.(**Card)
				}
				if retStackPtr != nil && *retStackPtr.(**Stack) != nil {
					retStack = *retStackPtr.(**Stack)
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
				
				substack.Lambda(lambda, retStack, retCard, retVarPtr, workingMem, deepSearchType, transformedHeight, passType, otherInfo, coords)
				if retCardPtr != nil && *retCardPtr.(**Card) != nil {
					retCard = *retCardPtr.(**Card)
				}
				if retStackPtr != nil && *retStackPtr.(**Stack) != nil {
					retStack = *retStackPtr.(**Stack)
				}
			}

		} else { // card is not substack

			if (passType == PASS_Both || passType == PASS_Cards) && passLayer {

				callLambda(lambda, card, stack, false, coords, toTypeStack(retStack), toTypeCard(retCard), &retVarPtr, []any {&card, &stack, retStackPtr, retCardPtr}, workingMem.([]any))
				if retCardPtr != nil && *retCardPtr.(**Card) != nil {
					retCard = *retCardPtr.(**Card)
				}
				if retStackPtr != nil && *retStackPtr.(**Stack) != nil {
					retStack = *retStackPtr.(**Stack)
				}

				// update properties
				stack.setStackProperties()
				if retStack != nil {
					retStack.(*Stack).setStackProperties()
				}

			}

		}

	}

	return stack, toTypeStack(retStack), toTypeCard(retCard), retVarPtr

}

/** Iterates through `stack` calling your lambda function on each card, returning `stack`, `retStack`, `retCard`, and `retVarPtr`

 stack.LambdaThis(
    lambda func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        retStack *Stack,
        retCard *Card,
        retVarPtr any,
        otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
		},
        workingMem ...any),
    retStack *Stack [nil],
    retCard *Card [nil],
    retVarPtr any [nil],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}],
    deepSearchType DEEPSEARCH [DEEPSEARCH_True],
    depth int [-1],
    passType PASS [PASS_Both],
    otherInfo []any {
        retStackPtr,
        retCardPtr
    } []any [[]any {nil, nil}],
 ) (stack)

 @ensures
 | IF a version for `lambda` is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 | IF you would like to manage more than 10 variables via `workingMem`:
 |   you must pass an []any array into `workingMem` when you call this function
 | IF you would like to reference the object address of `retStack` or `retCard`:
 |   pass the addresses of `retStack` or `retCard` into `otherInfo`
 @examples
 | myStack := MakeStackMatrix([]int {1, 3, 2, 4}, nil, []int {2, 2}).LambdaThis(func(card *Card) {
 |   if card.Idx == 0 && card.Val.(int) % 2 == 0 {
 |     card.Key = "Marker"	
 |   }
 | }) // Stack{nil:1, nil:3, "Marker":2, nil:4}
 */
func (stack *Stack) LambdaThis(lambda any, arguments ...any) *Stack {
	stack.Lambda(lambda, arguments...)
	return stack
}

/** Iterates through `stack` calling your lambda function on each card, returning `stack`, `retStack`, `retCard`, and `retVarPtr`

 stack.LambdaStack(
    lambda func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        retStack *Stack,
        retCard *Card,
        retVarPtr any,
        otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
		},
        workingMem ...any),
    retStack *Stack [nil],
    retCard *Card [nil],
    retVarPtr any [nil],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}],
    deepSearchType DEEPSEARCH [DEEPSEARCH_True],
    depth int [-1],
    passType PASS [PASS_Both],
    otherInfo []any {
        retStackPtr,
        retCardPtr
    } []any [[]any {nil, nil}],
 ) (retStack)

 @ensures
 | IF a version for `lambda` is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 | IF you would like to manage more than 10 variables via `workingMem`:
 |   you must pass an []any array into `workingMem` when you call this function
 | IF you would like to reference the object address of `retStack` or `retCard`:
 |   pass the addresses of `retStack` or `retCard` into `otherInfo`
 @examples
 | myStack := MakeStackMatrix([]int {1, 3, 2, 4}, nil, []int {2, 2}).LambdaThis(func(card *Card) {
 |   if card.Idx == 0 && card.Val.(int) % 2 == 0 {
 |     card.Key = "Marker"	
 |   }
 | }) // Stack{nil:1, nil:3, "Marker":2, nil:4}
 */
func (stack *Stack) LambdaStack(lambda any, arguments ...any) *Stack {
	_, thisStack, _, _ := stack.Lambda(lambda, arguments...)
	return thisStack
}

/** Iterates through `stack` calling your lambda function on each card, returning `stack`, `retStack`, `retCard`, and `retVarPtr`

 stack.LambdaCard(
    lambda func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        retStack *Stack,
        retCard *Card,
        retVarPtr any,
        otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
		},
        workingMem ...any),
    retStack *Stack [nil],
    retCard *Card [nil],
    retVarPtr any [nil],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}],
    deepSearchType DEEPSEARCH [DEEPSEARCH_True],
    depth int [-1],
    passType PASS [PASS_Both],
    otherInfo []any {
        retStackPtr,
        retCardPtr
    } []any [[]any {nil, nil}],
 ) (retCard)

 @ensures
 | IF a version for `lambda` is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 | IF you would like to manage more than 10 variables via `workingMem`:
 |   you must pass an []any array into `workingMem` when you call this function
 | IF you would like to reference the object address of `retStack` or `retCard`:
 |   pass the addresses of `retStack` or `retCard` into `otherInfo`
 @examples
 | myStack := MakeStackMatrix([]int {1, 3, 2, 4}, nil, []int {2, 2}).LambdaThis(func(card *Card) {
 |   if card.Idx == 0 && card.Val.(int) % 2 == 0 {
 |     card.Key = "Marker"	
 |   }
 | }) // Stack{nil:1, nil:3, "Marker":2, nil:4}
 */
func (stack *Stack) LambdaCard(lambda any, arguments ...any) *Card {
	_, _, thisCard, _ := stack.Lambda(lambda, arguments...)
	return thisCard
}

/** Iterates through `stack` calling your lambda function on each card, returning `stack`, `retStack`, `retCard`, and `retVarPtr`

 stack.LambdaVarAdr(
    lambda func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        retStack *Stack,
        retCard *Card,
        retVarPtr any,
        otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
		},
        workingMem ...any),
    retStack *Stack [nil],
    retCard *Card [nil],
    retVarPtr any [nil],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}],
    deepSearchType DEEPSEARCH [DEEPSEARCH_True],
    depth int [-1],
    passType PASS [PASS_Both],
    otherInfo []any {
        retStackPtr,
        retCardPtr
    } []any [[]any {nil, nil}],
 ) (retVarPtr)

 @ensures
 | IF a version for `lambda` is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 | IF you would like to manage more than 10 variables via `workingMem`:
 |   you must pass an []any array into `workingMem` when you call this function
 | IF you would like to reference the object address of `retStack` or `retCard`:
 |   pass the addresses of `retStack` or `retCard` into `otherInfo`
 @examples
 | myStack := MakeStackMatrix([]int {1, 3, 2, 4}, nil, []int {2, 2}).LambdaThis(func(card *Card) {
 |   if card.Idx == 0 && card.Val.(int) % 2 == 0 {
 |     card.Key = "Marker"	
 |   }
 | }) // Stack{nil:1, nil:3, "Marker":2, nil:4}
 */
func (stack *Stack) LambdaVarAdr(lambda any, arguments ...any) any {
	_, _, _, retVarPtr := stack.Lambda(lambda, arguments...)
	return retVarPtr
}

/** Sets `stack` to a set of cards from specified parameters in `stack`, returning `stack`

 stack.Filter(
    findType FIND [FIND_All],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    returnType RETURN [RETURN_Cards],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
 func (stack *Stack) Filter(arguments ...any) *Stack {
	*stack = *stack.GetMany(arguments...)
	return stack
}

/** Adds `insert` to `stack` before/after first found card and returns `stack`, or nil if invalid find

 stack.Add(
    insert any|[]any|*Card|*Stack,
    orderType ORDER [ORDER_After],
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideInsert OVERRIDE [OVERRIDE_False],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideInsert` == OVERRIDE_True:
 |   insert `insert` itself, rather than the elements within `insert` (assuming it is a stack or array)
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) Add(insert any, arguments ...any) *Stack {
	
	return stack.addHandler(false, insert, arguments...)

}

/** Adds `insert` to `stack` before/after each found card and returns `stack`, or nil if invalid find

 stack.AddMany(
    insert any|[]any|*Card|*Stack,
    orderType ORDER [ORDER_After],
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideInsert OVERRIDE [OVERRIDE_False],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideInsert` == OVERRIDE_True:
 |   insert `insert` itself, rather than the elements within `insert` (assuming it is a stack or array)
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) AddMany(insert any, arguments ...any) *Stack {
	
	return stack.addHandler(true, insert, arguments...)

}

/** Moves one card to before/after another card and returns `stack`

 stack.Move(
    orderType ORDER [ORDER_After],
    findTypeFrom FIND [FIND_First],
    findTypeTo FIND [FIND_Last],
    findDataFrom any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    findDataTo any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchTypeFrom DEEPSEARCH [DEEPSEARCH_False],
    deepSearchTypeTo DEEPSEARCH [DEEPSEARCH_False],
    depthFrom int [-1],
    depthTo int [-1],
    passTypeFrom PASS [PASS_Both],
    passTypeTo PASS [PASS_Both],
    dereferenceTypeFrom DEREFERENCE [DEREFERENCE_None],
    dereferenceTypeTo DEREFERENCE [DEREFERENCE_None],
    overrideFindDataFrom OVERRIDE [OVERRIDE_False],
    overrideFindDataTo OVERRIDE [OVERRIDE_False],
    workingMemFrom []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
    workingMemTo []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideFindDataX` == OVERRIDE_True:
 |   compare whether each element is equal to `findDataX` itself, rather than each element inside of `findDataX` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) Move(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var orderType, findTypeFrom, findTypeTo, findDataFrom, findDataTo, deepSearchTypeFrom, deepSearchTypeTo, depthFrom, depthTo, passTypeFrom, passTypeTo, dereferenceTypeFrom, dereferenceTypeTo, overrideFindDataFrom, overrideFindDataTo, workingMemFrom, workingMemTo any
	gogenerics.UnpackVariadic(arguments, &orderType, &findTypeFrom, &findTypeTo, &findDataFrom, &findDataTo, &deepSearchTypeFrom, &deepSearchTypeTo, &depthFrom, &depthTo, &passTypeFrom, &passTypeTo, &dereferenceTypeFrom, &dereferenceTypeTo, &overrideFindDataFrom, &overrideFindDataTo, &workingMemFrom, &workingMemTo)
	if findTypeFrom == nil {findTypeFrom = FIND_First}
	setORDERDefaultIfNil(&orderType)

	// main
	cardTo := stack.Get(findTypeTo, findDataTo, deepSearchTypeTo, depthTo, passTypeTo, dereferenceTypeTo, overrideFindDataTo, workingMemTo)

	cardFrom := stack.Extract(findTypeFrom, findDataFrom, deepSearchTypeFrom, depthFrom, passTypeFrom, dereferenceTypeFrom, overrideFindDataFrom, workingMemFrom)
	
	// return
	return stack.Add(cardFrom, orderType, FIND_Card, cardTo, DEEPSEARCH_True)

}

/** Swaps one card with another and returns `stack`

 stack.Swap(
    findType1 FIND [FIND_First],
    findType2 FIND [FIND_Last],
    findData1 any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    findData2 any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType1 DEEPSEARCH [DEEPSEARCH_False],
    deepSearchType2 DEEPSEARCH [DEEPSEARCH_False],
    depth1 int [-1],
    depth2 int [-1],
    passType1 PASS [PASS_Both],
    passType2 PASS [PASS_Both],
    dereferenceType1 DEREFERENCE [DEREFERENCE_None],
    dereferenceType2 DEREFERENCE [DEREFERENCE_None],
    overrideFindData1 OVERRIDE [OVERRIDE_False],
    overrideFindData2 OVERRIDE [OVERRIDE_False],
    workingMem1 []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
    workingMem2 []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideFindDataX` == OVERRIDE_True:
 |   compare whether each element is equal to `findDataX` itself, rather than each element inside of `findDataX` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) Swap(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var findType1, findType2, findData1, findData2, deepSearchType1, deepSearchType2, depth1, depth2, passType1, passType2, dereferenceType1, dereferenceType2, overrideFindData1, overrideFindData2, workingMem1, workingMem2 any
	gogenerics.UnpackVariadic(arguments, &findType1, &findType2, &findData1, &findData2, &deepSearchType1, &deepSearchType2, &depth1, &depth2, &passType1, &passType2, &dereferenceType1, &dereferenceType2, &overrideFindData1, &overrideFindData2, &workingMem1, &workingMem2)
	if findType1 == nil {findType1 = FIND_First}

	// get card1, add a placeholder right after card1, then remove card1
	card1Placeholder := MakeCard()
	card1 := stack.Get(findType1, findData1, deepSearchType1, depth1, passType1, dereferenceType1, overrideFindData1, workingMem1)
	stack.Add(card1Placeholder, ORDER_After, FIND_Card, card1, DEEPSEARCH_True)
	stack.Remove(FIND_Card, card1, DEEPSEARCH_True)
	
	// get card2, insert card1 after card2, then remove card2
	card2 := stack.Get(findType2, findData2, deepSearchType2, depth2, passType2, dereferenceType2, overrideFindData2, workingMem2)
	stack.Add(card1, ORDER_After, FIND_Card, card2, DEEPSEARCH_True)
	stack.Remove(FIND_Card, card2, DEEPSEARCH_True)

	// insert card2 after the placeholder, then remove the placeholder
	stack.Add(card2, ORDER_After, FIND_Card, card1Placeholder, DEEPSEARCH_True)
	stack.Remove(FIND_Card, card1Placeholder, DEEPSEARCH_True)
	
	// return
	return stack

}

/** Returns a bool for whether a card was found in `stack`

 stack.Has(
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (foundACard bool)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) Has(arguments ...any) bool {

	// return
	return stack.Get(arguments...) != nil

}

/** Gets the first card from specified parameters in `stack`, or nil if does not exist

 stack.Get(
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (foundCard *Card)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
*/
 func (stack *Stack) Get(arguments ...any) *Card {
	
	// unpack arguments into optional parameters
	var findType, findData, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &deepSearchType, &depth, &passType, &dereferenceType, &overrideFindData, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if overrideFindData == nil {overrideFindData = OVERRIDE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passType == nil {passType = PASS_Both}

	// get card
	out := stack.LambdaCard(func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarPtr any, otherInfo []any, workingMem ...any) {
		
		if selectCard(findType, findData, dereferenceType, overrideFindData.(OVERRIDE), card, parentStack, isSubstack, coords, retStack, retCard, retVarPtr, workingMem...) && retCard.Idx == -1 {

			*otherInfo[3].(**Card) = *otherInfo[0].(**Card)

		}

	}, nil, nil, nil, workingMem.([]any), deepSearchType, depth, passType)

	// return nil if no card found, else return card
	if out.Idx == -1 {
		return nil
	} else {
		return out
	}

}

/** Gets a stack of cards from specified parameters in `stack`

 stack.GetMany(
    findType FIND [FIND_All],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    returnType RETURN [RETURN_Cards],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (newStack *Stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) GetMany(arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var findType, findData, returnType, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &returnType, &deepSearchType, &depth, &passType, &dereferenceType, &overrideFindData, &workingMem)
	if findType == nil {findType = FIND_All}
	if returnType == nil {returnType = RETURN_Cards}
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if overrideFindData == nil {overrideFindData = OVERRIDE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passType == nil {passType = PASS_Both}

	// make new stack and return
	return stack.LambdaStack(func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarPtr any, otherInfo []any, workingMem ...any) {
		
		if selectCard(findType, findData, dereferenceType, overrideFindData.(OVERRIDE), card, parentStack, isSubstack, coords, retStack, retCard, retVarPtr, workingMem...) {

			switch returnType {
			case RETURN_Keys:
				retStack.Cards = append(retStack.Cards, MakeCard(card.Key))
			case RETURN_Vals:
				retStack.Cards = append(retStack.Cards, MakeCard(card.Val))
			case RETURN_Idxs:
				retStack.Cards = append(retStack.Cards, MakeCard(card.Idx))
			case RETURN_Cards:
				retStack.Cards = append(retStack.Cards, card.Clone())
			case RETURN_Adrs:
				retStack.Cards = append(retStack.Cards, MakeCard(fmt.Sprintf("%p", card)))
			case RETURN_Stacks:
				retStack.Cards = append(retStack.Cards, card.Val.(*Stack).Cards...)
			}
			
		}

	}, nil, nil, nil, workingMem.([]any), deepSearchType, depth, passType)

}

/** Returns a clone of the first found card from specified parameters in `stack`
 stack.Replace(
    replaceType REPLACE,
    replaceWith any|[]any|*Stack|func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        workingMem ...any,
    ),
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (replacedCard *Card)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) Replace(replaceType REPLACE, replaceWith any, arguments ...any) *Card {

	// unpack arguments into optional parameters
	var findType, findData, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &deepSearchType, &depth, &passType, &dereferenceType, &overrideFindData, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if overrideFindData == nil {overrideFindData = OVERRIDE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passType == nil {passType = PASS_Both}
	if findType == nil {findType = FIND_Last}

	// main
	return stack.LambdaCard(func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarPtr any, otherInfo []any, workingMem ...any) {
		
		if selectCard(findType, findData, dereferenceType, overrideFindData.(OVERRIDE), card, parentStack, isSubstack, coords, retStack, retCard, retVarPtr, workingMem...) && retCard.Idx == -1 {

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
				
				callLambdaReplaceWith(replaceWith, card, stack, isSubstack, coords, workingMem)
				
			}

		}

	}, nil, nil, false, workingMem.([]any), deepSearchType, depth, passType, dereferenceType)

}

/** Returns a stack of clones of the found cards from specified parameters in `stack`
 stack.ReplaceMany(
    replaceType REPLACE,
    replaceWith any|[]any|*Stack|func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        workingMem ...any,
    ),
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    returnType RETURN [RETURN_Cards],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (replacedCards *Stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
 func (stack *Stack) ReplaceMany(replaceType REPLACE, replaceWith any, arguments ...any) *Stack {

	// unpack arguments into optional parameters
	var findType, findData, returnType, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &returnType, &deepSearchType, &depth, &passType, &dereferenceType, &overrideFindData, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if overrideFindData == nil {overrideFindData = OVERRIDE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passType == nil {passType = PASS_Both}
	if returnType == nil {returnType = RETURN_Cards}
	if findType == nil {findType = FIND_Last}

	// main
	return stack.LambdaStack(func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarPtr any, otherInfo []any, workingMem ...any) {
		
		if selectCard(findType, findData, dereferenceType, overrideFindData.(OVERRIDE), card, parentStack, isSubstack, coords, retStack, retCard, retVarPtr, workingMem...) {

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
				
				callLambdaReplaceWith(replaceWith, card, stack, isSubstack, coords, workingMem)

			}

		}

	}, nil, nil, false, workingMem.([]any), deepSearchType, depth, passType).GetMany(FIND_All, nil, returnType)

}

/** Removes and returns a found card, or nil if not found

 stack.Extract(
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (extractedCard *Card)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) Extract(arguments ...any) *Card {

	// return the original value
	return stack.Replace(REPLACE_Card, nil, arguments...)

}

/** Removes and returns a stack of found cards

 stack.ExtractMany(
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    returnType RETURN [RETURN_Cards],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (extractedData *Stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) ExtractMany(arguments ...any) *Stack {

	// return the original value
	return stack.ReplaceMany(REPLACE_Card, nil, arguments...)

}

/** Removes a card from and returns `stack`

 stack.Remove(
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) Remove(arguments ...any) *Stack {

	// remove the card
	stack.Replace(REPLACE_Card, nil, arguments...)

	// return stack
	return stack

}

/** Removes a set of cards from and returns `stack`

 stack.RemoveMany(
    findType FIND [FIND_All],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
func (stack *Stack) RemoveMany(arguments ...any) *Stack {

	// unpack variadic
	var findType, findData, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &deepSearchType, &depth, &passType, &dereferenceType, &overrideFindData, &workingMem)

	// remove the card
	stack.ReplaceMany(REPLACE_Card, nil, findType, findData, nil, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem)

	// return stack
	return stack

}

/** Updates a card in and returns `stack`

 stack.Update(
    replaceType REPLACE,
    replaceWith any|[]any|*Stack|func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        workingMem ...any,
    ),
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
 func (stack *Stack) Update(replaceType REPLACE, replaceWith any, arguments ...any) *Stack {

	// update stack
	stack.Replace(replaceType, replaceWith, arguments...)

	// return the original stack
	return stack

}

/** Updates all matched cards in and returns `stack`
 
 stack.UpdateMany(
    replaceType REPLACE,
    replaceWith any|[]any|*Stack|func(
        card *Card,
        parentStack *Stack,
        isSubstack bool,
        coords *Stack,
        workingMem ...any,
    ),
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)

  Updates all matched cards in and returns `stack`

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
 func (stack *Stack) UpdateMany(replaceType REPLACE, replaceWith any, arguments ...any) *Stack {

	// unpack variadic
	var findType, findData, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &deepSearchType, &depth, &passType, &dereferenceType, &overrideFindData, &workingMem)

	// update stack
	stack.ReplaceMany(replaceType, replaceWith, findType, findData, nil, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem)

	// return the original stack
	return stack

}

/** Retrieves a stack containing the coordinates of the first found card, or empty stack if doesn't exist

 stack.Coordinates(
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_True],
    depth int [-1],
    passType PASS [PASS_Cards],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (foundCardCoords *Stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
 func (stack *Stack) Coordinates(arguments ...any) (*Stack) {
	
	// unpack arguments into optional parameters
	var findType, findData, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &deepSearchType, &depth, &passType, &dereferenceType, &overrideFindData, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if overrideFindData == nil {overrideFindData = OVERRIDE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_True}
	if passType == nil {passType = PASS_Cards}
	if findType == nil {findType = FIND_Last}

	// return
	found := false
	return stack.LambdaStack(func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarPtr any, otherInfo []any, workingMem ...any) {
		
		if !found && selectCard(findType, findData, dereferenceType, overrideFindData.(OVERRIDE), card, parentStack, isSubstack, coords, retStack, retCard, retVarPtr, workingMem...) {

			*retStack = *coords
			found = true

		}
	}, nil, nil, nil, workingMem, deepSearchType, depth, passType)

}

/** Retrieves a stack containing a set of stacks containing the coordinates of each found card

 stack.Coordinates(
    findType FIND [FIND_All],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_True],
    depth int [-1],
    passType PASS [PASS_Cards],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (foundCardsCoords *Stack)

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 */
 func (stack *Stack) CoordinatesMany(arguments ...any) (*Stack) {
	
	// unpack arguments into optional parameters
	var findType, findData, deepSearchType, depth, passType, dereferenceType, overrideFindData, workingMem any
	gogenerics.UnpackVariadic(arguments, &findType, &findData, &deepSearchType, &depth, &passType, &dereferenceType, &overrideFindData, &workingMem)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if overrideFindData == nil {overrideFindData = OVERRIDE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_True}
	if passType == nil {passType = PASS_Cards}
	if findType == nil {findType = FIND_All}

	// return
	return stack.LambdaStack(func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarPtr any, otherInfo []any, workingMem ...any) {
		
		if selectCard(findType, findData, dereferenceType, overrideFindData.(OVERRIDE), card, parentStack, isSubstack, coords, retStack, retCard, retVarPtr, workingMem...) {

			retStack.Add(coords, nil, nil, nil, nil, nil, nil, nil, OVERRIDE_True)

		}
	}, nil, nil, nil, workingMem, deepSearchType, depth, passType)

}

/** Takes a CSV at a given file path and returns it as a StackMatrix
 
 CSVToStackMatrix(inPath string) (newStackMatrix *Stack)

 @requires
 | `inPath` points to valid CSV file
 */
 func CSVToStackMatrix(inPath string) *Stack {
	file, _ := os.Open(inPath)
	csvReader := csv.NewReader(file)
	records, _ := csvReader.ReadAll()
	return MakeStackMatrix(records)
}

/** Creates a CSV at a given file path given a StackMatrix

 stack.ToCSV(outPath string) (csvFile *os.File)
 
 @requires
 | `outPath` points to valid location
 */
 func (stack *Stack) ToCSV(outPath string) *os.File {
	file, _ := os.Create(outPath)
	csvWriter := csv.NewWriter(file)
	var csv [][]string
	for i, substack := range stack.ToArray() {
		csv = append(csv, []string {})
		for _, val := range substack.([]any) {
			csv[i] = append(csv[i], val.(string))
		}
	}
	csvWriter.WriteAll(csv)
	return file
}
