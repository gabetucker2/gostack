package main

// VARIABLES

var showTestText = false // false for cleaner console, true for debugging

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

func _gostack_case_AddFirst(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	stack.AddFirst(testCard3).AddFirst(testCard2).AddFirst(testCard1)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		stack.cards[2] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_AddLast(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	stack.AddLast(testCard1).AddLast(testCard2).AddLast(testCard3)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		stack.cards[2] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_1_ExtractFirst(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_back_SampleStack()
	extract := stack.ExtractFirst()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 2),
		extract == testCard1,
		stack.cards[0] == testCard2,
		stack.cards[1] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_2_ExtractFirst(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	extract := stack.ExtractFirst()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		extract == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_1_ExtractLast(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_back_SampleStack()
	extract := stack.ExtractLast()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 2),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		extract == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_2_ExtractLast(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	extract := stack.ExtractLast()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		extract == nil,
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
	_gostack_case_AddLast("stack.AddLast") // regular case

	// layer three
	_gostack_case_Empty("stack.Empty")                    // regular case
	_gostack_case_AddFirst("stack.AddFirst")              // regular case
	_gostack_case_1_ExtractFirst("T1:stack.ExtractFirst") // regular case
	_gostack_case_2_ExtractFirst("T2:stack.ExtractFirst") // edge case
	_gostack_case_1_ExtractLast("T1:stack.ExtractLast")   // regular case
	_gostack_case_2_ExtractLast("T2:stack.ExtractLast")   // edge case
	_gostack_case_Has("stack.Has")                        // regular and edge cases
	_gostack_case_IndexOf("stack.IndexOf")                // regular and edge cases

}
