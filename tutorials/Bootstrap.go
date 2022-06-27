package tutorials

import (
<<<<<<< HEAD
	. "github.com/gabetucker2/gostack"
	"fmt"
)


func bootstrap_pop() {

	// > `stack.Extract(RETURN_Card, FINDBY_First)`
	// >> *removes and returns the first card in the stack*
	
	//
	stack.Extract(RETURN_Card, POSITION_First)
}



func bootstrap_push(card Card*, stack *Stack) bool {

	// > `stack.Add(insert, ORDER_BEFORE, POSITION_First)`
	// >> *adds a card to the beginning of the stack*


	//
	stack.Add(insert, ORDER_BEFORE, POSITION_First)

}

<<<<<<< HEAD
func bootstrap_indexOf(stack *Stack, val interface{}) int {

=======
func bootstrap_indexOf() {
>>>>>>> 49345842b396ff65dc60d24c4e7e492d97a10225

	//return stack.Get(FINDBY_Val, val).Idx
	// >> *returns the index of the first found matching card*
	return stack.Get(POSITION_Card, val).Idx
}

func bootstrap_keyset(stack *Stack) *Stack{
	/**
	@Param: stack
	@Return: stack
	*/
	return stack.GetMany(RETURN_Key, POSITION_All)
}

func bootstrap_valset(stack Stack*) Stack*{
	return stack.GetMany(RETURN_VAL, POSITION_All)
}

/** Executes the Bootstrap.go tutorial */
func Bootstrap() {

	_gostack_tutorials_pop()
	_gostack_tutorials_push()
	_gostack_tutorials_indexOf()
	_gostack_tutorials_keyset()
	_gostack_tutorials_valset()

	return
}
