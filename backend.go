package gostack

import "reflect"

/** Returns a clone of this interface

@param `toClone` type{interface{}}
@returns type{interface{}}
*/
func cloneInterface(toClone interface{}) interface{} {
	return reflect.New(reflect.ValueOf(toClone).Elem().Type()).Interface()
}

/** Returns out1 if test is true; else return out2
 
 @param `test` type{bool}
 @param `out1` type{interface{}}
 @param `out2` type{interface{}}
 @returns interface{} `out1` or `out2`
 @requires neither param yields a syntax error
 */
func ifElse(test bool, out1, out2 interface{}) interface{} {
	if test { return out1 } else { return out2 }
}

/** Returns stack if testIfEmpty is not empty; else returns nil
 
 @receiver `stack` type{*Stack}
 @param `testIfEmpty` type{[]*Card}
 @returns `stack` or nil
 */
func (stack *Stack) returnNilIfEmpty(testIfEmpty []*Card) *Stack {
	if len(testIfEmpty) == 0 {
		return nil
	} else {
		return stack
	}
}

/** Sets a set of variables to the variable set passed into a variadic parameter

@param `variadic` type{...[]interface{}}
@param `var1, var2, ..., varN` type{any}
@updates `var1, var2, ..., varN` are set to each of the values in the variadic array, or nil if undefined, respectively
*/
func unpackVariadic(variadic []interface{}, into ...*interface{}) {
	vLen := len(variadic)
	for i, v := range into {
		if i < vLen {
			*v = variadic[i]
		} else {
			*v = nil
		}
	}
}

/** Removes the cards from a stack for which lambda(card) is false
 
 @param `stack` type{*Stack}
 @param `lambda` type{func(*Card, workingMemory) bool}
 @returns `stack`
 @updates `stack.Cards` to a new set of Cards filtered using `lambda`
 @ensures each card in `stack.Cards` will not be affected by lambda updates
 */
func getIterator(stack *Stack, lambda func(*Card, ...interface{}) bool) {
	var filteredCards []*Card
	for i := range stack.Cards {
		card := stack.Cards[i]
		if lambda(card.Clone(), stack) { // use a clone card
			filteredCards = append(filteredCards, card)
		}
	}
	stack.Cards = filteredCards
}

/** Passes each card into the lambda function iteratively
 
 @param `stack` type{*Stack}
 @param `lambda` type{func(*Card, ...workingMemory)}
 @updates `stack.Cards` to whatever the `lambda` function specifies
 */
func generalIterator(stack *Stack, lambda func(*Card, ...interface{})) {
	for i := range stack.Cards {
		// use the card object so that card can be updated by the lambda expression
		lambda(stack.Cards[i], stack)
	}
}

/** Passes each card into the lambda function iteratively
 
 @param `stack` type{*Stack}
 @param `lambda` type{func(*Card, *Stack, ...workingMemory) (ORDER, int)}
 @updates `stack.Cards` to whatever the `lambda` function specifies
 */
func sortIterator(stack *Stack, lambda func(*Card, *Stack, ...interface{}) (ORDER, int)) {
	for i := range stack.Cards {
		// iterate, get the new index from the sorter
		newOrder, newIdx := lambda(stack.Cards[i], stack)
		// move from the old position to the new position
		stack.Move(FIND_Idx, newOrder, FIND_Idx, i, newIdx)
	}
}

/** Returns an []int of indices representing the targeted position(s) in a stack
 
 @param `getFirst` type{bool}
 @param `stack` type{*Stack} no pass-by-reference
 @param `findType` type{FIND}
 @param `findData` type{interface{}}
 @returns the []int of targeted positions
 @constructor creates a new []int
 @requires
  * `MakeStack()` and `MakeCard()` have been implemented
  * Inputted `findData` is of expected type (see documentation on FIND) 
 @ensures
   IF search finds no cards in `stack`
     return []int {}
   
   IF `getFirst`
     return an array of the first found element
   ELSE
     return an array of all found elements
 */
func getPositions(getFirst bool, stack *Stack, findType FIND, findData interface{}, matchByType MATCHBY) (targets []int) {

	/** Returns a bool for whether the matchBy yields a true result */
	matchByObjectOrReference := func(x1, x2 interface{}) bool {
		return (matchByType == MATCHBY_Object    &&  x1 ==  x2) ||
			   (matchByType == MATCHBY_Reference && &x1 == &x2)
	}

	switch findType {

	case FIND_First:
		if stack.Size > 0 {targets = append(targets, 0)}

	case FIND_Last:
		if stack.Size > 0 {targets = append(targets, stack.Size - 1)}

	case FIND_Idx:
		thisIdx := findData.(int)
		if stack.Size > thisIdx {targets = append(targets, thisIdx)}

	case FIND_Idxs:
		theseIdxs := findData.([]int)
		for testI := range stack.Cards {
			for _, targetI := range theseIdxs {
				if testI == targetI {
					targets = append(targets, testI)
					if getFirst { break }
				}
			}
		}

	case FIND_IdxsStack:
		if getFirst {
			targets = append(targets, findData.(*Stack).Cards[0].Val.(int))
		} else {
			for _, c := range findData.(*Stack).Cards {
				targets = append(targets, c.Val.(int))
			}
		}

	case FIND_Key:
		for i := range stack.Cards {
			testKey := stack.Cards[i].Key
			if matchByObjectOrReference(testKey, findData) {
				targets = append(targets, i)
				if getFirst { break }
			}
		}

	case FIND_Keys:
		keyArray := findData.([]interface{})
		for i := range stack.Cards {
			testKey := stack.Cards[i].Key
			for j := range keyArray {
				targetKey := keyArray[j]
				if matchByObjectOrReference(testKey, targetKey) {
					targets = append(targets, i)
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
					targets = append(targets, i)
					if getFirst { break }
				}
			}
		}

	case FIND_Val:
		for i := range stack.Cards {
			testVal := stack.Cards[i].Val
			if matchByObjectOrReference(testVal, findData) {
				targets = append(targets, i)
				if getFirst { break }
			}
		}

	case FIND_Vals:
		valArray := findData.([]interface{})
		for i := range stack.Cards {
			testVal := stack.Cards[i].Val
			for j := range valArray {
				targetVal := valArray[j]
				if matchByObjectOrReference(testVal, targetVal) {
					targets = append(targets, i)
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
					targets = append(targets, i)
					if getFirst { break }
				}
			}
		}

	case FIND_Card:
		for i := range stack.Cards {
			testCard := stack.Cards[i]
			if matchByObjectOrReference(testCard, findData.(*Card)) {
				targets = append(targets, i)
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
					targets = append(targets, i)
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
					targets = append(targets, i)
					if getFirst { break }
				}
			}
		}

	case FIND_Slice:
		slice := findData.([2]int)
		if stack.Size > 0 && 0 <= slice[0] && 0 <= slice[1] && slice[0] < stack.Size && slice[1] < stack.Size {
			targets = append(targets, slice[0])
			if !getFirst {
				for i := 0; i < slice[1] - slice[0]; {
					targets = append(targets, i+slice[0])
					i = ifElse(slice[1] > slice[0], i+1, i-1).(int)
				}
			}
		}

	case FIND_All:
		for i := range stack.Cards {
			targets = append(targets, i)
		}

	case FIND_Lambda:
		filterStack := stack.Clone() // so that no changes can be made to the original stack from FIND_Lambda functions
		getIterator(filterStack, findData.(func(*Card, ...interface{}) bool))
		for i := range filterStack.Cards {
			targets = append(targets, i)	
			if getFirst { break }
		}

	}

	return

}

/** Updates a target's field or value to new values based on replaceType

 @param setStack type{*Stack}
 @param replaceType type{REPLACE}
 @param replaceData type{interface{}}
 @param target type{*Card}
 @updates `setStack` or `target`
 @ensures if `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 */
func updateRespectiveField(setStack *Stack, replaceType REPLACE, replaceData interface{}, target *Card) {

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
		generalIterator(setStack, replaceData.(func(*Card, ...interface{})))

	}

}

/** Recursively add elements from 1D array to stack of matrix shape resembling `matrixShape`
 
 @receiver stack type{*Stack}
 @param matrixShape type{[]int}
 @param keys type{interface{}}
 @param vals type{interface{}}
 @param globalI type{*int} used because: extracting from 1-D arrays into N-D matrix, so need to track our position in the 1-D arrays between different recursive calls
 @returns type{*Stack}
 @requires
  * `MakeStack()` and `MakeCard()` have been implemented
  * |keys| == |vals| if neither are nil
  * |keys| or |vals| == product of ints in matrixShape
*/
func (stack *Stack) makeStackMatrixFrom1D(matrixShape []int, keys interface{}, vals interface{}, globalI *int) (ret *Stack) {

	// make stack
	if len(matrixShape) > 1 {

		for i := 0; i < matrixShape[0]; i++ {
			// return new stack of stack ... of stack whose vals are nil
			ret := MakeStack().makeStackMatrixFrom1D(matrixShape[1:], keys, vals, globalI)
			// insert this return value into a card of our current stack
			stack.Cards = append(stack.Cards, MakeCard(ret, nil, i))
		}

	// no more stacks to make, insert elements into and return current stack
	} else {
		
		for i := 0; i < matrixShape[0]; i++ {
			c := MakeCard()
			if keys != nil {
				c.Key = keys.([]interface{})[*globalI]
			}
			if vals != nil {
				c.Val = vals.([]interface{})[*globalI]
			}
			if keys != nil || vals != nil {
				*globalI++
			}
			stack.Cards = append(stack.Cards, c)
		}
		ret = stack

	}

	// return
	return

}

/** Recursively add elements to stack of matrix shape resembling the inputs
 
 @receiver stack type{*Stack}
 @param keys type{interface{}}
 @param vals type{interface{}}
 @returns type{*Stack}
 @requires
  * `MakeStack()` and `MakeCard()` have been implemented
  * |keys| == |vals| if neither are nil
  * at least one of `keys` or `vals` are not nil
*/
func (stack *Stack) makeStackMatrixFromND(keys interface{}, vals interface{}) (ret *Stack) {

	// initialize variable to use as reference for the matrix shape
	// just because we don't know which input is not nil
	var referenceArr []interface{}
	// one of these conditions are guaranteed to be true per the ensures clause
	if keys != nil {
		referenceArr = keys.([]interface{})
	} else if vals != nil {
		referenceArr = vals.([]interface{})
	}
	
	// main loop
	for i := range referenceArr {
		switch referenceArr[i].(type) {

		// add substack to stack
		case []interface{}:
			stack.Cards = append(
				stack.Cards,
				MakeCard(MakeStack().makeStackMatrixFromND(
					ifElse(keys != nil, keys, nil).([]interface{}),
					ifElse(vals != nil, vals, nil).([]interface{}),
				)),
			)

		// add element to stack
		default:
			c := MakeCard()
			if keys != nil {
				c.Key = keys.([]interface{})[i]
			}
			if vals != nil {
				c.Val = vals.([]interface{})[i]
			}
			stack.Cards = append(stack.Cards, c)
		}
	}

	// return
	return stack

}
