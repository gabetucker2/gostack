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

 > None, Both, Found, This

Assuming FIND is FIND_Key or FIND_Val, this enumerator which decides whether to dereference the key/val of the found cards data or the input data before checking for equality.  For instance, given `stack.Get(..., ..., ..., ..., ..., ..., dereferenceType DEREFERENCE)`:

```
init1 := 1
intValA = &init1
init2 := 2
intValB = &init2
init3 := 3
intValC = &init3
init4 := 2
intValD = &init4

mainStack := MakeStack([]any {intValA, intValB, intValC})

mainStack.Clone().Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_None) // gets intValB since intValB == intValB
mainStack.Clone().Get(FIND_Val, intValD, nil, nil, nil, DEREFERENCE_None) // doesn't get intValB since intValB != intValD

mainStack.Clone().Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_Both) // gets intValB since 2 == 2
mainStack.Clone().Get(FIND_Val, intValD, nil, nil, nil, DEREFERENCE_Both) // gets intValB since 2 == 2

mainStack.Clone().Get(FIND_Val, 2, nil, nil, nil, DEREFERENCE_Found) // gets intValB since *intValB == 2

MakeStack([]any {1, 2, 3}).Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_This) // gets 2 since 2 == *intValB
```

Enumerators:
 > DEREFERENCE_None
 
 > DEREFERENCE_Both

 > DEREFERENCE_Found

 > DEREFERENCE_This

<h2 name = "CLONE">CLONE</h2>

 > True, False

Provides information to a function on whether to clone the key/val of a substack or card.  For example, given `stack.Clone(..., ..., cloneCardKeys, cloneCardVals, cloneSubstackKeys, cloneSubstackVals CLONE)`:

```
myStack.Clone(nil, nil, CLONE_True, CLONE_False, CLONE_True, CLONE_True)
// returns a clone of myStack such that each val of a card in the new stack contains the same object val as the original card, but each card's key is a clone of the original key and each substacks key or val is a clone of the original key or val
```

Enumerators:
 > CLONE_True
 
 > CLONE_False

<h2 name = "DEEPSEARCH">DEEPSEARCH</h2>

 > True, False

Many functions have "deep search functionality".  This means they A) have a DEEPSEARCH parameter called `deepSearchType` and B) have an int/[]int/Stack{ints} parameter called `depth` which will act as a guide for which layers to consider in the deep search.  Sample:

 > enumerator type's name
 >> how the function will behave in its search

Enumerators:
 > DEEPSEARCH_True
 >> the function will listen to the `depth` parameter input in considering how deep it searches

 > DEEPSEARCH_False
 >> the function will set `depth` to 1, only considering the immediate children of this stack

<h2 name = "COMPARE">COMPARE</h2>

 > True, False

An enumerator which allows you to configure whether or not you compare certain aspects of a card in a `.Equals()` equality test.

Enumerators:
 > COMPARE_True

 > COMPARE_False

<h2 name = "PASS">PASS</h2>

 > True, False

An enumerator which allows you to configure whether or not you pass a certain type of element (either a card or a substack) into a search.  This allows you to filter out substacks without needing to use an if-statement inside of lambda logic.

Enumerators:
 > PASS_True

 > PASS_True

<h2 name = "OVERRIDE">OVERRIDE</h2>

 > True, False

An enumerator which allows you to configure whether or not you "override" a certain default argument interpretation.

For instance, if you pass a card as your `insert` argument into the `stack.Add()` function, then the function will automatically add that card to your stack.  If you would like to "override" this parameter and instead have the function add a card to your stack whose val is the `insert` card (i.e., add a card pointing to your card), then you would set override to true.

Enumerators:
 > OVERRIDE_True

 > OVERRIDE_False
 
 ---

 [> Return to glossary](../README.md)