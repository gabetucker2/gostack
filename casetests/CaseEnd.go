package casetests

import (
	"fmt"

	. "github.com/gabetucker2/gostack"
)

// variables
var showTestText bool

// case functions
// TEMPLATE:
/*
func case_MyFunction(funcName string) {

	test_Start(funcName, showTestText)

	// call MyFunction here

	// make miscellaneous variables to test whether MyFunction worked properly here if necessary

	// add failure conditions here for the case test
	// ensure test_LenAndSize is the first if and only if testing method with Stack type receiver
	conditions := []bool{
		
	}

	test_End(funcName, conditions)

}
*/

func case_MakeCard(funcName string) {

	test_Start(funcName, showTestText)

	card1 := MakeCard()
	card2 := MakeCard("Card 2")
	card3 := MakeCard("Card 3", card2)
	card4 := MakeCard(card1, 8, 2)

	conditions := []bool{
		card1.Idx == -1,
		card2.Idx == -1,
		card3.Idx == -1,
		card4.Idx == 2,
		card1.Key == nil,
		card2.Key == nil,
		card3.Key == card2,
		card4.Key == 8,
		card1.Val == nil,
		card2.Val == "Card 2",
		card3.Val == "Card 3",
		card4.Val == card1,
	}

	test_End(funcName, conditions)

}

func case_MakeStack(funcName string) {

	test_Start(funcName, showTestText)

	// to stacks (in order of conditions listed in documentation)
	stack1 := MakeStack(map1)
	stack2 := MakeStack(arrVals)
	stack3 := MakeStack(arrKeys, arrVals)
	stack4 := MakeStack(nil, arrVals)
	stack5 := MakeStack()

	conditions := []bool{
		test_IdxsAreGood(stack1),
		test_IdxsAreGood(stack2),
		test_IdxsAreGood(stack3),
		test_IdxsAreGood(stack4),
		test_IdxsAreGood(stack5),
		test_LenAndSize(stack1, 3),
		test_LenAndSize(stack2, 3),
		test_LenAndSize(stack3, 3),
		test_LenAndSize(stack4, 3),
		test_LenAndSize(stack5, 0),
	}

	test_End(funcName, conditions)

}

/** Executes all case tests */
func Run(_showTestText bool) {

	showTestText = _showTestText
	test_Setup()

	fmt.Println("- BEGINNING TESTS (fix failures/errors in descending order)")

	case_MakeCard("MakeCard")
	case_MakeStack("MakeStack")

}
