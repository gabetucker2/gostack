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
	stack.Add(testCard3, Position_First).Add(testCard2, Position_First).Add(testCard1, Position_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		stack.cards[2] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Add_Last(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	stack.Add(testCard1, Position_Last).Add(testCard2, Position_Last).Add(testCard3, Position_Last)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 3),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		stack.cards[2] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_1_Extract_First(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()
	extract := stack.Extract(Position_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 2),
		extract == testCard1,
		stack.cards[0] == testCard2,
		stack.cards[1] == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_2_Extract_First(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	extract := stack.Extract(Position_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		extract == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_1_Extract_Last(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()
	extract := stack.Extract(Position_Last)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 2),
		stack.cards[0] == testCard1,
		stack.cards[1] == testCard2,
		extract == testCard3,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_2_Extract_Last(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := MakeStack()
	extract := stack.Extract(Position_Last)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		extract == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Has(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()

	conditions := []bool{
		stack.Has(testCard1),
		stack.Has(testCard2),
		stack.Has(testCard3),
		!stack.Has(testCard4),
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_IndexCard(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()

	conditions := []bool{
		stack.IndexCard(testCard1) == 0,
		stack.IndexCard(testCard2) == 1,
		stack.IndexCard(testCard3) == 2,
		stack.IndexCard(testCard4) == -1,
	}

	_gostack_test_End(funcName, conditions)

}

// MAIN FUNCTION

func main() {

	println("- BEGINNING TESTS")

	// constructor tests
	_gostack_case_MakeStack("MakeStack") // regular case

	// other function tests
	_gostack_case_Add_Last("stack.Add_Last")                // regular case
	_gostack_case_Empty("stack.Empty")                      // regular case
	_gostack_case_Add_First("stack.Add_First")              // regular case
	_gostack_case_1_Extract_First("C1:stack.Extract_First") // regular case
	_gostack_case_2_Extract_First("C2:stack.Extract_First") // edge case
	_gostack_case_1_Extract_Last("C1:stack.Extract_Last")   // regular case
	_gostack_case_2_Extract_Last("C2:stack.Extract_Last")   // edge case
	_gostack_case_Has("stack.Has")                          // regular and edge cases
	_gostack_case_IndexCard("stack.IndexCard")              // regular and edge cases

}
