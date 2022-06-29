package gostack

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
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

	// set default Idx to -1
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
  * `MakeCard()` has been implemented
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

/** Removes all cards from `stack` which share the same field value as another card before

 @receiver `stack` type{Stack}
 @param `typeType` type{TYPE}
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates `stack` to have no repeating values between field `typeType`
 */
func (stack *Stack) Unique(typeType TYPE, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var matchByType interface{}
	unpackVariadic(variadic, &matchByType)

	// initialize array
	var newCards []*Card

	// remove all repeats
	for i := range stack.Cards {
		oldCard := stack.Cards[i]
		addToNewCards := true
		for j := range newCards {
			newCard := newCards[j]
			if (matchByType == MATCHBY_Object    &&  oldCard ==  newCard) ||
			   (matchByType == MATCHBY_Reference && &oldCard == &newCard) {
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

/** Shuffles the order of `stack` cards

 @receiver `stack` type{Stack}
 @returns `stack`
 @updates
  * `stack` card ordering is randomized
  * rand.Seed is updated to time.Now().UnixNano()
 */
func (stack *Stack) Shuffle() *Stack {

	// pseudo-randomize seed
	rand.Seed(time.Now().UnixNano())

	// shuffle
	rand.Shuffle(stack.Size, func(i, j int) { stack.Cards[i], stack.Cards[j] = stack.Cards[j], stack.Cards[i] })

	// return
	return stack

}

/** Flips the ordering of `stack.Cards`
 
 @receiver `stack` type{Stack}
 @returns `stack`
 @updates `stack` to have its ordering reversed
 */
func (stack *Stack) Flip() *Stack {

	// new card stack
	var newCards []*Card

	// flip it
	for i := range stack.Cards {
		newCards = append(newCards, stack.Cards[i])
	}

	// update
	stack.Cards = newCards

	// return
	return stack

}

/** Prints information regarding `card` to the console
 
 @receiver `card` type{Card}
 @updates terminal logs
 */
func (card *Card) Print() {

	fmt.Println("gostack: PRINTING CARD")
	fmt.Printf("- card.Idx: %v\n", card.Idx)
	fmt.Printf("- card.Key: %v\n", card.Key)
	fmt.Printf("- card.Val: %v\n", card.Idx)

}

/** Prints information regarding `stack` to the console
 
 @receiver `stack` type{Stack}
 @updates terminal logs
 @requires `card.Print()` has been implemented
 */
func (stack *Stack) Print() {

	fmt.Println("gostack: PRINTING STACK")
	fmt.Printf("- stack.Size: %v\n", stack.Size)
	for i := range stack.Cards {
		stack.Cards[i].Print()
	}

}

/** Order the cards contingent on some attribute they contain
 
 @receiver `stack` type{Stack}
 @param `lambda` type{func(*Card, *Stack, ...interface{}) (ORDER, int)}
 @requires
  * `lambda` returns the order (before/after) and index to which to move your card in the stack
  * `lambda` does not update `stack` itself
 @ensures each card in `stack` is passed into your lambda function
 */
func (stack *Stack) Sort(lambda func(*Card, *Stack, ...interface{}) (ORDER, int)) {
	sortIterator(stack, lambda)
}

/** Iterate through a stack calling your lambda function on each card
 
 @receiver `stack` type{Stack}
 @param `lambda` type{func(*Card, ...interface{})}
 @ensures
  * Each card in `stack` is passed into your lambda function
  * `stack` is the first argument passed into your variadic parameter on the first call
 */
func (stack *Stack) Lambda(lambda func(*Card, ...interface{})) {
	generalIterator(stack, lambda)
}

/** Adds to a stack of cards or a cards at (each) position(s) and returns `stack`
 
 @receiver `stack` type{Stack}
 @param `insert` type{Card, Stack}
 @param optional `orderType` type{ORDER} default ORDER_Before
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates `stack` to have new cards before/after each designated position
 */
func (stack *Stack) Add(insert interface{}, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var orderType, findType, findData, matchByType interface{}
	unpackVariadic(variadic, &orderType, &findType, &findData, &matchByType)

	// set types to default values
	setORDERDefaultIfNil(&orderType)
	setFINDDefaultIfNil(&findType)
	setMATCHBYDefaultIfNil(&matchByType)

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
	targets := getPositions(false, stack, findType.(FIND), findData, matchByType.(MATCHBY))

	// fill the array
	for i := range stack.Cards {
		for _, j := range targets {

			existingCard := stack.Cards[i]

			if j == i { // we are on a target, add `insert`

				// add cards to stack before or after existing element, based on orderType
				if orderType == ORDER_After { cardsWithAdded = append(cardsWithAdded, existingCard) }
				for _, c := range cardsIn {
					cardsWithAdded = append(cardsWithAdded, c)
				}
				if orderType == ORDER_Before { cardsWithAdded = append(cardsWithAdded, existingCard) }

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

/** Moves one element or slice of cards to before or after another element or slice of cards
 
 @receiver `stack` type{Stack}
 @param `findType_from` type{FIND}
 @param `orderType` type{ORDER}
 @param `findType_to` type{FIND}
 @param optional `findData_from` type{interface{}} default nil
 @param optional `findData_to` type{interface{}} default nil
 @param optional `matchByType_from` type{MATCHBY} default MATCHBY_Object
 @param optional `matchByType_to` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @ensures IF `findType_to` or `findType_from` get over one position, method doesn't perform move and prints invalid argument (FIND_Slice is the sole exception to this rule)
 */
func (stack *Stack) Move(findType_from FIND, orderType ORDER, findType_to FIND, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var findData_from, findData_to, matchByType_from, matchByType_to interface{}
	unpackVariadic(variadic, &findData_from, &findData_to, &matchByType_from, &matchByType_to)

	// set types to default values
	setORDERDefaultIfNil(&orderType)
	setMATCHBYDefaultIfNil(&matchByType_from)
	setMATCHBYDefaultIfNil(&matchByType_to)

	// initialize positions
	fromArr := getPositions(false, stack, findType_from, findData_from, matchByType_from.(MATCHBY))
	toArr := getPositions(false, stack, findType_to, findData_to, matchByType_to.(MATCHBY))

	// do main function only if ensures clause is fulfilled
	if (len(fromArr) == 1 || findType_from == FIND_Slice) && (len(toArr) == 1 || findType_to == FIND_Slice) {

		// set up to
		to := toArr[0]
		if findType_to == FIND_Slice && orderType == ORDER_After {
			to = toArr[1]
		}

		// initialize new cards
		var newCards []*Card

		// fill newCards with cards to add
		var from []*Card
		for i := range stack.Cards {
			if i == fromArr[0] { // on from
				for _, j := range fromArr {
					// add to from, remove from stack
					from = append(from, stack.Cards[j])
				}
			} else if i == to - len(from) { // on to
				// add from to stack before or after existing element, based on orderType
				if orderType == ORDER_After { newCards = append(newCards, stack.Cards[i]) }
				for j := range from {
					newCards = append(newCards, from[j])
				}
				if orderType == ORDER_Before { newCards = append(newCards, stack.Cards[i]) }
			} else { // on regular
				newCards = append(newCards, stack.Cards[i])
			}
		}

		stack.Cards = newCards

	} else {
		fmt.Printf("ERROR - gostack: stack.Move(...) function argument does not fulfill ensures clause")
	}

	// return
	return stack

}

/** Returns a boolean representing whether a search exists in the stack

 @receiver `stack` type{Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns true IF successful search, false IF unsuccessful search
 @requires `stack.Get()` has been implemented
 */
func (stack *Stack) Has(variadic ...interface{}) bool {

	// unpack variadic into optional parameters
	var findType, findData, matchByType interface{}
	unpackVariadic(variadic, &findType, &findData, &matchByType)

	// return
	return stack.Get(findType, findData, matchByType) != nil
}

/** Gets a card from specified parameters in a stack, or nil if does not exist

 @receiver `stack` type{Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `clonesType_card` type{CLONES} default CLONE_False
 @param optional `clonesType_keys` type{CLONES} default CLONE_False
 @param optional `clonesType_vals` type{CLONES} default CLONE_False
 @returns type{*Card} the found card OR nil if invalid FIND
 @ensures
  * CLONE_True for `clonesType_card` means the returned card object itself is a clone
  * CLONE_True for `clonesType_key` means the returned card key is a clone
  * CLONE_True for `clonesType_val` means the returned card val is a clone
 */
func (stack *Stack) Get(variadic ...interface{}) (ret *Card) {

	// unpack variadic into optional parameters
	var findType, findData, matchByType, clonesType_card, clonesType_key, clonesType_val interface{}
	unpackVariadic(variadic, &findType, &findData, &matchByType, &clonesType_card, &clonesType_key, &clonesType_val)

	// set types to default values
	setFINDDefaultIfNil(&findType)
	setMATCHBYDefaultIfNil(&matchByType)
	setCLONEDefaultIfNil(&clonesType_card)
	setCLONEDefaultIfNil(&clonesType_key)
	setCLONEDefaultIfNil(&clonesType_val)

	// get targeted card OR nil
	targets := getPositions(true, stack, findType.(FIND), findData, matchByType.(MATCHBY))
	if len(targets) > 0 {
		ret = stack.Cards[targets[0]]
		// clone if necessary
		if clonesType_card == CLONE_True {
			ret = ret.Clone()
		}
		if clonesType_key == CLONE_True {
			ret.Key = cloneInterface(ret.Key)
		}
		if clonesType_val == CLONE_True {
			ret.Val = cloneInterface(ret.Val)
		}
	} else {
		ret = nil
	}
	
	// return
	return

}

/** Gets a stack from specified parameters in a stack
 
 @receiver `stack` type{Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `clonesType` type{CLONES} default CLONE_False
 @param optional `clonesType_keys` type{CLONES} default CLONE_False
 @param optional `clonesType_vals` type{CLONES} default CLONE_False
 @returns type{*Stack} the new stack
 @constructs type{*Stack} new stack of specified values from specified cards in `stack`
 @requires
  * `MakeStack()` has been implemented
  * `clonesType_keys` and `clonesType_vals` are only passed if `returnType` == RETURN_Cards
 @ensures
  * CLONE_True means the vals of cards in the returned stack are not the original object that was gotten
  * CLONE_True for `clonesType_keys` means the cards in the returned stack keys are clones
  * CLONE_True for `clonesType_vals` means the cards in the returned stack vals are clones
 */
func (stack *Stack) GetMany(findType FIND, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType, returnType, clonesType, clonesType_keys, clonesType_vals interface{}
	unpackVariadic(variadic, &findData, &matchByType, &returnType, &clonesType, &clonesType_keys, &clonesType_vals)

	// set types to default values
	setMATCHBYDefaultIfNil(&matchByType)
	setRETURNDefaultIfNil(&returnType)
	setCLONEDefaultIfNil(&clonesType)
	setCLONEDefaultIfNil(&clonesType_keys)
	setCLONEDefaultIfNil(&clonesType_vals)

	// create new stack which returns the searched-for cards
	newStack := MakeStack()

	// get targeted cards
	targets := getPositions(false, stack, findType, findData, matchByType.(MATCHBY))

	// fill new stack with targeted cards
	for _, i := range targets {

		newCard := new(Card)
		oldCard := stack.Cards[i]

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

		// clone if necessary
		if clonesType == CLONE_True {
			newCard.Val = cloneInterface(newCard.Val)
		}
		if returnType == RETURN_Cards {
			if clonesType_keys == CLONE_True {
				newCard.Val.(*Card).Key = cloneInterface(newCard.Key)
			}
			if clonesType_vals == CLONE_True {
				newCard.Val.(*Card).Val = cloneInterface(newCard.Val)
			}
		}

		newStack.Cards = append(newStack.Cards, newCard)

	}

	// return
	return newStack

}

/** Returns a clone of a found card before its respective field is updated to `replaceData` (OR nil if not found)
 
 @receiver `stack` type{Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{interface{}}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns type{*Card} a clone of extracted card OR nil if found no cards
 @updates first found card to `replaceData`
 @requires `stack.Get()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 */
func (stack *Stack) Replace(replaceType REPLACE, replaceData interface{}, findType FIND, variadic ...interface{}) (ret *Card) {

	// unpack variadic into optional parameters
	var findData, matchByType interface{}
	unpackVariadic(variadic, &findData, &matchByType)

	// set types to default values
	setMATCHBYDefaultIfNil(&matchByType)

	// get deep copy of targeted card OR nil
	ret = stack.Get(findType, findData, matchByType, CLONE_True, CLONE_True, CLONE_True)
	// target is reference to card OR nil
	target := stack.Get(findType, findData, matchByType)

	// set targeted card field to replaceData if was found (updateRespectiveField fulfills our ensures clause)
	if target != nil {
		updateRespectiveField(stack, replaceType, replaceData, target)
	}

	// return
	return

}

/** Returns a stack whose values are clones of the original fields updated to `replaceData`
 
 @receiver `stack` type{Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{interface{}}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @returns type{*Stack} a stack whose values are the extracted cards pre-update
 @updates all found cards to `replaceData`
 @requires `stack.GetMany()` has been implemented
 @ensures if `replaceData` is nil and `replaceType is REPLACE_Card`, the cards found will be removed from `stack`
 */
func (stack *Stack) ReplaceMany(replaceType REPLACE, replaceData interface{}, findType FIND, variadic ...interface{}) (ret *Stack) {

	// unpack variadic into optional parameters
	var findData, matchByType, returnType interface{}
	unpackVariadic(variadic, &findData, &matchByType, &returnType)

	// set types to default values
	setMATCHBYDefaultIfNil(&matchByType)
	setRETURNDefaultIfNil(&returnType)

	// get deep copy of targeted cards to return
	ret = stack.GetMany(findType, findData, matchByType, returnType, CLONE_True, CLONE_True, CLONE_True)
	// target is reference to cards OR nil
	target := stack.GetMany(findType, findData, matchByType, returnType)

	// set targeted cards' fields to replaceData if was found (updateRespectiveField fulfills our ensures clause)
	if target != nil {
		for i := range target.Cards {
			updateRespectiveField(stack, replaceType, replaceData, target.Cards[i])
		}
	}

	// return
	return

}

/** Updates a card in and returns `stack`
 
 @receiver `stack` type{Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates the found card in `stack`
 @requires `stack.Replace()` has been implemented
 */
func (stack *Stack) Update(findType FIND, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType interface{}
	unpackVariadic(variadic, &findData, &matchByType)

	// update stack
	stack.Replace(REPLACE_Card, nil, findType, findData, matchByType)

	// return the original stack
	return stack

}

/** Updates cards in and returns `stack`
 
 @receiver `stack` type{Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates  the found cards in `stack`
 @requires `stack.ReplaceMany()` has been implemented
 */
func (stack *Stack) UpdateMany(findType FIND, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType, returnType interface{}
	unpackVariadic(variadic, &findData, &matchByType, &returnType)

	// update stack
	stack.ReplaceMany(REPLACE_Card, nil, findType, findData, matchByType, returnType)

	// return the original stack
	return stack

}

/** Gets and removes a card from `stack`, or returns nil if does not exist
 
 @receiver `stack` type{Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns type{*Card} the extracted card OR nil if invalid FIND
 @updates `stack` to no longer have found card
 @requires `stack.Replace()` has been implemented
 */
func (stack *Stack) Extract(findType FIND, variadic ...interface{}) *Card {

	// unpack variadic into optional parameters
	var findData, matchByType interface{}
	unpackVariadic(variadic, &findData, &matchByType)

	// return the original value
	return stack.Replace(REPLACE_Card, nil, findType, findData, matchByType)

}

/** Gets and removes a set of cards from `stack`
 
 @receiver `stack` type{Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @returns type{*Stack} the extracted card OR nil if invalid FIND
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 */
func (stack *Stack) ExtractMany(findType FIND, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType, returnType interface{}
	unpackVariadic(variadic, &findData, &matchByType, &returnType)

	// return the original value
	return stack.ReplaceMany(REPLACE_Card, nil, findType, findData, matchByType, returnType)

}

/** Removes a card from and returns `stack`
 
 @receiver `stack` type{Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates `stack` to no longer have found card
 @requires `stack.Replace()` has been implemented
 */
func (stack *Stack) Remove(findType FIND, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType interface{}
	unpackVariadic(variadic, &findData, &matchByType)

	// remove the card
	stack.Replace(REPLACE_Card, nil, findType, findData, matchByType)

	// return stack
	return stack

}

/** Removes a set of cards from and returns `stack`
 
 @receiver `stack` type{Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 */
func (stack *Stack) RemoveMany(findType FIND, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var findData, matchByType interface{}
	unpackVariadic(variadic, &findData, &matchByType)

	// remove the cards
	stack.ReplaceMany(REPLACE_Card, nil, findType, findData, matchByType)

	// return stack
	return stack

}
