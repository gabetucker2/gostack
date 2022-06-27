package tutorials

func _gostack_tutorials_pop(stack *Stack) *Card {

	// > `stack.Extract(RETURN_Card, POSITION_First)`
	// >> *removes and returns the first card in the stack*
	
	//
	stack.Extract(RETURN_Card, POSITION_First)


}

func _gostack_tutorials_push(card Card*, stack *Stack) bool {

	// > `stack.Add(insert, ORDER_BEFORE, POSITION_First)`
	// >> *adds a card to the beginning of the stack*

	//
	stack.Add(insert, ORDER_BEFORE, POSITION_First)

}

func _gostack_tutorials_indexOf(stack *Stack, card *Card) int{

	// > `stack.Get(RETURN_Idx, POSITION_Card, cardToMatch, MATCH_Object)`
	// >> *returns the index of the first found matching card*
	return stack.Get(POSITION_Card, card).Idx
}

func _gostack_tutorials_keyset(stack *Stack) *Stack{
	/**
	@Param: stack
	@Return: stack
	*/
	return stack.GetMany(RETURN_Key, POSITION_All)
}

func _gostack_tutorials_valset(stack Stack*) Stack*{
	return stack.GetMany(RETURN_VAL, POSITION_All)
}

func Bootstrap() {

	_gostack_tutorials_pop()
	_gostack_tutorials_push()
	_gostack_tutorials_indexOf()
	_gostack_tutorials_keyset()
	_gostack_tutorials_valset()


}
