package gostack

import (
	"fmt"
	"reflect"

	"github.com/gabetucker2/gogenerics"
)

func needleInHaystack(needle any, haystack []any) bool {
	for _, hay := range haystack {
		if needle == hay {
			return true
		}
	}
	return false
}

/** return {"element", "slice", or "stack"} */
func getType(in any, override bool) string {
	if in == nil || override {
		return "element"
	} else {
		_, isStack := in.(*Stack)
		if isStack {
			return "stack"
		} else {
			isArr := reflect.TypeOf(in).Kind() == reflect.Slice
			if isArr {
				return "slice"
			} else {
				return "element"
			}
		}
	}
}

/** hatstack type{any, []any, *Stack} */
func match(needle any, haystack any, override bool) bool {
	switch getType(haystack, override) {
	case "element":
		_, isCard := needle.(*Card)
		if isCard {
			return needle.(*Card).Equals(haystack.(*Card))
		} else {
			return needle == haystack
		}
	case "slice":
		return needleInHaystack(needle, gogenerics.UnpackArray(haystack))
	case "stack":
		return needleInHaystack(needle, haystack.(*Stack).ToArray())
	default:
		return false
	}
}

func selectCard(findType any, findData any, pointerType any, findCompareRaw COMPARE, card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, wmadrs ...any) bool {

	// set defaults
	setFINDDefaultIfNil(&findType)
	setDEREFERENCEDefaultIfNil(&pointerType)

	override := findCompareRaw == COMPARE_True

	getPointer := func(ptr any, variadic ...any) any {

		// unpack variadic into optional parameters
		var returnIfNotPointer any
		gogenerics.UnpackVariadic(variadic, &returnIfNotPointer)
		// set default
		if returnIfNotPointer == nil {
			returnIfNotPointer = true
		}
	
		_, isPtrInf := ptr.(*any)
		if gogenerics.IsPointer(ptr) && isPtrInf {
			return *ptr.(*any)
		} else if gogenerics.IsPointer(ptr) && !isPtrInf {
			return ptr
		} else {
			if returnIfNotPointer.(bool) {
				return ptr
			} else {
				return nil
			}
		}
	}

	// main
	switch findType {
	case FIND_First:
		return card.Idx == 0
	case FIND_Last:
		return card.Idx == parentStack.Size - 1
	case FIND_Idx: // -1 = stack.Size - 1
		haystack := []any {}
		switch getType(findData, false) {
		case "element":
			haystack = append(haystack, findData)
		case "slice":
			haystack = append(haystack, gogenerics.UnpackArray(findData)...)
		case "stack":
			haystack = append(haystack, findData.(*Stack).ToArray()...)
		}
		for i := range haystack {
			if haystack[i] == -1 {
				haystack[i] = parentStack.Size - 1
			}
		}
		return needleInHaystack(card.Idx, haystack)
	case FIND_Key:
		switch pointerType {
		case DEREFERENCE_True:
			return match(getPointer(card.Key, false), findData, override)
		case DEREFERENCE_False:
			return match(card.Key, findData, override)
		}
	case FIND_Val:
		switch pointerType {
		case DEREFERENCE_True:
			return match(getPointer(card.Val, false), findData, override)
		case DEREFERENCE_False:
			return match(card.Val, findData, override)
		}
	case FIND_Card:
		return fmt.Sprintf("%p", card) == fmt.Sprintf("%p", findData.(*Card))
	case FIND_Size:
		return match(card.Val.(*Stack).Size, findData, false)
	case FIND_Height:
		return match(card.Val.(*Stack).Height, findData, false)
	case FIND_Slice: // [inclusive, inclusive]; -1 => stack.Size - 1
		var slice []any
		switch getType(findData, false) {
		case "slice":
			slice = gogenerics.UnpackArray(findData)
		case "stack":
			slice = findData.(*Stack).ToArray()
		}
		start := slice[0].(int)
		end := slice[1].(int)
		if start == -1 {
			start = parentStack.Size - 1
		}
		if end == -1 {
			end = parentStack.Size - 1
		}
		end += 1
		slice = []any {}
		if start < end {
			for i := start; i < end; i++ {
				slice = append(slice, i)
			}
		} else if start > end {
			for i := start; i >= end - 1; i-- {
				slice = append(slice, i)
			}
		} else { // start == end
			slice = append(slice, start)
		}
		return needleInHaystack(card.Idx, slice)
	case FIND_All:
		return true
	case FIND_Lambda:

		conversion1, success := findData.(func() (bool))
		if success {return conversion1()}
		conversion2, success := findData.(func(*Card) (bool))
		if success {return conversion2(card)}
		conversion3, success := findData.(func(*Card, *Stack) (bool))
		if success {return conversion3(card, parentStack)}
		conversion4, success := findData.(func(*Card, *Stack, bool) (bool))
		if success {return conversion4(card, parentStack, isSubstack)}

		// specific to card return/N
		conversion5, success := findData.(func(*Card, *Stack, bool, ...any) (bool))
		if success {return conversion5(card, parentStack, isSubstack, wmadrs...)}

		// specific to stack return/NMany
		conversion6, success := findData.(func(*Card, *Stack, bool, *Stack) (bool))
		if success {return conversion6(card, parentStack, isSubstack, retStack)}
		conversion7, success := findData.(func(*Card, *Stack, bool, *Stack, ...any) (bool))
		if success {return conversion7(card, parentStack, isSubstack, retStack, wmadrs...)}
	}
	return false

}

/** Recursively add elements from 1D array to stack of matrix shape resembling `matrixShape`
 
 @receiver stack type{*Stack}
 @param matrixShape type{[]int}
 @param keys type{[]any}
 @param vals type{[]any}
 @param globalI type{*int} used because: extracting from 1-D arrays into N-D matrix, so need to track our position in the 1-D arrays between different recursive calls
 @param overrideCards type{OVERRIDE}
 @returns type{*Stack}
 @requires
  * `MakeStack()` and `MakeCard()` have been implemented
  * |keys| == |vals| if neither are nil
  * |keys| or |vals| == product of ints in matrixShape
*/
func (stack *Stack) makeStackMatrixFrom1D(matrixShape []int, keys []any, vals []any, globalI *int, overrideCards OVERRIDE) (ret *Stack) {

	// make stack
	if len(matrixShape) > 1 {
		for i := 0; i < matrixShape[0]; i++ {
			// insert this return value into a card of our current stack
			stack.Cards = append(stack.Cards, MakeCard(
				MakeStack().makeStackMatrixFrom1D(matrixShape[1:], keys, vals, globalI, overrideCards), nil, i))
		}

		ret = stack

	// no more stacks to make, insert elements into and return current stack
	} else {

		ret = stack

		makeNewCard := func() {
			// make new card initialized to vals determined in val array `vals` (do same for keys)
			c := MakeCard()
			updated := false
			if len(vals) > 0 {
				updated = true
				c.Val = vals[*globalI]
			}
			if len(keys) > 0 {
				updated = true
				c.Key = keys[*globalI]
			}
			if updated {
				*globalI++
			}
			ret.Cards = append(ret.Cards, c)
		}

		for i := 0; i < matrixShape[0]; i++ {

			if overrideCards == OVERRIDE_True {
				makeNewCard()
			} else {

				// if inserting vals
				if len(vals) != 0 {
					switch vals[i].(type) {
					case *Card:
						// set to existing card in card array `vals`
						ret.Cards = append(ret.Cards, vals[i].(*Card))
					default:
						makeNewCard()
					}

				// if inserting only keys
				} else if len(keys) != 0 {
					switch keys[i].(type) {
					case *Card:
						// set to existing card in card array `keys`
						ret.Cards = append(ret.Cards, keys[i].(*Card))
					default:
						makeNewCard()
					}
				} else {
					makeNewCard()
				}
			}
		}

	}

	// update properties in this layer
	stack.setStackProperties()

	// return
	return

}

/** Recursively add elements to stack of matrix shape resembling the inputs
 
 @receiver stack type{*Stack}
 @param keys type{any}
 @param vals type{any}
 @returns type{*Stack}
 @requires
  * `MakeStack()` and `MakeCard()` have been implemented
  * |keys| == |vals| if neither are nil
  * at least one of `keys` or `vals` are not nil
*/
func (stack *Stack) makeStackMatrixFromND(keys, vals any) (ret *Stack) {

	// initialize variable to use as reference for the matrix shape
	// just because we don't know which input is not nil
	var referenceArr []any
	// one of these conditions are guaranteed to be true per the ensures clause
	if keys != nil {
		referenceArr = gogenerics.UnpackArray(keys)
	} else if vals != nil {
		referenceArr = gogenerics.UnpackArray(vals)
	}
	
	// main loop
	for i := range referenceArr {
		switch reflect.TypeOf(referenceArr[i]).Kind() {

		// add substack to stack
		case reflect.Slice:
			var keysForward, valsForward any
			if keys != nil {keysForward = gogenerics.UnpackArray(keys)[i]} else {keysForward = nil}
			if vals != nil {valsForward = gogenerics.UnpackArray(vals)[i]} else {valsForward = nil}
			stack.Cards = append(
				stack.Cards,
				MakeCard(MakeStack().makeStackMatrixFromND(
					keysForward,
					valsForward,
				)),
			)

		// add element to stack
		default:
			c := MakeCard()
			if keys != nil {
				c.Key = gogenerics.UnpackArray(keys)[i]
			}
			if vals != nil {
				c.Val = gogenerics.UnpackArray(vals)[i]
			}
			stack.Cards = append(stack.Cards, c)
		}
	}

	// update properties in this layer
	stack.setStackProperties()

	// return
	return stack

}

/** I am too lazy to write documentation for this right now... TODO: add later
 */
 func (stack *Stack) makeStackMatrixFromNDMap(m any) {

	hasSubmap := reflect.TypeOf(m).Elem().Kind() == reflect.Map
	keys, vals := gogenerics.GetKeysValsFromMap(m)

	for i, k := range keys {
		c := new(Card)
		stack.Cards = append(stack.Cards, c)

		c.Key = k

		if hasSubmap {
			c.Val = MakeStack()
			c.Val.(*Stack).makeStackMatrixFromNDMap(vals[i])
		} else {
			c.Val = vals[i]
		}
	}

	// update properties in this layer
	stack.setStackProperties()

}

/** Updates a stack's Size, Height, and Card Indices */
func (stack *Stack) setStackProperties() {
	stack.Size = len(stack.Cards)
	

	/** Assuming normally-shaped matrix, returns the height of this stack */
	var getStackHeight func (stack *Stack) (height int) // so that it can be recursive and nested
	getStackHeight = func (stack *Stack) (height int) {
		
		if stack.Size > 0 {

			c := stack.Cards[0]

			isStack := false
			switch c.Val.(type) {
			case *Stack:
				isStack = true
				height = getStackHeight(c.Val.(*Stack)) + 1
			}

			if !isStack {
				height = 1
			}

		}

		if height == 0 {
			height = 1
		}

		return

	}
	stack.Height = getStackHeight(stack)
	for i := range stack.Cards {
		stack.Cards[i].Idx = i
	}
}

/** Prints some number of - outputs based on height. */
func heightPrinter(height int) (out string) {
	for i := 0; i < height; i++ {
		out += "-"
	}
	return out
}

/** Casts to *Stack, except ensures if `stack` is nil, then doesn't attempt to cast it */
func toTypeStack(stack any) *Stack {
	if stack == nil {
		return nil
	} else {
		return stack.(*Stack)
	}
}

/** Casts to *Card, except ensures if `card` is nil, then doesn't attempt to cast it */
func toTypeCard(card any) *Card {
	if card == nil {
		return nil
	} else {
		return card.(*Card)
	}
}

/** Exists to remove redundancy from writing Add() and AddMany() twice minus one condition */
func (stack *Stack) addHandler(allNotFirst bool, insert any, variadic ...any) *Stack {
	
	// unpack variadic into optional parameters
	var orderType, findType, findData, findCompareRaw, overrideCards, deepSearchType, height, pointerType, passSubstacks, passCards, workingMem any
	gogenerics.UnpackVariadic(variadic, &orderType, &findType, &findData, &findCompareRaw, &overrideCards, &deepSearchType, &height, &pointerType, &passSubstacks, &passCards, &workingMem)
	setOVERRIDEDefaultIfNil(&overrideCards)
	setORDERDefaultIfNil(&orderType)
	setFINDDefaultIfNil(&findType)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if findCompareRaw == nil {findCompareRaw = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passSubstacks == nil {passSubstacks = PASS_True}

	// initialize foundCard to false so you can determine whether valid find and
	// initialize variables
	foundCard := false
	insertArr := []any {}
	insertCards := []*Card {}

	// set up insertArr
	switch getType(insert, false) {
	case "element":
		insertArr = append(insertArr, insert)
	case "slice":
		insertArr = gogenerics.UnpackArray(insert)
	case "stack":
		insertArr = insert.(*Stack).ToArray(RETURN_Cards)
	}

	// set up insertCards
	for _, ins := range insertArr {
		insCard, isCard := ins.(*Card)
		if isCard && overrideCards == OVERRIDE_False {
			insertCards = append(insertCards, insCard) // insert this card
		} else {
			insertCards = append(insertCards, MakeCard(ins)) // insert a card whose val is ins
		}
	}

	// add cards to empty stack
	if stack.Size == 0 {

		foundCard = true
		stack.Cards = append(stack.Cards, insertCards...)
		stack.setStackProperties()

	// else add based on find
	} else {
		stack.Lambda(func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any,  wmadrs ...any) {
		
			// only do add to the first match if ACTION_First, otherwise do for every match
			if (allNotFirst && selectCard(findType, findData, pointerType, findCompareRaw.(COMPARE), card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs...)) || (!allNotFirst && selectCard(findType, findData, pointerType, findCompareRaw.(COMPARE), card, parentStack, isSubstack, retStack, retCard, retVarAdr, wmadrs...) && !foundCard) {
	
				// update foundCard
				foundCard = true
	
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
	
		}, nil, nil, nil, workingMem.([]any), deepSearchType, height, passSubstacks, passCards)
	}

	// return nil if no add was made, else return card
	if !foundCard {
		return nil
	} else {
		return stack
	}

}
