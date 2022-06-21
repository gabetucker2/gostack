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
	stack.Add(testCardC, Position_First).Add(testCardB, Position_First).Add(testCardA, Position_First)

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
	stack.Add(testCardA, Position_Last).Add(testCardB, Position_Last).Add(testCardC, Position_Last)

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
	extract := stack.Extract(Position_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		extract == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Extract_First(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()
	extract := stack.Extract(Position_First)

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
	extract := stack.Extract(Position_Last)

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
	extract := stack.Replace(testCardA, Position_First)

	conditions := []bool{
		_gostack_back_LenAndSize(stack, 0),
		extract == nil,
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_Replace_First(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()
	extract := stack.Replace(testCardC, Position_First)

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
	extract := stack.Replace(testCardA, Position_Last)

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
		stack.Has(testCardA),
		stack.Has(testCardB),
		stack.Has(testCardC),
		!stack.Has(testCardD),
	}

	_gostack_test_End(funcName, conditions)

}

func _gostack_case_IndexCard(funcName string) {

	_gostack_test_Start(funcName, showTestText)

	stack := _gostack_test_SampleStack()

	conditions := []bool{
		stack.IndexCard(testCardA) == 0,
		stack.IndexCard(testCardB) == 1,
		stack.IndexCard(testCardC) == 2,
		stack.IndexCard(testCardD) == -1,
	}

	_gostack_test_End(funcName, conditions)

}

// MAIN FUNCTION

func main() {

	println("- BEGINNING TESTS")

	// constructor tests
	_gostack_case_MakeStack("MakeStack") // regular case

	// other function tests
	_gostack_case_Empty("stack.Empty")                 // regular case
	_gostack_case_Add_First("stack.Add_First")         // regular case
	_gostack_case_Add_Last("stack.Add_Last")           // regular case
	_gostack_case_Extract_Empty("stack.Extract_Empty") // edge case
	_gostack_case_Extract_First("stack.Extract_First") // regular case
	_gostack_case_Extract_Last("stack.Extract_Last")   // regular case
	_gostack_case_Replace_Empty("stack.Replace_Empty") // edge case
	_gostack_case_Replace_First("stack.Replace_First") // regular case
	_gostack_case_Replace_Last("stack.Replace_Last")   // regular case
	_gostack_case_Has("stack.Has")                     // regular and edge cases
	_gostack_case_IndexCard("stack.IndexCard")         // regular and edge cases

}
