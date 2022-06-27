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

	stack.Size = len(stack.Cards)

	return stack

}

/** Makes a card with inputted vals and keys

 @receiver `stack` type{Stack}
 @returns `stack`
 @updates `stack.Cards` to be empty
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
 @returns `stack`
 @updates `stack.Cards` to have new cards before/after each designated position
 */
func (stack *Stack) Add(insert interface{}, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var orderType, positionType, positionData interface{}
	unpackVariadic(variadic, &orderType, &positionType, &positionData)

	// set types to default values
	setORDERDefaultIfNil(orderType)
	setPOSITIONDefaultIfNil(positionType)

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
	targets := getPositions(stack, positionType.(POSITION), positionData)

	// fill the array
	for i := 0; i < stack.Size; i++ {
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

	// return
	return stack

}

/** Gets a stack of specified values from specified card(s) at (each) position
 
 @receiver `stack` type{Stack}
 @param `positionType` type{POSITION} default ORDER_After
 @param optional `positionData` type{interface{}} default nil
 @param optional `matchType` type{MATCH} default MATCH_Object
 @returns type{*Stack} the new stack
 @constructs type{*Stack} new stack of specified values from specified cards in `stack`
 @requires `stack.Clone()` has been implemented
 */
func (stack *Stack) Get(positionType POSITION, variadic ...interface{}) *Card {

	// unpack variadic into optional parameters
	var positionData, matchType interface{}
	unpackVariadic(variadic, &positionData, &matchType)

	// set types to default values
	setMATCHDefaultIfNil(matchType)

	// create new stack which returns the searched-for cards
	gotCardsStack := MakeStack()

	// get targeted cards
	targets := getPositions(stack, positionType, positionData)
	
	// get whether to get first of or all of matching cards
	getOne := isSingular(returnType)

	// fill new stack with gotten cards
	if getOne {
		for _, i := range targets {
	
			gotCardsStack.Cards = append(gotCardsStack.Cards, stack.Cards[i])
	
		}
	}

	// return
	return gotCardsStack

}

/** Gets a stack of specified values from specified card(s) at (each) position
 
 @receiver `stack` type{Stack}
 @param `returnType` type{RETURN}
 @param `positionType` type{POSITION} default ORDER_After
 @param optional `positionData` type{interface{}} default nil
 @param optional `matchType` type{MATCH} default MATCH_Object
 @returns type{*Stack} the new stack
 @constructs type{*Stack} new stack of specified values from specified cards in `stack`
 @requires `stack.Clone()` has been implemented
 */
func (stack *Stack) GetMany(returnType RETURN, positionType POSITION, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var positionData, matchType interface{}
	unpackVariadic(variadic, &positionData, &matchType)

	// set types to default values
	setMATCHDefaultIfNil(matchType)

	// create new stack which returns the searched-for cards
	gotCardsStack := MakeStack()

	// get targeted cards
	targets := getPositions(stack, positionType, positionData)
	
	// get whether to get first of or all of matching cards
	getOne := isSingular(returnType)

	// fill new stack with gotten cards
	if getOne {
		for _, i := range targets {
	
			gotCardsStack.Cards = append(gotCardsStack.Cards, stack.Cards[i])
	
		}
	}

	// return
	return gotCardsStack

}
