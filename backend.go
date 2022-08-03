package gostack

import (
	"reflect"
)

/** Performs the function using a uniform framework for performing deepSearches

@shorthand Just pass the proper variables (or nil) into this function from Library.go, and this function will handle the rest
*/
func (stack *Stack) deepSearchHandler(callFrom string, getFirst bool, findType, findData, returnType, matchByType, deepSearchType, depth, typeType, uniqueType, insert, orderType, findData_to, findType_to, matchByType_to, cloneType1, cloneType2, cloneType3 any) (ret *Stack) {

	// 0) set defaults
	setORDERDefaultIfNil(&orderType)
	setFINDDefaultIfNil(&findType)
	setMATCHBYDefaultIfNil(&returnType)
	setMATCHBYDefaultIfNil(&matchByType)
	setDEEPSEARCHDefaultIfNil(&deepSearchType)
	setDepthDefaultIfNil(&depth)
	setCLONEDefaultIfNil(&cloneType1)
	setCLONEDefaultIfNil(&cloneType2)
	setCLONEDefaultIfNil(&cloneType3)

	// 1) clone the stack
	var stackClone *Stack
	if callFrom == "Get" || callFrom == "GetMany" {
		stackClone = new(Stack)
	} else {
		stackClone = stack.Clone()
	}

	// 2) get position data from clone
	targetIndices, targetCards, targetStacks := stackClone.getPositions(getFirst, findType.(FIND), findData, matchByType.(MATCHBY), deepSearchType.(DEEPSEARCH), depth.(int))
	
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
					if (typeType == TYPE_Card &&
						(matchByType == MATCHBY_Object && targetCard == newCard) ||
						(matchByType == MATCHBY_Reference && &targetCard == &newCard) ) || 
						(typeType == TYPE_Key &&
						(matchByType == MATCHBY_Object && targetCard.Key == newCard.Key) ||
						(matchByType == MATCHBY_Reference && &targetCard.Key == &newCard.Key) ) ||
						(typeType == TYPE_Val &&
						(matchByType == MATCHBY_Object && targetCard.Val == newCard.Val) ||
						(matchByType == MATCHBY_Reference && &targetCard.Val == &newCard.Val) ) {
							
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
					newCards = append(newCards, insert.(*Stack).Cards...)
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
						insertCard.Key = cloneInterface(insertCard.Key)
					}
					if cloneType3 == CLONE_True {
						insertCard.Val = cloneInterface(insertCard.Val)
					}

				case RETURN_Idxs:

					insertCard = new(Card)
					insertCard.Val = i
					if cloneType1 == CLONE_True {
						insertCard.Val = cloneInterface(insertCard.Val)
					}

				case RETURN_Keys:

					insertCard = new(Card)
					insertCard.Val = targetCard.Key
					if cloneType1 == CLONE_True {
						insertCard.Val = cloneInterface(insertCard.Val)
					}

				case RETURN_Vals:

					insertCard = new(Card)
					insertCard.Val = targetCard.Val
					if cloneType1 == CLONE_True {
						insertCard.Val = cloneInterface(insertCard.Val)
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

/** Returns a clone of this interface

@param `toClone` type{any}
@returns type{any}
*/
func cloneInterface(toClone any) any {
	return reflect.New(reflect.ValueOf(toClone).Elem().Type()).Interface()
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

/** Returns out1 if test is true; else return out2
 
 @param `test` type{bool}
 @param `out1` type{any}
 @param `out2` type{any}
 @returns any `out1` or `out2`
 @requires neither param yields a syntax error
 */
func ifElse(test bool, out1, out2 any) any {
	if test { return out1 } else { return out2 }
}

/** Sets a set of variables to the variable set passed into a variadic parameter

 @param `variadic` type{...[]any}
 @param `var1, var2, ..., varN` type{any}
 @updates `var1, var2, ..., varN` are set to each of the values in the variadic array, or nil if undefined, respectively
 */
func unpackVariadic(variadic []any, into ...*any) {
	vLen := len(variadic)
	for i, v := range into {
		if i < vLen {
			*v = variadic[i]
		} else {
			*v = nil
		}
	}
}

/** Removes the cards from a stack for which lambda(card) is false, updating to a new 1D stack
 
 @param `stack` type{*Stack}
 @param `lambda` type{func(*Card, workingMemory) bool}
 @param `deepSearchType` type{DEEPSEARCH}
 @param `depth` type{int}
 @returns `stack`
 @updates `stack.Cards` to a new set of Cards filtered using `lambda`
 @ensures each card in `stack.Cards` will not be affected by lambda updates
 @requires `stack.GetMany()` is implemented
 */
func getIterator(stack *Stack, lambda func(*Card, ...any) bool, deepSearchType DEEPSEARCH, depth int) {
	subStack := stack.GetMany(FIND_All, nil, nil, nil, nil, nil, nil, deepSearchType, depth)
	var filteredCards []*Card
	for i := range subStack.Cards {
		card := subStack.Cards[i]
		if lambda(card.Clone(), subStack) { // use a clone card
			filteredCards = append(filteredCards, card)
		}
	}
	stack.Cards = filteredCards
	stack.setStackProperties()

}

/** Passes each card into the lambda function iteratively
 
 @param `stack` type{*Stack}
 @param `lambda` type{func(*Card, ...workingMemory)}
 @param `deepSearchType` type{DEEPSEARCH}
 @param `depth` type{int}
 @updates `stack.Cards` to whatever the `lambda` function specifies
 @requires `stack.GetMany()` is implemented
 */
func generalIterator(stack *Stack, lambda func(*Card, ...any), deepSearchType DEEPSEARCH, depth int, otherStacks ...*Stack) {
	subStack := stack.GetMany(FIND_All, nil, nil, nil, nil, nil, nil, deepSearchType, depth)
	for i := range subStack.Cards {
		// use the card object so that card can be updated by the lambda expression
		lambda(subStack.Cards[i], subStack)
		subStack.setStackProperties()
	}
	subStack.setStackProperties()
}

/** Passes each card into the lambda function iteratively, updating to a new 1D stack
 
 @param `stack` type{*Stack}
 @param `lambda` type{func(*Card, *Stack, ...workingMemory) (ORDER, int)}
 @param `deepSearchType` type{DEEPSEARCH}
 @param `depth` type{int}
 @updates `stack.Cards` to whatever the `lambda` function specifies
 */
func sortIterator(stack *Stack, lambda func(*Card, *Stack, ...any) (ORDER, int), deepSearchType DEEPSEARCH, depth int) {
	subStack := stack.GetMany(FIND_All, nil, nil, nil, nil, nil, nil, deepSearchType, depth)
	for i := range subStack.Cards {
		// iterate, get the new index from the sorter
		newOrder, newIdx := lambda(subStack.Cards[i], subStack)
		// move from the old position to the new position
		subStack.Move(FIND_Idx, newOrder, FIND_Idx, i, newIdx)
	}
	stack.setStackProperties()
	
}

/** Returns an [][]int of index vertices representing the order of indices needed to access targeted position(s) in `stack`, with []*Card for pure card values
 
 @param `getFirst` type{bool}
 @param `stack` type{*Stack} no pass-by-reference
 @param `findType` type{FIND}
 @param `findData` type{any}
 @param `matchByType` type{MATCHBY} no pass-by-reference
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
func (stack *Stack) getPositions(getFirst bool, findType FIND, findData any, matchByType MATCHBY, deepSearchType DEEPSEARCH, depth int) (targetIdxs [][]int, targetCards []*Card, targetStacks []*Stack) {

	/** Returns a bool for whether the matchBy yields a true result */
	matchByObjectOrReference := func(x1, x2 any) bool {
		return (matchByType == MATCHBY_Object    &&  x1 ==  x2) ||
			   (matchByType == MATCHBY_Reference && &x1 == &x2)
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
				if matchByObjectOrReference(testKey, findData) {
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
					if matchByObjectOrReference(testKey, targetKey) {
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
					if matchByObjectOrReference(testKey, targetKey) {
						filteredList = append(filteredList, i)
						if getFirst { break }
					}
				}
			}
	
		case FIND_Val:
			for i := range stack.Cards {
				testVal := stack.Cards[i].Val
				if matchByObjectOrReference(testVal, findData) {
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
					if matchByObjectOrReference(testVal, targetVal) {
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
					if matchByObjectOrReference(testVal, targetVal) {
						filteredList = append(filteredList, i)
						if getFirst { break }
					}
				}
			}
	
		case FIND_Card:
			for i := range stack.Cards {
				testCard := stack.Cards[i]
				if matchByObjectOrReference(testCard, findData.(*Card)) {
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
					if matchByObjectOrReference(testCard, targetCard) {
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
					if matchByObjectOrReference(testCard, targetCard) {
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
						i = ifElse(slice[1] > slice[0], i+1, i-1).(int)
					}
				}
			}
	
		case FIND_All:
			for i := range stack.Cards {
				filteredList = append(filteredList, i)
			}
	
		case FIND_Lambda:
			filterStack := stack.Clone() // so that no changes can be made to the original stack from FIND_Lambda functions
			getIterator(filterStack, findData.(func(*Card, ...any) bool), deepSearchType, depth)
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
		generalIterator(setStack, replaceData.(func(*Card, ...any)), DEEPSEARCH_False, -1)

	}

}

/** Returns, from any array type, a version of that array which is converted to type []any

 @param `arr` type{any}
 @return type{[]any}
 @requires `arr` is an array
 */
func unpackArray(arr any) []any {
    valType := reflect.ValueOf(arr)
    new := make([]any, valType.Len())
    for i := 0; i < valType.Len(); i++ {
        new[i] = valType.Index(i).Interface()
    }
    return new
}

/** Returns, from any map type, a version of that map which is converted to type map[any]any

 @param `arr` type{any}
 @return type{[]any}
 @requires `arr` is an array
 */
 func unpackMap(s any) map[any]any {
	v := reflect.ValueOf(s)
    m := make(map[any]any, v.Len())
    for _, k := range v.MapKeys() {
		m[k.Interface()] = v.MapIndex(k).Interface()
    }
    return m
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
		
		for i := 0; i < matrixShape[0]; i++ {
			c := MakeCard()
			updated := false
			if len(keys) > 0 {
				updated = true
				c.Key = keys[*globalI]
			}
			if len(vals) > 0 {
				updated = true
				c.Val = vals[*globalI]
			}
			if updated {
				*globalI++
			}
			ret.Cards = append(ret.Cards, c)
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
		referenceArr = keys.([]any)
	} else if vals != nil {
		referenceArr = vals.([]any)
	}
	
	// main loop
	for i := range referenceArr {
		switch referenceArr[i].(type) {

		// add substack to stack
		case []any:
			stack.Cards = append(
				stack.Cards,
				MakeCard(MakeStack().makeStackMatrixFromND(
					ifElse(keys != nil, keys, nil).([]any),
					ifElse(vals != nil, vals, nil).([]any),
				)),
			)

		// add element to stack
		default:
			c := MakeCard()
			if keys != nil {
				c.Key = keys.([]any)[i]
			}
			if vals != nil {
				c.Val = vals.([]any)[i]
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
