package main

func gs_MakeStack() Stack {

	var stack Stack

	stack.len = 0

	return stack

}

func gs_Push(stack Stack, element interface{}) {

}

/*
func IndexOf(stack Stack, needle interface{}) (ret int) {

	ret = -1
	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			ret = i
			break
		}
	}

	return

}
*/
