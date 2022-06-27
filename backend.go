package gostack

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
 
 @param `stack` type{Stack}
 @param `lambda` type{func(*Stack, *Card) bool}
 @returns `stack`
 @updates `stack.Cards` to a new set of Cards filtered using `lambda`
 */
func iterator(stack *Stack, lambda func(*Card, ...interface{}) bool) {
	var filteredCards []*Card
	for _, card := range stack.Cards {
		if lambda(card, stack) {
			filteredCards = append(filteredCards, card)
		}
	}
	stack.Cards = filteredCards
}

/** Returns an []int of indices representing the targeted position(s) in a stack
 
 @param `getFirst` type{bool}
 @param `stack` type{Stack} no pass-by-reference
 @param `positionType` type{POSITION}
 @param `positionData` type{interface{}}
 @returns the []int of targeted positions
 @constructor creates a new []int
 @ensures
   IF `positionType` is singular
     return idx/idxs of cards whose field matches `positionData` field
   ELSE IF `positionType` is plural
	 return idx/idxs of cards whose field matches any of `positionData` fields
   
   IF `getFirst`
     return idx
   ELSE
     return idxs
 */
func getPositions(getFirst bool, stack *Stack, positionType POSITION, positionData interface{}, matchType MATCH) (targets []int) {

	switch positionType {

	case POSITION_First:
		targets = append(targets, 0)

		//... and so on

	case POSITION_Keys:
		keyArr := positionData.([]interface{})
		for i := 0; i < len(keyArr); i++ {
			for j, c := range stack.Cards {
				if (matchType == MATCH_Object    &&  keyArr[i] ==  c.Key) ||
				   (matchType == MATCH_Reference && &keyArr[i] == &c.Key) {
					targets = append(targets, j)
					if getFirst { break }
				}
			}
		}

	//... and so on

	case POSITION_Lambda:
		filterStack := stack.Clone()
		iterator(filterStack, positionData.(func(*Card, ...interface{}) bool))
		for i := range filterStack.Cards {
			targets = append(targets, i)	
			if getFirst { break }
		}

	}

	return

}

/** Returns a new card `newCard` whose value is the specified field of `oldCard` specified by `returnType`
 
 @param `newCard` type{Card}
 @param `oldCard` type{Card}
 @param `returnType` type{RETURN}
 @updates the `newCard` value to `oldCard` field defined by `returnType`
 */
func setCardVal(newCard *Card, oldCard *Card, returnType RETURN) {

	switch returnType {

	case RETURN_Idxs:
		newCard.Val = oldCard.Idx

	case RETURN_Keys:
		newCard.Val = oldCard.Key

	case RETURN_Vals:
		newCard.Val = oldCard.Val

	case RETURN_Cards:
		newCard.Val = *oldCard

	}

}
