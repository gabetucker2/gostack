package tutorials

<<<<<<< HEAD
import (
	. "github.com/gabetucker2/gostack" //lint:ignore ST1001 Ignore warning
)

/** Executes the Bootstrap.go tutorial */
func Bootstrap() {

	// makeSampleStack := func() *Stack {
	// 	return MakeStack([]int{2, 10, 11, 12, 40}, []int{0, 90, 4, 2, 20})
	// }

	// myStack := makeSampleStack()
	// bootstrap_

=======
/** Executes the Bootstrap.go tutorial */
func Bootstrap() {

>>>>>>> e208db9d0fdff8169bdf79d8cff564711cdce50f
}

// QUESTIONS FOR GABE!
//am i supposed to pass in an array or a stack to these functions?
//how to read function parameters/definition (what is the style called?)
//Similarly, I honestly have no idea the order of inputs lol
//Thought I understand the enum, but now I am confused
//How can i make this executable so i can test output and make sure its doing what i want it to?

func MakeArr() []int {
	return []int{1, 2, 3}
}

/*
Adds a card/value to the front of an array/any and returns a 'stack'

input is an array/any and a number you want to append to the front
*/
func Push(myArray any, myVal any) *Stack {
	return MakeStack(myArray).Add(MakeCard(myVal), ORDER_Before, FIND_First)
}

//create an array with elements 1, 2, 3 then push the value 4 to the front and store its value in myStack
// myStack := MakeArr().Push(4) // {1, 2, 3, 4}

/*
Removes the first card from and returns 'stack'

input is array/any
*/
func Pop(myArray any) *Card {
	return MakeStack(myArray).Extract(FIND_First)
}

//create an array with elements 1 then remove the first value from the front and store its value in myStack
// firstCard := MakeArr().Pop() // 1

/*
find and return a cards value at an index
*/
func getValAtIdx(myStack *Stack, myIdx int) any {
	return myStack.Get(FIND_Idx, myIdx).Val
}

//make a stack with 3 cards (and the cards have no keys) - Butterfly, Praying Mantis, and Beetle
// myStack := MakeStack([]string{"Butterfly", "Praying Mantis", "Beetle"})
// myCard := myStack.getValAtIdx(2) // myCard = "Beetle"

/*
add a value to the end of a stack
*/
func Append(myStack *Stack, myVal any) *Stack {
	return myStack.Add(MakeCard(myVal), ORDER_After, FIND_Last)
}

//make a stack with 3 cards (and the cards have no keys) - Butterfly, Praying Mantis, and Beetle
// myStack := MakeStack([]string{"Butterfly", "Praying Mantis", "Beetle"})
// myStack = myStack.Append("Ant") // Adds the card "Ant" to the end of myStack, leaving myStack = Butterfly, Praying Mantis, Beetle, and Ant

/*
Return the last card of a stack
Would you make this a real function, or would you just type whats inside the function if you wanted to excrete something?
*/
func Excrete(myStack *Stack) *Card {
	return myStack.Extract(FIND_Last)
}

//make a stack with 3 cards (and the cards have no keys) - Butterfly, Praying Mantis, and Beetle
// myStack := MakeStack([]string{"Butterfly", "Praying Mantis", "Beetle"})
// beetleCard = myStack.Excrete() // removes the last card (Beetle) leaving myStack = Butterfly, Praying Mantis

/*
	multiply all the values of a card in a stack by the inputted value
	question - this is probably wrong anyway, but are the cards in myStack updated and passed back with pointers according to this method?
	seems easy to pull and manipulate single cards, but how do I apply something to every card?
*/

// func (myStack *Stack) MultiplyList(myVal any) (myStack *Stack) {
// 	for i := range stack.Cards {
// 		myStack.Card.Val[i] *= myVal //get the value of a card at [i] and multiply it by 5
// 	}
// 	return
// }

// func (myStack *Stack) MultiplyList(myVal int) (myStack *Stack) {
// 	return myStack.Replace(REPLACE_Lambda, func(card *Card) {
// 		card.Val = card.Val.(int) * myVal // multiply selected cards by myVal
// 	}, FIND_All)
// }

// //make a stack with 3 cards with values 1, 2, and 3
// // myStack := MakeStack([]*Card{MakeCard(1), MakeCard(2), MakeCard(3)})
// // myStack = myStack.MultiplyList(5) //myStack is now = 5, 10, 15

// /*
// moves all values one to the left
// again, how do i iterate through many cards at once without a loop in gostack?
// */
// func (myStack *Stack) MoveLeft() (myStack *Stack) {
// 	return myStack.Move(FIND_First, ORDER_After, FIND_Last)
// }

// func (myStack *Stack) MoveLeft(val any) (myStack *Stack) {
// 	return myStack.Move(FIND_Val, ORDER_Before, FIND_Idx, val, card.Idx-1)
// }

// // myStack.MoveLeft("Praying Mantis")

// // //make a stack with 4 cards with values 1, 2, 3, and 4
// // myStack := MakeStack([]*Card{MakeCard(1), MakeCard(2), MakeCard(3), MakeCard(4)})
// // myStack = myStack.MoveLeft() //myStack is now = {2, 3, 4, 1}

// // func (myStack *Stack) SortDescending () (myStack *Stack) {

// // }
