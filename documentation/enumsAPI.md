![Banner](../images/gostack_Smaller.png)

<h1>Enumerator Documentation</h1>

> [FIND](#FIND)

> [REPLACE](#REPLACE)

> [RETURN](#RETURN)

> [TYPE](#TYPE)

> [ORDER](#ORDER)

> [DEREFERENCE](#DEREFERENCE)

> [COMPARE](#COMPARE)

> [CLONE](#CLONE)

> [DEEPSEARCH](#DEEPSEARCH)

> [COMPARE](#COMPARE)

> [PASS](#PASS)

> [OVERRIDE](#OVERRIDE)

<h2 name = "FIND">FIND</h2>

 > First, Last, Idx, Key, Val, Card, Size, Height, Slice, All, Lambda

Many functions have "find functionality".  This means they A) have a FIND parameter called `findType` and B) have an interface parameter called `findData` whose dynamic type will vary based on the input for `findType`.  Given `stack.Get(findType FIND, findData any)`:

```
stack.Get(FIND_Last) // doesn't require `findData` to be passed, so pass nothing
stack.Get(FIND_Last, nil) // doesn't require `findData` to be passed, so pass nil
stack.Get(FIND_Key, "MyKey") // does require `findData` to be passed, so pass argument of the data type corresponding to FIND_Idx (in this case, a string)
stack.Get(FIND_Idx, 1) // get first card whose index is 1
stack.Get(FIND_Idx, []int {2, 7}) // get first card whose index is either 2 or 7
stack.Get(FIND_Idx, MakeStack([]int {2, 7})) // get first card whose index is either 2 or 7
stack.Get(FIND_Lambda, func(card *Card) (bool) {return card.Val.(int) % 2 == 0}) // (assuming all vals are ints) get all cards whose vals are multiples of 2
```

Sample:
 > enumerator type's name
 >> the accepted data type(s), one of which must be passed into `findData`
 >
 >> the condition in which the function "finds" a given card

Enumerators:
 > FIND_First
 >> `nil`
 >
 >> Idx is `0`
 >
 > FIND_Last
 >> `nil`
 >
 >> Idx is `parentStack.Size - 1`
 >
 > FIND_Idx
 >> `int` / `[]int` / `Stack{ints}`
 >
 >> each card whose index == at least one value in findData
 >
 > FIND_Key
 >> any
 >
 >> each card whose Key is findData
 >
 > FIND_Val
 >> any
 >
 >> each card whose Val is findData
 >
 > FIND_Card
 >> *Card
 >
 >> each card whose object address is equal to the object address of findData
 >
 > FIND_Size
 >> `int` / `[]int` / `Stack{ints}`
 >
 >> each card which holds a substack whose Size == at least one value in findData
 >
 > FIND_Height
 >> `int` / `[]int` / `Stack{ints}`
 >
 >> each card which holds a substack whose Height == at least one value in findData
 >
 > FIND_Slice
 >> `[]int {startIdx, endIdx}` / `Stack{startIdx, endIdx}`
 >
 >> each card whose index is inclusively between startIdx and endIdx
 >
 > FIND_All
 >> `nil`
 >
 >> each card
 >
 > FIND_Lambda
 >> if function named N:
 >>> `func(card *Card, parentStack *Stack, isSubstack bool, workingMemAdrs ...any) (returnVal bool)`
 >> 
 >>> where you can pass a function containing between 0 and all of these parameters (assuming no parameter's order is changed)
 >>
 >> if function named NMany, any of the following are acceptable:
 >>> `func(card *Card, parentStack *Stack, isSubstack bool, retStack *Stack, workingMemAdrs ...any) (returnVal bool)`
 >> 
 >>> where you can pass a function containing between 0 and all of these parameters (assuming no parameter's order is changed)
 >
 >> each card for whom the findData function returns true

<h2 name = "REPLACE">REPLACE</h2>

 > Key, Val, Card, Lambda

Many functions have "replace functionality".  This means they A) have a REPLACE parameter called `replaceType` and B) have an interface parameter called `replaceWith` whose dynamic type will vary based on the input for `replaceType`.  Given `stack.Replace(..., ..., ..., replaceType)`:

```
stack.ReplaceMany(FIND_All, nil, nil, REPLACE_Val, "NewVal") // replaces the val of all cards with "NewVal"
stack.ReplaceMany(FIND_All, nil, nil, REPLACE_Lambda, func (*Card card) {card.Val = card.Val.(int) * 3}) // (assuming each found card's val is an int) multiplies the val of every card by 3
```

Sample:
 > enumerator type's name
 >> the accepted data type(s), one of which must be passed into `replaceWith`
 >
 >> the target(s) being replaced by `replaceWith`

Enumerators:
 > REPLACE_Key
 >> `any`
 >
 >> each found card's Key
 >
 > REPLACE_Val
 >> `any`
 >
 >> each found card's Val
 >
 > REPLACE_Card
 >> `*Card`
 >
 >> each found card
 >
 > FIND_Lambda
 >>> `func(*Card card, *Stack parentStack, bool isSubstack, ...any workingMemAdrs)`
 >>
 >>> where you can pass a function containing between 0 and all of these parameters (assuming no parameter's order is changed)
 >
 >> there is no build-in target; the lambda function passed in is responsible for updating whatever target it would like responsibly

<h2 name = "RETURN">RETURN</h2>

 > Idxs, Keys, Vals, Cards, Stacks

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
 > RETURN_Stacks
 >> a stack whose vals are clones of the cards within a found card's substack

<h2 name = "TYPE">TYPE</h2>

 > Key, Val

A enumerator unique to the Unique() function which allows you to control whether you are uniquifying a card by keys or vals.

Enumerators:
 > UNIQUE_Key
 
 > UNIQUE_Val

<h2 name = "ORDER">ORDER</h2>

 > Before, After

A enumerator which allows you to control whether an inserted card or set of cards is added before or after a selection in a stack.

Enumerators:
 > ORDER_Before
 
 > ORDER_After

<h2 name = "DEREFERENCE">DEREFERENCE</h2>

 > False, True

A enumerator which, if set to true, will attempt to dereference the key/val in question before checking for equality.  For instance, if you stored the object address of an int in an array.  For instance:

```
intValA := gogenerics.MakeInterface(1)
intValB := gogenerics.MakeInterface(2)
intValC := gogenerics.MakeInterface(3)
card14 := MakeStack([]any {&intValA, &intValB, &intValC}).Get(FIND_Val, &intValB) // finds the first card that matches
card15 := MakeStack([]any {&intValA, &intValB, &intValC}).Get(FIND_Val, intValB, COMPARE_True, nil, nil, DEREFERENCE_True)
```

Enumerators:
 > DEREFERENCE_True
 
 > DEREFERENCE_False
 
 ---

 [> Return to glossary](../README.md)