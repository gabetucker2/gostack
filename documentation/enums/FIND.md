![Banner](../../images/gostack_SmallerTransparent.png)

<h2>FIND</h2>

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
 >>> `func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, workingMemAdrs ...any) (returnVal bool)`
 >> 
 >>> where you can pass a function containing between 0 and all of these parameters (assuming no parameter's order is changed)
 >>
 >> if function named NMany, any of the following are acceptable:
 >>> `func(card *Card, parentStack *Stack, isSubstack bool, coords *Stack, retStack *Stack, workingMemAdrs ...any) (returnVal bool)`
 >> 
 >>> where you can pass a function containing between 0 and all of these parameters (assuming no parameter's order is changed)
 >
 >> each card for whom the findData function returns true

 ---

 [> Return to enumerators](../enumsAPI.md)