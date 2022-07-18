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

	// YOUR STUFF HERE

	conditions := []bool{
		
	}

	test_End(funcName, conditions)

}
*/

func case_MakeCard(funcName string) {

	test_Start(funcName, showTestText)

	// initialize
	card1 := MakeCard()
	card2 := MakeCard("Card 2")
	card3 := MakeCard("Card 3", card2)
	card4 := MakeCard(card1, 8, 2)

	// test whether updating fields does so by object
	card2.Val = "Card 4"
	card3.Key = 7

	conditions := []bool{
		card1.Idx == -1,
		card2.Idx == -1,
		card3.Idx == -1,
		card4.Idx == 2,
		card1.Key == nil,
		card2.Key == nil,
		card3.Key == 7,
		card4.Key == 8,
		card1.Val == nil,
		card2.Val == "Card 4",
		card3.Val == "Card 3",
		card4.Val == card1,
	}

	test_End(funcName, conditions)

}

func case_MakeStack(funcName string) {

	test_Start(funcName, showTestText)

	// initialize variables
	map1 := map[string]int {"Alexander" : 111, "Breton" : 222, "Charles" : 333}
	arrKeys := []string {"Alex", "Bre", "Charlie"}
	arrVals := []int {11, 22, 33}

	// make array of arrVals times three
	var arrValsTimesThree []interface{}
	for i := 0; i < 3; i++ {
		for j := range arrVals {
			arrValsTimesThree = append(arrValsTimesThree, arrVals[j])
		}
	}

	// to stacks (in order of conditions listed in documentation)
	stack1 := MakeStack(map1)
	stack2 := MakeStack(arrVals)
	stack3 := MakeStack(arrKeys, arrVals)
	stack4 := MakeStack(nil, arrKeys)
	stack5 := MakeStack(arrVals, nil, 3)
	stack6 := MakeStack()

	conditions := []bool{
		test_IdxsAreGood(stack1),
		test_IdxsAreGood(stack2),
		test_IdxsAreGood(stack3),
		test_IdxsAreGood(stack4),
		test_IdxsAreGood(stack5),
		test_IdxsAreGood(stack6),
		test_LenAndSize(stack1, 3),
		test_LenAndSize(stack2, 3),
		test_LenAndSize(stack3, 3),
		test_LenAndSize(stack4, 3),
		test_LenAndSize(stack5, 9),
		test_LenAndSize(stack6, 0),
		test_StackEqualArray(stack1, nil, nil, map1),
		test_StackEqualArray(stack2, arrVals, nil, nil),
		test_StackEqualArray(stack3, arrVals, arrKeys, nil),
		test_StackEqualArray(stack4, nil, arrKeys, nil),
		test_StackEqualArray(stack5, arrValsTimesThree, nil, nil),
	}

	test_End(funcName, conditions)

}

func case_MakeStackMatrix(funcName string) {

	test_Start(funcName, showTestText)

	// initialize variables
/*
	// to stacks (in order of conditions listed in documentation)
	stack1  := MakeStackMatrix(map1) // BAD
	stack2  := MakeStackMatrix(arrVals) // BAD
	stack3  := MakeStackMatrix(arrKeys, arrVals) // BAD
	stack4  := MakeStackMatrix(nil, arrKeys) // BAD
	stack5  := MakeStackMatrix()
	stack6  := MakeStackMatrix() // BAD
	stack7  := MakeStackMatrix() // BAD
	stack8  := MakeStackMatrix() // BAD
	stack9  := MakeStackMatrix() // BAD
	stack10 := MakeStackMatrix() // BAD
*/
	conditions := []bool{
		
	}

	test_End(funcName, conditions)

}

/** Executes all case tests */
func Run(_showTestText bool) {

	showTestText = _showTestText

	fmt.Println("- BEGINNING TESTS (fix failures/errors in descending order)")

	case_MakeCard("MakeCard")
	case_MakeStack("MakeStack")
	//case_MakeStackMatrix("MakeStackMatrix")

}
