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
 
 @param `stack` type{Stack} no pass-by-reference
 @param `positionType` type{POSITION}
 @param `data` type{interface{}}
 @returns the []int of targeted positions
 @constructor creates a new []int
 @ensures
  * SWITCH `positionType`
	  case POSITION_First
	    return 0
	  case POSITION_Last
	    return len(stack)
	  case POSITION_Idx
	    return first idx with card.Idx == data or any in data
	  case POSITION_Idxs
	    return all idxs with card.Idx == data or any in data
	  case POSITION_Key
	    return first idx with card.Key == data or any in data
	  case POSITION_Keys
	    return all idxs with card.Key == data or any in data
	  case POSITION_Val
	    return first idx with card.Val == data or any in data
	  case POSITION_Vals
	    return all idxs with card.Val == data or any in data
	  case POSITION_Card
	    return first idx with card == data or any in data
	  case POSITION_Cards
	    return all idxs with card == data or any in data
	  case POSITION_All
	    return all idxs
	  case POSITION_Lambda
	    return all idxs where lambda(card) == true
 */
func getPositions(stack *Stack, positionType POSITION, data interface{}) (targets []int) {

	switch positionType {

	case POSITION_First:
		targets = append(targets, 0)

	//... and so on

	case POSITION_Lambda:
		filterStack := stack.Clone()
		iterator(filterStack, data.(func(*Card, ...interface{}) bool))
		for i := range filterStack.Cards {
			targets = append(targets, i)	
		}

	}

	return

}

/** Returns a new stack of fields from a stack of cards based on `returnType`
 
 @param `stack` type{Stack}
 @param `returnType` type{RETURN}
 @returns the stack of the fetched values
 @constructor creates a new Stack
 @ensures
  * SWITCH `returnType`
	  case RETURN_Idxs
	    return stack of Idxs of each card in `stack`
	  case RETURN_Keys
	    return stack of Keys of each card in `stack`
	  case RETURN_Vals
	    return stack of Vals of each card in `stack`
	  case RETURN_Cards
	    return `stack`
 */
func getReturns(input *Stack, returnType RETURN) (ret interface{}) {

	switch returnType {

		// TODO: implement

	}

	return

}
