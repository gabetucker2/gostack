package casetests

import (
	"fmt"

	. "github.com/gabetucker2/gostack"//lint:ignore ST1001 Ignore warning
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

	// to stacks (in order of conditions listed in documentation)
	stack1 := MakeStack(map1)
	stack2 := MakeStack(arrVals)
	stack3 := MakeStack(arrKeys, arrVals)
	stack4 := MakeStack(nil, arrKeys)
	stack5 := MakeStack(arrVals, nil, 3)
	stack6 := MakeStack()

	// make array of arrVals times three (what stack5 should yield)
	var arrValsTimesThree []any
	for i := 0; i < 3; i++ {
		for j := range arrVals {
			arrValsTimesThree = append(arrValsTimesThree, arrVals[j])
		}
	}

	conditions := []bool{
		test_StackProperties(stack6, []int{0}),
		test_IdxsAreGood(stack1),
		test_IdxsAreGood(stack2),
		test_IdxsAreGood(stack3),
		test_IdxsAreGood(stack4),
		test_IdxsAreGood(stack5),
		test_StackProperties(stack1, []int{3}),
		test_StackProperties(stack2, []int{3}),
		test_StackProperties(stack3, []int{3}),
		test_StackProperties(stack4, []int{3}),
		test_StackProperties(stack5, []int{9}),
		test_StackEqualArrayOrMap(stack1, nil, nil, map1),
		test_StackEqualArrayOrMap(stack2, arrVals, nil, nil),
		test_StackEqualArrayOrMap(stack3, arrVals, arrKeys, nil),
		test_StackEqualArrayOrMap(stack4, nil, arrKeys, nil),
		test_StackEqualArrayOrMap(stack5, arrValsTimesThree, nil, nil),
	}

	test_End(funcName, conditions)

}

func case_MakeStackMatrix(funcName string) {

	test_Start(funcName, showTestText)

	// initialize variables
	matrixShape := []int { 3, 2 }

	/*shallowMap := map[string]int{
		"Alex": 111,
		"Bre": 222,
		"Charles": 333,
		"David": 444,
		"Elliot": 555,
		"Ferguson": 666,
    }*/

	arrShallowKeys := []string {"Alex", "Bre", "Charles", "David", "Elliot", "Ferguson"}
	arrShallowVals := []int {111, 222, 333, 444, 555, 666}

	/*deepMap := map[string]map[string]int{
        "First": {
            "Alex": 111,
            "Bre": 222,
        },
        "Second": {
            "Charles": 333,
            "David": 444,
        },
        "Third": {
            "Elliot": 555,
            "Ferguson": 666,
        },
    }*/

	//arrDeepKeys := [][]string {{"Alex", "Bre"}, {"Charles", "David"}, {"Elliot", "Ferguson"}}
	//arrDeepVals := [][]int {{111, 222}, {333, 444}, {555, 666}}

	// to stacks (in order of conditions listed in documentation)

	//TODO: fix map support
	//TODO: implement deep support

	//stack1  := MakeStackMatrix(deepMap) // BAD
	//stack2  := MakeStackMatrix(arrDeepVals) // BAD
	//stack3  := MakeStackMatrix(arrDeepKeys, arrDeepVals) // BAD
	//stack4  := MakeStackMatrix(nil, arrDeepKeys) // BAD
	stack5  := MakeStackMatrix()
	//stack6  := MakeStackMatrix(shallowMap, nil, matrixShape) // BAD
	stack7  := MakeStackMatrix(arrShallowVals, nil, matrixShape) // CHECK
	stack8  := MakeStackMatrix(arrShallowKeys, arrShallowVals, matrixShape) // CHECK
	stack9  := MakeStackMatrix(nil, arrShallowKeys, matrixShape) // CHECK
	stack10 := MakeStackMatrix(nil, nil, matrixShape) // CHECK

	conditions := []bool{
		test_IdxsAreGood(stack7),
		test_IdxsAreGood(stack8),
		test_IdxsAreGood(stack9),
		test_IdxsAreGood(stack10),
		test_StackProperties(stack5, []int{0}),
		test_StackProperties(stack7, matrixShape),
		test_StackProperties(stack8, matrixShape),
		test_StackProperties(stack9, matrixShape),
		test_StackProperties(stack10, matrixShape),
	}

	test_End(funcName, conditions)

}

func case_stack_StripStackMatrix(funcName string) {

	test_Start(funcName, showTestText)

	stackMatrix := test_SampleStackMatrix()

	conditions := []bool{
		stackMatrix.StripStackMatrix() == stackMatrix,
		stackMatrix.StripStackMatrix(0, 0).Cards[0] == testCardA,
		stackMatrix.StripStackMatrix(0, 1).Cards[0] == testCardB,
		stackMatrix.StripStackMatrix(1, 0).Cards[0] == testCardC,
		stackMatrix.StripStackMatrix(1, 1).Cards[0] == testCardD,
		stackMatrix.StripStackMatrix(0) == MakeStack([]*Card {testCardA, testCardB}),
		stackMatrix.StripStackMatrix(1) == MakeStack([]*Card {testCardC, testCardD}),
		stackMatrix.StripStackMatrix(0, []int {0, 1}) == MakeStack([]*Card {testCardA, testCardB}),
		stackMatrix.StripStackMatrix(1, []int {0, 1}) == MakeStack([]*Card {testCardC, testCardD}),
		stackMatrix.StripStackMatrix([]int {0, 1}, 0) == MakeStack([]*Card {testCardA, testCardC}),
		stackMatrix.StripStackMatrix([]int {0, 1}, 1) == MakeStack([]*Card {testCardB, testCardD}),
	}

	test_End(funcName, conditions)
	
}

func case_stack_ToArray(funcName string) {

	test_Start(funcName, showTestText)

	array := test_SampleStack().ToArray()

	conditions := []bool{
		len(array) == 3,
		array[0] == testCardA,
		array[1] == testCardB,
		array[2] == testCardC,
	}

	test_End(funcName, conditions)
	
}

func case_stack_ToMap(funcName string) {

	test_Start(funcName, showTestText)

	m := test_SampleStack().ToMap()

	conditions := []bool{
		len(m) == 3,
		m["Key1"] == "Card A",
		m["Key2"] == "Card B",
		m["Key3"] == "Card C",
	}

	test_End(funcName, conditions)
	
}

func case_stack_ToMatrix(funcName string) {

	test_Start(funcName, showTestText)

	mat1 := test_SampleStackMatrix().ToMatrix(1).([]any)
	mat2 := test_SampleStackMatrix().ToMatrix(2).([][]any)
	mat3 := test_SampleStackMatrix().ToMatrix(-1).([][]any)

	conditions := []bool{
		len(mat1) == 2,
		len(mat2) == 2,
		len(mat3) == 2,
		len(mat1[0].([]any)) == 0,
		len(mat1[1].([]any)) == 0,
		len(mat2[0]) == 2,
		len(mat2[1]) == 2,
		len(mat3[0]) == 2,
		len(mat3[1]) == 2,
		// TODO: test using Equals function here
	}

	test_End(funcName, conditions)
	
}

func case_stack_Empty(funcName string) {

	test_Start(funcName, showTestText)

	stack1 := test_SampleStack().Empty()
	stack2 := test_SampleStackMatrix().Empty()

	conditions := []bool{
		test_StackProperties(stack1, []int {0}, 1),
		test_StackProperties(stack2, []int {0}, 1),
	}

	test_End(funcName, conditions)
	
}

func case_card_Clone(funcName string) {

	test_Start(funcName, showTestText)

	cardA := MakeCard("Original", "Original")
	cardAClone := cardA.Clone(CLONE_True, CLONE_True)
	cardAClone.Key = "New"
	cardAClone.Val = "New"
	
	cardB := MakeCard("Original", "Original")
	cardBClone := cardB.Clone(CLONE_True, CLONE_False)
	cardBClone.Key = "New"
	cardBClone.Val = "New"
	
	cardC := MakeCard("Original", "Original")
	cardCClone := cardC.Clone(CLONE_False, CLONE_True)
	cardCClone.Key = "New"
	cardCClone.Val = "New"
	
	cardD := MakeCard("Original", "Original")
	cardDClone := cardD.Clone(CLONE_False, CLONE_False)
	cardDClone.Key = "New"
	cardDClone.Val = "New"

	conditions := []bool{
		cardA.Idx == -1,
		cardA.Key == "Original",
		cardA.Val == "Original",
		cardAClone.Idx == -1,
		cardAClone.Key == "New",
		cardAClone.Val == "New",

		cardB.Idx == -1,
		cardB.Key == "Original",
		cardB.Val == "New",
		cardBClone.Idx == -1,
		cardBClone.Key == "New",
		cardBClone.Val == "New",

		cardC.Idx == -1,
		cardC.Key == "New",
		cardC.Val == "Original",
		cardCClone.Idx == -1,
		cardCClone.Key == "New",
		cardCClone.Val == "New",

		cardD.Idx == -1,
		cardD.Key == "New",
		cardD.Val == "New",
		cardDClone.Idx == -1,
		cardDClone.Key == "New",
		cardDClone.Val == "New",
	}

	test_End(funcName, conditions)
	
}

func case_stack_Clone(funcName string) {

	test_Start(funcName, showTestText)

	// if card.Clone() works, we expect stack.Clone() to work since it calls card.Clone(), meaning we only need to test for non parameter-related functionality
	stackA := MakeStack([]string {"Original", "Original"}, []string {"Original", "Original"})
	stackAClone := stackA.Clone(CLONE_True, CLONE_False)
	stackAClone.Get(FIND_First).Key = "New"
	stackAClone.Get(FIND_Last).Key = "New"
	stackAClone.Get(FIND_First).Val = "New"
	stackAClone.Get(FIND_Last).Val = "New"

	conditions := []bool{
		test_StackProperties(stackA, []int {2}, 1),

		stackA.Get(FIND_First).Idx == 0,
		stackA.Get(FIND_Last).Idx == 1,
		stackA.Get(FIND_First).Key == "Original",
		stackA.Get(FIND_Last).Key == "Original",
		stackA.Get(FIND_First).Val == "New",
		stackA.Get(FIND_Last).Val == "New",

		stackAClone.Get(FIND_First).Idx == 0,
		stackAClone.Get(FIND_Last).Idx == 1,
		stackAClone.Get(FIND_First).Key == "New",
		stackAClone.Get(FIND_Last).Key == "New",
		stackAClone.Get(FIND_First).Val == "New",
		stackAClone.Get(FIND_Last).Val == "New",
	}

	test_End(funcName, conditions)
	
}

func case_stack_Unique(funcName string) {

	test_Start(funcName, showTestText)

	myStackKeys := MakeStack(nil, []string {"Person", "Place", "Person", "Thing", "Person"})
	myStackVals := MakeStack([]string {"Person", "Place", "Person", "Thing", "Person"})

	filteredByKey := myStackKeys.Clone().Unique(TYPE_Key)
	filteredByVal := myStackVals.Clone().Unique(TYPE_Val)

	conditions := []bool{
		filteredByKey.Size == 3,
		filteredByVal.Size == 3,
		filteredByKey.Cards[0].Key == "Person",
		filteredByKey.Cards[1].Key == "Place",
		filteredByKey.Cards[2].Key == "Thing",
		filteredByVal.Cards[0].Val == "Person",
		filteredByVal.Cards[1].Val == "Place",
		filteredByVal.Cards[2].Val == "Thing",
	}

	test_End(funcName, conditions)
	
}

func case_card_Equals(funcName string) {

	test_Start(funcName, showTestText)

	keyVar := "MyKey"
	valVar := "MyVal"

	card1 := MakeCard("MyKey", "MyVal") // Idx == -1
	card2 := MakeCard("MyKey", "MyVal", 0)

	card3 := MakeCard(keyVar, valVar)
	card4 := MakeCard(keyVar, valVar)

	conditions := []bool{

		// compare by object
		card1.Equals(card2),

		// test whether idx parameter works
		card1.Equals(card2, nil, nil, COMPARE_False),
		!card1.Equals(card2, nil, nil, COMPARE_True),

		// test whether matchByTypes work
		card1.Equals(card3),
		card1.Equals(card3, MATCHBY_Object, MATCHBY_Object),
		!card1.Equals(card3, MATCHBY_Object, MATCHBY_Reference),
		!card1.Equals(card3, MATCHBY_Reference, MATCHBY_Object),
		!card1.Equals(card3, MATCHBY_Reference, MATCHBY_Reference),
		card3.Equals(card4, MATCHBY_Object, MATCHBY_Reference),
		card3.Equals(card4, MATCHBY_Reference, MATCHBY_Object),
		card3.Equals(card4, MATCHBY_Reference, MATCHBY_Reference),
		
	}

	test_End(funcName, conditions)
	
}

func case_stack_Equals(funcName string) {

	test_Start(funcName, showTestText)

	// since we've already tested the properties of card.Equals(), and stack invokes card.Equals(),
	// we don't need as thorough of a test for non-stack-specific parameters

	stack1 := MakeStack([]string {"Hello", "Hey"})
	stack2 := MakeStack([]string {"Hello", "Hey"})
	stack3 := stack1
	stack4 := MakeStack([]string {"Hi", "Hey"})
	
	deep1 := MakeStackMatrix([]string {"Hello", "Hey", "Howdy", "Hi"}, nil, []int {2, 2})
	deep2 := MakeStackMatrix([]string {"Hello", "Hey", "Howdy", "Hi"}, nil, []int {2, 2})
	deep3 := deep1
	deep4 := MakeStackMatrix([]string {"Hello", "Hey", "Howdy", "Heyo"}, nil, []int {2, 2})

	string1 := "Hi"
	string2 := "Hello"
	objRefStack1 := MakeStack([]string {"Hi", "Hello"})
	objRefStack2 := MakeStack([]string {string1, string2})

	deeper1 := MakeStack(objRefStack1)
	deeper2 := MakeStackMatrix(objRefStack2)
	deeper3 := MakeStackMatrix(objRefStack2)

	conditions := []bool{
		
		// test whether stack compare object vs reference works
		stack1.Equals(stack2, COMPARE_False),
		stack1.Equals(stack2, COMPARE_True),
		stack1.Equals(stack2, COMPARE_True, MATCHBY_Object),
		!stack1.Equals(stack2, COMPARE_True, MATCHBY_Reference),
		stack1.Equals(stack3, COMPARE_True, MATCHBY_Reference),
		!stack1.Equals(stack4, COMPARE_True, MATCHBY_Object),
		!stack1.Equals(stack4, COMPARE_True, MATCHBY_Reference),

		// test whether the same tests hold true for a deepsearch-true equivalent
		deep1.Equals(deep2, COMPARE_False, nil, DEEPSEARCH_True),
		deep1.Equals(deep2, COMPARE_True, nil, DEEPSEARCH_True),
		deep1.Equals(deep2, COMPARE_True, MATCHBY_Object, DEEPSEARCH_True),
		!deep1.Equals(deep2, COMPARE_True, MATCHBY_Reference, DEEPSEARCH_True),
		deep1.Equals(deep3, COMPARE_True, MATCHBY_Reference, DEEPSEARCH_True),
		!deep1.Equals(deep4, COMPARE_True, MATCHBY_Object, DEEPSEARCH_True),
		!deep1.Equals(deep4, COMPARE_True, MATCHBY_Reference, DEEPSEARCH_True),

		// test depth
		deeper1.Equals(deeper2, nil, nil, DEEPSEARCH_True, -1),
		deeper1.Equals(deeper2, nil, nil, DEEPSEARCH_True, 2),
		deeper1.Equals(deeper2, nil, nil, DEEPSEARCH_True, 1),
		!deeper1.Equals(deeper2, nil, nil, DEEPSEARCH_True, 1, nil, MATCHBY_Reference),
		deeper2.Equals(deeper3, nil, nil, DEEPSEARCH_True, 1, nil, MATCHBY_Reference),

	}

	test_End(funcName, conditions)
	
}

func case_stack_Shuffle(funcName string) {

	test_Start(funcName, showTestText)

	

	conditions := []bool{
		
		

	}

	test_End(funcName, conditions)
	
}

func case_stack_Flip(funcName string) {

	test_Start(funcName, showTestText)

	conditions := []bool{
		
		

	}

	test_End(funcName, conditions)
	
}

func case_card_Print(funcName string) {

	test_Start(funcName, showTestText)

	MakeCard("CardKey", "CardVal", 420).Print()
	
	conditions := []bool{
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_stack_Print(funcName string) {

	test_Start(funcName, showTestText)

	MakeStack([]string {"ShallowKeyFirst", "ShallowKeySecond"}, []string {"ShallowValFirst", "ShallowValSecond"}).Print()

	MakeStackMatrix([]string {"DeepKeyFirst", "DeepKeySecond", "DeepKeyThird", "DeepKeyFourth"}, []string {"DeepValFirst", "DeepValSecond", "DeepValThird", "DeepValFourth"}, []int {2, 2}).Print()
	
	conditions := []bool{
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_stack_Lambda(funcName string) {

	test_Start(funcName, showTestText)

	stackToFlip := MakeStack([]int {1, 2, 3, 4})
	stackToCountKeysOver30 := MakeStack(nil, []int {5, 10, 20, 25, 50})
	stackToAdd := MakeStack([]int {1, 2, 3, 4})

	// flipper
	stackToFlip.Lambda(func(card *Card, stack *Stack, _ ...any) {
		// moves each card, from first to last, to the first position in the stack
		newCards := []*Card {card}
		for _, c := range stack.Cards {
			if c != card {
				newCards = append(newCards, c)
			}
		}
		stack.Cards = newCards
	})

	// get amount of cards with keys under 30
	keyCount := stackToCountKeysOver30.Lambda(func(card *Card, stack *Stack, ret any) {
		if card.Key.(int) < 30 { ret = ret.(int) + 1 }
	})

	// add each card by its previous value
	stackToAdd.Lambda(func(card *Card, stack *Stack, _ any, previousVal any) {
		card.Val = card.Val.(int) + previousVal.(int)
		previousVal = card.Val
	})
	
	conditions := []bool{
		stackToFlip.Equals(MakeStack([]int {4, 3, 2, 1})),
		keyCount == 4,
		stackToAdd.Equals(MakeStack(1, 3, 6, 10)),
	}

	test_End(funcName, conditions)
	
}

/** Executes all case tests */
func Run(_showTestText bool) {

	showTestText = _showTestText

	fmt.Println("- BEGINNING TESTS (fix failures/errors in descending order)")

	// NON-GENERALIZED FUNCTIONS
	case_MakeCard("MakeCard") // GOOD
	case_MakeStack("MakeStack") // BAD
	case_MakeStackMatrix("MakeStackMatrix") // BAD
	case_stack_StripStackMatrix("stack.StripStackMatrix") // BAD
	case_stack_ToArray("stack.ToArray") // BAD
	case_stack_ToMap("stack.ToMap") // BAD
	case_stack_ToMatrix("stack.ToMatrix") // BAD
	case_stack_Empty("stack.Empty") // BAD
	case_card_Clone("card.Clone") // BAD
	case_stack_Clone("stack.Clone") // BAD
	case_stack_Unique("stack.Unique") // BAD
	case_card_Equals("card.Equals") // BAD
	case_stack_Equals("stack.Equals") // BAD
	case_stack_Shuffle("stack.Shuffle") // BAD
	case_stack_Flip("stack.Flip") // BAD
	case_card_Print("card.Print") // BAD
	case_stack_Print("stack.Print") // BAD
	case_stack_Lambda("stack.Lambda") // BAD
	
	// GENERALIZED FUNCTIONS
	/*case_stack_Add("stack.Add") // BAD
	case_stack_Move("stack.Move") // BAD
	case_stack_Has("stack.Has") // BAD
	case_stack_Get("stack.Get") // BAD
	case_stack_GetMany("stack.GetMany") // BAD
	case_stack_Replace("stack.Replace") // BAD
	case_stack_ReplaceMany("stack.ReplaceMany") // BAD
	case_stack_Update("stack.Update") // BAD
	case_stack_UpdateMany("stack.UpdateMany") // BAD
	case_stack_Extract("stack.Extract") // BAD
	case_stack_ExtractMany("stack.ExtractMany") // BAD
	case_stack_Remove("stack.Remove") // BAD
	case_stack_RemoveMany("stack.RemoveMany") // BAD*/

}
