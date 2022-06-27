package tutorials

import (
	"fmt"

	. "github.com/gabetucker2/gostack"
)

// TEMPLATE:
/*
func gostack_tutorials_lambda_NameHere(card *Card, workingMemory interface{}) (ret bool) {

	if workingMemory == nil { // first run setup
		workingMemory[0] = MakeStack()

		// first run stuff here (e.g., for loop)

		// workingMemory.Add(MakeCard(dataToAccess)) // add to working memory
	}

	// stuff here

	return

}
*/

func lambda_ValInRange(card *Card, workingMemory ...interface{}) bool {
	v := card.Val.(int)
	return 5 < v && v < 14 && v%2 == 0
}

func lambda_KeyInRange(card *Card, workingMemory ...interface{}) bool {
	k := card.Key.(int)
	return k%5 == 0
}

func lambda_BothInRange(card *Card, workingMemory ...interface{}) bool {
	return lambda_ValInRange(card) && lambda_KeyInRange(card)
}

// (card, wm[0] = stack, wm[1] = workingMemory)
func lambda_Max(card *Card, workingMemory ...interface{}) bool {

	if workingMemory[1] == nil { // first run setup
		var workingMax int
		workingMemory[1] = MakeStack()
		for i, c := range workingMemory[0].(Stack).Cards {
			v := c.Val.(int)
			if i == 0 || workingMax > v {
				workingMax = v
			}
		}
		//workingMemory[1].Add(MakeCard(workingMax))
	}

	return false//workingMemory[1].Get(RETURN_Card, FINDBY_First) == card

}

func makeSampleStack() *Stack { // very rough ugly outline
	return MakeStack([]int{2, 10, 11, 12, 40}, []int{0, 90, 4, 2, 20})
}

/** Executes the Lambda.go tutorial */
func Lambda() {

	fmt.Println("tutorials.Lambda()")

	//////////////////////////////////

	// val in range
	makeSampleStack().Get(func(card *Card) bool {
		v := card.Val.(int)
		return 5 < v && v < 14 && v%2 == 0
	}) // 10, 12

	// SAME AS

	makeSampleStack().GetMany(FINDBY_Lambda, lambda_ValInRange) // 10, 12

	//////////////////////////////////

	makeSampleStack().GetMany(FINDBY_Lambda, lambda_KeyInRange) // 2, 10, 40

	makeSampleStack().GetMany(FINDBY_Lambda, lambda_BothInRange) // 10

	makeSampleStack().GetMany(FINDBY_Lambda, lambda_Max) // 40

}
