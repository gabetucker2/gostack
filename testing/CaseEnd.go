package testing

import (
	"fmt"

	"github.com/gabetucker2/gogenerics"
	. "github.com/gabetucker2/gostack" //lint:ignore ST1001 Ignore warning
)

// variables
var showTestText bool

// case functions
func case_MakeCard(funcName string) {

	test_Start(funcName, showTestText)

	// initialize
	card1 := MakeCard()
	card2 := MakeCard("Card 2")
	card3 := MakeCard(card2, "Card 3")
	card4 := MakeCard(8, card1, 2)
	myStr := gogenerics.MakeInterface("Hello")
	card5 := MakeCard(&myStr)
	card6 := MakeCard(&myStr)
	card7 := MakeCard(myStr)
	gogenerics.SetPointer(card6.Val, "Hi")

	// test whether updating fields does so by object
	card2.Val = "Card 4"
	card3.Val = 7

	conditions := []bool {
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
		gogenerics.GetPointer(card5.Val) == "Hi",
		gogenerics.GetPointer(card6.Val) == "Hi",
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
	stack9 := MakeStack()

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

	conditions := []bool {
		test_IdxsAreGood(stack1), // 1
		test_IdxsAreGood(stack2), // 2
		test_IdxsAreGood(stack3), // 3
		test_IdxsAreGood(stack4), // 4
		test_IdxsAreGood(stack5), // 5
		test_IdxsAreGood(stack6), // 6
		test_IdxsAreGood(stack7), // 7

		test_StackProperties(stack1, []int{3}), // 8
		test_StackProperties(stack2, []int{3}), // 9
		test_StackProperties(stack3, []int{3}), // 10
		test_StackProperties(stack4, []int{3}), // 11
		test_StackProperties(stack5, []int{9}), // 12
		test_StackProperties(stack6, []int{10}), // 13
		test_StackProperties(stack7, []int{3}), // 14

		test_StackEqualArrayOrMap(stack1, nil, nil, map1), // 15
		test_StackEqualArrayOrMap(stack2, arrVals, nil, nil), // 16
		test_StackEqualArrayOrMap(stack3, arrVals, arrKeys, nil), // 17
		test_StackEqualArrayOrMap(stack4, nil, arrKeys, nil), // 18
		test_StackEqualArrayOrMap(stack5, arrValsTimesThree, nil, nil), // 19
		test_StackEqualArrayOrMap(stack6, nil, nil, nil), // 20
		test_StackEqualArrayOrMap(stack7, arrVals, nil, nil), // 21
		
		test_StackProperties(stack9, []int{0}), // 22

		gogenerics.GetPointer(stack12.Cards[0].Val) == "Henry", // 23
		gogenerics.GetPointer(stack13.Cards[0].Val) == "Henry", // 24

		stack14.Cards[0].Key == "Hello1", // 25
		stack14.Cards[0].Val == "Hello2", // 26
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

	heightMap := map[string]map[string]int{
        "First": {
            "Alex": 111,
            "Bre": 222,
        },
    }

	arrHeightKeys := []any {[]any {"Alex", "Bre"}, []string {"Charles", "David"}, []any {"Elliot", "Ferguson"}}
	arrHeightVals := []any {[]int {111, 222}, []any {333, 444}, []any {555, 666}}
	
	irregularHeight := []any {10, []any {20, 30}, []any {[]int {40, 50}, []any {60, 70}}}

	// to stacks (in order of conditions listed in documentation)
	
	correctStack := MakeStack([]*Stack {MakeStack([]string {"Alex", "Bre"}, []int {111, 222}), MakeStack([]string {"Charles", "David"}, []int {333, 444}), MakeStack([]string {"Elliot", "Ferguson"}, []int {555, 666})})

	stack1 := MakeStackMatrix(heightMap)
	stack2 := MakeStackMatrix(arrHeightVals)
	stack3 := MakeStackMatrix(arrHeightKeys, arrHeightVals)
	stack4 := MakeStackMatrix(nil, arrHeightKeys)

	// shallow stacks
	stack5 := MakeStackMatrix()
	stack6 := MakeStackMatrix(shallowMap, nil, []int {1, 2})
	stack7 := MakeStackMatrix(arrShallowVals, nil, matrixShape)
	stack8 := MakeStackMatrix(arrShallowKeys, arrShallowVals, matrixShape)
	stack9 := MakeStackMatrix(nil, arrShallowKeys, matrixShape)
	stack10 := MakeStackMatrix(nil, nil, matrixShape)

	// irregular depth
	stack11 := MakeStackMatrix(irregularHeight)

	conditions := []bool {

		// depth tests
		stack1.Equals(MakeStack([]string {"First"}, []*Stack {MakeStack([]string {"Alex", "Bre"}, []int {111, 222})})) || stack1.Equals(MakeStack([]string {"First"}, []*Stack {MakeStack([]string {"Bre", "Alex"}, []int {222, 111})})), // 1
		stack2.Equals(correctStack, nil, nil, COMPARE_False, COMPARE_True), // 2
		stack3.Equals(correctStack, nil, nil, COMPARE_True, COMPARE_True), // 3
		stack4.Equals(correctStack, nil, nil, COMPARE_True, COMPARE_False), // 4

		// shallow tests
		stack5.Equals(MakeStack()), // 5
		stack6.Equals(MakeStack([]*Stack {MakeStack([]string {"Alex", "Bre"}, []int {111, 222})})) || stack6.Equals(MakeStack([]*Stack {MakeStack([]string {"Bre", "Alex"}, []int {222, 111})})), // 6
		stack7.Equals(correctStack, nil, nil, COMPARE_False, COMPARE_True), // 7
		stack8.Equals(correctStack, nil, nil, COMPARE_True, COMPARE_True), // 8
		stack9.Equals(correctStack, nil, nil, COMPARE_True, COMPARE_False), // 9
		stack10.Equals(MakeStack([]*Stack {MakeStack(nil, nil, 2), MakeStack(nil, nil, 2), MakeStack(nil, nil, 2)})), // 10

		// irregular depth
		stack11.Equals(MakeStack([]any {10, MakeStack([]int {20, 30}), MakeStack([]*Stack {MakeStack([]int {40, 50}), MakeStack([]any {60, 70})} ) } )), // 11

	}

	test_End(funcName, conditions)

}

func case_stack_StripStackMatrix(funcName string) {

	test_Start(funcName, showTestText)

	// main
	cardA := MakeCard("HeyA")
	cardB := MakeCard("HeyB")
	cardC := MakeCard("HeyC")
	cardD := MakeCard("HeyD")
	stack1 := MakeStack([]*Card {cardA, cardB})
	stack2 := MakeStack([]*Card {cardC, cardD})
	matrix := MakeStack([]*Stack {stack1, stack2})

	stackA := matrix.Clone().StripStackMatrix()
	stackB := matrix.Clone().StripStackMatrix(0)
	stackC := matrix.Clone().StripStackMatrix(1)
	stackD := matrix.Clone().StripStackMatrix([]int {0, 1})
	stackE := matrix.Clone().StripStackMatrix(MakeStack([]int {0, 1}))

	conditions := []bool {

		// main
		stackA.Equals(MakeStack([]*Card {cardA, cardB, cardC, cardD})), // 1
		stackB.Equals(stack1), // 2
		stackC.Equals(stack2), // 3
		stackD.Equals(MakeStack([]*Card {cardA, cardB, cardC, cardD})), // 4
		stackE.Equals(MakeStack([]*Card {cardA, cardB, cardC, cardD})), // 5

	}

	test_End(funcName, conditions)
	
}

func case_stack_ToArray(funcName string) {

	test_Start(funcName, showTestText)

	testCardA := MakeCard("Key1", "Card A")
	testCardB := MakeCard("Key2", "Card B")
	testCardC := MakeCard("Key3", "Card C")

	sampleStack := func() *Stack {
		return MakeStack([]*Card {testCardA.Clone(), testCardB.Clone(), testCardC.Clone()})
	}

	arrayVals1 := sampleStack().ToArray()
	arrayVals2 := sampleStack().ToArray(RETURN_Vals)
	arrayKeys := sampleStack().ToArray(RETURN_Keys)
	arrayIdxs := sampleStack().ToArray(RETURN_Idxs)
	arrayCards := MakeStack([]*Card {testCardA, testCardB, testCardC}).ToArray(RETURN_Cards)
	
	conditions := []bool {
		len(arrayVals1) == 3, // 1
		len(arrayVals2) == 3, // 2
		len(arrayKeys) == 3, // 3
		len(arrayIdxs) == 3, // 4
		len(arrayCards) == 3, // 5

		arrayVals1[0] == testCardA.Val, // 6
		arrayVals1[1] == testCardB.Val, // 7
		arrayVals1[2] == testCardC.Val, // 8

		arrayVals2[0] == testCardA.Val, // 9
		arrayVals2[1] == testCardB.Val, // 10
		arrayVals2[2] == testCardC.Val, // 11

		arrayKeys[0] == testCardA.Key, // 12
		arrayKeys[1] == testCardB.Key, // 13
		arrayKeys[2] == testCardC.Key, // 14

		arrayIdxs[0] == testCardA.Idx, // 15
		arrayIdxs[1] == testCardB.Idx, // 16
		arrayIdxs[2] == testCardC.Idx, // 17

		arrayCards[0].(*Card) == testCardA, // 18
		arrayCards[1].(*Card) == testCardB, // 19
		arrayCards[2].(*Card) == testCardC, // 20
	}

	test_End(funcName, conditions)
	
}

func case_stack_ToMap(funcName string) {

	test_Start(funcName, showTestText)

	testCardA := MakeCard("Key1", "Card A")
	testCardB := MakeCard("Key2", "Card B")
	testCardC := MakeCard("Key3", "Card C")

	m := MakeStack([]*Card {testCardA.Clone(), testCardB.Clone(), testCardC.Clone()}).ToMap()

	conditions := []bool {
		len(m) == 3, // 1
		m["Key1"] == "Card A", // 2
		m["Key2"] == "Card B", // 3
		m["Key3"] == "Card C", // 4
		m["Key4"] == nil, // 5
	}

	test_End(funcName, conditions)
	
}

func case_stack_ToMatrix(funcName string) {

	test_Start(funcName, showTestText)
	
	// test for shallow on shallow
	matA := MakeStack([]string {"Hi", "Hello", "Hey"}).ToMatrix()
	matACorrect := []any {"Hi", "Hello", "Hey"}

	// test for depth on depth
	matB := MakeStackMatrix([]string {"Hi", "Hey", "Hoy", "Hiya"}, nil, []int {2, 2}).ToMatrix()
	matBCorrect := []any {[]any {"Hi", "Hey"}, []any {"Hoy", "Hiya"}}

	// test for shallow on depth
	matC := MakeStackMatrix([]string {"Hi", "Hey", "Hoy", "Hiya"}, nil, []int {2, 2}).ToMatrix(nil, 1)
	matCCorrect := []any {[]any {}, []any {}}

	// test for irregular depth
	matD := MakeStack([]any {"Hi", MakeStack([]string {"Hoy", "Hiya"})}).ToMatrix()
	matDCorrect := []any {"Hi", []any {"Hoy", "Hiya"}}

	// test for heightsearchfalse <=> heightsearchtrue | depth: 1
	matE := MakeStackMatrix([]string {"Hi", "Hey", "Hoy", "Hiya"}, nil, []int {2, 2}).ToMatrix(nil, 1)
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


	conditions := []bool {

		// test for shallow on shallow
		MakeStackMatrix(matA).Equals(MakeStackMatrix(matACorrect)), // 1

		// test for depth on depth
		MakeStackMatrix(matB).Equals(MakeStackMatrix(matBCorrect)), // 2

		// test for shallow on depth
		MakeStackMatrix(matC).Equals(MakeStackMatrix(matCCorrect)), // 3

		// test for irregular depth
		MakeStackMatrix(matD).Equals(MakeStackMatrix(matDCorrect)), // 4

		// test for heightsearchfalse <=> heightsearchtrue | depth: 1
		MakeStackMatrix(matE).Equals(MakeStackMatrix(matECorrect)), // 5
		
		// test for different return types
		MakeStackMatrix(matKeys).Equals(MakeStackMatrix(matKeysCorrect)), // 6
		MakeStackMatrix(matIdxs).Equals(MakeStackMatrix(matIdxsCorrect)), // 7
		MakeStackMatrix(matCards).Equals(MakeStackMatrix(matCardsCorrect)), // 8
		
	}

	test_End(funcName, conditions)
	
}

func case_stack_Shape(funcName string) {

	test_Start(funcName, showTestText)

	stack1 := MakeStackMatrix([]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, nil, []int {2, 2, 5})
	stack2 := MakeStack()
	stack3 := MakeStack([]string {"Hi", "Hey"})

	conditions := []bool {
		gogenerics.SlicesEqual(stack1.Shape(), []int {2, 2, 5}), // 1
		gogenerics.SlicesEqual(stack2.Shape(), []int {0}), // 2
		gogenerics.SlicesEqual(stack3.Shape(), []int {2}), // 3
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

	conditions := []bool {
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

	conditions := []bool {
		stack1.Equals(MakeStack()), // 1
		stack2.Equals(MakeStack([]string {"Hey", "Hi"})), // 2
		stack3.Equals(MakeStack([]string {"Hey", "Hi", "Hey", "Hi"})), // 3
	}

	test_End(funcName, conditions)
	
}

func case_stack_Empty(funcName string) {

	test_Start(funcName, showTestText)

	stack1 := MakeStack([]string {"Hey", "Hi", "gdjifjgdfoigj"}).Empty()
	stack2 := MakeStackMatrix([]int {1, 2, 3, 4, 5, 6}, []string {"Hi", "Hey", "Hoy", "Ciao", "Heyy", "Hiya"}, []int {3, 2}).Empty()

	conditions := []bool {
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

	conditions := []bool {
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

	// depth cloning
	stackB := MakeStackMatrix([]string {"Original", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "Original"}, []int{2, 2})
	stackBClone := stackB.Clone()
	stackBClone.Cards[0].Val.(*Stack).Cards[0].Key = "New"
	stackBClone.Cards[1].Val.(*Stack).Cards[1].Val = "New"

	// shallow clone stackmatrix
	stackC := MakeStackMatrix([]string {"Original", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "Original"}, []int{2, 2})
	stackCClone := stackC.Clone(DEEPSEARCH_False)
	stackCClone.Cards[0].Val.(*Stack).Cards[0].Key = "New"
	stackCClone.Cards[1].Val.(*Stack).Cards[1].Val = "New"

	// test no cloning
	stackD := MakeStack([]string {"StackAKey", "StackBKey"}, []*Stack {MakeStack([]string {"Hey", "Hey"}, []string {"Original", "Original"}), MakeStack([]string {"Hey", "Hey"}, []string {"Original", "Original"})})
	stackDClone1 := stackD.Clone(nil, nil, CLONE_False, CLONE_True, CLONE_True, CLONE_True)
	stackDClone2 := stackD.Clone(nil, nil, CLONE_True, CLONE_False, CLONE_True, CLONE_True)
	stackDClone3 := stackD.Clone(nil, nil, CLONE_True, CLONE_True, CLONE_False, CLONE_True)
	stackDClone4 := stackD.Clone(nil, nil, CLONE_True, CLONE_True, CLONE_True, CLONE_False)

	conditions := []bool {

		// shallow cloning
		stackA.Equals(MakeStack([]string {"Original", "Original"}, []string {"Original", "Original"})), // 1
		stackAClone.Equals(MakeStack([]string {"New", "New"}, []string {"New", "Original"})), // 2
		
		// depth cloning
		stackB.Equals(MakeStackMatrix([]string {"Original", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "Original"}, []int{2, 2})), // 3
		stackBClone.Equals(MakeStackMatrix([]string {"New", "Original", "Original", "Original"}, []string {"Original", "Original", "Original", "New"}, []int{2, 2})), // 4

		// shallow clone stackmatrix
		stackC.Equals(stackCClone), // 5

		// test no cloning
		stackDClone1.Equals(MakeStack([]string {"StackAKey", "StackBKey"}, []*Stack {MakeStack([]string {"Original", "Original"}), MakeStack([]string {"Original", "Original"})})), // 6
		stackDClone2.Equals(MakeStack([]string {"StackAKey", "StackBKey"}, []*Stack {MakeStack(nil, []string {"Hey", "Hey"}), MakeStack(nil, []string {"Hey", "Hey"})})), // 7
		stackDClone3.Equals(MakeStack([]*Stack {MakeStack([]string {"Hey", "Hey"}, []string {"Original", "Original"}), MakeStack([]string {"Hey", "Hey"}, []string {"Original", "Original"})})), // 8
		stackDClone4.Equals(MakeStack([]string {"StackAKey", "StackBKey"}, []any {nil, nil})), // 9

	}

	test_End(funcName, conditions)
	
}

func case_stack_Unique(funcName string) {

	test_Start(funcName, showTestText)

	// test for type conditions
	myStackKeys := MakeStack(nil, []string {"Person", "Place", "Person", "Thing", "Person"}).Unique(TYPE_Key)
	myStackVals := MakeStack([]string {"Person", "Place", "Person", "Thing", "Person"}).Unique(TYPE_Val)
	myStackBoth := MakeStack([]string {"Person", "Place", "Person", "Thing", "Person"}, []string {"Person", "Place", "Person", "Thing", "Person"}).Unique(TYPE_Val)
	
	conditions := []bool {
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

	conditions := []bool {

		// compare by object
		card1.Equals(card2), // 1

		// test whether idx parameter works
		card1.Equals(card2, COMPARE_False, COMPARE_False, COMPARE_False), // 2
		!card1.Equals(card2, COMPARE_True, COMPARE_False, COMPARE_False), // 3

		// test whether keys parameter works
		card1.Equals(card3, COMPARE_False, COMPARE_False, COMPARE_False), // 4
		!card5.Equals(card6, COMPARE_False, COMPARE_True, COMPARE_False), // 5

		// test whether val parameter works
		card1.Equals(card5, COMPARE_False, COMPARE_False, COMPARE_False), // 6
		!card3.Equals(card4, COMPARE_False, COMPARE_False, COMPARE_True), // 7

		// test whether dereferenceTypes work for keys
		cardA.Equals(cardB, nil, nil, nil, nil, DEREFERENCE_None), // 8
		!cardA.Equals(cardB, nil, nil, nil, nil, DEREFERENCE_Both), // 9
		!cardB.Equals(cardC, nil, nil, nil, nil, DEREFERENCE_None), // 10
		!cardB.Equals(cardC, nil, nil, nil, nil, DEREFERENCE_Both), // 11
		cardC.Equals(cardD, nil, nil, nil, nil, DEREFERENCE_None), // 12
		cardC.Equals(cardD, nil, nil, nil, nil, DEREFERENCE_Both), // 13

		// test whether dereferenceTypes work for vals
		cardE.Equals(cardF, nil, nil, nil, nil, nil, DEREFERENCE_None), // 14
		!cardE.Equals(cardF, nil, nil, nil, nil, nil, DEREFERENCE_Both), // 15
		!cardF.Equals(cardG, nil, nil, nil, nil, nil, DEREFERENCE_None), // 16
		!cardF.Equals(cardG, nil, nil, nil, nil, nil, DEREFERENCE_Both), // 17
		cardG.Equals(cardH, nil, nil, nil, nil, nil, DEREFERENCE_None), // 18
		cardG.Equals(cardH, nil, nil, nil, nil, nil, DEREFERENCE_Both), // 19
		
		// test whether we can compare not-by-pointer given we set the card keys
		// with a non-interface (v := "Hey" RATHER THAN var v any; v = "Hey")
		cardUno.Equals(cardDos), // 20

		// test compare object vs non
		!card1.Equals(card2, nil, nil, nil, COMPARE_True), // 21
		card1.Equals(card2, nil, nil, nil, COMPARE_False), // 22

		// test additional dereference options
		cardE.Equals(cardG, nil, nil, nil, nil, nil, DEREFERENCE_Found), // 23
		!cardE.Equals(cardG, nil, nil, nil, nil, nil, DEREFERENCE_Both), // 24
		!cardE.Equals(cardG, nil, nil, nil, nil, nil, DEREFERENCE_This), // 25
		!cardE.Equals(cardG, nil, nil, nil, nil, nil, DEREFERENCE_None), // 26
		!cardG.Equals(cardE, nil, nil, nil, nil, nil, DEREFERENCE_Found), // 27
		!cardG.Equals(cardE, nil, nil, nil, nil, nil, DEREFERENCE_Both), // 28
		cardG.Equals(cardE, nil, nil, nil, nil, nil, DEREFERENCE_This), // 29
		!cardG.Equals(cardE, nil, nil, nil, nil, nil, DEREFERENCE_None), // 30
		
	}

	test_End(funcName, conditions)
	
}

func case_stack_Equals(funcName string) {

	// it is crucial that we do not rely on MakeStackMatrix, since MakeStackMatrix uses this function for testing; we cannot have a reciprocal dependency
	// this is why we do weird stuff here with MakeStack(MakeStack) rather than doing MakeStackMatrix

	test_Start(funcName, showTestText)

	// since we already tested pointer comparisons in card.Equals, and since our DEREFERENCE parameters merely implement those in card.Equals, we need not test our DEREFERENCE parameters here

	// test for shallow-on-shallow equality
	sos1 := MakeStack([]string {"Hello", "Hey"})
	sos2 := MakeStack([]string {"Hello", "Hey"})
	sos3 := MakeStack([]string {"Hi", "Hey"})

	// test for depth-on-depth equality
	dod1 := MakeStack([]*Stack {MakeStack([]string {"Hi1", "Hi2"}), MakeStack([]string {"Hi3", "Hi4"})})
	dod2 := MakeStack([]*Stack {MakeStack([]string {"Hi1", "Hi2"}), MakeStack([]string {"Hi3", "Hi4"})})
	dod3 := MakeStack([]*Stack {MakeStack([]string {"Hi1", "Hello2"}), MakeStack([]string {"Hi3", "Hi4"})})

	// test for shallow-on-depth equality
	sod1 := MakeStack([]*Stack {MakeStack([]string {"Hi1", "Hi2"}), MakeStack([]string {"Hi3", "Hi4"})})
	sod2 := MakeStack([]*Stack {MakeStack([]string {"Hi1", "Hello2"}), MakeStack([]string {"Hi3", "Hi4"})})
	sod3 := MakeStack([]*Stack {MakeStack([]string {"Hi1"}), MakeStack([]string {"Hi3", "Hi4"})})
	sod4 := MakeStack([]*Stack {MakeStack([]string {"Hi1"})})

	// test for compare card filters
	ccf1 := MakeStack([]string {"Hi"}, []int {4})
	ccf2 := MakeStack([]string {"Hi"}, []int {4})
	ccf3 := MakeStack([]string {"Hey"}, []int {4})
	ccf4 := MakeStack([]string {"Hi"}, []int {2})
	ccf5 := MakeStack([]string {"Hey"}, []int {2})

	// test for compare substack filters
	csfStack1 := MakeStack([]string {"Hi1", "Hi2"})
	csfStack2 := MakeStack([]string {"Hi3", "Hi4"})
	csf1 := MakeStack([]string {"StackA", "StackB"}, []*Stack {csfStack1, csfStack2})
	csf2 := MakeStack([]string {"StackA", "StackB"}, []*Stack {csfStack1, csfStack2})
	csf3 := MakeStack([]string {"Stack420", "StackB"}, []*Stack {csfStack1, csfStack2})
	csf4 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]string {"Hi1", "Hello2"}), csfStack2})
	csf5 := MakeStack([]string {"Stack420", "StackB"}, []*Stack {MakeStack([]string {"Hi1", "Hello2"}), csfStack2})

	// test for depth search filters
	dsf1 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]string {"Hi1", "Hi2"}), MakeStack([]string {"Hi3", "Hi4"})})
	dsf2 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]string {"Hi1", "Hi2"}), MakeStack([]string {"Hi3", "Hi4"})})
	dsf3 := MakeStack([]string {"Stack420", "StackB"}, []*Stack {MakeStack([]string {"Hi1", "Hi2"}), MakeStack([]string {"Hi3", "Hi4"})})
	dsf4 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]string {"Hi1", "Hello2"}), MakeStack([]string {"Hi3", "Hi4"})})
	dsf5 := MakeStack([]string {"Stack420", "StackB"}, []*Stack {MakeStack([]string {"Hi1", "Hello2"}), MakeStack([]string {"Hi3", "Hi4"})})
	dsf6 := MakeStack([]string {"StackA"}, []*Stack {MakeStack([]string {"Hi1", "Hi2"})})

	conditions := []bool {
		
		// test for shallow-on-shallow equality
		sos1.Equals(sos2, DEEPSEARCH_False), // 1
		!sos1.Equals(sos3, DEEPSEARCH_False), // 2

		// test for depth-on-depth equality
		dod1.Equals(dod2, DEEPSEARCH_True), // 3
		!dod1.Equals(dod3, DEEPSEARCH_True), // 4

		// test for shallow-on-depth equality
		sod1.Equals(sod2, DEEPSEARCH_False), // 5
		sod1.Equals(sod3, DEEPSEARCH_False), // 6
		!sod1.Equals(sod4, DEEPSEARCH_False), // 7

		// test for compare card filters
		ccf1.Equals(ccf2, nil, nil, COMPARE_True, COMPARE_True), // 8
		ccf1.Equals(ccf2, nil, nil, COMPARE_False, COMPARE_True), // 9
		ccf1.Equals(ccf2, nil, nil, COMPARE_True, COMPARE_False), // 10
		ccf1.Equals(ccf2, nil, nil, COMPARE_False, COMPARE_False), // 11

		!ccf1.Equals(ccf3, nil, nil, COMPARE_True, COMPARE_True), // 12
		ccf1.Equals(ccf3, nil, nil, COMPARE_False, COMPARE_True), // 13
		!ccf1.Equals(ccf3, nil, nil, COMPARE_True, COMPARE_False), // 14
		ccf1.Equals(ccf3, nil, nil, COMPARE_False, COMPARE_False), // 15

		!ccf1.Equals(ccf4, nil, nil, COMPARE_True, COMPARE_True), // 16
		!ccf1.Equals(ccf4, nil, nil, COMPARE_False, COMPARE_True), // 17
		ccf1.Equals(ccf4, nil, nil, COMPARE_True, COMPARE_False), // 18
		ccf1.Equals(ccf4, nil, nil, COMPARE_False, COMPARE_False), // 19

		!ccf1.Equals(ccf5, nil, nil, COMPARE_True, COMPARE_True), // 20
		!ccf1.Equals(ccf5, nil, nil, COMPARE_False, COMPARE_True), // 21
		!ccf1.Equals(ccf5, nil, nil, COMPARE_True, COMPARE_False), // 22
		ccf1.Equals(ccf5, nil, nil, COMPARE_False, COMPARE_False), // 23

		// test for compare substack keys
		csf1.Equals(csf2, nil, nil, nil, nil, COMPARE_True), // 24
		csf1.Equals(csf2, nil, nil, nil, nil, COMPARE_False), // 25
		csf1.Equals(csf2, nil, nil, nil, nil, COMPARE_True), // 26
		csf1.Equals(csf2, nil, nil, nil, nil, COMPARE_False), // 27

		!csf1.Equals(csf3, nil, nil, nil, nil, COMPARE_True), // 28
		csf1.Equals(csf3, nil, nil, nil, nil, COMPARE_False), // 29
		!csf1.Equals(csf3, nil, nil, nil, nil, COMPARE_True), // 30
		csf1.Equals(csf3, nil, nil, nil, nil, COMPARE_False), // 31

		true, // 32 // TODO: remove/replace the trues
		true, // 33
		csf1.Equals(csf4, nil, nil, nil, COMPARE_False, COMPARE_True), // 34
		csf1.Equals(csf4, nil, nil, nil, COMPARE_False, COMPARE_False), // 35

		!csf1.Equals(csf5, nil, nil, nil, nil, COMPARE_True), // 36
		!csf1.Equals(csf5, nil, nil, nil, nil, COMPARE_False), // 37
		!csf1.Equals(csf5, nil, nil, nil, nil, COMPARE_True), // 38
		csf1.Equals(csf5, nil, nil, nil, COMPARE_False, COMPARE_False), // 39

		// test for depth search filters
		dsf1.Equals(dsf2, nil, 0), // 40

		dsf1.Equals(dsf2, nil, -1), // 41
		dsf2.Equals(dsf1, nil, -1), // 42
		dsf1.Equals(dsf2, nil, 2), // 43
		dsf1.Equals(dsf2, nil, 3), // 44
		dsf1.Equals(dsf2, nil, []int {1, 2}), // 45

		dsf1.Equals(dsf2, DEEPSEARCH_False), // 46
		dsf1.Equals(dsf2, nil, []int {1}), // 47
		dsf1.Equals(dsf2, nil, 1), // 48
		
		dsf1.Equals(dsf2, nil, []int {2}), // 49
		
		
		dsf1.Equals(dsf3, nil, 0), // 50

		!dsf1.Equals(dsf3, nil, -1), // 51
		!dsf1.Equals(dsf3, nil, 2), // 52
		!dsf1.Equals(dsf3, nil, 3), // 53
		!dsf1.Equals(dsf3, nil, []int {1, 2}), // 54

		!dsf1.Equals(dsf3, DEEPSEARCH_False), // 55
		!dsf1.Equals(dsf3, nil, []int {1}), // 56
		!dsf1.Equals(dsf3, nil, 1), // 57

		dsf1.Equals(dsf3, nil, []int {2}), // 58
		
		
		dsf1.Equals(dsf4, nil, 0), // 59

		!dsf1.Equals(dsf4, nil, -1), // 60
		!dsf1.Equals(dsf4, nil, 2), // 61
 		!dsf1.Equals(dsf4, nil, 3), // 62
		!dsf1.Equals(dsf4, nil, []int {1, 2}), // 63

		dsf1.Equals(dsf4, DEEPSEARCH_False), // 64
		dsf1.Equals(dsf4, nil, []int {1}), // 65
		dsf1.Equals(dsf4, nil, 1), // 66
		
		!dsf1.Equals(dsf4, nil, []int {2}), // 67
		
		
		dsf1.Equals(dsf5, nil, 0), // 68

		!dsf1.Equals(dsf5, nil, -1), // 69
		!dsf1.Equals(dsf5, nil, 2), // 70
		!dsf1.Equals(dsf5, nil, 3), // 71
		!dsf1.Equals(dsf5, nil, []int {1, 2}), // 72

		!dsf1.Equals(dsf5, DEEPSEARCH_False), // 73
		!dsf1.Equals(dsf5, nil, []int {1}), // 74
		!dsf1.Equals(dsf5, nil, 1), // 75
		
		!dsf1.Equals(dsf5, nil, []int {2}), // 76
		
		
		dsf1.Equals(dsf6, nil, 0), // 77

		!dsf1.Equals(dsf6, nil, -1), // 78
		!dsf6.Equals(dsf1, nil, -1), // 79
		!dsf1.Equals(dsf6, nil, 2), // 80
		!dsf1.Equals(dsf6, nil, 3), // 81
		!dsf1.Equals(dsf6, nil, []int {1, 2}), // 82

		!dsf1.Equals(dsf6, DEEPSEARCH_False), // 83
		!dsf1.Equals(dsf6, nil, []int {1}), // 84
		!dsf1.Equals(dsf6, nil, 1), // 85
		
		!dsf1.Equals(dsf6, nil, []int {2}), // 86

		// test Stack heights
		dsf1.Equals(dsf2, nil, MakeStack([]int {1, 2})), // 87
		!dsf1.Equals(dsf6, nil, MakeStack([]int {1, 2})), // 88

		// test stack adr compare
		dsf1.Equals(dsf1, nil, nil, nil, nil, nil, nil, nil, nil, COMPARE_True), // 89
		!dsf1.Equals(dsf1.Clone(), nil, nil, nil, nil, nil, nil, nil, nil, COMPARE_True), // 90

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
	stackB := MakeStack([]int {1, 2}).Shuffle(REPEAT_True)

	conditions := []bool {
		
		// test for probable shuffle
		!stackA.Equals(MakeStack([]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10})), // 1

		// test for definite shuffle
		stackB.Equals(MakeStack([]int {2, 1})), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_Flip(funcName string) {

	test_Start(funcName, showTestText)

	// main cases
	stack1 := MakeStack([]int {1, 2, 3}).Flip()
	stack2 := MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2}).Flip()

	conditions := []bool {
		
		// main cases
		stack1.Equals(MakeStack([]int {3, 2, 1})), // 1
		stack2.Equals(MakeStackMatrix([]int {3, 4, 1, 2}, nil, []int {2, 2})), // 2

	}

	test_End(funcName, conditions)
	
}

func case_card_SwitchKeyVal(funcName string) {

	test_Start(funcName, showTestText)

	// main cases
	cardA := MakeCard("MyKey", "MyVal").SwitchKeyVal()

	conditions := []bool {
		
		// main cases
		cardA.Equals(MakeCard("MyVal", "MyKey")), // 1

	}

	test_End(funcName, conditions)
	
}

func case_stack_SwitchKeysVals(funcName string) {

	test_Start(funcName, showTestText)

	// main cases
	stack1 := MakeStack([]int {1, 3, 2, 4}, []string {"Guy", "Girl", "Guy", "Girl"}).SwitchKeysVals(FIND_All)
	stack2 := MakeStackMatrix([]int {1, 3, 2, 4}, []string {"Guy", "Girl", "Guy", "Girl"}, []int {2, 2}).SwitchKeysVals(FIND_All, nil, nil, DEEPSEARCH_True)

	conditions := []bool {
		
		// main cases
		stack1.Equals(MakeStack([]string {"Guy", "Girl", "Guy", "Girl"}, []int {1, 3, 2, 4})), // 1
		stack2.Equals(MakeStackMatrix([]string {"Guy", "Girl", "Guy", "Girl"}, []int {1, 3, 2, 4}, []int {2, 2})), // 2

	}

	test_End(funcName, conditions)
	
}

func case_card_Print(funcName string) {

	test_Start(funcName, showTestText)
	
	conditions := []bool {
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_stack_Print(funcName string) {

	test_Start(funcName, showTestText)
	
	conditions := []bool {
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_CSVToStackMatrix(funcName string) {

	test_Start(funcName, showTestText)
	
	conditions := []bool {
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_stack_ToCSV(funcName string) {

	test_Start(funcName, showTestText)
	
	conditions := []bool {
		true, // unfortunately, we have to check manually
	}

	test_End(funcName, conditions)
	
}

func case_stack_Lambdas(funcName string) {

	test_Start(funcName, showTestText)

	// test stack updating, multiply each by 5
	stack1 := MakeStack([]int {1, 5, 20}).LambdaThis(func(card *Card, _ *Stack, _ bool, _ *Stack, _ *Stack, _ *Card, _ any, _ []any, _ ...any) {
		card.Val = card.Val.(int) * 5
	})

	// test retStack output, get all in range
	stack2 := MakeStack([]int {1, 5, 20, 41, 92, 4104}).LambdaStack(func(card *Card, _ *Stack, _ bool, _ *Stack, retStack *Stack, _ *Card, _ any, _ []any, _ ...any) {
		if 5 < card.Val.(int) && card.Val.(int) < 4104 {
			retStack.Cards = append(retStack.Cards, card.Clone())
		}
	})

	// test retCard output, get last in range
	card1 := MakeStack([]int {1, 5, 20, 41, 92, 4104}).LambdaCard(func(card *Card, _ *Stack, _ bool, _ *Stack, _ *Stack, retCard *Card, _ any, _ []any, _ ...any) {
		if 5 < card.Val.(int) && card.Val.(int) < 4104 {
			*retCard = *card
		}
	})

	// test retOther output, get max
	maxAdr := MakeStack([]int {50, 2, 45, 140, 42}).LambdaVarAdr(func(card *Card, _ *Stack, _ bool, _ *Stack, _ *Stack, _ *Card, maxAdr any, _ []any, _ ...any) {
		if gogenerics.GetPointer(maxAdr).(int) < card.Val.(int) {
			gogenerics.SetPointer(maxAdr, card.Val)
		}
	}, nil, nil, 0) // initialize max to 0

	// test workingMemAdr, get all which are under or equal to 15 away from stack average
	stack3 := MakeStack([]int {50, 2, 45, 140, 42}).LambdaStack(func(card *Card, stack *Stack, _ bool, _ *Stack, retStack *Stack, _ *Card, _ any, _ []any, workingMem ...any) {
		if workingMem[0] == nil {
			sum := gogenerics.MakeInterface(0)
			for _, c := range stack.Cards {
				sum = sum.(int) + c.Val.(int)
			}
			sum = sum.(int) / stack.Size
			workingMem[0] = &sum
		}
		avg := gogenerics.GetPointer(workingMem[0]).(int)
		if avg - 15 <= card.Val.(int) && card.Val.(int) <= avg + 15 {
			retStack.Cards = append(retStack.Cards, card.Clone())
		}
	})

	// test heightStacks, multiply each by 5, and test coords
	coordSequence := MakeStack([]*Stack {MakeSubstack([]int {0, 0}), MakeSubstack([]int {0, 1}), MakeSubstack([]int {1, 0}), MakeSubstack([]int {1, 1})})
	stack4 := MakeStackMatrix([]int {1, 5, 20, 2}, nil, []int {2, 2}).LambdaThis(func(card *Card, _ *Stack, _ bool, coords *Stack) {
		if coords.Equals(coordSequence.Get(FIND_First).Val.(*Stack)) { // only do the right thing if coords are working
			card.Val = card.Val.(int) * 5
			coordSequence.Remove(FIND_First)
		}
	})

	// test passSubstacks true passCards false, multiply each stack.Key by 5
	stack5 := MakeStack([]int {4, 7}, []*Stack {MakeStack([]int {1, 5}), MakeStack([]int {20, 2})}).LambdaThis(func(card *Card) {
		card.Key = card.Key.(int) * 5
	}, nil, nil, nil, nil, nil, nil, PASS_Substacks)

	// test that all init values work
	this, stack, card, varAdr := MakeStack([]string {"Heyy"}).Lambda(func() {}, MakeStack([]int {666}), MakeCard("Howdy"), 420)

	// test that various heightStack depth options work AND that you can abstract away parameters
	stack6 := MakeStackMatrix([]int {1, 5, 20, 2}, nil, []int {2, 2}).LambdaThis(func(card *Card) {
		if card.Idx == 0 {
			card.Key = "Marker"	
		}
	}, nil, nil, nil, nil, nil, []int {2}, PASS_Both)
	
	conditions := []bool {

		// test stack updating, multiply each by 5
		stack1.Equals(MakeStack([]int {5, 25, 100})), // 1

		// test retStack output, get all in range
		stack2.Equals(MakeStack([]int {20, 41, 92})), // 2

		// test retCard output, get last in range
		card1.Equals(MakeCard(92)), // 3

		// test retOther output, get max
		gogenerics.GetPointer(maxAdr) == 140, // 4

		// test workingMemAdr, get average and return all in average's range
		stack3.Equals(MakeStack([]int {50, 45, 42})), // 5

		// test heightStacks, multiply each by 5
		stack4.Equals(MakeStackMatrix([]int {5, 25, 100, 10}, nil, []int {2, 2})), // 6

		// test passSubstacks true passCards false, multiply each stack.Key by 5
		stack5.Equals(MakeStack([]int {20, 35}, []*Stack {MakeStack([]int {1, 5}), MakeStack([]int {20, 2})})), // 7

		// test that all init values work
		this.Equals(MakeStack([]string {"Heyy"})), // 8
		stack.Equals(MakeStack([]int {666})), // 9
		card.Equals(MakeCard("Howdy")), // 10
		gogenerics.GetPointer(varAdr) == 420, // 11

		// test that various heightStack depth options work
		stack6.Equals(MakeStackMatrix([]any {"Marker", nil, "Marker", nil}, []int {1, 5, 20, 2}, []int {2, 2})), // 12

	}

	test_End(funcName, conditions)
	
}

func case_stack_Filter(funcName string) {

	test_Start(funcName, showTestText)

	// expansively test functionality
	stack := MakeStack([]int {1, 2, 3})
	stack.Filter(FIND_Idx, []int {1, 2})

	conditions := []bool {

		// expansively test functionality
		stack.Equals(MakeStack([]int {2, 3})),

	}

	test_End(funcName, conditions)
	
}

func case_stack_Get(funcName string) {

	test_Start(funcName, showTestText)

	// expansively test functionality
	card1 := MakeStack([]int {1, 2, 3}).Get()
	card2 := MakeStack([]int {1, 2, 3}).Get(FIND_First)
	card3 := MakeStack([]int {1, 2, 3}).Get(FIND_Idx, 1)
	card4 := MakeStack([]int {1, 2, 3}).Get(FIND_Idx, []int {5, 1})
	card5 := MakeStack([]int {1, 2, 3}).Get(FIND_Idx, MakeStack([]int {5, 1}))
	card6 := MakeStack([]int {1, 2, 3}).Get(FIND_Val, []int {5, 2})
	card7 := MakeStack([]int {1, 2, 3}).Get(FIND_Key, nil)
	card8 := MakeStack([]int {1, 2, 3}).Get(FIND_Val, nil)
	cardA := MakeCard(2)
	card9 := MakeStack([]*Card {MakeCard(1, 2), cardA, MakeCard(3)}).Get(FIND_Card, cardA)
	card10 := MakeStack([]*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get().Val.(*Stack).Get()
	card11 := MakeStack([]*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get(nil, nil, DEEPSEARCH_True, nil, PASS_Cards)
	stackA := MakeStack([]int {3, 6})
	card12 := MakeStack([]*Stack {stackA, MakeStack([]int {9, 12})}).Get(FIND_Val, stackA, nil, nil, nil, nil, OVERRIDE_True)
	card13 := MakeStack([]*Stack {stackA, MakeStack([]int {9, 12})}).Get(FIND_Val, MakeStack([]int {3, 6}), nil, nil, nil, nil, OVERRIDE_True)
	var intValA, intValB, intValC, intValD *int
	init1 := 1
	intValA = &init1
	init2 := 2
	intValB = &init2
	init3 := 3
	intValC = &init3
	init4 := 2
	intValD = &init4
	card14 := MakeStack([]any {intValA, intValB, intValC}).Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_None)
	card15 := MakeStack([]any {intValA, intValB, intValC}).Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_Both)
	card16 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get(FIND_All, nil, DEEPSEARCH_True, []int {1})
	card17 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get(FIND_All, nil, DEEPSEARCH_True, []int {2})
	card18 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get(FIND_All, nil, DEEPSEARCH_True, 2)
	card19 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get(FIND_All, nil, DEEPSEARCH_True, []int {1, 2})
	card20 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get(FIND_All, nil, DEEPSEARCH_True, MakeStack([]int {1, 2}))
	card21 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get(FIND_Lambda, func(card *Card) (bool) {
		return card.Val == 6
	}, DEEPSEARCH_True)
	card22 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).Get(FIND_Lambda, func(card *Card, stack *Stack, isSubstack bool) (bool) {
		return !isSubstack && card.Val.(int) < stack.Size*2
	}, DEEPSEARCH_True)
	initVal := gogenerics.MakeInterface(0)
	card23 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {4, 7}), MakeStack([]int {10, 14})}).Get(FIND_Lambda, func(card *Card, stack *Stack, isSubstack bool, _ *Stack, workingMem ...any) (bool) {
		test := gogenerics.GetPointer(workingMem[0]).(int) + 3 == card.Clone().Val
		gogenerics.SetPointer(workingMem[0], card.Val)
		return test
	}, DEEPSEARCH_True, nil, PASS_Cards, nil, nil, []any {&initVal})

	// test for card comparison
	cardB := MakeCard("Hey")
	card24 := MakeStack([]*Card {cardB}).Get(FIND_Card, MakeCard("Hey"))
	card25 := MakeStack([]*Card {cardB}).Get(FIND_Card, cardB)

	// more pointer tests
	card26 := MakeStack([]any {intValA, intValB, intValC}).Get(FIND_Val, intValD, nil, nil, nil, DEREFERENCE_None)
	card27 := MakeStack([]any {intValA, intValB, intValC}).Get(FIND_Val, intValD, nil, nil, nil, DEREFERENCE_Both)
	card28 := MakeStack([]any {intValA, intValB, intValC}).Get(FIND_Val, 2, nil, nil, nil, DEREFERENCE_Found)
	card29 := MakeStack([]any {1, 2, 3}).Get(FIND_Val, intValD, nil, nil, nil, DEREFERENCE_This)
	card30 := MakeStack([]any {intValA, intValB, intValC}).Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_Found)
	card31 := MakeStack([]any {intValA, intValB, intValC}).Get(FIND_Val, intValD, nil, nil, nil, DEREFERENCE_This)

	conditions := []bool {

		// expansively test functionality
		card1.Equals(MakeCard(3, nil, 2), COMPARE_True), // 1
		card2.Equals(MakeCard(1, nil, 0), COMPARE_True), // 2
		card3.Equals(MakeCard(2, nil, 1), COMPARE_True), // 3
		card4.Equals(MakeCard(2, nil, 1), COMPARE_True), // 4
		card5.Equals(MakeCard(2, nil, 1), COMPARE_True), // 5
		card6.Equals(MakeCard(2, nil, 1), COMPARE_True), // 6
		card7.Equals(MakeCard(1, nil, 0), COMPARE_True), // 7
		card8 == nil, // 8
		card9.Equals(cardA), // 9
		card10.Equals(MakeCard(12, nil, 1), COMPARE_True), // 10
		card11.Equals(MakeCard(6, nil, 1), COMPARE_True), // 11
		card12.Equals(MakeCard(stackA, nil, 0), COMPARE_True), // 12
		card13 == nil, // 13
		card14.Equals(MakeCard(intValB, nil, 1), COMPARE_True), // 14
		card15.Equals(MakeCard(intValB, nil, 1), COMPARE_True), // 15
		card16.Key == "StackA", // 16
		card17.Val == 3, // 17
		card18.Key == "StackA", // 18
		card19.Key == "StackA", // 19
		card20.Key == "StackA", // 20
		card21.Equals(MakeCard(6, nil, 1), COMPARE_True), // 21
		card22.Equals(MakeCard(3, nil, 0), COMPARE_True), // 22
		card23.Equals(MakeCard(7, nil, 1), COMPARE_True), // 23

		// test for card comparison
		card24 == nil, // 24
		card25.Equals(cardB, nil, nil, nil, COMPARE_True), // 25

		// more pointer tests
		!card26.Equals(MakeCard(intValB, nil, 1), COMPARE_True), // 26
		card27.Equals(MakeCard(intValB, nil, 1), COMPARE_True), // 27
		card28.Equals(MakeCard(intValB, nil, 1), COMPARE_True), // 28
		card29.Equals(MakeCard(2, nil, 1), COMPARE_True), // 29
		card30 == nil, // 30
		card31 == nil, // 31

	}

	test_End(funcName, conditions)
	
}

func case_stack_GetMany(funcName string) {

	test_Start(funcName, showTestText)

	// since we already tested Get(), and since GetMany() is nearly identical to Get(), we do not need to test identical cases for GetMany

	// test base functionality
	stack1 := MakeStack([]int {1, 2, 3}).GetMany(FIND_All)
	stack2 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {3, 6}), MakeStack([]int {9, 12})}).GetMany(FIND_Last, nil, nil, DEEPSEARCH_True, []int {2})
	
	// test slices
	stack3 := MakeStack([]int {1, 2, 3, 4, 5, 6}).GetMany(FIND_Slice, MakeStack([]int {1, 4}))
	stack4 := MakeStack([]int {1, 2, 3, 4, 5, 6}).GetMany(FIND_Slice, []int {1, 4})
	stack5 := MakeStack([]int {1, 2, 3, 4, 5, 6}).GetMany(FIND_Slice, []int {0, -1})
	stack6 := MakeStack([]int {1, 2, 3, 4, 5, 6}).GetMany(FIND_Slice, []int {-1, 1})

	// test lambda
	stack7 := MakeStack([]string {"StackA", "StackB"}, []*Stack {MakeStack([]int {4, 7}), MakeStack([]int {10, 14})}).GetMany(FIND_Lambda, func(card *Card, _ *Stack, _ bool, _ *Stack, workingMem ...any) (bool) {
		prevCardVal := 0
		if workingMem[0] == nil {
			prevCardVal = 0
		} else {
			prevCardVal = workingMem[0].(*Card).Val.(int)
		}
		workingMem[0] = card
		return prevCardVal + 3 == card.Val
	}, nil, DEEPSEARCH_True, nil, PASS_Cards)
	
	stack8 := MakeStack([]int {1, 2, 3, 4}).GetMany(FIND_Lambda, func(card *Card, _ *Stack, _ bool, _ *Stack, workingStack *Stack) (bool) {
		if workingStack.Size == 0 {
			return true
		} else {
			return workingStack.Cards[workingStack.Size - 1].Val.(int) * 2 == card.Val
		}
	})

	// other
	stack9 := MakeStack([]int {1, 2, 3, 4, 5, 6}).GetMany(FIND_Idx, []int {0, -1})
	cardA := MakeCard("Hey")
	stack10 := MakeStack([]*Card {cardA}).GetMany(FIND_Card, cardA)
	stack11 := MakeStack([]*Card {cardA}).GetMany(FIND_Card, MakeCard("Hey"))
	stack12 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3}).GetMany(FIND_All, nil, RETURN_Idxs)
	stack13 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3}).GetMany(FIND_All, nil, RETURN_Vals)
	stack14 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3}).GetMany(FIND_All, nil, RETURN_Keys)
	stack15 := MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2}).GetMany(FIND_Idx, 0, RETURN_Stacks)
	stack16 := MakeStack([]*Stack {MakeStack([]int {1, 2, 3, 4}), MakeStack([]int {5, 6}), MakeStack([]int {7, 8, 9})}).GetMany(FIND_Size, []int {3, 4})

	conditions := []bool {

		// test base functionality
		stack1.Equals(MakeStack([]int {1, 2, 3})), // 1
		stack2.Equals(MakeStack([]int {6, 12})), // 2

		// test slices
		stack3.Equals(MakeStack([]int {2, 3, 4, 5})), // 3
		stack4.Equals(MakeStack([]int {2, 3, 4, 5})), // 4
		stack5.Equals(MakeStack([]int {1, 2, 3, 4, 5, 6})), // 5
		stack6.Equals(MakeStack([]int {2, 3, 4, 5, 6})), // 6
		
		// test lambda
		stack7.Equals(MakeStack([]int {7, 10})), // 7
		stack8.Equals(MakeStack([]int {1, 2, 4})), // 8

		// test other
		stack9.Equals(MakeStack([]int {1, 6})), // 9
		stack10.Equals(MakeStack([]*Card {cardA})), // 10
		!stack11.Equals(MakeStack([]string {"Hey"})), // 11
		stack12.Equals(MakeStack([]int {0, 1, 2})), // 12
		stack13.Equals(MakeStack([]int {1, 2, 3})), // 13
		stack14.Equals(MakeStack([]string {"KeyA", "KeyB", "KeyC"})), // 14
		stack15.Equals(MakeStack([]int {1, 2})), // 15
		stack16.Equals(MakeStack([]*Stack {MakeStack([]int {1, 2, 3, 4}), MakeStack([]int {7, 8, 9})})), // 16

	}

	test_End(funcName, conditions)
	
}

func case_stack_Add(funcName string) {

	test_Start(funcName, showTestText)

	// since we already tested core lambda functionality in Lambdas and Get and GetMany, we need not do it here

	// test base functionality
	stack3 := MakeStack([]int {1, 2, 3}).Add(0, ORDER_Before, FIND_All, nil, nil)
	stack5 := MakeStack([]int {1, 2, 3}).Add(MakeStack([]int {0}), ORDER_Before, FIND_All, nil, nil, nil, nil, nil, OVERRIDE_False)
	stack6 := MakeStack([]int {1, 2, 3}).Add(MakeStack([]int {0}), ORDER_Before, FIND_All, nil, nil, nil, nil, nil, OVERRIDE_True)
	stack7 := MakeStack([]int {1, 2, 3}).Add(4)
	stack8 := MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2}).Add(5, nil, FIND_Val, 2, DEEPSEARCH_True)
	stack9 := MakeStackMatrix([]int {1, 2, 3, 4}, []string {"Hi", "Hey", "Hi", "Hey"}, []int {2, 2}).Add(MakeCard(5, "He"), nil, FIND_Lambda, func(card *Card) (bool) {
		return card.Key.(int) < 4 && card.Val == "Hi"
	}, DEEPSEARCH_True, nil, PASS_Cards)
	stack10 := MakeStack().Add(MakeCard())

	// test simplification of paramateriazation for FIND_Lambda
	stack11 := MakeStackMatrix([]int {1, 2, 3, 4}, []string {"Hi", "Hey", "Hi", "Hey"}, []int {2, 2}).Add(MakeCard(5, "He"), nil, FIND_Lambda, func(card *Card) (bool) {
		return card.Key.(int) < 4 && card.Val == "Hi"
	}, DEEPSEARCH_True, nil, PASS_Cards)

	conditions := []bool {

		// test base functionality
		stack3.Equals(MakeStack([]int {0, 1, 2, 3})), // 1
		stack5.Equals(MakeStack([]int {0, 1, 2, 3})), // 2
		stack6.Equals(MakeStack([]any {MakeStack([]int {0}), 1, 2, 3})), // 3
		stack7.Equals(MakeStack([]int {1, 2, 3, 4})), // 4
		stack8.Equals(MakeStack([]*Stack {MakeStack([]int {1, 2, 5}), MakeStack([]int {3, 4})})), // 5
		stack9.Equals(MakeStack([]*Stack {MakeStack([]int {1, 5, 2}, []string {"Hi", "He", "Hey"}), MakeStack([]int {3, 4}, []string {"Hi", "Hey"})})), // 6
		stack10.Equals(MakeStack(nil, nil, 1)), // 7
		
		// test simplification of paramateriazation for FIND_Lambda
		stack11.Equals(stack9), // 8

	}

	test_End(funcName, conditions)
	
}

func case_stack_AddMany(funcName string) {

	test_Start(funcName, showTestText)

	// since we already tested core lambda functionality in Lambdas and Get and GetMany, we need not do it here

	// test base functionality
	stack1 := MakeStack([]int {1, 2, 3}).AddMany(4)
	stack2 := MakeStack([]int {1, 2, 3}).AddMany(0, ORDER_Before, FIND_First)
	stack4 := MakeStack([]int {1, 2, 3}).AddMany(0, ORDER_Before, FIND_All, nil, nil)

	conditions := []bool {

		// test base functionality
		stack1.Equals(MakeStack([]int {1, 2, 3, 4})), // 1
		stack2.Equals(MakeStack([]int {0, 1, 2, 3})), // 2
		stack4.Equals(MakeStack([]int {0, 1, 0, 2, 0, 3})), // 3

	}

	test_End(funcName, conditions)
	
}

func case_stack_Has(funcName string) {

	test_Start(funcName, showTestText)

	// don't need to test all cases like you did for Get since Has 100% implements Get
	cardA := MakeCard("Hey")
	stack1 := MakeStack([]*Card {cardA})
	
	conditions := []bool {

		// test base functionality
		MakeStack([]int {1, 5, 9}).Has(FIND_Val, 1), // 1
		MakeStack([]int {1, 5, 9}).Has(FIND_Val, 5), // 2
		MakeStack([]int {1, 5, 9}).Has(FIND_Val, 9), // 3
		!MakeStack([]int {1, 5, 9}).Has(FIND_Val, 10), // 4
		!MakeStack().Has(), // 5
		MakeStack([]int {1}).Has(), // 6
		stack1.Has(FIND_Card, cardA), // 7

	}

	test_End(funcName, conditions)
	
}

func case_stack_Move(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stack1 := MakeStack([]int {1, 2, 3}).Move()
	stack2 := MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2}).Move(nil, FIND_Val, FIND_Val, 1, 4, nil, nil, DEEPSEARCH_True, DEEPSEARCH_True)

	conditions := []bool {

		// test base functionality
		stack1.Equals(MakeStack([]int {2, 3, 1})), // 1
		stack2.Equals(MakeStack([]*Stack {MakeStack([]int {2}), MakeStack([]int {3, 4, 1})})), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_Swap(funcName string) {

	test_Start(funcName, showTestText)

	stack1 := MakeStack([]int {1, 2, 3, 4}).Swap()
	stack2 := MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2}).Swap(FIND_Val, FIND_Val, 1, 4, nil, nil, DEEPSEARCH_True, DEEPSEARCH_True)

	conditions := []bool {

		// test base functionality
		stack1.Equals(MakeStack([]int {4, 2, 3, 1})), // 1
		stack2.Equals(MakeStackMatrix([]int {4, 2, 3, 1}, nil, []int {2, 2})), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_Replace(funcName string) {

	test_Start(funcName, showTestText)

	// no need to test parameters non-unique to replace since these were tested in Add and Get

	// test REPLACE configurations
	stack1 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	card1 := stack1.Replace(REPLACE_Val, 4)
	stack2 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	card2 := stack2.Replace(REPLACE_Card, MakeCard("KeyD", 4))
	stack3 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	card3 := stack3.Replace(REPLACE_Key, "KeyD")
	stack4 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	card4 := stack4.Replace(REPLACE_Card, nil)
	stack5 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	stack5A := stack5.Clone()
	card5 := stack5A.Replace(REPLACE_Lambda, func(card *Card) {
		card.Val = card.Val.(int) * 2
	})

	// test replaceWith data types
	stack6 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	card6 := stack6.Replace(REPLACE_Card, []*Card {MakeCard("KeyD", 4), MakeCard("KeyE", 5)})
	stack7 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	card7 := stack7.Replace(REPLACE_Card, MakeStack([]string {"KeyD", "KeyE"}, []int {4, 5}))
	
	// ensure only works on one card
	stack5B := stack5.Clone()
	card8 := stack5B.Replace(REPLACE_Card, nil, FIND_All)

	// test simplification of paramaterization for REPLACE_Lambda
	card9 := stack5.Clone().Replace(REPLACE_Lambda, func(card *Card) {
		card.Val = card.Val.(int) * 2
	})

	conditions := []bool {

		// test REPLACE configurations
		stack1.Equals(MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 4})), // 1
		card1.Equals(MakeCard("KeyC", 3, 2), nil, nil, COMPARE_True), // 2
		stack2.Equals(MakeStack([]string {"KeyA", "KeyB", "KeyD"}, []int {1, 2, 4})), // 3
		card2.Equals(MakeCard("KeyC", 3, 2), nil, nil, COMPARE_True), // 4
		stack3.Equals(MakeStack([]string {"KeyA", "KeyB", "KeyD"}, []int {1, 2, 3})), // 5
		card3.Equals(MakeCard("KeyC", 3, 2), nil, nil, COMPARE_True), // 6
		stack4.Equals(MakeStack([]string {"KeyA", "KeyB"}, []int {1, 2})), // 7
		card4.Equals(MakeCard("KeyC", 3, 2), nil, nil, COMPARE_True), // 8
		stack5A.Equals(MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 6})), // 9
		card5.Equals(MakeCard("KeyC", 3, 2), nil, nil, COMPARE_True), // 10

		// test replaceWith data types
		stack6.Equals(MakeStack([]string {"KeyA", "KeyB", "KeyD", "KeyE"}, []int {1, 2, 4, 5})), // 11
		card6.Equals(MakeCard("KeyC", 3, 2), nil, nil, COMPARE_True), // 12
		stack7.Equals(MakeStack([]string {"KeyA", "KeyB", "KeyD", "KeyE"}, []int {1, 2, 4, 5})), // 13
		card7.Equals(MakeCard("KeyC", 3, 2), nil, nil, COMPARE_True), // 14

		// ensure only works on one card
		stack5B.Equals(MakeStack([]string {"KeyB", "KeyC"}, []int {2, 3})), // 15
		card8.Equals(MakeCard("KeyA", 1, 0), nil, nil, COMPARE_True), // 16

		// test simplification of paramaterization for REPLACE_Lambda
		card9.Equals(card5), // 17

	}

	test_End(funcName, conditions)
	
}

func case_stack_ReplaceMany(funcName string) {

	test_Start(funcName, showTestText)

	// don't need as many test cases since Replace covers most of them

	// test base functionality
	originalStack1 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	stack1 := originalStack1.ReplaceMany(REPLACE_Card, nil, FIND_All)
	originalStack2 := MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})
	stack2 := originalStack2.ReplaceMany(REPLACE_Card, MakeStack([]string {"KeyD", "KeyE"}, []int {4, 5}), FIND_All, nil, RETURN_Vals)

	conditions := []bool {

		// test base functionality
		originalStack1.Equals(MakeStack()), // 1
		stack1.Equals(MakeStack([]string {"KeyA", "KeyB", "KeyC"}, []int {1, 2, 3})), // 2
		originalStack2.Equals(MakeStack([]string {"KeyD", "KeyE"}, []int {4, 5}, 3)), // 3
		stack2.Equals(MakeStack([]int {1, 2, 3})), // 4

	}

	test_End(funcName, conditions)
	
}

func case_stack_Extract(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stackA := MakeStack([]string {"KeyB", "KeyC"}, []int {2, 3})
	cardA := stackA.Extract()

	conditions := []bool {

		// test base functionality
		cardA.Equals(MakeCard("KeyC", 3, 1), nil, nil, COMPARE_True), // 1
		stackA.Equals(MakeStack([]string {"KeyB"}, []int {2})), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_ExtractMany(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stackA := MakeStack([]string {"KeyB", "KeyC"}, []int {2, 3})
	stack1 := stackA.ExtractMany(FIND_All)

	conditions := []bool {

		// test base functionality
		stack1.Equals(MakeStack([]string {"KeyB", "KeyC"}, []int {2, 3})), // 1
		stackA.Equals(MakeStack()), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_Remove(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stackA := MakeStack([]string {"KeyB", "KeyC"}, []int {2, 3})
	stack1 := stackA.Remove()

	conditions := []bool {

		// test base functionality
		stackA.Equals(MakeStack([]string {"KeyB"}, []int {2})), // 1
		stack1.Equals(stackA), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_RemoveMany(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stackA := MakeStack([]string {"KeyB", "KeyC"}, []int {2, 3})
	stack1 := stackA.RemoveMany(FIND_All)

	conditions := []bool {

		// test base functionality
		stackA.Equals(MakeStack()), // 1
		stack1.Equals(stackA), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_Update(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stackA := MakeStack([]string {"KeyB", "KeyC"}, []int {2, 3})
	stack1 := stackA.Update(REPLACE_Card, MakeCard("KeyD", 4))

	conditions := []bool {

		// test base functionality
		stackA.Equals(MakeStack([]string {"KeyB", "KeyD"}, []int {2, 4})), // 1
		stack1.Equals(stackA), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_UpdateMany(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stackA := MakeStack([]string {"KeyB", "KeyC"}, []int {2, 3})
	stack1 := stackA.UpdateMany(REPLACE_Card, MakeCard("KeyD", 4), FIND_All)

	conditions := []bool {

		// test base functionality
		stackA.Equals(MakeStack([]string {"KeyD", "KeyD"}, []int {4, 4})), // 1
		stack1.Equals(stackA), // 2

	}

	test_End(funcName, conditions)
	
}

func case_stack_Coordinates(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stack := MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2})

	conditions := []bool {

		// test base functionality
		stack.Coordinates(FIND_First).Equals(MakeStack([]int {0, 0})), // 1
		stack.Coordinates(FIND_Last).Equals(MakeStack([]int {0, 1})), // 2
		stack.Coordinates(FIND_Last).Equals(MakeStack([]int {0, 1})), // 3
		stack.Coordinates(FIND_First, nil, nil, nil, PASS_Substacks).Equals(MakeStack([]int {0})), // 4
		stack.Coordinates(FIND_Last, nil, nil, nil, PASS_Substacks).Equals(MakeStack([]int {1})), // 5

	}

	test_End(funcName, conditions)
	
}

func case_stack_CoordinatesMany(funcName string) {

	test_Start(funcName, showTestText)

	// test base functionality
	stack := MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2})

	conditions := []bool {

		// test base functionality
		stack.CoordinatesMany(FIND_All, nil, nil, nil, PASS_Both).Equals(MakeStack([]*Stack {MakeStack([]int {0}), MakeStack([]int {0, 0}), MakeStack([]int {0, 1}), MakeStack([]int {1}), MakeStack([]int {1, 0}), MakeStack([]int {1, 1})})), // 1

	}

	test_End(funcName, conditions)
	
}

/** Executes all case tests */
func Run(_showTestText bool) {

	showTestText = _showTestText
	gogenerics.RemoveUnusedError(case_MakeCard, case_MakeStack, case_MakeStackMatrix, case_stack_StripStackMatrix, case_stack_Filter, case_stack_ToArray, case_stack_ToMap, case_stack_ToMatrix, case_stack_IsRegular, case_stack_Shape, case_stack_Duplicate, case_stack_Empty, case_card_Clone, case_stack_Clone, case_stack_Unique, case_card_Equals, case_stack_Equals, case_stack_Shuffle, case_stack_Flip, case_card_Print, case_stack_Print, case_stack_Lambdas, case_stack_Get, case_stack_GetMany, case_stack_Add, case_stack_AddMany, case_stack_Has, case_stack_Move, case_stack_Replace, case_stack_ReplaceMany, case_stack_Extract, case_stack_ExtractMany, case_stack_Remove, case_stack_RemoveMany, case_stack_Update, case_stack_UpdateMany, case_stack_Swap, case_CSVToStackMatrix, case_stack_ToCSV, case_stack_Coordinates, case_stack_CoordinatesMany)

	fmt.Println("- BEGINNING TESTS")

	// NON-GENERALIZED FUNCTIONS (NOT DEPENDENT ON GENERALIZED FUNCTIONS)
	case_MakeCard("MakeCard") // GOOD
	case_card_Equals("card.Equals") // GOOD
	case_MakeStack("MakeStack") // GOOD
	case_stack_Equals("stack.Equals") // GOOD
	case_MakeStackMatrix("MakeStackMatrix") // GOOD
	case_stack_Lambdas("stack.Lambdas") // GOOD
	case_stack_ToArray("stack.ToArray") // GOOD
	case_stack_ToMap("stack.ToMap") // GOOD
	case_stack_ToMatrix("stack.ToMatrix") // GOOD
	case_stack_IsRegular("stack.IsRegular") // GOOD
	case_stack_Shape("stack.Shape") // GOOD
	case_stack_Duplicate("stack.Duplicate") // GOOD
	case_stack_Empty("stack.Empty") // GOOD
	case_card_Clone("card.Clone") // GOOD
	case_stack_Clone("stack.Clone") // GOOD
	case_stack_Shuffle("stack.Shuffle") // GOOD
	case_card_Print("card.Print") // GOOD
	case_stack_Print("stack.Print") // GOOD
	case_CSVToStackMatrix("CSVToStackMatrix") // GOOD
	case_stack_ToCSV("stack.ToCSV") // GOOD
	case_stack_Coordinates("stack.Coordinates") // GOOD
	case_stack_CoordinatesMany("stack.CoordinatesMany") // GOOD
	case_stack_Flip("stack.Flip") // GOOD
	case_card_SwitchKeyVal("card.SwitchKeyVal") // GOOD
	case_stack_SwitchKeysVals("stack.SwitchKeysVals") // GOOD
	
	// GENERALIZED FUNCTIONS
	case_stack_Get("stack.Get") // GOOD
	case_stack_GetMany("stack.GetMany") // GOOD
	case_stack_Add("stack.Add") // GOOD
	case_stack_AddMany("stack.AddMany") // GOOD
	case_stack_Replace("stack.Replace") // GOOD
	case_stack_ReplaceMany("stack.ReplaceMany") // GOOD
	case_stack_Update("stack.Update") // GOOD
	case_stack_UpdateMany("stack.UpdateMany") // GOOD
	case_stack_Extract("stack.Extract") // GOOD
	case_stack_ExtractMany("stack.ExtractMany") // GOOD
	case_stack_Remove("stack.Remove") // GOOD
	case_stack_RemoveMany("stack.RemoveMany") // GOOD
	case_stack_Has("stack.Has") // GOOD
	case_stack_Move("stack.Move") // GOOD
	case_stack_Swap("stack.Swap") // GOOD
	case_stack_Filter("stack.Filter") // GOOD
	
	// NON-GENERALIZED FUNCTIONS (DEPENDENT ON GENERALIZED FUNCTIONS)
	case_stack_StripStackMatrix("stack.StripStackMatrix") // GOOD
	case_stack_Unique("stack.Unique") // GOOD

}
