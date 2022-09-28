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

	// override card support
	card1 := MakeCard("Hey")
	card2 := MakeCard(card1)
	stack10 := MakeStack([]*Card {card1}, nil, nil, false)
	stack11 := MakeStack([]*Card {card1}, nil, nil, true)

	// pointer storage support
	var name any//lint:ignore S1021 Ignore warning
	name = "Josh"
	stack12 := MakeStack([]any {&name})
	stack13 := MakeStack([]any {&name})
	gogenerics.SetPointer(stack12.Cards[0].Val, "Henry")

	// stack input to stack tests
	stack14 := MakeStack(MakeStack([]string {"Hi1"}, []string {"Hello1"}), MakeStack([]string {"Hi2"}, []string {"Hello2"}))

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
		test_IdxsAreGood(stack1), // 1
		test_IdxsAreGood(stack2), // 2
		test_IdxsAreGood(stack3), // 3
		test_IdxsAreGood(stack4), // 4
		test_IdxsAreGood(stack5), // 5
		test_IdxsAreGood(stack6), // 6
		test_IdxsAreGood(stack7), // 7
		test_IdxsAreGood(stack8), // 8

		test_StackProperties(stack1, []int{3}), // 9
		test_StackProperties(stack2, []int{3}), // 10
		test_StackProperties(stack3, []int{3}), // 11
		test_StackProperties(stack4, []int{3}), // 12
		test_StackProperties(stack5, []int{9}), // 13
		test_StackProperties(stack6, []int{10}), // 14
		test_StackProperties(stack7, []int{3}), // 15
		test_StackProperties(stack8, []int{3}), // 16

		test_StackEqualArrayOrMap(stack1, nil, nil, map1), // 17
		test_StackEqualArrayOrMap(stack2, arrVals, nil, nil), // 18
		test_StackEqualArrayOrMap(stack3, arrVals, arrKeys, nil), // 19
		test_StackEqualArrayOrMap(stack4, nil, arrKeys, nil), // 20
		test_StackEqualArrayOrMap(stack5, arrValsTimesThree, nil, nil), // 21
		test_StackEqualArrayOrMap(stack6, nil, nil, nil), // 22
		test_StackEqualArrayOrMap(stack7, arrVals, nil, nil), // 23
		test_StackEqualArrayOrMap(stack8, stack2.Cards, nil, nil), // 24
		
		test_StackProperties(stack9, []int{0}), // 25

		stack10.Cards[0] == card1, // 26
		stack11.Cards[0].Equals(card2), // 27

		gogenerics.GetPointer(stack12.Cards[0].Val) == "Henry", // 28
		gogenerics.GetPointer(stack13.Cards[0].Val) == "Henry", // 29

		stack14.Cards[0].Key == "Hello1", // 30
		stack14.Cards[0].Val == "Hello2", // 31
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
    }

	arrShallowKeys := []string {"Alex", "Bre", "Charles", "David", "Elliot", "Ferguson"}
	arrShallowVals := []int {111, 222, 333, 444, 555, 666}

	deepMap := map[string]map[string]int{
        "First": {
            "Alex": 111,
            "Bre": 222,
        },
    }

	arrDeepKeys := []any {[]any {"Alex", "Bre"}, []string {"Charles", "David"}, []any {"Elliot", "Ferguson"}}
	arrDeepVals := []any {[]int {111, 222}, []any {333, 444}, []any {555, 666}}
	
	irregularDepth := []any {10, []any {20, 30}, []any {[]int {40, 50}, []any {60, 70}}}

	// to stacks (in order of conditions listed in documentation)
	
	correctStack := MakeStack([]*Stack {MakeStack([]string {"Alex", "Bre"}, []int {111, 222}), MakeStack([]string {"Charles", "David"}, []int {333, 444}), MakeStack([]string {"Elliot", "Ferguson"}, []int {555, 666})})

	stack1 := MakeStackMatrix(deepMap)
	stack2 := MakeStackMatrix(arrDeepVals)
	stack3 := MakeStackMatrix(arrDeepKeys, arrDeepVals)
	stack4 := MakeStackMatrix(nil, arrDeepKeys)

	// shallow stacks
	stack5 := MakeStackMatrix()
	stack6 := MakeStackMatrix(shallowMap, nil, []int {1, 2})
	stack7 := MakeStackMatrix(arrShallowVals, nil, matrixShape)
	stack8 := MakeStackMatrix(arrShallowKeys, arrShallowVals, matrixShape)
	stack9 := MakeStackMatrix(nil, arrShallowKeys, matrixShape)
	stack10 := MakeStackMatrix(nil, nil, matrixShape)

	// irregular depth
	stack11 := MakeStackMatrix(irregularDepth)

	conditions := []bool{

		// deep tests
		stack1.Equals(MakeStack([]string {"First"}, []*Stack {MakeStack([]string {"Alex", "Bre"}, []int {111, 222})})) || stack1.Equals(MakeStack([]string {"First"}, []*Stack {MakeStack([]string {"Bre", "Alex"}, []int {222, 111})})), // 1
		stack2.Equals(correctStack, COMPARE_False, COMPARE_True), // 2
		stack3.Equals(correctStack, COMPARE_True, COMPARE_True), // 3
		stack4.Equals(correctStack, COMPARE_True, COMPARE_False), // 4

		// shallow tests
		stack5.Equals(MakeStack()), // 5
		stack6.Equals(MakeStack([]*Stack {MakeStack([]string {"Alex", "Bre"}, []int {111, 222})})) || stack6.Equals(MakeStack([]*Stack {MakeStack([]string {"Bre", "Alex"}, []int {222, 111})})), // 6
		stack7.Equals(correctStack, COMPARE_False, COMPARE_True), // 7
		stack8.Equals(correctStack, COMPARE_True, COMPARE_True), // 8
		stack9.Equals(correctStack, COMPARE_True, COMPARE_False), // 9
		stack10.Equals(MakeStack([]*Stack {MakeStack(nil, nil, 2), MakeStack(nil, nil, 2), MakeStack(nil, nil, 2)})), // 10

		// irregular depth
		stack11.Equals(MakeStack([]any {10, MakeStack([]int {20, 30}), MakeStack([]*Stack {MakeStack([]int {40, 50}), MakeStack([]any {60, 70})} ) } )), // 11

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
	
	// test for shallow on shallow
	matA := MakeStack([]string {"Hi", "Hello", "Hey"}).ToMatrix()
	matACorrect := []any {"Hi", "Hello", "Hey"}

	// test for deep on deep
	matB := MakeStackMatrix([]string {"Hi", "Hey", "Hoy", "Hiya"}, nil, []int {2, 2}).ToMatrix()
	matBCorrect := []any {[]any {"Hi", "Hey"}, []any {"Hoy", "Hiya"}}

	// test for shallow on deep
	matC := MakeStackMatrix([]string {"Hi", "Hey", "Hoy", "Hiya"}, nil, []int {2, 2}).ToMatrix(nil, DEEPSEARCH_False)
	matCCorrect := []any {[]any {}, []any {}}

	// test for irregular deep
	matD := MakeStack([]any {"Hi", MakeStack([]string {"Hoy", "Hiya"})}).ToMatrix()
	matDCorrect := []any {"Hi", []any {"Hoy", "Hiya"}}

	// test for deepsearchfalse <=> deepsearchtrue | depth: 1
	matE := MakeStackMatrix([]string {"Hi", "Hey", "Hoy", "Hiya"}, nil, []int {2, 2}).ToMatrix(nil, DEEPSEARCH_True, 1)
	matECorrect := []any {[]any {}, []any {}}

	// test for different return types
	matKeys := MakeStack([]string {"Hi", "Hello", "Hey"}, []string {"Hi", "He", "H"}).ToMatrix(RETURN_Keys)
	matKeysCorrect := []any {"Hi", "Hello", "Hey"}
	
	matIdxs := MakeStack([]string {"Hi", "He", "H"}, []string {"Hi", "Hello", "Hey"}).ToMatrix(RETURN_Idxs)
	matIdxsCorrect := []any {0, 1, 2}
	
	c1 := MakeCard("Hey")
	c2 := MakeCard("Hoy")
	matCards := MakeStack([]*Card {c1, c2}).ToMatrix(RETURN_Cards)
	matCardsCorrect := []any {c1, c2}


	conditions := []bool{

		// test for shallow on shallow
		MakeStackMatrix(matA).Equals(MakeStackMatrix(matACorrect)), // 1

		// test for deep on deep
		MakeStackMatrix(matB).Equals(MakeStackMatrix(matBCorrect)), // 2

		// test for shallow on deep
		MakeStackMatrix(matC).Equals(MakeStackMatrix(matCCorrect)), // 3

		// test for irregular deep
		MakeStackMatrix(matD).Equals(MakeStackMatrix(matDCorrect)), // 4

		// test for deepsearchfalse <=> deepsearchtrue | depth: 1
		MakeStackMatrix(matE).Equals(MakeStackMatrix(matECorrect)), // 5
		
		// test for different return types
		MakeStackMatrix(matKeys).Equals(MakeStackMatrix(matKeysCorrect)), // 6
		MakeStackMatrix(matIdxs).Equals(MakeStackMatrix(matIdxsCorrect)), // 7
		MakeStackMatrix(matCards).Equals(MakeStackMatrix(matCardsCorrect)), // 8
		
	}

	test_End(funcName, conditions)
	
}

func case_stack_IsRegular(funcName string) {

	test_Start(funcName, showTestText)

	// {{1, 2}, 3} == irregular/false
	stack1 := MakeStackMatrix([]any {[]any {1, 2}, 3})

	// {{1, 2}, {3}} == irregular/false
	stack2 := MakeStackMatrix([]any {[]any {1, 2}, []any {3}})

	// {{1, 2}, {3, 4}} == regular/true
	stack3 := MakeStackMatrix([]any {[]any {1, 2}, []any {3, 4}})

	// {1, 3} == regular/true
	stack4 := MakeStackMatrix([]any {1, 3})

	// {} == regular/true
	stack5 := MakeStackMatrix([]any {})

	conditions := []bool{
		!stack1.IsRegular(), // 1
		!stack2.IsRegular(), // 2
		stack3.IsRegular(), // 3
		stack4.IsRegular(), // 4
		stack5.IsRegular(), // 5
	}

	test_End(funcName, conditions)
	
}

func case_stack_Duplicate(funcName string) {

	test_Start(funcName, showTestText)

	stack1 := MakeStack([]string {"Hey", "Hi"}).Duplicate(0)
	stack2 := MakeStack([]string {"Hey", "Hi"}).Duplicate(1)
	stack3 := MakeStack([]string {"Hey", "Hi"}).Duplicate(2)

	conditions := []bool{
		stack1.Equals(MakeStack()), // 1
		stack2.Equals(MakeStack([]string {"Hey", "Hi"})), // 2
		stack3.Equals(MakeStack([]string {"Hey", "Hi", "Hey", "Hi"})), // 3
	}

	test_End(funcName, conditions)
	
}

func case_stack_Empty(funcName string) {

	test_Start(funcName, showTestText)

	stack1 := test_SampleStack(true).Empty()
	stack2 := MakeStackMatrix([]int {1, 2, 3, 4, 5, 6}, []string {"Hi", "Hey", "Hoy", "Ciao", "Heyy", "Hiya"}, []int {3, 2}).Empty()

	conditions := []bool{
		test_StackProperties(stack1, []int {0}, 1),
		test_StackProperties(stack2, []int {0}, 1),
	}

	test_End(funcName, conditions)
	
}

func case_card_Clone(funcName string) {

	test_Start(funcName, showTestText)

	cardA := MakeCard("Original", "Original", 3)
	cardAClone := cardA.Clone()
	cardAClone.Key = "New"

	conditions := []bool{
		cardA.Idx == 3, // 1
		cardA.Key == "Original", // 2
		cardA.Val == "Original", // 3
		
		cardAClone.Idx == 3, // 4
		cardAClone.Key == "New", // 5
		cardAClone.Val == "Original", // 6
	}

	test_End(funcName, conditions)
	
}

func case_stack_Clone(funcName string) {

	test_Start(funcName, showTestText)

	// shallow cloning
	stackA := MakeStack([]string {"Original", "Original"}, []string {"Original", "Original"})
	stackAClone := stackA.Clone()
	stackAClone.Cards[0].Key = "New"
	stackAClone.Cards[1].Key = "New"
	stackAClone.Cards[0].Val = "New"

	// deep cloning
	stackB := MakeStackMatrix([]string {"Original", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "Original"}, []int{2, 2})
	stackBClone := stackB.Clone()
	stackBClone.Cards[0].Val.(*Stack).Cards[0].Key = "New"
	stackBClone.Cards[1].Val.(*Stack).Cards[1].Val = "New"

	// shallow clone stackmatrix
	stackC := MakeStackMatrix([]string {"Original", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "Original"}, []int{2, 2})
	stackCClone := stackC.Clone(DEEPSEARCH_False)
	stackCClone2 := stackC.Clone(DEEPSEARCH_True, 1) // should equal stackCClone since deepsearchfalse <=> deepsearchtrue | depth: 1
	stackCClone.Cards[0].Val.(*Stack).Cards[0].Key = "New"
	stackCClone.Cards[1].Val.(*Stack).Cards[1].Val = "New"

	conditions := []bool{

		// shallow cloning
		stackA.Equals(MakeStack([]string {"Original", "Original"}, []string {"Original", "Original"})), // 1
		stackAClone.Equals(MakeStack([]string {"New", "New"}, []string {"New", "Original"})), // 2
		
		// deep cloning
		stackB.Equals(MakeStackMatrix([]string {"Original", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "Original"}, []int{2, 2})), // 3
		stackBClone.Equals(MakeStackMatrix([]string {"New", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "New"}, []int{2, 2})), // 4

		// shallow clone stackmatrix
		stackC.Equals(MakeStackMatrix([]string {"New", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "New"}, []int{2, 2})), // 5
		stackCClone.Equals(MakeStackMatrix([]string {"New", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "New"}, []int{2, 2})), // 6
		stackCClone2.Equals(stackCClone), // 7

	}

	test_End(funcName, conditions)
	
}

func case_stack_Unique(funcName string) {

	test_Start(funcName, showTestText)

	// test for type conditions
	myStackKeys := MakeStack(nil, []string {"Person", "Place", "Person", "Thing", "Person"}).Unique(TYPE_Key)
	myStackVals := MakeStack([]string {"Person", "Place", "Person", "Thing", "Person"}).Unique(TYPE_Val)
	myStackBoth := MakeStack([]string {"Person", "Place", "Person", "Thing", "Person"}, []string {"Person", "Place", "Person", "Thing", "Person"}).Unique(TYPE_Val)

	conditions := []bool{
		myStackKeys.Equals(MakeStack(nil, []string {"Person", "Place", "Thing"})), // 1
		myStackVals.Equals(MakeStack([]string {"Person", "Place", "Thing"})), // 2
		myStackBoth.Equals(MakeStack([]string {"Person", "Place", "Thing"}, []string {"Person", "Place", "Thing"})), // 3
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

	// extra tests to ensure that if it's not a pointer test, then you don't need to use an interface argument input to test
	nonInfVar1 := "Hey"
	nonInfVar2 := "Hey"
	cardUno := MakeCard(nonInfVar1)
	cardDos := MakeCard(nonInfVar2)

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
		!cardA.Equals(cardB, POINTER_True), // 9
		!cardB.Equals(cardC, POINTER_False), // 10
		!cardB.Equals(cardC, POINTER_True), // 11
		cardC.Equals(cardD, POINTER_False), // 12
		cardC.Equals(cardD, POINTER_True), // 13

		// test whether pointerTypes work for vals
		cardE.Equals(cardF, nil, POINTER_False), // 14
		!cardE.Equals(cardF, nil, POINTER_True), // 15
		!cardF.Equals(cardG, nil, POINTER_False), // 16
		!cardF.Equals(cardG, nil, POINTER_True), // 17
		cardG.Equals(cardH, nil, POINTER_False), // 18
		cardG.Equals(cardH, nil, POINTER_True), // 19
		
		// test whether we can compare not-by-pointer given we set the card keys
		// with a non-interface (v := "Hey" RATHER THAN var v any; v = "Hey")
		cardUno.Equals(cardDos), // 20
		
	}

	test_End(funcName, conditions)
	
}

func case_stack_Equals(funcName string) {

	// it is crucial that we do not rely on MakeStackMatrix, since MakeStackMatrix uses this function for testing; we cannot have a reciprocal dependency
	// this is why we do weird stuff here with MakeStack(MakeStack) rather than doing MakeStackMatrix

	test_Start(funcName, showTestText)

	// since we've already tested the compare/pointer key/val parameters in card.Equals(),
	// and since they're just passed forward in our stack.Equals() function, we don't need to test these again here

	// shallow
	shallow1 := MakeStack([]string {"Hello", "Hey"})
	shallow2 := MakeStack([]string {"Hello", "Hey"})
	shallow3 := MakeStack([]string {"Hi", "Hey"})
	
	// deep
	deep1 := MakeStack([]*Stack {MakeStack([]string {"Hello", "Hey"}), MakeStack([]string {"Howdy", "Hi"})})
	deep2 := MakeStack([]*Stack {MakeStack([]string {"Hello", "Hey"}), MakeStack([]string {"Howdy", "Hi"})})
	deep3 := MakeStack([]*Stack {MakeStack([]string {"Hello", "Hey"}), MakeStack([]string {"Howdy", "Heyo"})})
	deep4 := MakeStack([]*Stack {MakeStack(), MakeStack()})

	// stack pointers
	sub1 := MakeStack([]string {"Hello", "Hey"})
	sub2 := MakeStack([]string {"Howdy", "Hi"})
	sub3 := MakeStack([]string {"Howdy", "Heyo"})
	sub4 := MakeStack([]string {"Howdy", "Hi"})
	ptrs1 := MakeStack([]*Stack {sub1, sub2})
	ptrs2 := MakeStack([]*Stack {sub1, sub2})
	ptrs3 := MakeStack([]*Stack {sub1, sub3})
	ptrs4 := MakeStack([]*Stack {sub1, sub4})

	// keys to stacks
	kts1 := MakeStack([]string {"Stack A", "Stack B"}, []*Stack {sub1, sub2})
	kts2 := MakeStack([]string {"Stack A", "Stack B"}, []*Stack {sub1, sub2})
	kts3 := MakeStack([]string {"Stack C", "Stack D"}, []*Stack {sub1, sub2})
	kts4 := MakeStack([]string {"Stack A", "Stack B"}, []*Stack {sub1, sub3})

	conditions := []bool{
		
		// test for shallow equality
		shallow1.Equals(shallow2, DEEPSEARCH_False), // 1
		shallow2.Equals(shallow1, DEEPSEARCH_False), // 2
		!shallow1.Equals(shallow3, DEEPSEARCH_False), // 3

		// test for deep equality
		deep1.Equals(deep2, nil, nil, DEEPSEARCH_True), // 4
		deep2.Equals(deep1, nil, nil, DEEPSEARCH_True), // 5
		!deep1.Equals(deep3, nil, nil, DEEPSEARCH_True), // 6
		deep1.Equals(deep2, nil, nil, DEEPSEARCH_True, 1), // 7
		deep1.Equals(deep3, nil, nil, DEEPSEARCH_True, 1), // 8
		deep1.Equals(deep2, nil, nil, DEEPSEARCH_True, 0), // 9

		// test for smaller not auto being equal to larger if stack is missing
		!deep4.Equals(deep1, nil, nil, DEEPSEARCH_True), // 10

		// test for same shape different val comparison
		deep1.Equals(deep3, COMPARE_True, COMPARE_False), // 11

		// test for stack pointers
		ptrs1.Equals(ptrs2, nil, nil, nil, nil, nil, nil, POINTER_True), // 12
		!ptrs1.Equals(ptrs3, nil, nil, nil, nil, nil, nil, POINTER_True), // 13
		!ptrs1.Equals(ptrs4, nil, nil, nil, nil, nil, nil, POINTER_True), // 14

		// test for empty equality
		MakeStack().Equals(MakeStack()), // 14

		// test for keys pointing to stacks
		kts1.Equals(kts2), // 15
		!kts1.Equals(kts3), // 16
		!kts1.Equals(kts4), // 17

		// test that deepsearchfalse <=> deepsearchtrue | depth: 1
		shallow1.Equals(shallow2, DEEPSEARCH_True, 1), // 18

	}

	test_End(funcName, conditions)
	
}

func case_stack_Shuffle(funcName string) {

	test_Start(funcName, showTestText)

	// if there's no issue, there's a 1/10! chance of a false positive
	// test for probable shuffle
	stackA := MakeStack([]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Shuffle()

	// if there's an issue, you will have to output a few times to catch it
	// test for definite shuffle
	stackB := MakeStack([]int {1, 2}).Shuffle(true)

	conditions := []bool{
		
		// test for probable shuffle
		!stackA.Equals(MakeStack([]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10})), // 1

		// test for definite shuffle
		stackB.Equals(MakeStack([]int {2, 1})), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_Inverse(funcName string) {

	test_Start(funcName, showTestText)

	conditions := []bool{
		
		false, // temp

	}

	test_End(funcName, conditions)
	
}

func case_card_Print(funcName string) {

	test_Start(funcName, showTestText)
	
	conditions := []bool{
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_stack_Print(funcName string) {

	test_Start(funcName, showTestText)
	
	conditions := []bool{
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_stack_Lambda(funcName string) {

	test_Start(funcName, showTestText)

	stackToInverse := MakeStack([]int {1, 2, 3, 4})
	stackToCountKeysOver30 := MakeStack(nil, []int {5, 10, 20, 25, 50})
	stackToAdd := MakeStack([]int {1, 2, 3, 4})

	// inverseper
	stackToInverse.Lambda(func(card *Card, stack *Stack, _ ...any) {
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
		stackToInverse.Equals(MakeStack([]int {4, 3, 2, 1})),
		keyCount == 4,
		stackToAdd.Equals(MakeStack(1, 3, 6, 10)),
	}

	test_End(funcName, conditions)
	
}

/** Executes all case tests */
func Run(_showTestText bool) {

	showTestText = _showTestText
	gogenerics.RemoveUnusedError(case_MakeCard, case_MakeStack, case_MakeStackMatrix, case_stack_StripStackMatrix, case_stack_ToArray, case_stack_ToMap, case_stack_ToMatrix, case_stack_IsRegular, case_stack_Duplicate, case_stack_Empty, case_card_Clone, case_stack_Clone, case_stack_Unique, case_card_Equals, case_stack_Equals, case_stack_Shuffle, case_stack_Inverse, case_card_Print, case_stack_Print, case_stack_Lambda)

	fmt.Println("- BEGINNING TESTS (fix failures/errors in descending order)")

	// NON-GENERALIZED FUNCTIONS (NOT DEPENDENT ON GENERALIZED FUNCTIONS)
	case_MakeCard("MakeCard") // GOOD
	case_card_Equals("card.Equals") // GOOD
	case_MakeStack("MakeStack") // GOOD
	case_stack_Equals("stack.Equals") // GOOD
	case_MakeStackMatrix("MakeStackMatrix") // GOOD
	case_stack_ToArray("stack.ToArray") // GOOD
	case_stack_ToMap("stack.ToMap") // GOOD
	case_stack_ToMatrix("stack.ToMatrix") // GOOD
	case_stack_IsRegular("stack.IsRegular") // GOOD
	case_stack_Duplicate("stack.Duplicate") // GOOD
	case_stack_Empty("stack.Empty") // GOOD
	case_card_Clone("card.Clone") // GOOD
	case_stack_Clone("stack.Clone") // GOOD
	case_stack_Shuffle("stack.Shuffle") // GOOD
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
	
	// NON-GENERALIZED FUNCTIONS (DEPENDENT ON GENERALIZED FUNCTIONS)
	// case_stack_StripStackMatrix("stack.StripStackMatrix") // BAD - update to just use the get() function
	// case_stack_Inverse("stack.Inverse") // BAD
	// case_stack_Unique("stack.Unique") // BAD

}
