package main

func MakeStack() *Stack {

	// creates new stack
	var stack *Stack

	// initialized new stack's variables
	stack.len = 0

	// return
	return stack

}

func (stack *Stack) Push(element interface{}) *Stack {

	// pushes all previous values to i + 1
	for i := stack.len; i > 0; i-- {
		stack.vals[i+1] = stack.vals[i]
	}

	// pushes new element into stack
	stack.vals[0] = element

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
