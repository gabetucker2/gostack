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
 @ensures each card in `stack.Cards` will not be affected by lambda updates
 */
func getIterator(stack *Stack, lambda func(*Card, ...interface{}) bool) {
	var filteredCards []*Card
	for i := range stack.Cards {
		card := stack.Cards[i]
		if lambda(card.Clone(), stack) {
			filteredCards = append(filteredCards, card)
		}
	}
	stack.Cards = filteredCards
}

/** Passes each card into the lambda function iteratively
 
 @param `stack` type{Stack}
 @param `lambda` type{func(*Stack, *Card)}
 @updates `stack.Cards` to whatever the `lambda` function specifies
 */
func setIterator(stack *Stack, lambda func(*Card, ...interface{})) {
	for i := range stack.Cards {
		// use the original iterator so that card can be updated by the lambda expression
		lambda(stack.Cards[i], stack)
	}
}

/** Returns an []int of indices representing the targeted position(s) in a stack
 
 @param `getFirst` type{bool}
 @param `stack` type{Stack} no pass-by-reference
 @param `findByType` type{FINDBY}
 @param `findByData` type{interface{}}
 @returns the []int of targeted positions
 @constructor creates a new []int
 @ensures
   IF `findByType` is singular
     return idx/idxs of cards whose field matches `findByData` field
   ELSE IF `findByType` is plural
	 return idx/idxs of cards whose field matches any of `findByData` fields
   
   IF `getFirst`
     return idx
   ELSE
     return idxs
 */
func getPositions(getFirst bool, stack *Stack, findByType FINDBY, findByData interface{}, matchByType MATCHBY) (targets []int) {

	switch findByType {

	case FINDBY_First:
		targets = append(targets, 0)

		//... and so on

	case FINDBY_Keys:
		keyArr := findByData.([]interface{})
		for i := 0; i < len(keyArr); i++ {
			for j, c := range stack.Cards {
				if (matchByType == MATCHBY_Object    &&  keyArr[i] ==  c.Key) ||
				   (matchByType == MATCHBY_Reference && &keyArr[i] == &c.Key) {
					targets = append(targets, j)
					if getFirst { break }
				}
			}
		}

	//... and so on

	case FINDBY_Lambda:
		filterStack := stack.Clone()
		getIterator(filterStack, findByData.(func(*Card, ...interface{}) bool))
		for i := range filterStack.Cards {
			targets = append(targets, i)	
			if getFirst { break }
		}

	}

	return

}

/**
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
		setIterator(setStack, replaceData.(func(*Card, ...interface{})))

	}

}
