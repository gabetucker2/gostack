package gostack

import (
	"fmt"
	"reflect"

	"github.com/gabetucker2/gogenerics"
)

func needleInHaystack(needle any, haystack []any, dereferenceType DEREFERENCE) bool {
	for _, hay := range haystack {
		if compareDereference(dereferenceType, needle, hay) {
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

func compareDereference(dereferenceType DEREFERENCE, this, found any) bool {
	switch dereferenceType {
	case DEREFERENCE_None:
		return found == this
	case DEREFERENCE_Both:
		return gogenerics.PointersEqual(found, this)
	case DEREFERENCE_Found:
		return gogenerics.IsPointer(found) && reflect.DeepEqual(reflect.ValueOf(found).Elem().Interface(), this)
	case DEREFERENCE_This:
		return gogenerics.IsPointer(this) && reflect.DeepEqual(found, reflect.ValueOf(this).Elem().Interface())
	}
	return false
}

/** hatstack type{any, []any, *Stack} */
func match(needle any, haystack any, override bool, dereferenceType DEREFERENCE) bool {
	switch getType(haystack, override) {
	case "element":
		_, isCard := needle.(*Card)
		if isCard {
			return needle.(*Card).Equals(haystack.(*Card))
		} else {
			return compareDereference(dereferenceType, haystack, needle)
		}
	case "slice":
		return needleInHaystack(needle, gogenerics.UnpackArray(haystack), dereferenceType)
	case "stack":
		return needleInHaystack(needle, haystack.(*Stack).ToArray(), dereferenceType)
	}
	return false
}

func selectCard(findType any, findData any, dereferenceType any, overrideFindData COMPARE, card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarAdr any, workingMem ...any) bool {

	// set defaults
	setFINDDefaultIfNil(&findType)
	setDEREFERENCEDefaultIfNil(&dereferenceType)

	override := overrideFindData == COMPARE_True

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
		return needleInHaystack(card.Idx, haystack, DEREFERENCE_None)
	case FIND_Key:
		return match(getPointer(card.Key, true), findData, override, dereferenceType.(DEREFERENCE))
	case FIND_Val:
		return match(getPointer(card.Val, true), findData, override, dereferenceType.(DEREFERENCE))
	case FIND_Card:
		return fmt.Sprintf("%p", card) == fmt.Sprintf("%p", findData.(*Card))
	case FIND_Size:
		return match(card.Val.(*Stack).Size, findData, false, DEREFERENCE_None)
	case FIND_Height:
		return match(card.Val.(*Stack).Height, findData, false, DEREFERENCE_None)
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
		return needleInHaystack(card.Idx, slice, DEREFERENCE_None)
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
		conversion5, success := findData.(func(*Card, *Stack, bool, *Stack) (bool))
		if success {return conversion5(card, parentStack, isSubstack, coords)}

		// specific to card return/N
		conversion6, success := findData.(func(*Card, *Stack, bool, *Stack, ...any) (bool))
		if success {return conversion6(card, parentStack, isSubstack, coords, workingMem...)}

		// specific to stack return/NMany
		conversion7, success := findData.(func(*Card, *Stack, bool, *Stack, *Stack) (bool))
		if success {return conversion7(card, parentStack, isSubstack, coords, retStack)}
		conversion8, success := findData.(func(*Card, *Stack, bool, *Stack, *Stack, ...any) (bool))
		if success {return conversion8(card, parentStack, isSubstack, coords, retStack, workingMem...)}

	}
	return false

}

/** Recursively add elements from 1D array to stack of matrix shape resembling `matrixShape`
 
 @receiver stack type{*Stack}
 @param matrixShape type{[]int}
 @param keys type{[]any}
 @param vals type{[]any}
 @param globalI type{*int} used because: extracting from 1-D arrays into N-D matrix, so need to track our position in the 1-D arrays between different recursive calls
 @param overrideInsert type{OVERRIDE}
 @returns type{*Stack}
 @requires
  * `MakeStack()` and `MakeCard()` have been implemented
  * |keys| == |vals| if neither are nil
  * |keys| or |vals| == product of ints in matrixShape
*/
func (stack *Stack) makeStackMatrixFrom1D(matrixShape []int, keys []any, vals []any, globalI *int) (ret *Stack) {

	// make stack
	if len(matrixShape) > 1 {
		for i := 0; i < matrixShape[0]; i++ {
			// insert this return value into a card of our current stack
			stack.Cards = append(stack.Cards, MakeCard(
				MakeStack().makeStackMatrixFrom1D(matrixShape[1:], keys, vals, globalI), nil, i))
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
	

	/** Assuming normally-shaped matrix, returns the depth of this stack */
	var getStackHeight func (stack *Stack) (depth int) // so that it can be recursive and nested
	getStackHeight = func (stack *Stack) (depth int) {
		
		if stack.Size > 0 {

			c := stack.Cards[0]

			isStack := false
			switch c.Val.(type) {
			case *Stack:
				isStack = true
				depth = getStackHeight(c.Val.(*Stack)) + 1
			}

			if !isStack {
				depth = 1
			}

		}

		if depth == 0 {
			depth = 1
		}

		return

	}
	stack.Height = getStackHeight(stack)
	for i := range stack.Cards {
		stack.Cards[i].Idx = i
	}
}

/** Prints some number of - outputs based on depth. */
func depthPrinter(depth int) (out string) {
	for i := 0; i < depth; i++ {
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
	var orderType, findType, findData, overrideFindData, overrideInsert, deepSearchType, depth, dereferenceType, passType, passCards, workingMem any
	gogenerics.UnpackVariadic(variadic, &orderType, &findType, &findData, &overrideFindData, &overrideInsert, &deepSearchType, &depth, &dereferenceType, &passType, &passCards, &workingMem)
	setOVERRIDEDefaultIfNil(&overrideInsert)
	setORDERDefaultIfNil(&orderType)
	setFINDDefaultIfNil(&findType)
	if workingMem == nil {workingMem = []any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}}
	if overrideFindData == nil {overrideFindData = COMPARE_False}
	if deepSearchType == nil {deepSearchType = DEEPSEARCH_False}
	if passType == nil {passType = PASS_Cards}

	// initialize foundCard to false so you can determine whether valid find and
	// initialize variables
	foundCard := false
	insertCards := []*Card {}

	// set up insertArr
	if overrideInsert == OVERRIDE_False {
		insertArr := []any {}
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
			insCard, insIsCard := ins.(*Card)
			if insIsCard {
				insertCards = append(insertCards, insCard) // insert a card whose val is ins
			} else {
				insertCards = append(insertCards, MakeCard(ins)) // insert a card whose val is ins
			}
		}

	} else {
		insertCards = append(insertCards, MakeCard(insert)) // insert a card whose val is whatever the input is
	}

	// add cards to empty stack
	if stack.Size == 0 {

		foundCard = true
		stack.Cards = append(stack.Cards, insertCards...)
		stack.setStackProperties()

	// else add based on find
	} else {
		stack.Lambda(func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any,  workingMem ...any) {
		
			// only do add to the first match if ACTION_First, otherwise do for every match
			if (allNotFirst && selectCard(findType, findData, dereferenceType, overrideFindData.(COMPARE), card, parentStack, isSubstack, coords, retStack, retCard, retVarAdr, workingMem...)) || (!allNotFirst && selectCard(findType, findData, dereferenceType, overrideFindData.(COMPARE), card, parentStack, isSubstack, coords, retStack, retCard, retVarAdr, workingMem...) && !foundCard) {
	
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
	
		}, nil, nil, nil, workingMem.([]any), deepSearchType, depth, passType, passCards)
	}

	// return nil if no add was made, else return card
	if !foundCard {
		return nil
	} else {
		return stack
	}

}

func callLambda(lambda any, card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, retCard *Card, retVarAdr any, otherInfo []any, workingMem []any) {
	
	conversion1, success := lambda.(func())
	if success {conversion1()}
	conversion2, success := lambda.(func(*Card))
	if success {conversion2(card)}
	conversion3, success := lambda.(func(*Card, *Stack))
	if success {conversion3(card, parentStack)}
	conversion4, success := lambda.(func(*Card, *Stack, bool))
	if success {conversion4(card, parentStack, isSubstack)}
	conversion5, success := lambda.(func(*Card, *Stack, bool, *Stack))
	if success {conversion5(card, parentStack, isSubstack, coords)}
	conversion6, success := lambda.(func(*Card, *Stack, bool, *Stack, *Stack))
	if success {conversion6(card, parentStack, isSubstack, coords, retStack)}
	conversion7, success := lambda.(func(*Card, *Stack, bool, *Stack, *Stack, *Card))
	if success {conversion7(card, parentStack, isSubstack, coords, retStack, retCard)}
	conversion8, success := lambda.(func(*Card, *Stack, bool, *Stack, *Stack, *Card, any))
	if success {conversion8(card, parentStack, isSubstack, coords, retStack, retCard, retVarAdr)}
	conversion9, success := lambda.(func(*Card, *Stack, bool, *Stack, *Stack, *Card, any, []any))
	if success {conversion9(card, parentStack, isSubstack, coords, retStack, retCard, retVarAdr, otherInfo)}
	conversion10, success := lambda.(func(*Card, *Stack, bool, *Stack, *Stack, *Card, any, []any, ...any))
	if success {conversion10(card, parentStack, isSubstack, coords, retStack, retCard, retVarAdr, otherInfo, workingMem...)}
}

func callLambdaReplaceWith(replaceWith any, card *Card, parentStack *Stack, isSubstack bool, coords *Stack, workingMem []any) {
	conversion1, success := replaceWith.(func())
	if success {conversion1()}
	conversion2, success := replaceWith.(func(*Card))
	if success {conversion2(card)}
	conversion3, success := replaceWith.(func(*Card, *Stack))
	if success {conversion3(card, parentStack)}
	conversion4, success := replaceWith.(func(*Card, *Stack, bool))
	if success {conversion4(card, parentStack, isSubstack)}
	conversion5, success := replaceWith.(func(*Card, *Stack, bool, *Stack))
	if success {conversion5(card, parentStack, isSubstack, coords)}
	conversion6, success := replaceWith.(func(*Card, *Stack, bool, *Stack, ...any))
	if success {conversion6(card, parentStack, isSubstack, coords, workingMem...)}
}
