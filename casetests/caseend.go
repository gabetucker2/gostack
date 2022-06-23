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

	stack := _gostack_test_SampleStack()
	stack2 := stack.Empty()

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		stack == stack2,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Add_First(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	stack.Add(testCardC, POSITION_First).Add(testCardB, POSITION_First).Add(testCardA, POSITION_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCardA,
		stack.cards[1] == testCardB,
		stack.cards[2] == testCardC,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Add_Last(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	stack.Add(testCardA, POSITION_Last).Add(testCardB, POSITION_Last).Add(testCardC, POSITION_Last)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCardA,
		stack.cards[1] == testCardB,
		stack.cards[2] == testCardC,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Extract_Empty(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	extract := stack.Extract(POSITION_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		extract == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Extract_First(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()
	extract := stack.Extract(POSITION_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 2),
		extract == testCardA,
		stack.cards[0] == testCardB,
		stack.cards[1] == testCardC,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Extract_Last(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()
	extract := stack.Extract(POSITION_Last)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 2),
		stack.cards[0] == testCardA,
		stack.cards[1] == testCardB,
		extract == testCardC,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Replace_Empty(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	extract := stack.Replace(testCardA, POSITION_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		extract == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Replace_First(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()
	extract := stack.Replace(testCardC, POSITION_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCardC,
		stack.cards[1] == testCardB,
		stack.cards[2] == testCardC,
		extract == testCardA,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Replace_Last(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()
	extract := stack.Replace(testCardA, POSITION_Last)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCardA,
		stack.cards[1] == testCardB,
		stack.cards[2] == testCardA,
		extract == testCardC,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Has(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()

	conditions := []bool{
		stack.Has(POSITION_Card, testCardA),
		stack.Has(POSITION_Card, testCardB),
		stack.Has(POSITION_Card, testCardC),
		stack.Has(POSITION_Val, testCardA.val),
		stack.Has(POSITION_Val, testCardB.val),
		stack.Has(POSITION_Val, testCardC.val),
		!stack.Has(POSITION_Card, testCardD),
		!stack.Has(POSITION_Val, testCardD.val),
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Index(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()

	conditions := []bool{
		stack.Index(POSITION_Card, testCardA) == 0,
		stack.Index(POSITION_Card, testCardB) == 1,
		stack.Index(POSITION_Card, testCardC) == 2,
		stack.Index(POSITION_Val, testCardA.val) == 0,
		stack.Index(POSITION_Val, testCardB.val) == 1,
		stack.Index(POSITION_Val, testCardC.val) == 2,
		stack.Index(POSITION_Card, testCardD) == -1,
		stack.Index(POSITION_Val, testCardD.val) == -1,
	}

	_gostack_test_End(funcName, conditions)

}

// MAIN FUNCTION

func main() {

	println("- BEGINNING TESTS")

	// constructor tests
	_gostack_case_MakeStack("MakeStack") // regular case

	// other function tests
	_gostack_case_Empty("stack.Empty") // regular case

	_gostack_case_Add_First("stack.Add_First")         // regular case
	_gostack_case_Add_Last("stack.Add_Last")           // regular case
	_gostack_case_Extract_Empty("stack.Extract_Empty") // edge case
	_gostack_case_Extract_First("stack.Extract_First") // regular case
	_gostack_case_Extract_Last("stack.Extract_Last")   // regular case
	_gostack_case_Replace_Empty("stack.Replace_Empty") // edge case
	_gostack_case_Replace_First("stack.Replace_First") // regular case
	_gostack_case_Replace_Last("stack.Replace_Last")   // regular case

	_gostack_case_Has("stack.Has")     // regular and edge cases
	_gostack_case_Index("stack.Index") // regular and edge cases

}
