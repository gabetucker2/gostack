package tutorials

import (
	"fmt"
)

// TEMPLATE:
/*
func gostack_NameHere(stack *Stack, card *Card, workingMemory *interface{}) (ret bool) {

	if workingMemory == nil { // first run setup
		workingMemory = MakeStack()

		// first run stuff here (e.g., for loop)

		// workingMemory.Add(MakeCard(dataToAccess)) // add to working memory
	}

	// stuff here

	return

}
*/

func gostack_lambda_ValInRange(stack *Stack, card *Card, workingMemory ...*Stack) bool {
	v := card.val.(int)
	return 5 < v && v < 14 && v%2 == 0
}

func gostack_lambda_KeyInRange(stack *Stack, card *Card, workingMemory ...*Stack) bool {
	k := card.key.(int)
	return k%5 == 0
}

func gostack_lambda_BothInRange(stack *Stack, card *Card, workingMemory ...*Stack) bool {
	return gostack_ValInRange(stack, card) && gostack_KeyInRange(stack, card)
}

func gostack_lambda_Max(stack *Stack, card *Card, workingMemory ...*Stack) bool {

	if workingMemory == nil { // first run setup
		workingMemory = MakeStack()
		for _, c := range stack.cards {
			v := c.val.(int)
			if workingMax > v {
				workingMax = v
			}
		}
		workingMemory.Add(MakeCard(workingMax))
	}

	return workingMemory.Get(RETURN_Card, POSITION_First) == card

}

func (stack *Stack) MainFunc(lambda func(*Stack, *Card) bool) { // TODO: delete

	stack.gostack_back_iterator(lambda)

	fmt.Println(" - vals after:")
	for _, card := range stack.cards {
		fmt.Println(card.val)
	}

	fmt.Println()

}

func makeSampleStack() *Stack { // very rough ugly outline
	ivals := []int{2, 10, 11, 12, 40}
	kvals := []int{0, 90, 4, 2, 20}
	stack := new(Stack)
	for i := range ivals {
		newC := new(Card)
		newC.val = ivals[i]
		newC.key = kvals[i]
		stack.cards = append(stack.cards, newC)
	}
	return stack
}

func Lambda() {

	// `stack.Get(RETURN_*, POSITION_*, ...POSITIONDATA, ...MATCH_*)`

	makeSampleStack().MainFunc(gostack_lambda_ValInRange) // 10, 12

	makeSampleStack().MainFunc(gostack_lambda_KeyInRange) // 2, 10, 40

	makeSampleStack().MainFunc(gostack_lambda_BothInRange) // 10

	makeSampleStack().Get(.., gostack_lambda_Max) // 40

}
