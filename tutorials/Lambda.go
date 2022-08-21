/**
TUTORIAL BEGINS ON LINE 32.
*/

package tutorials

import (
	"fmt"
	"math"

	"github.com/gabetucker2/gogenerics"
	. "github.com/gabetucker2/gostack"//lint:ignore ST1001 Ignore warning
)

func lambda_ValInRange(card *Card, workingMem ...any) bool {
	v := card.Val.(int)
	return 5 < v && v < 14 && v%2 == 0
}

func lambda_KeyInRange(card *Card, workingMem ...any) bool {
	k := card.Key.(int)
	return k%5 == 0
}

/** Executes the Lambda.go tutorial */
func Lambda() {

	// print that the function received the call
	fmt.Println("tutorials.Lambda()")

	/**
	//--------------------------------------------------------------------------------------//
	 Hi there.  Welcome to our lambda tutorial!  This is the most advanced tutorial in
	 	gostack, but it is also our most exciting because lambda functions are the most
		powerful tools in gostack.  If you are reading this, you should be at the point of
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
	 You can also use FIND_Lambda to filter for any other function, including UpdateMany:*/
	
	makeSampleStack().UpdateMany(REPLACE_Val, 2, FIND_Lambda, lambda_ValInRange).Print() // val: 10

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
	
	 Approach A is less grounded in gostack: you create a for loop, create a variable
		representing the current highest integer (probably initialized to the lowest
		representable integer), and, for each card in stack.Cards, if that card's
		value is greater than your variable, set your variable to that value.  After
		the loop, output the value.
	
	 What is disappointing here is how the more optimized approach requires you to write
	 	for loops.  But if you are strongly grounded in the principle that for loops should
		be abstracted away if possible, Approach B is an approach supported by gostack.

	 Remember earlier when I said to ignore the `...any` field?  Now, we are going to utilize
		it to track what we will call our working memory.  We will use our gogenerics library
		to help unpack variables from the working mem.  Working mem will allow us to keep
		track of information between each iteration over a card.  Otherwise, we would have to
		run the same for loop in each card iteration, in turn wasting a myriad of (computer)
		memory.
	*/

	makeSampleStack().Get(FIND_Lambda, func(card *Card, workingMem ...any) bool {

		// unpack variadic arguments
		var stack, maxIdx any
		gogenerics.UnpackVariadic(workingMem, &stack, &maxIdx)
	
		// if max == nil, that means the for loop determining the Idx of the highest card has
		//not yet been run (expect initial max input value of nil)
		if maxIdx == nil {
			currentMax := math.MinInt
			for _, card := range stack.(*Stack).Cards {
				if card.Val.(int) > currentMax {
					currentMax = card.Val.(int)
					maxIdx = card.Idx
				}
			}
		}
		// now, return whether this card's Idx == the idx of the card with the highest value
		return maxIdx == card.Idx
	
	}).Print() // val: 40

	/**
	 Great!  We still had to write a for loop, but at least it's localized to the inside
		of a gostack function.  This approach might be unpreferable when it comes to running
		a function only once, but if you formally declare your lambda function which utilizes
		working memory, you can call that function multiple times, in turn saving space.

	 For functions offering replace functionality, you can use REPLACE_Lambda in order to
	 	apply a custom transformation to the stack.  For a complex example, like the one
		above, you could manage working memory to multiply selected cards by the previous
		cards values.  For a simpler example, where you select some cards and multiply it by
		two:*/
	
	makeSampleStack().UpdateMany(REPLACE_Lambda, func(card *Card, workingMem ...any) {
		card.Val = card.Val.(int) * 2 // multiply selected cards by two
	}, FIND_All).Print() // vals: 4, 20, 22, 24, 80

	/**
	 If you wanted to get really fancy, you could use lambda support to replace AND to
	 	select cards.  In this case, let's apply the same *2 multiplier to only the cards
		which are even before the multiplication:*/
	
	makeSampleStack().UpdateMany(REPLACE_Lambda, func(card *Card, workingMem ...any) {
		card.Val = card.Val.(int) * 2 // multiply selected cards by two
	}, FIND_Lambda, func(card *Card, stack *Stack, _ ...any) bool {
		return card.Val.(int) % 2 == 0 // select even cards
	}).Print() // vals: 4, 20, 11, 24, 80

	/**
	 We have now shown you how to use lambda support in core gostack functions. But what
	 	if you wanted to sort a stack in a certain order?  This would require you to update
		the position of cards relative to other cards, which our core functions do not allow
		you to do.  So this is where our Lambda function comes into play.  Below is a behind-
		the-curtain example of how the Flip function, which reverses the order of cards in
		a stack, was implemented (using the Lambda function):*/

	makeSampleStack().Lambda(func(card *Card, stack *Stack, _ ...any) {
		// moves each card, from first to last, to the first position in the stack
		stack.Move(FIND_Card, ORDER_Before, FIND_Idx, card, 0)
	})

	/**
	 For a more complex example of using the Lambda function, e.g., to sort in descending
	 	order by each card's int value:*/

	 makeSampleStack().Lambda(func(card *Card, stack *Stack, _ ...any) {
		thisVal := card.Val.(int) // the current card
		for i := card.Idx+1; i < stack.Size; i++ {
			otherCard := stack.Cards[i] // the card being compared
			otherVal := otherCard.Val.(int)
			if otherVal >= thisVal {
				// if the card being compared is greater than this card, swap this card with the other card,
				// thereby making the smallest cards end up at the end of the list
				stack.Swap(FIND_Card, FIND_Card, card, otherCard)
			}
		}
	})

	/**
	 That's all you need to know about lambda support in gostack!  If this tutorial could
	 	benefit from any form of improvement, please email me at tucker.854@osu.edu to let
	 	me know.  I welcome your feedback and appreciate you reading this.*/

}
