package main

import "fmt"

func gostack_PrintTesting(thisFuncName string) {

	fmt.Println("TESTING " + thisFuncName)

}

func gostack_PrintOutcome(success bool, thisFuncName string) {

	var out = thisFuncName
	if success {
		out += " SUCCESS"
	} else {
		out += " FAILURE"
	}

	fmt.Println(out)

}

func main() {

	fmt.Println("BEGINNING TESTS")

	thisFuncName := "MakeStack"
	gostack_PrintTesting(thisFuncName)
	stack := MakeStack()
	success := stack != nil && len(stack.vals) == 0 && stack.len == 0
	gostack_PrintOutcome(success, thisFuncName)

	thisFuncName = "stack.Push"
	gostack_PrintTesting(thisFuncName)
	stack.Push("Element 2")
	stack.Push("Element 1")
	success = stack.vals[0] == "Element 1" && stack.vals[1] == "Element 2"
	gostack_PrintOutcome(success, thisFuncName)

	/*

		// IndexOf test 1
		haystack := []int{4, 1, 6}
		needle := 6
		expected := 2
		haystack.Test_IndexOf(needle, expected)

		// IndexOf test 2
		haystack = []int{4, 1, 6}
		needle = 5
		expected = -1
		haystack.Test_IndexOf(needle, expected)

	*/

}
