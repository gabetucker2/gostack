package casetests

import (
	"fmt"

	"github.com/gabetucker2/gogenerics"
	. "github.com/gabetucker2/gostack" //lint:ignore ST1001 Ignore warning
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
	var myStr any//lint:ignore S1021 Ignore warning
	myStr = "Hello"
	card5 := MakeCard(&myStr)
	card6 := MakeCard(&myStr)
	card7 := MakeCard(myStr)
	gogenerics.SetPointer(card6.Val, "Hi")

	// test whether updating fields does so by object
	card2.Val = "Card 4"
	card3.Val = 7

	conditions := []bool{
		card1.Idx == -1,
		card2.Idx == -1,
		card3.Idx == -1,
		card4.Idx == 2,
		card5.Idx == -1,
		card6.Idx == -1,
		card7.Idx == -1,
		
		card1.Key == nil,
		card2.Key == nil,
		card3.Key == card2,
		card4.Key == 8,
		card5.Key == nil,
		card6.Key == nil,
		card7.Key == nil,

		card1.Val == nil,
		card2.Val == "Card 4",
		card3.Val == 7,
		card4.Val == card1,
		*card5.Val.(*any) == "Hi",
		*card6.Val.(*any) == "Hi",
		card7.Val == "Hello",
	}

	test_End(funcName, conditions)

}

func case_MakeStack(funcName string) {

	test_Start(funcName, showTestText)

	// initialize variables
	map1 := map[string]int {"Alexander" : 111, "Breton" : 222, "Charles" : 333}
	arrKeys := []string {"Alex", "Bre", "Charlie"}
	arrVals := []int {11, 22, 33}
	arrCardToVals := []*Card {new(Card), new(Card), new(Card)}
	
	// to stacks (in order of conditions listed in documentation)
	stack1 := MakeStack(map1)
	stack2 := MakeStack(arrVals)
	stack3 := MakeStack(arrKeys, arrVals)
	stack4 := MakeStack(nil, arrKeys)
	stack5 := MakeStack(arrVals, nil, 3)
	stack6 := MakeStack(nil, nil, 10)
	stack7 := MakeStack(stack2.Cards) // should equal stack2
	stack8 := MakeStack(stack2.Cards, nil, nil, true) // should be a stack of cards pointing to the cards of stack2
	stack9 := MakeStack()

	// make array of arrVals times three (what stack5 should yield)
	var arrValsTimesThree []any
	for i := 0; i < 3; i++ {
		for j := range arrVals {
			arrValsTimesThree = append(arrValsTimesThree, arrVals[j])
		}
	}
	
	// stack of Cards of Cards
	for i := 0; i < 3; i++ {
		arrCardToVals[i].Val = stack2.Cards[i]
		arrCardToVals[i].Idx = i
	}

	conditions := []bool{
		test_IdxsAreGood(stack1),
		test_IdxsAreGood(stack2),
		test_IdxsAreGood(stack3),
		test_IdxsAreGood(stack4),
		test_IdxsAreGood(stack5),
		test_IdxsAreGood(stack6),
		test_IdxsAreGood(stack7),
		test_IdxsAreGood(stack8),

		test_StackProperties(stack1, []int{3}),
		test_StackProperties(stack2, []int{3}),
		test_StackProperties(stack3, []int{3}),
		test_StackProperties(stack4, []int{3}),
		test_StackProperties(stack5, []int{9}),
		test_StackProperties(stack6, []int{10}),
		test_StackProperties(stack7, []int{3}),
		test_StackProperties(stack8, []int{3}),

		test_StackEqualArrayOrMap(stack1, nil, nil, map1),
		test_StackEqualArrayOrMap(stack2, arrVals, nil, nil),
		test_StackEqualArrayOrMap(stack3, arrVals, arrKeys, nil),
		test_StackEqualArrayOrMap(stack4, nil, arrKeys, nil),
		test_StackEqualArrayOrMap(stack5, arrValsTimesThree, nil, nil),
		test_StackEqualArrayOrMap(stack6, nil, nil, nil),
		test_StackEqualArrayOrMap(stack7, arrVals, nil, nil),
		test_StackEqualArrayOrMap(stack8, stack2.Cards, nil, nil),
		
		test_StackProperties(stack9, []int{0}),
	}
	
	test_End(funcName, conditions)

}

func case_MakeStackMatrix(funcName string) {

	test_Start(funcName, showTestText)

	// initialize variables
	matrixShape := []int { 3, 2 }

	shallowMap := map[string]int{
		"Alex": 111,
		"Bre": 222,
		"Charles": 333,
		"David": 444,
		"Elliot": 555,
		"Ferguson": 666,
    }

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

	//TODO: add stackequalsarrayormap tests
	//TODO: fix map support
	//TODO: implement deep support
	//TODO: implement overrideCard support
	//TODO: ensure cases of irregularly-shaped matrices

	//stack1  := MakeStackMatrix(deepMap) // BAD
	//stack2  := MakeStackMatrix(arrDeepVals) // BAD
	//stack3  := MakeStackMatrix(arrDeepKeys, arrDeepVals) // BAD
	//stack4  := MakeStackMatrix(nil, arrDeepKeys) // BAD
	stack5  := MakeStackMatrix() // GOOD
	stack6  := MakeStackMatrix(shallowMap, nil, matrixShape) // BAD
	stack7  := MakeStackMatrix(arrShallowVals, nil, matrixShape) // BAD
	stack8  := MakeStackMatrix(arrShallowKeys, arrShallowVals, matrixShape) // BAD
	stack9  := MakeStackMatrix(nil, arrShallowKeys, matrixShape) // BAD
	stack10 := MakeStackMatrix(nil, nil, matrixShape) // BAD

	stack7Test := MakeStack(nil, nil, 3)
	for i := 0; i < 3; i++ {
		subStack := MakeStack(nil, nil, 2)
		stack7Test.Cards[i].Val = subStack
		for j := 0; j < 2; j++ {
			c := subStack.Cards[j]
			c.Val = arrShallowVals[i*2 + j]
		}
	}

	conditions := []bool{
		test_IdxsAreGood(stack6),
		test_IdxsAreGood(stack7),
		test_IdxsAreGood(stack8),
		test_IdxsAreGood(stack9),
		test_IdxsAreGood(stack10),

		test_StackProperties(stack6, matrixShape, 2),
		test_StackProperties(stack7, matrixShape, 2),
		test_StackProperties(stack8, matrixShape, 2),
		test_StackProperties(stack9, matrixShape, 2),
		test_StackProperties(stack10, matrixShape, 2),

		test_StackProperties(stack5, []int{0}),
	}

	test_End(funcName, conditions)

}

func case_stack_StripStackMatrix(funcName string) {

	test_Start(funcName, showTestText)

	stackMatrix := test_SampleStackMatrix()
	// fmt.Println("this is stackMatrix")
	// stackMatrix.Print()

	// fmt.Println("this is STRIPstackMatrix")
	// stackMatrix.StripStackMatrix().Print()
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

	arrayVals1 := test_SampleStack(true).ToArray()
	arrayVals2 := test_SampleStack(true).ToArray(RETURN_Vals)
	arrayKeys := test_SampleStack(true).ToArray(RETURN_Keys)
	arrayIdxs := test_SampleStack(true).ToArray(RETURN_Idxs)
	arrayCards := test_SampleStack(false).ToArray(RETURN_Cards)
	
	conditions := []bool{
		len(arrayVals1) == 3,
		len(arrayVals2) == 3,
		len(arrayKeys) == 3,
		len(arrayIdxs) == 3,
		len(arrayCards) == 3,

		arrayVals1[0] == testCardA.Val,
		arrayVals1[1] == testCardB.Val,
		arrayVals1[2] == testCardC.Val,

		arrayVals2[0] == testCardA.Val,
		arrayVals2[1] == testCardB.Val,
		arrayVals2[2] == testCardC.Val,

		arrayKeys[0] == testCardA.Key,
		arrayKeys[1] == testCardB.Key,
		arrayKeys[2] == testCardC.Key,

		arrayIdxs[0] == testCardA.Idx,
		arrayIdxs[1] == testCardB.Idx,
		arrayIdxs[2] == testCardC.Idx,

		arrayCards[0].(*Card) == testCardA,
		arrayCards[1].(*Card) == testCardB,
		arrayCards[2].(*Card) == testCardC,
	}

	test_End(funcName, conditions)
	
}

func case_stack_ToMap(funcName string) {

	test_Start(funcName, showTestText)

	m := test_SampleStack(true).ToMap()

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

	stack1 := test_SampleStack(true).Empty()
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
	fmt.Println(cardB.Val)
	fmt.Println(cardB.Val.(string) == "New")
	
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
		test_StackProperties(stackA, []int {2}, 1), // 1

		stackA.Get(FIND_First).Idx == 0, // 2
		stackA.Get(FIND_Last).Idx == 1, // 3
		stackA.Get(FIND_First).Key == "Original", // 4
		stackA.Get(FIND_Last).Key == "Original", // 5
		stackA.Get(FIND_First).Val == "New", // 6
		stackA.Get(FIND_Last).Val == "New", // 7

		stackAClone.Get(FIND_First).Idx == 0, // 8
		stackAClone.Get(FIND_Last).Idx == 1, // 9
		stackAClone.Get(FIND_First).Key == "New", // 10
		stackAClone.Get(FIND_Last).Key == "New", // 11
		stackAClone.Get(FIND_First).Val == "New", // 12
		stackAClone.Get(FIND_Last).Val == "New", // 13
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
		filteredByKey.Size == 3, // 1
		filteredByVal.Size == 3, // 2
		filteredByKey.Cards[0].Key == "Person", // 3
		filteredByKey.Cards[1].Key == "Place", // 4
		filteredByKey.Cards[2].Key == "Thing", // 5
		filteredByVal.Cards[0].Val == "Person", // 6
		filteredByVal.Cards[1].Val == "Place", // 7
		filteredByVal.Cards[2].Val == "Thing", // 8
	}

	test_End(funcName, conditions)
	
}

func case_card_Equals(funcName string) {

	test_Start(funcName, showTestText)

	// compare parameter tests
	card1 := MakeCard("MyKey", "MyVal") // Idx == -1
	card2 := MakeCard("MyKey", "MyVal", 0)

	card3 := MakeCard("MyKey", "MyVal1")
	card4 := MakeCard("MyKey", "MyVal2")

	card5 := MakeCard("MyKey1", "MyVal")
	card6 := MakeCard("MyKey2", "MyVal")

	// pointerKey parameter tests
	var keyVar any//lint:ignore S1021 Ignore warning
	keyVar = "MyKey"
	cardA := MakeCard(nil, keyVar)
	cardB := MakeCard(nil, keyVar)
	cardC := MakeCard(nil, &keyVar)
	cardD := MakeCard(nil, &keyVar)

	// pointerKey parameter tests
	var valVar any//lint:ignore S1021 Ignore warning
	valVar = "MyVal"
	cardE := MakeCard(valVar)
	cardF := MakeCard(valVar)
	cardG := MakeCard(&valVar)
	cardH := MakeCard(&valVar)

	conditions := []bool{

		// compare by object
		card1.Equals(card2), // 1

		// test whether idx parameter works
		card1.Equals(card2, nil, nil, COMPARE_False, COMPARE_False, COMPARE_False), // 2
		!card1.Equals(card2, nil, nil, COMPARE_True, COMPARE_False, COMPARE_False), // 3

		// test whether val parameter works
		card1.Equals(card3, nil, nil, COMPARE_False, COMPARE_False, COMPARE_False), // 4
		!card1.Equals(card4, nil, nil, COMPARE_False, COMPARE_True, COMPARE_False), // 5

		// test whether key parameter works
		card1.Equals(card5, nil, nil, COMPARE_False, COMPARE_False, COMPARE_False), // 6
		!card1.Equals(card6, nil, nil, COMPARE_False, COMPARE_False, COMPARE_True), // 7

		// test whether pointerTypes work for keys
		cardA.Equals(cardB, POINTER_False), // 8
		cardA.Equals(cardB, POINTER_True), // 9
		!cardB.Equals(cardC, POINTER_False), // 10
		cardB.Equals(cardC, POINTER_True), // 11
		cardC.Equals(cardD, POINTER_False), // 12
		cardC.Equals(cardD, POINTER_True), // 13

		// test whether pointerTypes work for vals
		cardE.Equals(cardF, nil, POINTER_False), // 14
		cardE.Equals(cardF, nil, POINTER_True), // 15
		!cardF.Equals(cardG, nil, POINTER_False), // 16
		cardF.Equals(cardG, nil, POINTER_True), // 17
		cardG.Equals(cardH, nil, POINTER_False), // 18
		cardG.Equals(cardH, nil, POINTER_True), // 19
		
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
		stack1.Equals(stack2, COMPARE_True, POINTER_False),
		!stack1.Equals(stack2, COMPARE_True, POINTER_True),
		stack1.Equals(stack3, COMPARE_True, POINTER_True),
		!stack1.Equals(stack4, COMPARE_True, POINTER_False),
		!stack1.Equals(stack4, COMPARE_True, POINTER_True),

		// test whether the same tests hold true for a deepsearch-true equivalent
		deep1.Equals(deep2, COMPARE_False, nil, DEEPSEARCH_True),
		deep1.Equals(deep2, COMPARE_True, nil, DEEPSEARCH_True),
		deep1.Equals(deep2, COMPARE_True, POINTER_False, DEEPSEARCH_True),
		!deep1.Equals(deep2, COMPARE_True, POINTER_True, DEEPSEARCH_True),
		deep1.Equals(deep3, COMPARE_True, POINTER_True, DEEPSEARCH_True),
		!deep1.Equals(deep4, COMPARE_True, POINTER_False, DEEPSEARCH_True),
		!deep1.Equals(deep4, COMPARE_True, POINTER_True, DEEPSEARCH_True),

		// test depth
		deeper1.Equals(deeper2, nil, nil, DEEPSEARCH_True, -1),
		deeper1.Equals(deeper2, nil, nil, DEEPSEARCH_True, 2),
		deeper1.Equals(deeper2, nil, nil, DEEPSEARCH_True, 1),
		!deeper1.Equals(deeper2, nil, nil, DEEPSEARCH_True, 1, nil, POINTER_True),
		deeper2.Equals(deeper3, nil, nil, DEEPSEARCH_True, 1, nil, POINTER_True),

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

	// comment out to clean up console:

	// MakeCard("CardKey", "CardVal", 420).Print()
	// MakeCard("CardKey", "CardVal", 420).Print(2)
	
	conditions := []bool{
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_stack_Print(funcName string) {

	test_Start(funcName, showTestText)

	// comment out to clean up console:

	// MakeStack([]string {"ShallowKeyFirst", "ShallowKeySecond"}, []string {"ShallowValFirst", "ShallowValSecond"}).Print()
	// MakeStackMatrix([]string {"DeepKeyFirst", "DeepKeySecond", "DeepKeyThird", "DeepKeyFourth"}, []string {"DeepValFirst", "DeepValSecond", "DeepValThird", "DeepValFourth"}, []int {2, 2}).Print()
	
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
	stackToAdd.Lambda(func(card *Card, stack *Stack, _ any, workingMem ...any) {
		var previousVal any
		gogenerics.UnpackVariadic(workingMem, &previousVal)
		
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
	gogenerics.RemoveUnusedError(case_MakeCard, case_MakeStack, case_MakeStackMatrix, case_stack_StripStackMatrix, case_stack_ToArray, case_stack_ToMap, case_stack_ToMatrix, case_stack_Empty, case_card_Clone,
								 case_stack_Clone, case_stack_Unique, case_card_Equals, case_stack_Equals, case_stack_Shuffle, case_stack_Flip, case_card_Print, case_stack_Print, case_stack_Lambda)

	fmt.Println("- BEGINNING TESTS (fix failures/errors in descending order)")

	// NON-GENERALIZED FUNCTIONS
	case_MakeCard("MakeCard") // GOOD
	case_card_Equals("card.Equals") // BAD
	// case_stack_Equals("stack.Equals") // BAD
	// case_MakeStack("MakeStack") // GOOD
	// case_MakeStackMatrix("MakeStackMatrix") // BAD
	// case_stack_StripStackMatrix("stack.StripStackMatrix") // BAD
	case_stack_ToArray("stack.ToArray") // GOOD
	case_stack_ToMap("stack.ToMap") // GOOD
	// case_stack_ToMatrix("stack.ToMatrix") // BAD
	case_stack_Empty("stack.Empty") // GOOD
	// case_card_Clone("card.Clone") // BAD
	// case_stack_Clone("stack.Clone") // BAD
	// case_stack_Unique("stack.Unique") // BAD
	// case_stack_Shuffle("stack.Shuffle") // BAD
	// case_stack_Flip("stack.Flip") // BAD
	case_card_Print("card.Print") // GOOD
	case_stack_Print("stack.Print") // GOOD
	// case_stack_Lambda("stack.Lambda") // BAD
	
	// GENERALIZED FUNCTIONS
	// case_stack_Add("stack.Add") // BAD
	// case_stack_Move("stack.Move") // BAD
	// case_stack_Has("stack.Has") // BAD
	// case_stack_Get("stack.Get") // BAD
	// case_stack_GetMany("stack.GetMany") // BAD
	// case_stack_Replace("stack.Replace") // BAD
	// case_stack_ReplaceMany("stack.ReplaceMany") // BAD
	// case_stack_Update("stack.Update") // BAD
	// case_stack_UpdateMany("stack.UpdateMany") // BAD
	// case_stack_Extract("stack.Extract") // BAD
	// case_stack_ExtractMany("stack.ExtractMany") // BAD
	// case_stack_Remove("stack.Remove") // BAD
	// case_stack_RemoveMany("stack.RemoveMany") // BAD

}
