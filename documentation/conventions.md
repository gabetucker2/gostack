![Banner](../images/gostack_Smaller.png)

<h1>Conventional Documentation</h1>

<h2>The Golden Rule</h2>

 Never *ever* enter the same **Card** object into more than one **Stack**.  Doing so will break your cards/stacks.  If you would like to have the same card in two stacks, you must first make a clone of the card.

 Given the following:

 ```
 stackA := MakeStack([]int {1, 2})
 cardA := stackA.Get(FIND_First)
 ```

 ...this is a crime:

 ```
 stackB := MakeStack([]*Card {cardA})
 ```

 ...and this is morally permissable:

 ```
 stackB := MakeStack([]*Card {cardA.Clone()})
 ```

 ---

 [> Return to glossary](../README.md)