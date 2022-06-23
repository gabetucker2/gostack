package main

import (
	"fmt"
	"github.com/gabetucker2/gostack/aorta/*"
)

type Card struct {
	key interface{}
	val interface{}
}

type Stack struct {
	cards []*Card
	size  int
}

type TYPE int

const (
	TYPE_Key TYPE = iota
	TYPE_Val
	TYPE_Card
)

// TEMPLATE:
/*
func gostack_NameHere(stack *Stack, card *Card, workingMemory *interface{}) (ret bool) { // TODO: look into workarounds for efficiency cost

	if workingMemory == nil { // first run setup
		workingMemory = MakeStack()

		// first run stuff here (e.g., for loop)

		// workingMemory.Add(MakeCard(dataToAccess)) // add to working memory
	}

	// stuff here

	return

}
*/

func gostack_Max(stack *Stack, card *Card, workingMemory ...*Stack) (ret bool) {

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

func gostack_ValInRange(stack *Stack, card *Card, workingMemory ...*Stack) (ret bool) {
	v := card.val.(int)
	return 5 < v && v < 14 && v%2 == 0
}

func gostack_KeyInRange(stack *Stack, card *Card, workingMemory ...*Stack) (ret bool) {
	k := card.key.(int)
	return k%5 == 0
}

func gostack_BothInRange(stack *Stack, card *Card, workingMemory ...*Stack) (ret bool) {
	return gostack_ValInRange(stack, card) && gostack_KeyInRange(stack, card)
}

// TODO: rename iterator
func (stack *Stack) Iterator(lambda func(*Stack, *Card) bool) {
	newStack := new(Stack)
	for _, card := range stack.cards {
		if lambda(stack, card) {
			newStack.cards = append(newStack.cards, card)
		}
	}
	stack.cards = newStack.cards
}

func (stack *Stack) MainFunc(lambda func(*Stack, *Card) bool) {

	stack.Iterator(lambda)

	fmt.Println(" - vals after:")
	for _, card := range stack.cards {
		fmt.Println(card.val)
	}

	fmt.Println()

}

func MakeSampleStack() *Stack { // very rough ugly outline
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

func main() {

	fmt.Println(" - vals before:")
	for _, card := range MakeSampleStack().cards {
		fmt.Println(card.val)
	}

	fmt.Println()

	fmt.Println("Max")
	MakeSampleStack().MainFunc(gostack_Max) // 40

	fmt.Println("ValInRange")
	MakeSampleStack().MainFunc(gostack_ValInRange) // 10, 12

	fmt.Println("KeyInRange")
	MakeSampleStack().MainFunc(gostack_KeyInRange) // 2, 10, 40

	fmt.Println("BothInRange")
	MakeSampleStack().MainFunc(gostack_BothInRange) // 10

	//stack.Get(......, gostack_ValInRange)

}
