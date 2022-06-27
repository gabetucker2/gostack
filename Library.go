package gostack

import (
	"reflect"
)

/** Makes a card with inputted vals and keys

 @param optional `val` type{any} default nil
 @param optional `key` type{any} default nil
 @param optional `idx` type{int} default -1 no pass-by-reference
 @returns type{*Card} the newly-constructed card
 @constructs type{*Card} a newly-constructed card
 @ensures the new card will have val `val`, key `key`, and idx `idx`
*/
func MakeCard(variadic ...interface{}) *Card {

	// unpack variadic into optional parameters
	var Val, Key, Idx interface{}
	unpackVariadic(variadic, &Val, &Key, &Idx)

	// TODO: Andy changes here

	var newIdx int
	if Idx == nil { newIdx = -1 } else { newIdx = Idx.(int) }

	// initialize and set new Card
	card := new(Card)
	card.Idx = newIdx // clones Idx
	card.Key = &Key
	card.Val = &Val

	// return
	return card

}

/** Makes a stack of cards with optional starting cards
 
 @param optional `input1` type{[]any, map[any]any} default nil
 @param optional `input2` type{[]any} default nil
 @param optional `repeats` type{int} default 1
 @returns type{*Stack} the newly-constructed stack of newly-constructed cards
 @constructs type{*Stack} a newly-constructed stack of newly-constructed cards
 @requires
  * `input1` is map and nil `input2`
      OR `input1` is an array and nil `input2`
	  OR `input1` is an array and `input2` is an array
  * IF `input1` and `input2` are both passed as arguments
      |`input1`| == |`input2`|
  * MakeCard() has been implemented
 @ensures
  * `repeats` (or, if nil or under 0, 1) amount of times
      IF `input1` is passed
	      IF `input1` is a map
            unpack the map into new cards with corresponding keys and vals
          ELSEIF `input1` is an array and `input2` is not passed
            unpack values from `input1` into new cards
          ELSEIF `input1` is an array and `input2` is an array
		    unpack keys from `input1` and values from `input2` into new cards
	  ELSE
	    the stack is empty
 */
func MakeStack(variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var input1, input2, repeats interface{}
	unpackVariadic(variadic, &input1, &input2, &repeats)

	// initialize stack
	stack := new(Stack)

	// `repeats` (or, if nil or under 0, 1) amount of times
	if repeats == nil || repeats.(int) < 0 { repeats = 1 }
	for i := 0; i < repeats.(int); i++ {

		// IF `input1` is passed
		if input1 != nil {

			input1Type := reflect.ValueOf(input1).Kind()
			switch input1Type {
			
			// IF `input1` is a map
			case reflect.Map:
				// unpack the map into new cards with corresponding keys and vals
				i := 0
				for k, v := range input1.(map[interface{}]interface{}) {
					stack.Cards = append(
						stack.Cards,
						MakeCard(&v, &k, i),
					)
					i++
				}
			
			case reflect.Array:
				input1Len := len(input1.([]interface{}))

				// ELSEIF `input1` is an array and `input2` is not passed
				if input2 != nil {
					// unpack values from `input1` into new cards
					for i := 0; i < input1Len; i++ {
						stack.Cards = append(
							stack.Cards,
							MakeCard(&input1.([]interface{})[i], nil, i),
						)
					}

				// ELSEIF `input1` is an array and `input2` is an array
				} else {
					// unpack keys from `input1` and values from `input2` into new cards
					for i := 0; i < input1Len; i++ {
						stack.Cards = append(
							stack.Cards,
							MakeCard(&input1.([]interface{})[i], &input2.([]interface{})[i], i),
						)
					}
				}
			}
		}
	}

	// set up stack size
	stack.Size = len(stack.Cards)

	// return
	return stack

}

/** Makes a card with inputted vals and keys

 @receiver `stack` type{Stack}
 @returns `stack`
 @updates `stack` to be empty
*/
func (stack *Stack) Empty() *Stack {

	// clear stack.Cards
	stack.Size = 0
	stack.Cards = []*Card{} // avoid replacing stack object

	// return
	return stack

}

/** Returns a clone of the given stack

 @receiver `stack` type{Stack}
 @returns type{*Stack} stack clone
 @constructs type{*Stack} clone of `stack`
 @ensures the stack clone has the same card pointers as `stack`
*/
func (stack *Stack) Clone() *Stack {

	// init
	clone := new(Stack)
	clone.Size = stack.Size
	clone.Cards = stack.Cards

	// return
	return clone

}

/** Returns a clone of the given card

 @receiver `card` type{Card}
 @returns type{*Card} card clone
 @constructs type{*Card} clone of `card`
*/
func (card *Card) Clone() *Card {

	// init
	clone := new(Card)
	clone.Idx = card.Idx
	clone.Key = card.Key
	clone.Val = card.Val

	// return
	return clone

}

/** Adds to a stack of cards or a cards at (each) position(s) 
 
 @receiver `stack` type{Stack}
 @param `insert` type{Card, Stack}
 @param optional `orderType` type{ORDER} default ORDER_After
 @param optional `positionType` type{POSITION} default POSITION_First
 @param optional `positionData` type{interface{}} default nil
 @param optional `matchType` type{MATCH} default MATCH_Object
 @returns `stack`
 @updates `stack` to have new cards before/after each designated position
 */
func (stack *Stack) Add(insert interface{}, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var orderType, positionType, positionData, matchType interface{}
	unpackVariadic(variadic, &orderType, &positionType, &positionData, &matchType)

	// set types to default values
	setORDERDefaultIfNil(orderType)
	setPOSITIONDefaultIfNil(positionType)
	setMATCHDefaultIfNil(matchType)

	// convert insert into slice of cards
	var cardsIn []*Card
	switch insert.(type) {
	case Card:
		cardsIn = append(cardsIn, insert.(*Card))
	case Stack:
		for _, c := range insert.(*Stack).Cards {
			cardsIn = append(cardsIn, c)
		}

	}

	// create new array in which to insert `insert`
	var cardsWithAdded []*Card

	// get targeted cards
	targets := getPositions(false, stack, positionType.(POSITION), positionData, matchType.(MATCH))

	// fill the array
	for i := range stack.Cards {
		for _, j := range targets {

			existingCard := stack.Cards[i]

			if j == i { // we are on a target, add `insert`

				if orderType == ORDER_Before {

					for _, c := range cardsIn {
						cardsWithAdded = append(cardsWithAdded, c)
					}
					cardsWithAdded = append(cardsWithAdded, existingCard)

				} else if orderType == ORDER_After {

					cardsWithAdded = append(cardsWithAdded, existingCard)
					for _, c := range cardsIn {
						cardsWithAdded = append(cardsWithAdded, c)
					}

				}

			} else { // we are on a non-target, just add

				cardsWithAdded = append(cardsWithAdded, existingCard)

			}

		}
	}

	// set old cards array to new cards array with added element(s)
	stack.Cards = cardsWithAdded
	stack.Size = len(stack.Cards)

	// return
	return stack

}

/** Removes all cards from `stack` which share the same field value as another card before

 @receiver `stack` type{Stack}
 @param `typeType` type{TYPE}
 @param optional `matchType` type{MATCH} default MATCH_Object
 @returns `stack`
 @updates `stack` to have no repeating values between field `typeType`
 */
func (stack *Stack) Unique(typeType TYPE, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var matchType interface{}
	unpackVariadic(variadic, &matchType)

	// initialize array
	var newCards []*Card

	// remove all repeats
	for i := range stack.Cards {
		oldCard := stack.Cards[i]
		addToNewCards := true
		for j := range newCards {
			newCard := newCards[j]
			if (matchType == MATCH_Object    &&  oldCard ==  newCard) ||
			   (matchType == MATCH_Reference && &oldCard == &newCard) {
				addToNewCards = false
				break
			}
		}
		if addToNewCards {
			newCards = append(newCards, oldCard)
		}
	}

	// update stack
	stack.Cards = newCards
	stack.Size = len(stack.Cards)
	
	// return
	return stack

}

/** Creates a new interface array from values of `stack`

 @receiver `stack` type{Stack}
 @returns type{[]interface{}} new array
 @ensures new array values correspond to `stack` values
 */
func (stack *Stack) ToArray() (arr []interface{}) {

	// add all card values in stack to arr
	for i := range stack.Cards {
		arr = append(arr, stack.Cards[i].Val)
	}

	// return
	return

}

/** Creates a new interface-interface map from values of `stack`

 @receiver `stack` type{Stack}
 @returns type{map[interface{}]interface{}} new map
 @ensures new map keys and values correspond to `stack` keys and values
 */
func (stack *Stack) ToMap() (m map[interface{}]interface{}) {

	// add all card keys and values in stack to m
	for i := range stack.Cards {
		c := stack.Cards[i]
		m[c.Key] = c.Val
	}

	// return
	return

}

/** Gets a card from specified parameters in a stack, or nil if does not exist

 @receiver `stack` type{Stack}
 @param optional `positionType` type{POSITION} default POSITION_First
 @param optional `positionData` type{interface{}} default nil
 @param optional `matchType` type{MATCH} default MATCH_Object
 @returns type{*Card} the found card OR nil
 */
func (stack *Stack) Get(variadic ...interface{}) (ret *Card) {

	// unpack variadic into optional parameters
	var positionType, positionData, matchType interface{}
	unpackVariadic(variadic, &positionType, &positionData, &matchType)

	// set types to default values
	setPOSITIONDefaultIfNil(positionType)
	setMATCHDefaultIfNil(matchType)

	// get targeted card OR nil
	positions := getPositions(true, stack, positionType.(POSITION), positionData, matchType.(MATCH))
	if len(positions) > 0 {
		ret = stack.Cards[positions[0]]
	} else {
		ret = nil
	}
	
	// return
	return

}

/** Gets a stack from specified parameters in a stack
 
 @receiver `stack` type{Stack}
 @param `positionType` type{POSITION}
 @param optional `positionData` type{interface{}} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `matchType` type{MATCH} default MATCH_Object
 @returns type{*Stack} the new stack
 @constructs type{*Stack} new stack of specified values from specified cards in `stack`
 */
func (stack *Stack) GetMany(positionType POSITION, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var positionData, returnType, matchType interface{}
	unpackVariadic(variadic, &positionData, &returnType, &matchType)

	// set types to default values
	setRETURNDefaultIfNil(returnType)
	setMATCHDefaultIfNil(matchType)

	// create new stack which returns the searched-for cards
	newStack := MakeStack()

	// get targeted cards
	targets := getPositions(false, stack, positionType, positionData, matchType.(MATCH))

	// fill new stack with targeted cards
	for _, i := range targets {

		newCard := new(Card)
		setCardVal(newCard, stack.Cards[i], returnType.(RETURN))

		newStack.Cards = append(newStack.Cards, newCard)

	}

	// return
	return newStack

}
