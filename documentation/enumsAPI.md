![Banner](../images/gostack_Smaller.png)

<h1>Enumerator Documentation</h1>

* [FIND](#FIND)

<h2 name = "FIND">FIND</h2>

 > First, Last, Idx, Key, Val, Card, Size, Height, Slice, All, Lambda

Many functions have "find functionality".  This means they A) have a FIND parameter called `findType`, and B) have an interface parameter called `findData` whose dynamic type will vary based on the input for `findType`.  Given `stack.Get(findType FIND, findData any)`:

```
stack.Get(FIND_Last) // doesn't require `findData` to be passed, so pass nothing
stack.Get(FIND_Last, nil) // doesn't require `findData` to be passed, so pass nil
stack.Get(FIND_Key, "MyKey") // does require `findData` to be passed, so pass argument of the data type corresponding to FIND_Idx (in this case, a string)
stack.Get(FIND_Idx, 1) // get first card whose Idx is 1
stack.Get(FIND_Idx, []int {2, 7}) // get first card whose Idx is either 2 or 7
stack.Get(FIND_Idx, MakeStack([]int {2, 7})) // get first card whose Idx is either 2 or 7
```

Sample:
 > Enumerator type's name
 >> The accepted data type(s), one of which must be passed into `findData`
 >
 >> The condition in which the function "finds" a given card

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
 >> Each card whose Idx == at least one value in findData
 >
 > FIND_Key
 >> any
 >
 >> Each card whose Key is findData
 >
 > FIND_Val
 >> any
 >
 >> Each card whose Val is findData
 >
 > FIND_Card
 >> *Card
 >
 >> Each card whose object address is equal to the object address of findData
 >
 > FIND_Size
 >> `int` / `[]int` / `Stack{ints}`
 >
 >> Each card which holds a substack whose Size == at least one value in findData
 >
 > FIND_Height
 >> `int` / `[]int` / `Stack{ints}`
 >
 >> Each card which holds a substack whose Height == at least one value in findData
 >
 > FIND_Slice
 >> `[]int {startIdx, endIdx}` / `Stack{startIdx, endIdx}`
 >
 >> Each card whose Idx is inclusively between startIdx and endIdx
 >
 > FIND_All
 >> `nil`
 >
 >> Each card
 >
 > FIND_Lambda
 >> If function named N, any of the following are acceptable:
 >>> `func(*Card card, *Stack parentStack, bool isSubstack, workingMem ...any) (bool)`
 >>
 >>> `func(*Card card, *Stack parentStack, bool isSubstack) (bool)`
 >>
 >>> `func(*Card card, *Stack parentStack) (bool)`
 >>
 >>> `func(*Card card) (bool)`
 >>
 >>> `func() (bool)`
 >>
 >> If function named NMany, any of the following are acceptable:
 >>> `func(*Card card, *Stack parentStack, bool isSubstack, retStack *Stack, workingMem ...any) (bool)`
 >>
 >>> `func(*Card card, *Stack parentStack, bool isSubstack, retStack *Stack) (bool)`
 >>
 >>> `func(*Card card, *Stack parentStack, bool isSubstack) (bool)`
 >>
 >>> `func(*Card card, *Stack parentStack) (bool)`
 >>
 >>> `func(*Card card) (bool)`
 >>
 >>> `func() (bool)`
 >
 >> Each card for whom the findData function returns true
 
 ---

 [> Return to glossary](../README.md)