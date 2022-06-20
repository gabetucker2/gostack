package main

func MakeStack() Stack {

	// creates new stack
	var stack Stack

	// initialized new stack's variables
	stack.len = 0

	// return
	return stack

}

func (stack *Stack) Push(elem interface{}) *Stack {

	// insert elem into new interface slice to satisfy append function
	newInterface := []interface{}{elem}

	// append each element in stack.elems to elem
	for _, e := range stack.elems {
		newInterface = append(newInterface, e)
	}

	// set stack.elems to our new interface
	stack.elems = newInterface

	// return
	return stack

}

func (stack *Stack) IndexOf(target interface{}) (idx int) {

	idx = -1
	for i, elem := range stack.elems {
		if elem == target {
			idx = i
			break
		}
	}

	return

}
