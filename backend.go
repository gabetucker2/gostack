package gostack

import (
	"reflect"

	"github.com/gabetucker2/gogenerics"
)

func selectCard(findType any, findData any, pointerType any, card *Card, stack *Stack, isSubstack bool, retStack *Stack, retCard *Card, retVarAdr any, wmadrs ...any) bool {

	// set defaults
	setFINDDefaultIfNil(&findType)
	setPOINTERDefaultIfNil(&pointerType)

	// local functions
	needleInHaystack := func(needle any, haystack []any) bool {
		for _, hay := range haystack {
			if needle == hay {
				return true
			}
		}
		return false
	}

	getType := func(in any) string {
		if in == nil {
			return "element"
		} else {
			_, isStack := in.(*Stack)
			if isStack {
				return "stack"
			} else {
				isArr := reflect.TypeOf(in).Kind() == reflect.Slice
				if isArr {
					return "array"
				} else {
					return "element"
				}
			}
		}
	}

	match := func(needle any, haystack any) bool {
		switch getType(findData) {
		case "element":
			return needle == haystack
		case "array":
			return needleInHaystack(needle, gogenerics.UnpackArray(haystack))
		case "stack":
			return needleInHaystack(needle, gogenerics.UnpackArray(haystack.(*Stack).ToArray()))
		default:
			return false
		}
	}

	// main
	switch findType {
	case FIND_First:
		return card.Idx == 0
	case FIND_Last:
		return card.Idx == stack.Size - 1
	case FIND_Idx:
		return match(card.Idx, findData)
	case FIND_Key:
		switch pointerType {
		case POINTER_True:
			return match(gogenerics.GetPointer(card.Key, false), findData)
		case POINTER_False:
			return match(card.Key, findData)
		}
	case FIND_Val:
		switch pointerType {
		case POINTER_True:
			return match(gogenerics.GetPointer(card.Val, false), findData)
		case POINTER_False:
			return match(card.Val, findData)
		}
	case FIND_Card:
		return match(card, findData)
	case FIND_Stack:
		return match(card.Val.(*Stack), findData)
	case FIND_Size:
		return match(card.Val.(*Stack).Size, findData)
	case FIND_Depth:
		return match(card.Val.(*Stack).Depth, findData)
	case FIND_Slice: // [inclusive, exclusive)
		test := false
		var slice []any
		switch getType(findData) {
		case "array":
			slice = gogenerics.UnpackArray(findData)
		case "stack":
			slice = findData.(*Stack).ToArray()
		}
		start := slice[0].(int)
		end := slice[1].(int)
		slice = []any {}
		if start < end {
			for i := start; i < end; i++ {
				slice = append(slice, i)
			}
		} else if start > end {
			for i := end - 1; i >= start; i++ {
				slice = append(slice, i)
			}
		}
		for _, idx := range slice {
			test = match(card.Idx, idx)
			if !test {break}
		}
		return test
	case FIND_All:
		return true
	case FIND_Lambda:
		return findData.(func(*Card, *Stack, bool, *Stack, *Card, any, ...any) (bool)) (card, stack, isSubstack, retStack, retCard, &retVarAdr, wmadrs...)
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
 @returns type{*Stack}
 @requires
  * `MakeStack()` and `MakeCard()` have been implemented
  * |keys| == |vals| if neither are nil
  * |keys| or |vals| == product of ints in matrixShape
*/
func (stack *Stack) makeStackMatrixFrom1D(matrixShape []int, keys []any, vals []any, globalI *int, overrideCards any) (ret *Stack) {
	
	// set defaults
	if overrideCards == nil {overrideCards = false}

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

			if overrideCards.(bool) {
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

/** Sets the card value to `to`'s address if it's already a pointer'; else, just set the card value to `to`. */
func setCardProps(c *Card, to any, valNotKey bool) {
	if valNotKey {
		if gogenerics.SetPointer(c.Val, to) == nil {
			c.Val = to
		}
	} else {
		if gogenerics.SetPointer(c.Key, to) == nil {
			c.Key = to
		}
	}
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
