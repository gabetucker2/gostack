package gostack

import (
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
		return needle == haystack
	case "slice":
		return needleInHaystack(needle, gogenerics.UnpackArray(haystack))
	case "stack":
		return needleInHaystack(needle, haystack.(*Stack).ToArray())
	default:
		return false
	}
}

func selectCard(findType any, findData any, pointerType any, findCompareRaw COMPARE, returnType string, card *Card, stack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, wmadrs ...any) bool {

	// set defaults
	setFINDDefaultIfNil(&findType)
	setPOINTERDefaultIfNil(&pointerType)

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
		return card.Idx == stack.Size - 1
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
				haystack[i] = stack.Size - 1
			}
		}
		return needleInHaystack(card.Idx, haystack)
	case FIND_Key:
		switch pointerType {
		case POINTER_True:
			return match(getPointer(card.Key, false), findData, override)
		case POINTER_False:
			return match(card.Key, findData, override)
		}
	case FIND_Val:
		switch pointerType {
		case POINTER_True:
			return match(getPointer(card.Val, false), findData, override)
		case POINTER_False:
			return match(card.Val, findData, override)
		}
	case FIND_Card:
		return match(card, findData, false)
	case FIND_Size:
		return match(card.Val.(*Stack).Size, findData, false)
	case FIND_Depth:
		return match(card.Val.(*Stack).Depth, findData, false)
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
			start = stack.Size - 1
		}
		if end == -1 {
			end = stack.Size - 1
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
		switch returnType {
		case "card":
			return findData.(func(*Card, *Stack, bool, ...any) (bool)) (card, stack, isSubstack, wmadrs...)
		case "stack":
			return findData.(func(*Card, *Stack, bool, *Stack, ...any) (bool)) (card, stack, isSubstack, retStack, wmadrs...)
		}
	}
	return false

}

/** Returns, from any map type, a version of that map which is converted to type deep map[any]any...

 @param `arr` type{any}
 @return type{[]any}
 @requires `arr` is an array
 */
/*func unpackDeepMapToKeysVals(input1 any, keys , vals []any) map[any]any {
    m := unpackMap(input1)
	for k, v := range m {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			unpackDeepMap(m)
		}
	}

	for k, v := range  { // TODO: deep unpack
		keys = append(keys, k)
		vals = append(vals, v)
	}
	return m
}*/

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

/** Updates a stack's Size, Depth, and Card Indices */
func (stack *Stack) setStackProperties() {
	stack.Size = len(stack.Cards)
	

	/** Assuming normally-shaped matrix, returns the depth of this stack */
	var getStackDepth func (stack *Stack) (depth int) // so that it can be recursive and nested
	getStackDepth = func (stack *Stack) (depth int) {
		
		if stack.Size > 0 {

			c := stack.Cards[0]

			isStack := false
			switch c.Val.(type) {
			case *Stack:
				isStack = true
				depth = getStackDepth(c.Val.(*Stack)) + 1
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
	stack.Depth = getStackDepth(stack)
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
