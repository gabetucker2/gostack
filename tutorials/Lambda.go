/**
TUTORIAL BEGINS ON LINE 72.
*/

package tutorials

import (
	"fmt"
	"math"

	"github.com/gabetucker2/gogenerics"
	. "github.com/gabetucker2/gostack"//lint:ignore ST1001 Ignore warning
)

// TEMPLATE:
/*
func gostack_tutorials_lambda_NameHere(card *Card, workingMem any) (ret bool) {

	if workingMem == nil { // first run setup
		workingMem[0] = MakeStack()

		// first run stuff here (e.g., for loop)

		// workingMem.Add(MakeCard(dataToAccess)) // add to working memory
	}

	// stuff here

	return

}
*/

func lambda_ValInRange(card *Card, workingMem ...any) bool {
	v := card.Val.(int)
	return 5 < v && v < 14 && v%2 == 0
}

func lambda_KeyInRange(card *Card, workingMem ...any) bool {
	k := card.Key.(int)
	return k%5 == 0
}

func lambda_Max(card *Card, workingMem ...any) bool {

	var stack, maxIdx any
	gogenerics.UnpackVariadic(workingMem, &stack, &maxIdx)

	// if max == nil, that means the for loop determining the Idx of the highest card has not been run yet
	// (expect initial max input value of nil)
	if maxIdx == nil {
		currentMax := math.MinInt
		for _, card := range stack.(*Stack).Cards {
			if card.Val.(int) > currentMax {
				currentMax = card.Val.(int);
				maxIdx = card.Idx
			}
		}
	}
	return maxIdx == card.Idx

}

/** Executes the Lambda.go tutorial */
func Lambda() {

	// print that the function received the call
	fmt.Println("tutorials.Lambda()")

	/**
	//--------------------------------------------------------------------------------------//
	 Hi there.  Welcome to our lambda tutorial!  This is the most advanced tutorial in
	 	gostack, but also our most exciting, because lambda functions are the most powerful
		tools in gostack.  If you are reading this, you should be at the point of
		understanding everything else in gostack in order to understand what is covered
		in this script.  By the time you are done reading, you will understand the most
		complex feature of gostack, meaning you will be capable of utilizing this
		library to its fullest potential.
		
	 Consider: if you wanted to do something simple like getting a stack whose values are
		the keys of another stack, gostack's traditional functions have you covered.  If you
		want to replace the 5th card in a stack with another card, core functions also have
		you covered.  The same is true of pulling information from the Nth dimension in
		stack matrices, cloning, or directing multiple stacks to reference the same object.
		
	 But if you wanted to create a custom filter for a stack (e.g., get all
		cards whose Vals are over 2, whose Keys are multiples of 3, and whose Idxs are
		even), you would traditionally have to write your own for loop with a nested if
		statement and extract the cards that match your condition from the old stack to a
		new stack.  gostack's support, which we will call lambda support, removes this
		necessity.

	 Below, we are creating a simple function to return a new stack with specified properties:*/

	makeSampleStack := func() *Stack {
		//                     keys                    vals
		return MakeStack([]int{0, 90, 4, 2, 20}, []int{2, 10, 11, 12, 40})
		// card idx:           0, 1,  2, 3, 4          0, 1,  2,  3,  4
	}

	/**
	 Reference this function in the future to get an idea of how it is being affected by
	 	our lambda functions.

	 Next, we are going to pass the FIND_Lambda argument to the GetMany() function.  Our
	 	goal will be to get a stack of cards from another stack based on which cards match
	 	our custom filter.  Our custom filter will be a function of structure
		`func(*Card, ...any) bool`.  (For now, don't worry about the `...any` field.)  Our
		function, which we will call lambda_ValInRange, is going to do something very
		simple: it will return true, or false, respective of whether the inputted card
		matches your filter.  Given v is card.Val, our filter will test whether v is
		between 5 and 14 and whether it's a multiple of 2.  See lambda_ValInRange on
		line 33.
		
	 Next, let us apply it below:*/
	
	makeSampleStack().GetMany(FIND_Lambda, lambda_ValInRange).Print() // vals: 10, 12
	
	/**
	 It prints the cards with vals 10, 12!  Because of 2, 10, 11, 12, and 40, 10 and 12
	 	are the only values that match the conditions defined in lambda_ValInRange!

	 One limitation of this approach is how golang does not provide support for nested
	 	functions, meaning lambda_ValInRange is in an awkward position being so far from
		its call.  If we wanted to make lambda_ValInRange an anonymous function (calls
		once) that's all in one location, we would do the following:*/
	
	makeSampleStack().GetMany(FIND_Lambda, func(card *Card) bool {
		v := card.Val.(int)
		return 5 < v && v < 14 && v%2 == 0
	}).Print() // vals: 10, 12

	/**
	 Notice how this approach is identical to our previous approach.

	 And, just for a reminder, calling Get, as opposed to GetMany, on lambda_ValInRange
	 	yields the first card found matching the filter:*/

	makeSampleStack().Get(FIND_Lambda, lambda_ValInRange).Print() // val: 10

	/**
	 Great!  Just for example's sake, let's make another function filtering by
	 	properties of each card's key:*/

	makeSampleStack().GetMany(FIND_Lambda, lambda_KeyInRange).Print() // keys: 0, 90, 20

	/**
	 Stunning.  Now, let's get only the cards whose vals match the lambda_ValInRange
		condition and whose keys match the lambda_KeyInRange condition:*/

	makeSampleStack().GetMany(FIND_Lambda, func(card *Card) bool {
		return lambda_ValInRange(card) && lambda_KeyInRange(card)
	}).Print() // card: {Idx: 1, Key: 90, Val: 10}

	/**
	 This should hopefully have given you an intuitive feel for how lambda functions
		work.  But what if you wanted your filter to return true if and only if the card
		is the maximum in your current stack?  This would require you to keep track of
		information outside the scope of your current card.  In this case, there are two
		approaches you could take.
	
	 Approach A, sadly, is more optimized: you create a for loop, create a variable
		representing the current highest integer (probably initialized to the lowest
		representable integer), and, for each card in stack.Cards, if that card's
		value is greater than your variable, set your variable to that value.  After
		the loop, output the value.
	
	 What is disappointing here is how the more optimized approach requires you to write
	 	for loops.  But if you are strongly grounded in the principle that for loops should
		be abstracted away if possible, Approach B is an approach supported by gostack.

	 Remember earlier when I said to ignore the `...any` field?  Now, we are going to utilize
		it to track what we will call our working memory.  This will allow us to keep track
		of information between each iteration over a card.
	*/

	makeSampleStack().Get(FIND_Lambda, lambda_Max).Print() // val: 40
	makeSampleStack().UpdateMany(REPLACE_Val, 5, FIND_Lambda, lambda_Max).Print() // vals: 2, 10, 11, 12, 5

	// do for matrices

	// varying outputs using gostack's Sort() function
	

	// varying outputs using gostack's Lambda() function
	makeSampleStack().Lambda(func(card *Card, _ ...any) {
		// if condition is not true, set value to -1
		v := card.Val.(int)
		if 5 < v && v < 14 && v%2 == 0 {
			card.Val = -1
		}
	}).Print() // 2, -1, 11, -1, 40
	// - another way of doing: makeSampleStack().UpdateMany(REPLACE_Val, -1, FIND_Lambda, lambda_ValInRange).Print()

}
