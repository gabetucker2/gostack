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
	elemToInterface := []interface{}{elem}

	// append stack.elems to our elem
	stack.elems = append(elemToInterface, stack.elems)

	// return
	return stack

}

/*
func IndexOf(stack Stack, element interface{}) (ret int) {

	ret = -1
	for i := 0; i < len(stack); i++ {
		if element == stack[i] {
			ret = i
			break
		}
	}

	return

}
*/
