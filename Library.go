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

 @receiver `stack` type{*Stack}
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

// stack.Add(insert, ...ORDER_*, ...POSITION_*, ...POSITIONDATA)
func (stack *Stack) Add(insert interface{}, variadic ...interface{}) *Stack {

	// unpack variadic into optional parameters
	var orderType, positionType, data interface{}
	unpackVariadic(variadic, &orderType, &positionType, &data)

	// create new array in which to insert `insert`
	var cardsWithAdded []*Card



	// set old cards array to new cards array with added element(s)
	stack.Cards = cardsWithAdded

	// return
	return stack

}
