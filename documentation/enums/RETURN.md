![Banner](../../media/gostack_SmallerTransparent.png)

<h2>RETURN</h2>

 > Idxs, Keys, Vals, Cards, Adrs, Stacks

Many functions have "return functionality".  This means they A) return a *Stack and B) will make the stack a set of cards whose vals are the return type specified (or whose cards are simply clones of the original cards).  Sample:
 > enumerator type's name
 >> what type of *Stack the function returns given this enumerator type input

Enumerators:
 > RETURN_Idxs
 >> a stack whose vals are the indices of the cards in stack
 >
 > RETURN_Keys
 >> a stack whose vals are the keys of the cards in stack
 >
 > RETURN_Vals
 >> a stack whose vals are the vals of the cards in stack
 >
 > RETURN_Cards
 >> a stack whose cards are clones of cards in stack
 >
 > RETURN_Adrs
 >> a stack whose vals are the object addresses of cards in stack
 >
 > RETURN_Stacks
 >> a stack whose vals are clones of the cards within a found card's substack

 ---

 [> Return to enumerators](../enumsAPI.md)