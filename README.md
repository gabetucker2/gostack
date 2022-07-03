 <h1 name = "preface">PREFACE</h1>

![Banner](images/gostack_Smaller.png)

 <h1 name = "introduction">Introduction</h1>

 Introducing **Stacks**—sets of **Card** elements (like a stack of cards)—***gostack*** serves as an all-in-one library for flexible, parsimonious, and elegant data management in *golang*.

 ***gostack***'s stacks...
 * ...replace maps, arrays, and matrices, expelling with the need for fetching or translating data between varying data structure types, all the while supporting smooth conversion between stacks and your existing data structures
 * ...offer the minimum functions needed for unlimited flexibility, allowing the user to seamlessly write what would previously have been a verbose monstrosity of 5 nested for-loops in a single line
 * ...allow the user to get and set based on reference or object with ease, preventing the user from having to worry about convoluted pointer/address management
 * ..., even when our built-in functions aren't enough, allow the user to effortlessly implement their own lambda functions to create novel stack mechanisms of their own design

 Is ***gostack*** really more efficient than ***classical go***?  To put this to the test, we created a race for the two; they each have to complete 3 data management tasks as quickly and efficiently as possible.  Whereas ***classical go*** took 45 lines to make it to the finish, ***gostack*** took roughly one fifth the amount of lines (merely 10)—[see for yourself!](/tutorials/race.md)

 To get a better feel of the library, feel free to take a look at some [examples](/tutorials/Bootstrap.go) of how ***gostack*** can substitute commonly-used functions.

<h1 name = "glossary">Glossary</h1>

 > [Files](#files)

 > [Preface](#preface)
 >> [Introduction](#introduction)
 >
 >> [Glossary](#glossary)
 >
 >> [File Explanations](#fileExplanations)
 >
 >> [Conventions](#conventions)

 > [Overview](#overview)
 >> [Brief Documentation](#briefDocumentation)
 >>> [Data Structures](#dataStructuresBrief)
 >>>> [structs](#structsBrief)
 >>>
 >>>> [enums](#enumsBrief)
 >>>
 >>> [Non-Generalized Functions](#nonGeneralizedFunctionsBrief)
 >>>
 >>> [Generalized Functions](#generalizedFunctionsBrief)
 >
 >> [Exhaustive Documentation](#exhaustiveDocumentation)
 >>> [Data Structures](#dataStructures)
 >>>> [structs](#structs)
 >>>>> [Stack](#stack)
 >>>>
 >>>>> [Card](#card)
 >>>>
 >>>> [enums](#enums)
 >>>>> [RETURN](#RETURN)
 >>>>
 >>>>> [FIND](#FIND)
 >>>>
 >>>>> [REPLACE](#REPLACE)
 >>>>
 >>>>> [TYPE](#TYPE)
 >>>>
 >>>>> [ORDER](#ORDER)
 >>>>
 >>>>> [MATCHBY](#MATCHBY)
 >>>>
 >>>>> [CLONE](#CLONE)
 >>>>
 >>>>> [DEEPSEARCH](#DEEPSEARCH)
 >>>
 >>> [Non-Generalized Functions](#nonGeneralizedFunctions)
 >>>> [MakeCard(...)](#MakeCard)
 >>>
 >>>> [MakeStack(...)](#MakeStack)
 >>>
 >>>> [MakeStackMatrix(...)](#MakeStackMatrix)
 >>>
 >>>> [stack.StripStackMatrix(...)](#StripStackMatrix)
 >>>
 >>>> [stack.ToArray()](#ToArray)
 >>>
 >>>> [stack.ToMap()](#ToMap)
 >>>
 >>>> [stack.ToMatrix()](#ToMatrix)
 >>>
 >>>> [stack.Empty()](#Empty)
 >>>
 >>>> [{card, stack}.Clone()](#Clone)
 >>>
 >>>> [stack.Unique(...)](#Unique)
 >>>
 >>>> [stack.Shuffle()](#Shuffle)
 >>>
 >>>> [stack.Flip()](#Flip)
 >>>
 >>>> [{card, stack}.Print()](#Print)
 >>>>
 >>>> [stack.Sort(...)](#Sort)
 >>>>
 >>>> [stack.Lambda(...)](#Lambda)
 >>>
 >>> [Generalized Functions](#generalizedFunctions)
 >>>> [stack.Add(...)](#Add)
 >>>
 >>>> [stack.Move(...)](#Move)
 >>>
 >>>> [stack.Has(...)](#Has)
 >>>
 >>>> [stack.Get(...)](#Get)
 >>>>
 >>>> [stack.GetMany(...)](#GetMany)
 >>>
 >>>> [stack.Replace(...)](#Replace)
 >>>>
 >>>> [stack.ReplaceMany(...)](#ReplaceMany)
 >>>
 >>>> [stack.Update(...)](#Update)
 >>>>
 >>>> [stack.UpdateMany(...)](#UpdateMany)
 >>>
 >>>> [stack.Extract(...)](#Extract)
 >>>>
 >>>> [stack.ExtractMany(...)](#ExtractMany)
 >>>
 >>>> [stack.Remove(...)](#Remove)
 >>>>
 >>>> [stack.RemoveMany(...)](#RemoveMany)
 >
 >> [Future Updates](#futureUpdates)
 >
 >> [Footer](#footer)

<h1 name = "fileExplanations">File Explanations</h1>

 > [gostack](#Files) [**.../gostack** package] any .go files that are direct children of this folder will be built to the **.../gostack** package
 >>
 >> [casetests](/casetests) [**.../gostack/casetests** package]
 >>> [CaseEnd.go](/casetests/CaseEnd.go) contains case tests for **library.go** functions
 >>
 >>> [Init.go](/casetests/Init.go) contains an empty function that prevents compiler errors when importing but not referencing this package
 >>
 >>> [testend.go](/casetests/testend.go) contains functions to implement **caseend.go** functions
 >>
 >>> [unaddedcases.txt](/casetests/unaddedcases.txt) is where obsolete data to be added into future case tests is stored, intended only for the developers
 >>
 >> [executive](/executive) [**main** package]
 >>> [executive.go](/executive.go) exists to call functions in this project, either for case testing or executing tutorials
 >>
 >> [images](/images)
 >>> **gostack_Smaller.png** is the banner image for this project
 >>
 >>> **packages.png** is a layout of the package dependencies/structure of this project
 >>
 >> [tutorials](/tutorials) [**.../gostack/tutorials** package]
 >>> [Bootstrap.go](/tutorials/Bootstrap.go) is a tutorial on how to implement some common functions using golang
 >>
 >>> [Init.go](/tutorials/Init.go) contains an empty function that prevents compiler errors when importing but not referencing this package
 >>
 >>> [Lambda.go](/tutorials/Lambda.go) is a tutorial on how to implement lambda functions
 >>
 >>> [Matrices.go](/tutorials/Matrices.go) is a tutorial on how to implement matrices in ***gostack***
 >>
 >>> [ObjectReference.go](/tutorials/ObjectReference.go) is a tutorial on how to use passing-by-object and -by-reference with respect to ***gostack*** functions
 >>
 >>> [race.md](/tutorials/race.md) showcases a race to complete the same set of tasks using ***classical go*** vs ***gostack***
 >>
 >>> [unaddedtutorials.txt](/tutorials/unaddedtutorials.txt) is where obsolete data to be added into future tutorials is stored, intended only for the developers
 >>
 >> [backend.go](/backend/backend.go) contains private functions to implement **Library.go**
 >>
 >> [DataStructures.go](/DataStructures.go) initializes structs and enums, as well as methods for conversion
 >>
 >> [go.mod](/go.mod) is to initialize the directories
 >>
 >> [frontend.go](/frontend.go)  contains public functions that the user will be calling
 >>
 >> [README.md](/README.md) is this file
 >>
 >> [TODO.txt](/TODO.txt) is a task list, intended only for the developers
 >>
 >> [unaddedgostack.txt](/unaddedgostack.txt) is where obsolete data to be added back into backend.go and Library.go

![Packages](images/packages.png)

<h1 name = "conventions">CONVENTIONS</h1>

 <h2>General</h2>

 Executing `go run executive/executive.go` in a terminal in the main directory, or executing `go run .` in the `executive` directory, will run whichever file(s) are being called by `executive.go`.

 **Generalized Functions** refer to functions that have a `findType` and `findData` parameter, meaning they perform a search through the stack for upon where to act.

 In ***gostack***, creating an array of cards is considered atrocious and immoral.  There is no functional support for passing arrays of cards as an argument.  Please only create a []\*Card array if it is a temporary variable to which you are assigning `stack.Cards`.  For instance, if you wanted to make your own `stack.Move(...)` function, you would create a temporary []\*Card variable, iteratively append that variable such that it "moves" whatever card(s) you want to move, then assign `stack.Cards` to that variable, never referencing the variable again.

 <h2>Naming</h2>

 * "FunctionName" functions are public functions, accessible to the user
 * "functionName" functions are private functions, hidden from the user
 * "FileName" files contain at least one public function (ideally which calls all private functions within that file)/struct, accessible to the user
 * "filename" files contain all private stuff, hidden from the user

 <h2>Design-By-Contract</h2>

 We use design-by-contract principles with JDoc annotations, as instructed by OSU's CSE department (http://web.cse.ohio-state.edu/software/2221/web-sw1/extras/slides/09.Design-by-Contract.pdf).

<h1 name = "overview">OVERVIEW</h1>

<h1 name = "briefDocumentation">Brief Documentation</h1>

<h2 name = "dataStructuresBrief">Data Structures</h2>

<h3 name = "structsBrief">Structs</h3>

 > **stack** *Stack*
 >> **Cards** *[]\*Card*
 >
 >> **Size** *int*

 > **card** *Card*
 >> **Idx** *int*
 >
 >> **Key** *any*
 >
 >> **Val** *any*

<h3 name = "enumsBrief">Enums</h3>

 > **RETURN**
 > * _RETURN_NotationSample *type that's returned*
 > * RETURN_Idxs *stack of ints*
 > * RETURN_Keys *stack of anys*
 > * RETURN_Vals *stack of anys*
 > * RETURN_Cards *stack of Cards*

 > **FIND**
 > * _FIND_NotationSample *findData argument type*
 > * FIND_First *NONE*
 > * FIND_Last *NONE*
 > * FIND_Idx *int*
 > * FIND_Idxs *ints*
 > * FIND_IdxsStack *stack whose vals are ints*
 > * FIND_Key *any*
 > * FIND_Keys *anys*
 > * FIND_KeysStack *stack whose vals are keys*
 > * FIND_Val *any*
 > * FIND_Vals *anys*
 > * FIND_ValsStack *stack whose vals are vals*
 > * FIND_Card *Card*
 > * FIND_Cards *Stack*
 > * FIND_CardsStack *stack whose vals are cards*
 > * FIND_Slice *[2]int*
 > * FIND_All *NONE*
 > * FIND_Lambda *lambda function*

 > **REPLACE**
 > * _REPLACE_NotationSample *setByData argument type*
 > * REPLACE_Key *interface{}*
 > * REPLACE_Val *interface{}*
 > * REPLACE_Card *Card*
 > * REPLACE_Stack *Stack*
 > * REPLACE_Lambda *lambda function*

 > **TYPE**
 > * TYPE_Key
 > * TYPE_Val
 > * TYPE_Card

 > **ORDER**
 > * ORDER_Before
 > * ORDER_After

 > **MATCHBY**
 > * MATCHBY_Object
 > * MATCHBY_Reference

 > **CLONE**
 > * CLONE_True
 > * CLONE_False

 > **DEEPSEARCH**
 > * DEEPSEARCH_True
 > * DEEPSEARCH_False

<h2 name = "nonGeneralizedFunctionsBrief">Non-Generalized Functions</h2>

 * **MakeCard(...idx, ...key, ...val)**
 * **MakeStack(...input1, ...input2, ...repeats)**
 * **MakeStackMatrix(...input1, ...input2, ...matrixShape)**
 * **stack.StripStackMatrix(target1, target2, ..., targetN)**
 * **stack.ToArray()**
 * **stack.ToMap()**
 * **stack.ToMatrix(...depth)**
 * **stack.Empty()**
 * **{card, stack}.Clone()**
 * **stack.Unique(typeType, ...matchByType)**
 * **stack.Shuffle()**
 * **stack.Flip()**
 * **{card, stack}.Print()**
 * **stack.Sort(lambda sort function)**
 * **stack.Lambda(lambda function)**

<h2 name = "generalizedFunctionsBrief">Generalized Functions</h2>

 * **stack.Add(insert, ...orderType, ...findType, ...findData, ...matchByType)**
 * **stack.Move(findType_from, orderType, findType_to, ...findData_from, ...findData_to, ...matchByType_from, ...matchByType_to)**
 * **stack.Has(returnType, findType, ...findData, ...matchByType)**
 * **stack.Get(...findType, ...findData, ...matchByType, ...clonesType_card, ...clonesType_keys, ...clonesType_vals)**
 * **stack.GetMany(findType, ...findData, ...matchByType, ...returnType, ...clonesType, ...clonesType_keys, ...clonesType_vals)**
 * **stack.Replace(replaceType, replaceData, findType, ...findData, ...matchByType)**
 * **stack.ReplaceMany(replaceType, replaceData, findType, ...findData, ...matchByType, ...returnType)**
 * **stack.Update(findType, ...findData, ...matchByType)**
 * **stack.UpdateMany(findType, ...findData, ...matchByType)**
 * **stack.Extract(findType, ...findData, ...matchByType)**
 * **stack.ExtractMany(findType, ...findData, ...matchByType, ...returnType)**
 * **stack.Remove(findType, ...findData, ...matchByType)**
 * **stack.RemoveMany(findType, ...findData, ...matchByType)**

<h1 name = "exhaustiveDocumentation">Exhaustive Documentation</h1>

<h2 name = "dataStructures">Data Structures</h2>

<h3 name = "structs">structs</h3>

<h4 name = "stack">Stack</h4>

 This is the main struct in the project.

 > `stack` *Stack{}*
 >> `Cards` *[]\*Card{}*
 >>> Returns an interface array to represent the elements in the Stack
 >> `Size` *int*
 >>> Returns the cardinality (i.e., `len(stack.cards)`) of this Stack

<h4 name = "card">Card</h4>

 This is a struct for our elements/maps within stacks.

 >> `Card` *Card{}*
 >>> `card.Idx` *int*
 >>>> The index of this card
 >>>
 >>> `card.Key` *any (interface{})*
 >>>> The key of this card (or nil if doesn't exist)
 >>>
 >>> `card.Val` *any (interface{})*
 >>>> The val of this card (or nil if doesn't exist)

<h3 name = "enums">enums</h3>

<h4 name = "RETURN">RETURN</h4>

 This is an enum intended to make it easy for the user to control the output type from a getter function.

 > ***RETURN***
 >> *_RETURN_NotationSample*
 >>> *If you input RETURN_Keys to stack.GetMany(), then you will get a new stack of cards whose values are the keys of the initial stack.*
 >>
 >> RETURN_Idxs
 >>> stack of ints
 >>
 >> RETURN_Keys
 >>> stack of interface{}s
 >>
 >> RETURN_Vals
 >>> stack of interface{}s
 >>
 >> RETURN_Cards
 >>>> default
 >>>
 >>> stack of Cards

<h4 name = "FIND">FIND</h4>

 This is an enum intended to make it easy to flexibly inform functions what the intended target is.

 > ***FIND***
 >> *_FIND_NotationSample*
 >>> *The type of the variable (called `data`) that needs to be passed into the function utilizing this constant*
 >>
 >>> *For instance, if you input `FIND_Keys`, you would need to pass a Stack whose values are the keys you want to find to your `data` parameter*
 >>
 >> FIND_First
 >>>> default
 >>>
 >>> *NONE*
 >>
 >> FIND_Last
 >>> *NONE*
 >>
 >> FIND_Idx
 >>> int
 >>
 >> FIND_Idxs
 >>> []int
 >>
 >> FIND_IdxsStack
 >>> stack where type{stack.card.Val} == int
 >>
 >> FIND_Key
 >>> any (interface{})
 >>
 >> FIND_Keys
 >>> []any ([]interface{})
 >>
 >> FIND_KeysStack
 >>> stack where type{stack.card.Val} == any (interface{})
 >>
 >> FIND_Val
 >>> any (interface{})
 >>
 >> FIND_Vals
 >>> []any ([]interface{})
 >>
 >> FIND_ValsStack
 >>> stack where type{stack.card.Val} == any (interface{})
 >>
 >> FIND_Card
 >>> *Card
 >>
 >> FIND_Cards
 >>> *Stack (input is the cards in this stack)
 >>
 >> FIND_CardsStack
 >>> stack where type{stack.card.Val} == *Card
 >>
 >> FIND_Slice
 >>> [2]int
 >>>> [startIndex, endIndex]
 >>>>
 >>>> *(if endIndex is lower than startIndex, the find will transpire in reverse order)*
 >>
 >> FIND_All
 >>> *NONE*
 >>
 >> FIND_Lambda
 >>> *lambda function*

<h4 name = "REPLACE">REPLACE</h4>

 This is an enum intended to make it easy to flexibly decide what part of an object to update.  For instance, you could replace a card with a card, replace a card's value with a value, etc, so you are not just limited to replacing cards.

 > ***REPLACE***
 >> *_REPLACE_NotationSample*
 >>> *The type of the variable (called `data`) that needs to be passed into the function utilizing this constant*
 >>
 >> REPLACE_Key
 >>> *interface{}*
 >>
 >> REPLACE_Val
 >>> *interface{}*
 >>
 >> REPLACE_Card
 >>> *Card*
 >>
 >> REPLACE_Stack
 >>> *Stack*
 >>>> replaces a card with each card in a stack of cards
 >>
 >> REPLACE_Lambda
 >>> interface{} *lambda function*
 
<h4 name = "TYPE">TYPE</h4>

 This is an enum intended to make it easy to tell certain functions the type of value being targeted.

> ***TYPE***
>> TYPE_Key
>
>> TYPE_Val
>
>> TYPE_Card
 
<h4 name = "ORDER">ORDER</h4>

 This is an enum intended to make it easy to tell certain functions whether to insert a value before or after the input index.

> ***ORDER***
>> ORDER_Before
>>> default
>>
>> ORDER_After

<h4 name = "MATCHBY">MATCHBY</h4>

 This is an enum intended to make it easy to target whether a function searching for a match between input data and data in the stack element is matching by having the same values (MATCHBY_Object) or the same memory address (MATCHBY_Reference).

 Matching by reference only works for Val, Key, and Card FIND types.  It would not make much sense to match an index that's managed on the backend by reference (FIND_Idx), to match a lambda expression (FIND_Lambda), or to match using a position that's not even comparing values (FIND_First, FIND_Last, FIND_All).

 Take care to note that all cases where objects are matching by reference will also be matching by object.

 > ***MATCHBY***
 >> MATCHBY_Object
 >>> default
 >>
 >> MATCHBY_Reference

<h4 name = "CLONE">CLONE</h4>

 This is an enum intended to make it possible to tell the function whether to return a clone of an object or a pointer of an object.

 > ***MATCHBY***
 >> CLONE_True
 >>
 >> CLONE_False
 >>> default

<h4 name = "DEEPSEARCH">DEEPSEARCH</h4>

 This is an enum intended to make it easy to call functions to perform on certain elements in matrices.

 > ***DEEPSEARCH***
 >> DEEPSEARCH_True
 >>
 >> DEEPSEARCH_False
 >>> default

<h2 name = "nonGeneralizedFunctions">Non-Generalized Functions</h2>

<h3 name = "MakeCard">MakeCard</h3>

 > `gostack.MakeCard(...val, ...key, ...idx)`
 ```
 Makes a card with inputted vals and keys

 @param optional `val` type{any} default nil
 @param optional `key` type{any} default nil
 @param optional `idx` type{int} default -1 no pass-by-reference
 @returns type{*Card} the newly-constructed card
 @constructs type{*Card} a newly-constructed card
 @ensures the new card will have val `val`, key `key`, and idx `idx`
 ```

<h3 name = "MakeStack">MakeStack</h3>

 `gostack.MakeStack(...input1, ...input2, ...repeats)`
 ```
 Creates a stack of cards with optional starting cards
 
 @param optional `input1` type{[]any, map[any]any} default nil
 @param optional `input2` type{[]any} default nil
 @param optional `repeats` type{int} default 1
 @returns type{*Stack} the newly-constructed stack of newly-constructed cards
 @constructs type{*Stack} a newly-constructed stack of newly-constructed cards
 @requires
  * `input1` is map and nil `input2`
      OR `input1` is an array and nil `input2`
	  OR `input1` is an array and `input2` is an array
	  OR `input1` is nil and `input2` is an array
  * IF `input1` and `input2` are both passed as arguments
      |`input1`| == |`input2`|
  * `MakeCard()` has been implemented
 @ensures
  * repeats the function's filling `repeats` (or, if nil or under 0, 1) amount of times
  * IF `input1` is passed
	    IF `input1` is a map
        unpack the map into new cards with corresponding keys and vals
      ELSEIF `input1` is an array and `input2` is not passed/nil
        unpack values from `input1` into new cards
      ELSEIF `input1` is an array and `input2` is an array
	      unpack keys from `input1` and values from `input2` into new cards
      ELSEIF `input1` is nil and `input2` is an array
	      unpack keys from `input2` into new cards
	  ELSE
	    the stack is empty
 ```

<h3 name = "MakeStackMatrix">MakeStackMatrix</h3>

 `MakeStackMatrix(...input1, ...input2, ...matrixShape)`
 ```
 Creates a new stack-within-stack-structured stack
 
 @param optional `input1` type{interface{}} default nil
 @param optional `input2` type{interface{}} default nil
 @param optional `matrixShape` type{[]int} default nil
  * an int array representing the shape of the matrix
  * the first int is the largest container
  * the last int is the container directly containing the inputted cards
 @requires
  * `MakeCard()` has been implemented
  * IF `input1` and `input2` are both passed as arguments
      |`input1`| == |`input2`|
 @ensures
  * IF no `matrixShape` is passed
      treating `input1`/`input2` as matrices/a map of matrices:
      IF `input1` is passed
        IF `input1` is a map
          unpack the map into matrix of shape `inputx` with corresponding keys and vals
        ELSEIF `input1` is an array and `input2` is not passed/nil
          unpack values from `input1` into matrix of shape `inputx`
        ELSEIF `input1` is an array and `input2` is an array
          unpack keys from `input1` and values from `input2` into matrix of shape `inputx`
        ELSEIF `input1` is nil and `input2` is an array
          unpack keys from `input2` into matrix of shape `inputx` 
      ELSEIF `input1` is not passed
        the stack is empty
	ELSEIF `matrixShape` is passed
	  treating `input1`/`input2` as 1D arrays:
	  IF `input1` is passed
        IF `input1` is a map
          unpack the map into matrix of shape `matrixShape` with corresponding keys and vals
        ELSEIF `input1` is an array and `input2` is not passed/nil
          unpack values from `input1` into matrix of shape `matrixShape`
        ELSEIF `input1` is an array and `input2` is an array
          unpack keys from `input1` and values from `input2` into matrix of shape `matrixShape`
        ELSEIF `input1` is nil and `input2` is an array
          unpack keys from `input2` into matrix of shape `matrixShape`
	  ELSEIF `input1` is not passed
	    create a StackMatrix of shape `matrixShape` whose deepest card vals are nil
 ```
 
<h3 name = "ToArray">ToArray</h3>
 
 `stack.ToArray()`
 ```
 Creates a new interface array from values of `stack`

 @receiver `stack` type{*Stack}
 @returns type{[]interface{}} new array
 @ensures new array values correspond to `stack` values
 ```
 
<h3 name = "ToMap">ToMap</h3>
 
 `stack.ToMap()`
 ```
 Creates a new interface-interface map from values of `stack`

 @receiver `stack` type{*Stack}
 @returns type{map[interface{}]interface{}} new map
 @ensures new map keys and values correspond to `stack` keys and values
 ```
 
<h3 name = "ToMatrix">ToMatrix</h3>
 
 `stack.ToMatrix(...depth)`
 ```
 Creates a new matrix from a stack by depth.  For instance, if depth = 2, then returns the stacks inside stack as an [][]interface{}

 @receiver `stack` type{*Stack}
 @param optional `depth` type{int} default -1
 @returns type{[]interface}
 @ensures
  * -1 depth means it will go as deep as it can
  * new map keys and values correspond to `stack` keys and values
  * example: Stack{Stack{"Hi"}, Stack{"Hello", "Hola"}, "Hey"} =>
      []interface{}{[]interface{}{"Hi"}, []interface{}{"Hola", "Hello"}, "Hey"}
 ```
 
<h3 name = "Empty">Empty</h3>
 
 `stack.Empty()`
 ```
 Makes a card with inputted vals and keys

 @receiver `stack` type{*Stack}
 @returns `stack`
 @updates `stack.Cards` to be empty
 ```

<h3 name = "Clone">Clone</h3>

 `stack.Clone()`
 ```
 Returns a clone of the given stack

 @receiver `stack` type{*Stack}
 @returns type{*Stack} stack clone
 @constructs type{*Stack} clone of `stack`
 @ensures the stack clone has the same card pointers as `stack`
 ```

 `card.Clone()`
 ```
 Returns a clone of the given card

 @receiver `card` type{*Card}
 @returns type{*Card} card clone
 @constructs clone of `card`
 ```
 
<h3 name = "Unique">Unique</h3>
 
 `stack.Unique(typeType, ...matchByType)`
 ```
 Removes all cards from `stack` which share the same field value as another card before

 @receiver `stack` type{*Stack}
 @param `typeType` type{TYPE}
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates `stack` to have no repeating values between field `typeType`
 ```
 
<h3 name = "Shuffle">Shuffle</h3>
 
 `stack.Shuffle()`
 ```
 Shuffles the order of `stack` cards

 @receiver `stack` type{*Stack}
 @returns `stack`
 @updates
  * `stack` card ordering is randomized
  * rand.Seed is updated to time.Now().UnixNano()
 ```
 
<h3 name = "Flip">Flip</h3>
 
 `stack.Flip()`
 ```
 Flips the ordering of `stack.Cards`
 
 @receiver `stack` type{*Stack}
 @returns `stack`
 @updates `stack` to have its ordering reversed
 ```
 
<h3 name = "Print">Print</h3>
 
 `card.Print()`
 ```
 Prints information regarding `card` to the console
 
 @receiver `card` type{*Card}
 @updates terminal logs
 ```
 
 `stack.Print()`
 ```
 Prints information regarding `stack` to the console
 
 @receiver `stack` type{*Stack}
 @updates terminal logs
 @requires card.Print() has been implemented
 ```
 
<h3 name = "Sort">Sort</h3>
 
 `stack.Sort(lambda sort function)`
 ```
 Order the cards contingent on some attribute they contain
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(*Card, *Stack, ...interface{}) (ORDER, int)}
 @requires
  * `lambda` returns the order (before/after) and index to which to move your card in the stack
  * `lambda` does not update `stack` itself
 @ensures each card in `stack` is passed into your lambda function
 ```
 
<h3 name = "Lambda">Lambda</h3>
 
 `stack.Lambda(lambda function)`
 ```
 Iterate through a stack calling your lambda function on each card
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(*Card, ...interface{})}
 @ensures
  * Each card in `stack` is passed into your lambda function
  * `stack` is the first argument passed into your variadic parameter on the first call
 ```

<h2 name = "generalizedFunctions">Generalized Functions</h2>
 
<h3 name = "Add">Add</h3>
 
 `stack.Add(insert, ...orderType, ...findType, ...findData, ...matchByType)`
 ```
 Adds to a stack of cards or a cards at (each) position(s) 
 
 @receiver `stack` type{*Stack}
 @param `insert` type{Card, Stack}
 @param optional `orderType` type{ORDER} default ORDER_Before
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates `stack.Cards` to have new cards before/after each designated position
 ```
 
<h3 name = "Move">Move</h3>
 
 `stack.Move(findType_from, orderType, findType_to, ...findData_from, ...findData_to, ...matchByType_from, ...matchByType_to)`
 ```
 Moves one element or slice of cards to before or after another element or slice of cards
 
 @receiver `stack` type{*Stack}
 @param `findType_from` type{FIND}
 @param `orderType` type{ORDER}
 @param `findType_to` type{FIND}
 @param optional `findData_from` type{interface{}} default nil
 @param optional `findData_to` type{interface{}} default nil
 @param optional `matchByType_from` type{MATCHBY} default MATCHBY_Object
 @param optional `matchByType_to` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @ensures IF `findType_to` or `findType_from` get over one position, method doesn't perform move and prints invalid argument (FIND_Slice is the sole exception to this rule)
 ```
 
<h3 name = "Has">Has</h3>
 
 `stack.Has(...findType, ...findData, ...matchByType)`
 ```
 Returns a boolean representing whether a search exists in the stack

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns true IF successful search, false IF unsuccessful search
 @requires `stack.Get()` has been implemented
 ```
 
<h3 name = "Get">Get</h3>
 
 `stack.Get(...findType, ...findData, ...matchByType, ...clonesType_card, ...clonesType_keys, ...clonesType_vals)`
 ```
 Gets a card from specified parameters in a stack, or nil if does not exist

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `clonesType_card` type{CLONES} default CLONE_FALSE
 @param optional `clonesType_keys` type{CLONES} default CLONE_FALSE
 @param optional `clonesType_vals` type{CLONES} default CLONE_FALSE
 @returns type{*Card} the found card OR nil if invalid FIND
 @ensures
  * CLONE_True for `clonesType_card` means the returned card object itself is a clone
  * CLONE_True for `clonesType_key` means the returned card key is a clone
  * CLONE_True for `clonesType_val` means the returned card val is a clone
 ```
 
<h3 name = "GetMany">GetMany</h3>
 
 `stack.GetMany(findType, ...findData, ...matchByType, ...returnType, ...clonesType, ...clonesType_keys, ...clonesType_vals)`
 ```
 Gets a stack from specified parameters in a stack
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `clonesType` type{CLONES} default CLONE_FALSE
 @param optional `clonesType_keys` type{CLONES} default CLONE_FALSE
 @param optional `clonesType_vals` type{CLONES} default CLONE_FALSE
 @returns type{*Stack} the new stack
 @constructs type{*Stack} new stack of specified values from specified cards in `stack`
 @requires
  * `MakeStack()` has been implemented
  * `clonesType_keys` and `clonesType_vals` are only passed if `returnType` == RETURN_Cards
 @ensures
  * CLONE_True means the vals of cards in the returned stack are not the original object that was gotten
  * CLONE_True for `clonesType_keys` means the cards in the returned stack keys are clones
  * CLONE_True for `clonesType_vals` means the cards in the returned stack vals are clones
 ```
 
<h3 name = "Replace">Replace</h3>
 
 `stack.Replace(replaceType, replaceData, findType, ...findData, ...matchByType)`
 ```
 Returns a clone of a found card before its respective field is updated to `replaceData` (OR nil if not found)
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{interface{}}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns type{*Card} a clone of extracted card OR nil if found no cards
 @updates first found card to `replaceData`
 @requires `stack.Get()` has been implemented
 @ensures if `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 ```
 
<h3 name = "ReplaceMany">ReplaceMany</h3>
 
 `stack.Replace(replaceType, replaceData, findType, ...findData, ...matchByType)`
 ```
  Returns a stack whose values are clones of the original fields updated to `replaceData`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{interface{}}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @returns type{*Stack} a stack whose values are the extracted cards pre-update
 @updates all found cards to `replaceData`
 @requires `stack.GetMany()` has been implemented
 @ensures if `replaceData` is nil and `replaceType is REPLACE_Card`, the cards found will be removed from `stack`
 ```
 
<h3 name = "Update">Update</h3>
 
 `stack.Update(findType, ...findData, ...matchByType)`
 ```
 Updates a card in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates the found card in `stack`
 @requires `stack.Replace()` has been implemented
 ```
 
<h3 name = "UpdateMany">UpdateMany</h3>
 
 `stack.UpdateMany(findType, ...findData, ...matchByType)`
 ```
 Updates cards in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates  the found cards in `stack`
 @requires `stack.ReplaceMany()` has been implemented
 ```
 
<h3 name = "Extract">Extract</h3>
 
 `stack.Extract(findType, ...findData, ...matchByType)`
 ```
 Gets and removes a card from `stack`, or returns nil if it does not exist
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns type{*Card} the extracted card OR nil if invalid FIND
 @updates `stack` to no longer have found card
 ```
 
<h3 name = "ExtractMany">ExtractMany</h3>
 
 `stack.ExtractMany(findType, ...findData, ...matchByType, ...returnType)`
 ```
 Gets and removes a set of cards from `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @returns type{*Stack} the extracted card OR nil if invalid FIND
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 ```
 
<h3 name = "Remove">Remove</h3>
 
 `stack.Remove(findType, ...findData, ...matchByType)`
 ```
 Removes a card from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates `stack` to no longer have found card
 @requires `stack.Replace()` has been implemented
 ```
 
<h3 name = "RemoveMany">RemoveMany</h3>
 
 `stack.RemoveMany(findType, ...findData, ...matchByType)`
 ```
 Removes a set of cards from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{interface{}} default nil
 @param optional `matchByType` type{MATCHBY} default MATCHBY_Object
 @returns `stack`
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 ```
 
<h2 name = "futureUpdates">Future Updates</h2>

 *To be added later*

<h2 name = "footer">Footer</h1>

This library was created by Gabe Tucker and Andy Chen.

If you have any suggestions, questions, or comments you would like to make in respect to this project, please email `tucker.854@osu.edu`.  I appreciate any feedback and will usually respond within 1-2 business days.

Feel free to visit my personal pages at `https://gabetucker.com` or `https://www.linkedin.com/in/gabetucker2/`.

[Return to Glossary](#glossary)
