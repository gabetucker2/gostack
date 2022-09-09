/**
TUTORIAL BEGINS ON LINE 20.
*/

package tutorials

import (
	"fmt"

	. "github.com/gabetucker2/gostack" //lint:ignore ST1001 Ignore warning
	"github.com/gabetucker2/gogenerics"
)

func Introduction() {

	// print that the function received the call
	fmt.Println("tutorials.Introduction()")

	/**
	//--------------------------------------------------------------------------------------//
	 Hi there.  Welcome to the gostack introductory tutorial!  This tutorial will familiarize
		you with the primary functions and concepts in gostack.  There are two main ideas
		that are essential to understand in gostack: stacks and cards.
	
	 Stacks have three attributes: Size [int], Depth [int], and an array (or "stack") of
	 	cards [[]*Cards].  Say you have a stack called "myStack"; if you wanted to access
		these attributes, you would do stack.Size, stack.Depth, or stack.Cards.  stack.Size
		is equal to len(stack.Cards); stack.Depth represents how many stacks within stacks
		there are; if you have a 1D stack (stack{card, card}), stack.Depth == 1; if you have
		a 2D stack/a stack of stacks (stack{stack{card, card}, stack{card, card}}),
		stack.Depth == 2.  You should never have to access stack.Cards unless you are writing
		your own lambda functions, which are covered in the Lambda.go tutorial.

	 Cards have three attributes: Idx [int], Key [interface/any], and Val [interface/any].
	 	Idx is an int value representing the index, or order (beginning at 0), of your card
		in the stack.Cards array.  Key and Val can be set to anything you'd like, but Key
		should almost always be a mutable (e.g., string, int, float) type so that it's easy
		to use as a reference when you're searching for Vals.  Val, on the other hand, should
		be the main information you're storing.  You do not need to define your Key or Val for
		a card, in which case their values are nil.  Also, if and only if a card is not in a
		stack, its Idx will be -1.  Further, the same card should *never* be in two different
		stacks, or else this will lead to internal bugs (for instance, if your card is in the
		first position of myStack1 and in the third position of myStack2, then to what value
		what should card.Idx be set?).
	
	 That was a lot of information, so let's move on to some more intuitive examples.

	 Let's start by making a stack.  This stack will be very simple and will look like the
	 	following:
		
		myStack Stack {
			Size: 0,
			Depth: 1,
			[]*Cards: { },
		}*/
	
	myStack := MakeStack()//lint:ignore SA4006 Ignore warning (please ignore these types of lint comments)

	/**
	 Great!  We have now made an empty stack.  Now let's make some cards we can use to add
	 	to the stack.  The first will look like the following:
		
		cardA Card {
			Idx: -1,
			Key: nil,
			Val: "Butterfly"
		}*/

	cardA := MakeCard("Butterfly")
	cardB := MakeCard("Praying Mantis")
	cardC := MakeCard("Beetle")

	/**
	 We have now created three cards.  But these cards are kind of awkwardly floating around
	 	in the void, so let's put them in our myStack.  There are several ways we could do this
		that are useful in different situations (feel free to ignore the alternative examples):*/
	
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// if you have already created your cards but have not yet created your stack:
	myStack = MakeStack([]*Card {cardA, cardB, cardC})//lint:ignore SA4006 Ignore warning

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// if you have neither created your cards nor your stack:
	 // makes a stack of cards, each of whom's vals are a string in an array of strings (recommended)
	myStack = MakeStack([]string {"Butterfly", "Praying Mantis", "Beetle"})//lint:ignore SA4006 Ignore warning
	
	//   ALTERNATIVE:
	// makes a stack by directly initializing it using cards
   myStack = MakeStack([]*Card {MakeCard("Butterfly"), MakeCard("Praying Mantis"), MakeCard("Beetle")})

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// if you have already created your cards and your stack and you don't want to replace your existing stack:
	 // creates a temporary stack and inserts its cards into myStack (recommended)
	myStack.Add(MakeStack([]*Card {cardA, cardB, cardC}))

	 //   ALTERNATIVES:
	 // repeatedly adds to the front of myStack in reverse order
	myStack.Add(cardC).Add(cardB).Add(cardA)
	 // repeatedly adds to the back of myStack
	myStack.Add(cardA, ORDER_After, FIND_Last).Add(cardB, ORDER_After, FIND_Last).Add(cardC, ORDER_After, FIND_Last)
	 // creates an array of cards and uses a loop to add them each to the back of myStack
	myCards := []*Card {cardA, cardB, cardC}; for _, c := range myCards { myStack.Add(c, ORDER_After, FIND_Last) };
	 // creates an array of cards and uses a loop to add them in reverse order to the front of myStack
	myCards = []*Card {cardA, cardB, cardC}; for i := range myCards { myStack.Add(myCards[len(myCards)-1-i]) };

	/**
	 Each of these lines of code are different ways of creating the same thing.  Let's
	 	draw this out visually:
		
		myStack Stack {
			Size: 3,
			Depth: 1,
			[]*Cards: {
				cardA Card {
					Idx: 0,
					Key: nil,
					Val: "Butterfly"
				},
				cardB Card {
					Idx: 1,
					Key: nil,
					Val: "Praying Mantis"
				},
				cardC Card {
					Idx: 2,
					Key: nil,
					Val: "Beetle"
				},
			},
		}
		
	 Simple enough!  Next, let's say we were sad enough to make a book about insects.
	 	This is a very strange book: it is just a list of insect names paired with the
		amount of years that have passed since their discovery.  This is a map structure:
		there can only be one of each insect name, but there can be many of the same year
		discovered.  As such, this is how we would make our data structure:*/

	myStack = MakeStack([]string {"Butterfly", "Praying Mantis", "Beetle", "Ant", "Bumble Bee"}, []int {20, 539, 539, 340, 11})

	/**
	 Great.  So now, let's say we wanted to access this structure in various ways.  We
	 	would do this using the stack.Get/GetMany(...) function.  For simplicity's sake,
		we will show only the keys of the card(s) returned.  Here's some examples:*/

	var gottenData any
	gogenerics.RemoveUnusedError(gottenData) // (please ignore this line)
	
	// to get the first card in myStack: card {Idx: 0, Key: "Butterfly", Val: 539}:
	gottenData = myStack.Get(FIND_First) // Butterfly
	gottenData = myStack.Get(FIND_First, nil) // Butterfly
	gottenData = myStack.Get(FIND_Idx, 0) // Butterfly

	// to get the first card with the val of 539:
	gottenData = myStack.Get(FIND_Val, 539) // Praying Mantis

	// to get a stack containing all cards with the val of 539:
	gottenData = myStack.GetMany(FIND_Val, 539) // Praying Mantis, Beetle

	// to get all cards whose vals are either 20 or 340
	gottenData = myStack.GetMany(FIND_Val, []int {20, 340}) // Butterfly, Ant
	gottenData = myStack.GetMany(FIND_ValsStack, MakeStack([]int {20, 340})) // Butterfly, Ant

	// to get all cards whose vals are under 30 (read more on this type of functionality in Lambda.go)
	gottenData = myStack.GetMany(FIND_Lambda, func(card *Card) bool {
		return card.Val.(int) < 30
	}) // Butterfly, Bumble Bee

	/**
	 Exciting!  So we have now have performed many different types of retrievals on our
	 	myStack stack.  But, in each output, we have gotten either a Card (from .Get()) or
		a stack of cards (from .GetMany()).  But what if we wanted to make a new stack
		whose values are the keys of myStack?  For instance, if you wanted to play around
		with the types of insects without messing up myStack, what would you do?*/
	
	gottenData = myStack.GetMany(FIND_All, nil, RETURN_Keys)

	/**
	 There we go!  We now have a stack whose cards' keys are nil and whose cards vals are
	 	myStack's keys.  You can do the equivalent for Idxs or Vals.
	
	 But what if we wanted to do something more complex, like removing cards that were found
	 	in the stack, or replacing them with another value?*/

	// remove found cards for gottenData from a clone of myStack
	gottenData = myStack.Clone().ExtractMany(FIND_Val, 539) // Praying Mantis, Beetle

	// set gottenData to myStack where myStack's first card is placed after its last card
	gottenData = myStack.Move(FIND_First, ORDER_After, FIND_Last)

	// set gottenData to a clone of myStack where its keys point to clones of old myStack key objects, but its values point to actual myStack val objects
	gottenData = myStack.Clone(CLONE_True, CLONE_False)

	// set gottenData to a clone of myStack which has another card appended to the end of its stack
	gottenData = myStack.Clone().Add(MakeCard("Moth", 400), ORDER_After, FIND_Last)

	// replace first card in myStack with a moth card, setting gottenData to the card that was replaced
	gottenData = myStack.Replace(REPLACE_Card, MakeCard("Moth", 400), FIND_First)

	// set gottenData to a clone of myStack where the first card and the second card are switched
	gottenData = myStack.Clone().Swap(FIND_First, FIND_Idx, nil, 1)

	// set gotten data to a clone of myStack which is put into a random order, flipped in the opposite order, and with any cards who have the same value as other cards removed
	gottenData = myStack.Clone().Shuffle().Flip().Unique(TYPE_Val)

	/**
	 There are many more functions than the ones displayed here that you are welcome
	 	to read about in the git's README.md, but the these examples should suffice for
		providing intuitions about how gostack works.  If this tutorial could benefit
		from any form of improvement, please email me at tucker.854@osu.edu to let me
		know.  I welcome your feedback and appreciate you reading this.*/

}
