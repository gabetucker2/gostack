package main

// VARIABLES

var showTestText = true // false for cleaner console, true for debugging

// CASE FUNCTIONS

func _gostack_case_MakeStack(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Empty(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_back_SampleStack()
	stack2 := stack.Empty()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		stack == stack2,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Push(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	stack.Push(testCard3).Push(testCard2).Push(testCard1)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		&stack.cards[0] == testCard1,
		&stack.cards[1] == testCard2,
		&stack.cards[2] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Append(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	stack.Append(testCard1).Append(testCard2).Append(testCard3)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		stack.cards[2] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_1_Pop(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_back_SampleStack()
	pop := stack.Pop()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 2),
		&stack.cards[0] == testCard1,
		&stack.cards[1] == testCard2,
		pop == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_2_Pop(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	pop := stack.Pop()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		pop == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_1_Behead(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_back_SampleStack()
	behead := stack.Behead()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 2),
		behead == testCard1,
		&stack.cards[0] == testCard2,
		&stack.cards[1] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_2_Behead(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	behead := stack.Pop()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		behead.value == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Has(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_back_SampleStack()

	conditions := []bool{
		stack.Has(testCard1),
		stack.Has(testCard2),
		stack.Has(testCard3),
		!stack.Has(testCard4),
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_IndexOf(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_back_SampleStack()

	conditions := []bool{
		stack.IndexOf(testCard1) == 0,
		stack.IndexOf(testCard2) == 1,
		stack.IndexOf(testCard3) == 2,
		stack.IndexOf(testCard4) == -1,
	}

	_gostack_test_End(funcName, conditions)

}

// MAIN FUNCTION

func main() {

	// layer two is dependent on layer one in case tests, layer three dependent on layer two, etc
	println("- BEGINNING TESTS")

	// layer one
	_gostack_case_MakeStack("MakeStack") // regular case

	// layer two
	_gostack_case_Append("stack.Append") // regular case

	// layer three
	_gostack_case_Empty("stack.Empty")        // regular case
	_gostack_case_Push("stack.Push")          // regular case
	_gostack_case_1_Pop("T1:stack.Pop")       // regular case
	_gostack_case_2_Pop("T2:stack.Pop")       // edge case
	_gostack_case_1_Behead("T1:stack.Behead") // regular case
	_gostack_case_2_Behead("T2:stack.Behead") // edge case
	_gostack_case_Has("stack.Has")            // regular and edge cases
	_gostack_case_IndexOf("stack.IndexOf")    // regular and edge cases

}
