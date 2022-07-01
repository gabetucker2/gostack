package tutorials

import (
	"fmt"

	. "github.com/gabetucker2/gostack"
)

/** Remove and return first card in stack

@param `stack` type{*Stack}
@returns `card` type{*Card}
@updates `stack` with designated card removed
*/
func bootstrap_pop(stack *Stack) *Card{
	return stack.Extract(FIND_First)
}

/** Push a card to the top of the stack of cards
 
 @param `stack` type{*Stack}
 @param `card` type Card 
 @returns `stack` after adding, type{*Stack}
 @updates `stack` to have new card on top of stack
 */
func bootstrap_push(stack *Stack, card *Card) *Stack{
	return stack.Add(card, ORDER_Before)
}

/** Get the index of a card in the stack
 
 @param `stack` type{*Stack}
 @param `card` type Card 
 @returns `idx` type{int}
 */
func bootstrap_indexOf(stack *Stack, card *Card) int {
	return stack.Get(FIND_First, card).Idx
}

/** Return list of all keys in the stack
 @param `stack` type{*Stack}
 @returns `stack` type{*Stack}
 */
func bootstrap_keyset(stack *Stack) *Stack{
	return stack.GetMany(FIND_Lambda, func(stack *Stack, card *Card) bool{
		v := card.Key
		return v != nil
	})
}

/** Return list of all values in the stack
 @param `stack` type{*Stack}
 @returns `stack` type{*Stack}
 */
func bootstrap_valset(stack *Stack) *Stack{
	return stack.GetMany(FIND_Lambda, func(stack *Stack, card *Card) bool{
		v := card.Val
		return v != nil
	})
}

/** Executes the Bootstrap.go tutorial */
func Bootstrap() {

	makeSampleStack := func() *Stack {
		return MakeStack([]int{2, 10, 11, 12, 40}, []int{0, 90, 4, 2, 20})
	}
	
	myStack := makeSampleStack()
	bootstrap_push(myStack, MakeCard(1,2,3))
	bootstrap_indexOf(myStack, MakeCard(1,2,3))
	bootstrap_keyset(myStack)
	bootstrap_valset(myStack)
	fmt.Println(bootstrap_pop(myStack))

	return
}
