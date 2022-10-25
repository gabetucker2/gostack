![Banner](../images/gostack_Smaller.png)

<h1>Conventional Documentation</h1>

<h2>The Golden Rule</h2>

 Never (*ever*) enter the same **Card** object into more than one **Stack**.  Doing so will break your cards/stacks.  If you would like to add a card to a new stack, you must either A) remove the card from the first stack, or B) make a clone of the card.

 Given the following:

 ```
 stack := MakeStack([]int {1, 2})
 card := stack.Get(FIND_First)
 ```

 ...this is a crime:

 ```
 newStack := MakeStack([]*Card {card})
 ```

 ...and these are morally permissable:

 ```
 stack.Remove(card)
 newStack := MakeStack([]*Card {card})
 ```

 ```
 newStack := MakeStack([]*Card {card.Clone()})
 ```

 It should be noted that functions which return a *Stack, like GetMany, automatically clone each card.

 ---

 [> Return to glossary](../README.md)