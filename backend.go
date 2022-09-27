package gostack

import (
	"reflect"
	"fmt"

	"github.com/gabetucker2/gogenerics"
)

/** Performs the function using a uniform framework for performing deepSearches

@shorthand Just pass the proper variables (or nil) into this function from Library.go, and this function will handle the rest
*/
func (stack *Stack) deepSearchHandler(callFrom string, getFirst bool, findType, findData, returnType, pointerType, deepSearchType, depth, typeType, uniqueType, insert, orderType, findData_to, findType_to, pointerType_to, cloneType1, cloneType2, cloneType3, overrideStackConversion any) (ret *Stack) {

	// 0) set defaults
	setORDERDefaultIfNil(&orderType)
	setFINDDefaultIfNil(&findType)
	setPOINTERDefaultIfNil(&returnType)
	setPOINTERDefaultIfNil(&pointerType)
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setDepthDefaultIfNil(&depth)
	setCLONEDefaultIfNil(&cloneType1)
	setCLONEDefaultIfNil(&cloneType2)
	setCLONEDefaultIfNil(&cloneType3)
	if overrideStackConversion == nil {overrideStackConversion = false}

	// 1) clone the stack
	var stackClone *Stack
	if callFrom == "Get" || callFrom == "GetMany" {
		stackClone = new(Stack)
	} else {
		stackClone = stack.Clone()
	}

	// 2) get position data from clone
	targetIndices, targetCards, targetStacks := stackClone.getPositions(getFirst, findType.(FIND), findData, pointerType.(POINTER), deepSearchType.(DEEPSEARCH), depth.(int))
	
	// 3) iterate through each card in targetCards
	if !(getFirst && len(targetCards) == 0) {
		for i := range targetCards {
			// 4) perform function on found card contingent on the caller function type, treating stackClone or targetStack (within stackClone) as the output in this function

			currentIdxSet := targetIndices[i] // current set of indices to get to target from stackClone
			targetStack := targetStacks[i] // parent stacks of each target
			var newCards []*Card // set of cards with which to replace targetStack.Cards (original stack)
			targetLocalIdx := currentIdxSet[len(currentIdxSet)-1]
			targetCard := targetStack.Cards[targetLocalIdx]

			switch callFrom {
			case "Unique":

				// where newCards is uniqueCards
				newCards = targetStack.Cards
				for i, newCard := range newCards {
					if (typeType == TYPE_Key &&
						(pointerType == POINTER_False && targetCard.Key == newCard.Key) ||
						(pointerType == POINTER_True && &targetCard.Key == &newCard.Key) ) ||
						(typeType == TYPE_Val &&
						(pointerType == POINTER_False && targetCard.Val == newCard.Val) ||
						(pointerType == POINTER_True && &targetCard.Val == &newCard.Val) ) {
							
							// target already exists in the card array, so remove it from the output card array
							removeIdx(newCards, i)
							break

					}
				}

				// set the local stack to the new stack after setting newCards
				targetStack.Cards = newCards

			case "Add":

				//// CARDS < targetLocalIdx
				// add the cards before targetCard
				for j := 0; j < targetLocalIdx; j++ {
					newCards = append(newCards, targetStack.Cards[j])
				}

				//// CARDS == targetLocalIdx
				// add the targetCard before insert if insert is Order_AFTER (insert ordered after targetCard)
				if orderType == ORDER_After { newCards = append(newCards, targetCard) }

				// add insert Card(s) before/after targetCard
				switch insert.(type) {
				case Card:
					newCards = append(newCards, insert.(*Card))
				case Stack:
					if overrideStackConversion.(bool) {
						newCards = append(newCards, MakeCard(insert.(*Stack)))
					} else {
						newCards = append(newCards, insert.(*Stack).Cards...)
					}
				}

				// add the targetCard after insert if insert is Order_BEFORE (insert ordered before targetCard)
				if orderType == ORDER_Before { newCards = append(newCards, targetCard) }

				//// CARDS > targetLocalIdx
				// add the cards after targetCard
				for j := targetLocalIdx+1; j < len(targetStack.Cards); j++ {
					newCards = append(newCards, targetStack.Cards[j])
				}

				// set the local stack to the new stack after setting newCards
				targetStack.Cards = newCards

			case "Get":

				var insertCard *Card

				switch returnType {
				case RETURN_Cards:

					// card which we will transform (if necessary) to insert
					insertCard = targetCard

					// clone if necessary
					if cloneType1 == CLONE_True {
						insertCard = insertCard.Clone()
					}
					if cloneType2 == CLONE_True {
						insertCard.Key = gogenerics.CloneInterface(insertCard.Key)
					}
					if cloneType3 == CLONE_True {
						insertCard.Val = gogenerics.CloneInterface(insertCard.Val)
					}

				case RETURN_Idxs:

					insertCard = new(Card)
					insertCard.Val = i
					if cloneType1 == CLONE_True {
						insertCard.Val = gogenerics.CloneInterface(insertCard.Val)
					}

				case RETURN_Keys:

					insertCard = new(Card)
					insertCard.Val = targetCard.Key
					if cloneType1 == CLONE_True {
						insertCard.Val = gogenerics.CloneInterface(insertCard.Val)
					}

				case RETURN_Vals:

					insertCard = new(Card)
					insertCard.Val = targetCard.Val
					if cloneType1 == CLONE_True {
						insertCard.Val = gogenerics.CloneInterface(insertCard.Val)
					}

				}

				// get targeted card OR nil
				stackClone.Cards = append(stackClone.Cards, insertCard)
				
			}

		}
		
		// finalize stackClone in preparation for return
		stackClone.setStackProperties()
		ret = stackClone

	} else {
		ret = nil
	}
	
	// 5) return nil if performing function on one card and failed to find any targets on which to perform that function, else return stackClone
	return

}

/** Sets every card's index in an array to a new index

 @param `cards` type{[]*Card}
 @updates `cards`
 */
func setIndices(cards []*Card) {
	for i := range cards {
		cards[i].Idx = i
	}
}

/** Removes an element at the index within an array (only works for cards)

 @param `arr` type{[]*Card{}}
 @param `idx` type{int}
 @returns new arr
 @constructs new arr
*/
func removeIdx(arr []*Card, idx int) []*Card {
	var newArr []*Card
	for i := range arr {
		if i != idx {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}



/** Removes the cards from a stack for which lambda(card) is false, updating to a new 1D stack
 
 @param `stack` type{*Stack}
 @param `lambda` type{func(*Card, *Stack, workingMem) bool}
 @param `deepSearchType` type{DEEPSEARCH}
 @param `depth` type{int}
 @returns `stack`
 @updates `stack.Cards` to a new set of Cards filtered using `lambda`
 @ensures each card in `stack.Cards` will not be affected by lambda updates
 @requires `stack.GetMany()` is implemented
 */
func getIterator(stack *Stack, lambda func(*Card, *Stack, ...any) bool, deepSearchType DEEPSEARCH, depth int) {
	subStack := stack.GetMany(FIND_All, nil, nil, nil, nil, nil, nil, deepSearchType, depth)
	var filteredCards []*Card
	for i := range subStack.Cards {
		card := subStack.Cards[i]
		if lambda(card.Clone(), subStack.Clone(CLONE_True, CLONE_True, CLONE_True)) { // use a clone card and stack
			filteredCards = append(filteredCards, card)
		}
	}
	stack.Cards = filteredCards
	stack.setStackProperties()

}

/** Passes each card into the lambda function iteratively
 
 @param `stack` type{*Stack}
 @param `lambda` type{func(*Card, *Stack, any, ...workingMem)}
 @param `deepSearchType` type{DEEPSEARCH}
 @param `depth` type{int}
 @updates `stack.Cards` to whatever the `lambda` function specifies
 @requires `stack.GetMany()` is implemented
 */
func generalIterator(stack *Stack, lambda func(*Card, *Stack, any, ...any), deepSearchType DEEPSEARCH, depth int, ret any, workingMem ...any) {
	subStack := stack.GetMany(FIND_All, nil, nil, nil, nil, nil, nil, deepSearchType, depth)
	for i := range subStack.Cards {
		// use the card object so that card can be updated by the lambda expression
		lambda(subStack.Cards[i], subStack, ret, workingMem...)
		subStack.setStackProperties()
	}
	subStack.setStackProperties()
}

/** Returns an [][]int of index vertices representing the order of indices needed to access targeted position(s) in `stack`, with []*Card for pure card values
 
 @param `getFirst` type{bool}
 @param `stack` type{*Stack} no pass-by-reference
 @param `findType` type{FIND}
 @param `findData` type{any}
 @param `pointerType` type{POINTER} no pass-by-reference
 @param `deepSearchType` type{DEEPSEARCH}
 @param `depth` type{int}
 @returns 3 arrays of data pertaining to the found cards:
  * type{[][]int} int array representing path to card from root stack
  * type{[]*Card} the card pointer itself
  * type{[]*Stack} the stack which is the direct parent of that card
 @constructor creates a new []int
 @requires
  * `MakeStack()` and `MakeCard()` have been implemented
  * Inputted `findData` is of expected type (see documentation on FIND) 
 @ensures
   IF search finds no cards in `stack`
     return [][]int {}
   
   IF `getFirst`
     return an array of the first found element
   ELSE
     return an array of all found elements
 */
func (stack *Stack) getPositions(getFirst bool, findType FIND, findData any, pointerType POINTER, deepSearchType DEEPSEARCH, depth int) (targetIdxs [][]int, targetCards []*Card, targetStacks []*Stack) {

	/** Returns a bool for whether the pointer yields a true result */
	comparePointer := func(x1, x2 any) bool {
		return (pointerType == POINTER_False    &&  x1 ==  x2) ||
			   (pointerType == POINTER_True && &x1 == &x2)
	}

	// setup main by deepening iteratively
	if deepSearchType == DEEPSEARCH_False { depth = 1 }
	workingCards := []*Card {}
	currentCards := []*Card {}
	for i := 0; i < depth; i++ {

		// first iteration
		if i == 0 {
			// fill first layer with ints representing indices
			for j := range stack.Cards {
				currentCards = stack.Cards
				targetIdxs = append(targetIdxs, []int{j})
			}

		// next iterations
		} else {
			for j, indexList := range targetIdxs {
				c := currentCards[indexList[i]]
				// if there is another stack within this stack, deepen
				switch c.Val.(type) {
				case *Stack:
					workingCards = append(workingCards, c)
					targetIdxs[j] = append(targetIdxs[j], j)
				}
			}
			currentCards = workingCards
		}
	}

	// main
	for i := range targetIdxs {
		filteredList := []int{}
		//subStack
		targetIdxs = append(targetIdxs, filteredList)

		switch findType {

		case FIND_First:
			if stack.Size > 0 {filteredList = append(filteredList, 0)}
	
		case FIND_Last:
			if stack.Size > 0 {filteredList = append(filteredList, stack.Size - 1)}
	
		case FIND_Idx:
			thisIdx := findData.(int)
			if stack.Size > thisIdx {filteredList = append(filteredList, thisIdx)}
	
		case FIND_Idxs:
			theseIdxs := findData.([]int)
			for testI := range stack.Cards {
				for _, targetI := range theseIdxs {
					if testI == targetI {
						filteredList = append(filteredList, testI)
						if getFirst { break }
					}
				}
			}
	
		case FIND_IdxsStack:
			if getFirst {
				filteredList = append(filteredList, findData.(*Stack).Cards[0].Val.(int))
			} else {
				for _, c := range findData.(*Stack).Cards {
					filteredList = append(filteredList, c.Val.(int))
				}
			}
	
		case FIND_Key:
			for i := range stack.Cards {
				testKey := stack.Cards[i].Key
				if comparePointer(testKey, findData) {
					filteredList = append(filteredList, i)
					if getFirst { break }
				}
			}
	
		case FIND_Keys:
			keyArray := findData.([]any)
			for i := range stack.Cards {
				testKey := stack.Cards[i].Key
				for j := range keyArray {
					targetKey := keyArray[j]
					if comparePointer(testKey, targetKey) {
						filteredList = append(filteredList, i)
						if getFirst { break }
					}
				}
			}
	
		case FIND_KeysStack:
			keyStack := findData.(*Stack)
			for i := range stack.Cards {
				testKey := stack.Cards[i].Key
				for j := range keyStack.Cards {
					targetKey := keyStack.Cards[j].Val
					if comparePointer(testKey, targetKey) {
						filteredList = append(filteredList, i)
						if getFirst { break }
					}
				}
			}
	
		case FIND_Val:
			for i := range stack.Cards {
				testVal := stack.Cards[i].Val
				if comparePointer(testVal, findData) {
					filteredList = append(filteredList, i)
					if getFirst { break }
				}
			}
	
		case FIND_Vals:
			valArray := findData.([]any)
			for i := range stack.Cards {
				testVal := stack.Cards[i].Val
				for j := range valArray {
					targetVal := valArray[j]
					if comparePointer(testVal, targetVal) {
						filteredList = append(filteredList, i)
						if getFirst { break }
					}
				}
			}
	
		case FIND_ValsStack:
			valStack := findData.(*Stack)
			for i := range stack.Cards {
				testVal := stack.Cards[i].Val
				for j := range valStack.Cards {
					targetVal := valStack.Cards[j].Val
					if comparePointer(testVal, targetVal) {
						filteredList = append(filteredList, i)
						if getFirst { break }
					}
				}
			}
	
		case FIND_Card:
			for i := range stack.Cards {
				testCard := stack.Cards[i]
				if comparePointer(testCard, findData.(*Card)) {
					filteredList = append(filteredList, i)
					if getFirst { break }
				}
			}
	
		case FIND_Cards:
			cardStack := findData.(*Stack)
			for i := range stack.Cards {
				testCard := stack.Cards[i]
				for j := range cardStack.Cards {
					targetCard := cardStack.Cards[j]
					if comparePointer(testCard, targetCard) {
						filteredList = append(filteredList, i)
						if getFirst { break }
					}
				}
			}
	
		case FIND_CardsStack:
			cardStack := findData.(*Stack)
			for i := range stack.Cards {
				testCard := stack.Cards[i]
				for j := range cardStack.Cards {
					targetCard := cardStack.Cards[j].Val
					if comparePointer(testCard, targetCard) {
						filteredList = append(filteredList, i)
						if getFirst { break }
					}
				}
			}
	
		case FIND_Slice:
			slice := findData.([2]int)
			if stack.Size > 0 && 0 <= slice[0] && 0 <= slice[1] && slice[0] < stack.Size && slice[1] < stack.Size {
				filteredList = append(filteredList, slice[0])
				if !getFirst {
					for i := 0; i < slice[1] - slice[0]; {
						filteredList = append(filteredList, i+slice[0])
						i = gogenerics.IfElse(slice[1] > slice[0], i+1, i-1).(int)
					}
				}
			}
	
		case FIND_All:
			for i := range stack.Cards {
				filteredList = append(filteredList, i)
			}
	
		case FIND_Lambda:
			filterStack := stack.Clone() // so that no changes can be made to the original stack from FIND_Lambda functions
			getIterator(filterStack, findData.(func(*Card, *Stack, ...any) bool), deepSearchType, depth)
			for i := range filterStack.Cards {
				filteredList = append(filteredList, i)	
				if getFirst { break }
			}
	
		}

		targetIdxs[i] = filteredList

		for _, indexList := range targetIdxs {
			substack := stack
			target := substack.Cards[indexList[0]]
			for j := 1; j < len(indexList); j++ {
				substack = target.Val.(*Stack)
				target = substack.Cards[j]
			}
			targetCards = append(targetCards, target)
			targetStacks = append(targetStacks, substack)
		}

	}

	return

}

/** Updates a target's field or value to new values based on replaceType

 @param setStack type{*Stack}
 @param replaceType type{REPLACE}
 @param replaceData type{any}
 @param target type{*Card}
 @updates `setStack` or `target`
 @ensures if `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 */
func (setStack *Stack) updateRespectiveField(replaceType REPLACE, replaceData any, target *Card) {

	switch replaceType {

	case REPLACE_Key:
		target.Key = replaceData

	case REPLACE_Val:
		target.Val = replaceData

	case REPLACE_Card:
		if replaceData == nil {
			// remove
			var newCards []*Card
			for i := range setStack.Cards {
				c := setStack.Cards[i]
				if c != target {
					newCards = append(newCards, c)
				}
			}
			setStack.Cards = newCards
		} else {
			*target = replaceData.(Card)
		}

	case REPLACE_Stack:
		// replace with new set of cards
		var newCards []*Card
		for i := range setStack.Cards {
			c := setStack.Cards[i]
			if c != target {
				newCards = append(newCards, c)
			} else {
				cardsIn := replaceData.(*Stack).Cards
				for j := range cardsIn {
					newCards = append(newCards, cardsIn[j])
				}
			}
		}
		setStack.Cards = newCards

	case REPLACE_Lambda:
		 // DEEPSEARCH_False since targeting cards that have already been filtered using Get()
		generalIterator(setStack, replaceData.(func(*Card, *Stack, any, ...any)), DEEPSEARCH_False, -1, nil)

	}

}

// TODO: FIX LATER
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

/** Assuming normally-shaped matrix, returns the depth of this stack */
func (stack *Stack) getStackDepth() (depth int) {
	
	if stack.Size > 0 {

		c := stack.Cards[0]

		isStack := false
		switch c.Val.(type) {
		case *Stack:
			isStack = true
			depth = c.Val.(*Stack).getStackDepth() + 1
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

/** Updates a stack's Size, Depth, and Card Indices */
func (stack *Stack) setStackProperties() {
	stack.Size = len(stack.Cards)
	stack.Depth = stack.getStackDepth()
	setIndices(stack.Cards)
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

func unpackDeepMapToKeysVals(input1 *any, keys, vals []any) {

    fmt.Println(reflect.TypeOf(*input1).Key().Kind())
    fmt.Println(reflect.TypeOf(*input1).Elem().Kind())

	for k, v := range gogenerics.UnpackMap(input1) {
		switch reflect.TypeOf(*input1).Elem().Kind() {
		case reflect.Map:
			newKArr := []any{}
			newVArr := []any{}
			keys = append(keys, newKArr)
			vals = append(keys, newVArr)
			unpackDeepMapToKeysVals(, newKArr, newVArr)
	
		default:
			keys = append(keys, k)
			vals = append(vals, v)
	
		}
	}

}
