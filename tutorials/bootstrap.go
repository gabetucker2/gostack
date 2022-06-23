package tutorials

func _gostack_tutorials_pop() {

	// > `stack.Extract(RETURN_Card, POSITION_First)`
	// >> *removes and returns the first card in the stack*

}

func _gostack_tutorials_push() {

	// > `stack.Add(insert, ORDER_BEFORE, POSITION_First)`
	// >> *adds a card to the beginning of the stack*

}

func _gostack_tutorials_indexOf() {

	// > `stack.Get(RETURN_Idx, POSITION_Card, cardToMatch, MATCH_Object)`
	// >> *returns the index of the first found matching card*

}

func Bootstrap() {

	_gostack_tutorials_pop()
	_gostack_tutorials_push()
	_gostack_tutorials_indexOf()

}
