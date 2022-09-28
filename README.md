 <h1 name = "preface">PREFACE</h1>

![Banner](images/gostack_Smaller.png)

 *"The purpose of abstraction is not to be vague, but to create a new semantic level in which one can be absolutely precise." - Edsger W. Dijkstra*

 <h1 name = "introduction">Introduction</h1>

 Introducing **Stacks**—sets of **Card** elements (like a stack of cards)—***gostack*** serves as an all-in-one library for flexible, parsimonious, and elegant scripting in *golang*.

 ***gostack***'s stacks...
 * ...replace arrays, maps, and matrices, eliminating the need for translating data between varying data types while supporting smooth conversion between stacks and your existing data structures
 * ...offer the minimum functions needed for unlimited flexibility, allowing the user to seamlessly write what would previously have been a verbose monstrosity of 5 nested for-loops in a single, yet concise, line
 * ...allow the user to get and set based on reference or object with ease, preventing the user from having to worry about convoluted pointer/address management
 * ...support the treatment of stacks as matrices, allowing the user to easily manage data tables, perform linear algebra, and control deep-versus-shallow matrix operations
 * ..., even when our built-in functions aren't enough, allow the user to effortlessly implement their own lambda functions to create novel stack mechanisms of their own design

 Is ***gostack*** really more efficient than ***classical go***?  To put this to the test, we created a race for the two; they each have to complete 3 data management tasks as quickly and efficiently as possible.  Whereas ***classical go*** took 61 lines to make it to the finish, ***gostack*** took merely 9—[see for yourself!](/tutorials/race.md)

 To get a better feel of the library, feel free to take a look at some [examples](/tutorials/Bootstrap.go) of how ***gostack*** can substitute commonly-used functions.  Alternatively, take a look at our beginner-friendly [introductory tutorial](/tutorials/Introduction.go)!

<h1 name = "glossary">Glossary</h1>

 > [Files](#files)

 > [Preface](#preface)
 >> [Introduction](#introduction)
 >
 >> [Glossary](#glossary)
 >
 >> [File Explanations](#fileExplanations)
 >
 >> [Package Schema](#packageSchema)
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
 >>>>> [POINTER](#POINTER)
 >>>>
 >>>>> [CLONE](#CLONE)
 >>>>
 >>>>> [DEEPSEARCH](#DEEPSEARCH)
 >>>>
 >>>>> [COMPARE](#COMPARE)
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
 >>>> [stack.ToMatrix(...)](#ToMatrix)
 >>>
 >>>> [card.ToPointer(...)](#ToPointer)
 >>>>
 >>>> [stack.ToPointer(...)](#ToPointer)
 >>>
 >>>> [card.ToObject(...)](#ToObject)
 >>>>
 >>>> [stack.ToObject(...)](#ToObject)
 >>>
 >>>> [stack.IsRegular()](#IsRegular)
 >>>
 >>>> [stack.Duplicate(...)](#Duplicate)
 >>>
 >>>> [stack.Empty()](#Empty)
 >>>
 >>>> [card.Clone(...)](#Clone)
 >>>>
 >>>> [stack.Clone(...)](#Clone)
 >>>
 >>>> [stack.Unique(...)](#Unique)
 >>>
 >>>> [card.Equals(...)](#Equals)
 >>>>
 >>>> [stack.Equals(...)](#Equals)
 >>>
 >>>> [stack.Shuffle()](#Shuffle)
 >>>
 >>>> [stack.Inverse()](#Inverse)
 >>>
 >>>> [card.Print()](#Print)
 >>>>
 >>>> [stack.Print()](#Print)
 >>>
 >>>> [stack.Lambda(...)](#Lambda)
 >>>
 >>> [Generalized Functions](#generalizedFunctions)
 >>>> [stack.Add(...)](#Add)
 >>>
 >>>> [stack.Move(...)](#Move)
 >>>
 >>>> [stack.Swap(...)](#Swap)
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
 >>> [Introduction.go](/tutorials/Introduction.go) contains a tutorial for absolute beginners to ***gostack***
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
 >> [unaddedgostack.txt](/unaddedgostack.txt) is where obsolete data to be added back into backend.go and Library.go

<h1 name = "packageSchema">Package Schema</h1>

![Packages](images/packages.png)

<h1 name = "conventions">CONVENTIONS</h1>

 <h2>General</h2>

 Executing `go run executive/executive.go` in a terminal in the main directory, or executing `go run .` in the `executive` directory, will run whichever file(s) are being called by `executive.go`.

 **Generalized Functions** refer to functions that have a `findType` and `findData` parameter, meaning they perform a search through the stack for upon where to act.

 In ***gostack***, creating an array of cards is considered atrocious and immoral.  There is no functional support for passing arrays of cards as an argument.  Please only create a []\*Card array if it is a temporary variable to which you are assigning `stack.Cards`.  For instance, if you wanted to make your own `stack.Move(...)` function, you would create a temporary []\*Card variable, iteratively append that variable such that it "moves" whatever card(s) you want to move, then assign `stack.Cards` to that variable, never referencing the variable again.

 Never insert the same card twice into the same stack.  Instead, insert the same value into two different cards.  If you insert two of the same card into the same stack, then the index property will become conflated between the two cards and functions will yield bugs while iterating through stacks.

 While you are technically permitted to have a substack within a card.Key, gostack functions always assume substacks are within card.Val's.  If you choose to do this, you are on your own.

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
 >
 >> **Depth** *int*

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
 > * FIND_Idxs Stack *stack whose vals are ints*
 > * FIND_Key *any*
 > * FIND_Keys *anys*
 > * FIND_Keys Stack *stack whose vals are keys*
 > * FIND_Val *any*
 > * FIND_Vals *anys*
 > * FIND_Vals Stack *stack whose vals are vals*
 > * FIND_Card *Card*
 > * FIND_Cards *Stack*
 > * FIND_Cards Stack *stack whose vals are cards*
 > * FIND_Slice *[2]int*
 > * FIND_All *NONE*
 > * FIND_Lambda *lambda function*

 > **REPLACE**
 > * _REPLACE_NotationSample *setByData argument type*
 > * REPLACE_Key *any*
 > * REPLACE_Val *any*
 > * REPLACE_Card *Card*
 > * REPLACE_Stack *Stack*
 > * REPLACE_Lambda *lambda function*

 > **TYPE**
 > * TYPE_Key
 > * TYPE_Val

 > **ORDER**
 > * ORDER_Before
 > * ORDER_After

 > **POINTER**
 > * POINTER_False
 > * POINTER_True

 > **CLONE**
 > * CLONE_True
 > * CLONE_False

 > **DEEPSEARCH**
 > * DEEPSEARCH_True
 > * DEEPSEARCH_False

 > **COMPARE**
 > * COMPARE_True
 > * COMPARE_False

<h2 name = "nonGeneralizedFunctionsBrief">Non-Generalized Functions</h2>

 * **MakeCard(...idx, ...key, ...val)**
 * **MakeStack(...input1, ...input2, ...repeats)**
 * **MakeStackMatrix(...input1, ...input2, ...matrixShape)**
 * **stack.StripStackMatrix(target1, target2, ..., targetN)**
 * **stack.ToArray()**
 * **stack.ToMap()**
 * **stack.ToMatrix(...depth)**
 * **card.ToPointer(.....todo: add)**
 * **stack.ToPointer(.....todo: add)**
 * **card.ToObject(.....todo: add)**
 * **stack.ToObject(.....todo: add)**
 * **stack.IsRegular()**
 * **stack.Duplicate(...n)**
 * **stack.Empty()**
 * **card.Clone(...cloneKey, ...cloneVal)**
 * **stack.Clone(...cloneCards, ...cloneKeys, ...cloneVals)**
 * **stack.Unique(typeType, ...pointerType, ...deepSearchType, ...depth)**
 * **card.Equals(Card, ...pointerTypeKey, ...pointerTypeVal, ...compareIdxs, ...printType)**
 * **stack.Equals(Stack, ...compareStacks, ...pointerTypeStack, ...deepSearchType, ...depth, ...pointerTypeKey, ...pointerTypeVal)**
 * **stack.Shuffle()**
 * **stack.Inverse()**
 * **card.Print()**
 * **stack.Print()**
 * **stack.Lambda(lambda function, ...deepSearchType, ...depth)**

<h2 name = "generalizedFunctionsBrief">Generalized Functions</h2>

 * **stack.Add(insert, ...orderType, ...findType, ...findData, ...pointerType, ...deepSearchType, ...depth, ...overrideStackConversion)**
 * **stack.Move(findType_from, orderType, findType_to, ...findData_from, ...findData_to, ...pointerType_from, ...pointerType_to, ...deepSearchType, ...depth)**
 * **stack.Swap(findType_from, findType_to, ...findData_from, ...findData_to, ...pointerType_from, ...pointerType_to, ...deepSearchType, ...depth)**
 * **stack.Has(returnType, findType, ...findData, ...pointerType, ...deepSearchType, ...depth)**
 * **stack.Get(...findType, ...findData, ...pointerType, ...clonesType_card, ...clonesType_keys, ...clonesType_vals, ...deepSearchType, ...depth)**
 * **stack.GetMany(findType, ...findData, ...pointerType, ...returnType, ...clonesType, ...clonesType_keys, ...clonesType_vals, ...deepSearchType, ...depth)**
 * **stack.Replace(replaceType, replaceData, findType, ...findData, ...pointerType, ...deepSearchType, ...depth)**
 * **stack.ReplaceMany(replaceType, replaceData, findType, ...findData, ...pointerType, ...returnType, ...deepSearchType, ...depth)**
 * **stack.Update(findType, ...findData, ...pointerType, ...deepSearchType, ...depth)**
 * **stack.UpdateMany(findType, ...findData, ...pointerType, ...deepSearchType, ...depth)**
 * **stack.Extract(findType, ...findData, ...pointerType, ...deepSearchType, ...depth)**
 * **stack.ExtractMany(findType, ...findData, ...pointerType, ...returnType, ...deepSearchType, ...depth)**
 * **stack.Remove(findType, ...findData, ...pointerType, ...deepSearchType, ...depth)**
 * **stack.RemoveMany(findType, ...findData, ...pointerType, ...deepSearchType, ...depth)**

<h1 name = "exhaustiveDocumentation">Exhaustive Documentation</h1>

<h2 name = "dataStructures">Data Structures</h2>

<h3 name = "structs">structs</h3>

<h4 name = "stack">Stack</h4>

 This is the main struct in the project.

 > `stack` *Stack{}*
 >> `Cards` *[]\*Card{}*
 >>> Returns an interface array to represent the elements in the Stack
 >>
 >> `Size` *int*
 >>> Returns the cardinality (i.e., `len(stack.Cards)`) of this Stack
 >>
 >> `Depth` *int*
 >>> Returns the dimensionality of this Stack, assuming it is uniform

<h4 name = "card">Card</h4>

 This is a struct for our elements/maps within stacks.

 >> `Card` *Card{}*
 >>> `card.Idx` *int*
 >>>> The index of this card
 >>>
 >>> `card.Key` *any*
 >>>> The key of this card (or nil if doesn't exist)
 >>>
 >>> `card.Val` *any*
 >>>> The val of this card (or nil if doesn't exist)

<h3 name = "enums">enums</h3>

Please note that enumerator defaults are default in *most cases*.  However, you should defer to each function's documentation to see the actual argument default values.

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
 >>> stack of anys
 >>
 >> RETURN_Vals
 >>> stack of anys
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
 >>> stack where type{stack.Card.Val} == int
 >>
 >> FIND_Key
 >>> any
 >>
 >> FIND_Keys
 >>> []any
 >>
 >> FIND_KeysStack
 >>> stack where type{stack.Card.Val} == any
 >>
 >> FIND_Val
 >>> any
 >>
 >> FIND_Vals
 >>> []any
 >>
 >> FIND_ValsStack
 >>> stack where type{stack.Card.Val} == any
 >>
 >> FIND_Card
 >>> *Card
 >>
 >> FIND_Cards
 >>> *Stack (input is the cards in this stack)
 >>
 >> FIND_CardsStack
 >>> stack where type{stack.Card.Val} == *Card
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
 >>> *any*
 >>
 >> REPLACE_Val
 >>> *any*
 >>
 >> REPLACE_Card
 >>> *Card*
 >>
 >> REPLACE_Stack
 >>> *Stack*
 >>>> replaces a card with each card in a stack of cards
 >>
 >> REPLACE_Lambda
 >>> any *lambda function*
 
<h4 name = "TYPE">TYPE</h4>

 This is an enum intended to make it easy to tell certain functions the type of value being targeted.

> ***TYPE***
>> TYPE_Key
>
>> TYPE_Val
 
<h4 name = "ORDER">ORDER</h4>

 This is an enum intended to make it easy to tell certain functions whether to insert a value before or after the input index.

> ***ORDER***
>> ORDER_Before
>>> default
>>
>> ORDER_After

<h4 name = "POINTER">POINTER</h4>

 This is an enum intended to make it easy to target whether a function searching for a match between input data and data in the stack element is matching by having the same values (POINTER_False) or the same memory address (POINTER_True).

 Matching by reference only works for Val, Key, and Card FIND types.  It would not make much sense to match an index that's managed on the backend by reference (FIND_Idx), to match a lambda expression (FIND_Lambda), or to match using a position that's not even comparing values (FIND_First, FIND_Last, FIND_All).

 Take care to note that all cases where objects are matching by reference will also be matching by object.

 > ***POINTER***
 >> POINTER_False
 >>> default
 >>
 >> POINTER_True

<h4 name = "CLONE">CLONE</h4>

 This is an enum intended to make it possible to tell the function whether to return a clone of an object or a pointer of an object.

 > ***POINTER***
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

<h4 name = "COMPARE">COMPARE</h4>

 This is an enum intended to make it easy to call functions to decide whether to compare things in Equals tests.

 > ***COMPARE***
 >> COMPARE_True
 >>
 >> COMPARE_False
 >>> default

<h2 name = "nonGeneralizedFunctions">Non-Generalized Functions</h2>

<h3 name = "MakeCard">MakeCard</h3>

 `gostack.MakeCard(...val, ...key, ...idx)`
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
 
 @param optional `input1` type{any} default nil
 @param optional `input2` type{any} default nil
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
 
<h3 name = "StripStackMatrix">StripStackMatrix</h3>
 
 `stack.StripStackMatrix()`
 ```
 Returns a stack representing a selection within a stack matrix
 
 @receiver `stack` type{*Stack}
 @param variadic `selections` type{int, []int} a set of args representing the indices being selected within an array
 @returns type{*Stack} a new Stack representing the selection
 @constructs type{*Stack} a new Stack representing the selection
 @requires `idx` arguments get valid index positions from the stack
 ```
 
<h3 name = "ToArray">ToArray</h3>
 
 `stack.ToArray()`
 ```
 Creates a new interface array from values of `stack`

 @receiver `stack` type{*Stack}
 @returns type{[]any} new array
 @ensures new array values correspond to `stack` values
 ```
 
<h3 name = "ToMap">ToMap</h3>
 
 `stack.ToMap()`
 ```
 Creates a new interface-interface map from values of `stack`

 @receiver `stack` type{*Stack}
 @returns type{map[any]any} new map
 @ensures new map keys and values correspond to `stack` keys and values
 ```
 
<h3 name = "ToMatrix">ToMatrix</h3>
 
 `stack.ToMatrix(...depth)`
 ```
 Creates a new matrix from a stack by depth.  For instance, if depth = 2, then returns the stacks inside stack as an [][]any

 @receiver `stack` type{*Stack}
 @param optional `depth` type{int} default -1
 @returns type{[]interface}
 @ensures
  * -1 depth means it will go as deep as it can
  * new map keys and values correspond to `stack` keys and values
  * example: Stack{Stack{"Hi"}, Stack{"Hello", "Hola"}, "Hey"} =>
      []any{[]any{"Hi"}, []any{"Hola", "Hello"}, "Hey"}
 ```
 
<h3 name = "ToPointer">ToPointer</h3>
 
 `card.ToPointer()`
 ```

 ```
 
 `stack.ToPointer()`
 ```

 ```
 
<h3 name = "ToObject">ToObject</h3>
 
 `card.ToObject()`
 ```

 ```
 
 `stack.ToObject()`
 ```

 ```
 
<h3 name = "IsRegular">IsRegular</h3>
 
 `stack.IsRegular()`
 ```
 Returns whether the matrix is of a regular shape

 @receiver `stack` type{*Stack}
 @returns type{bool}
 @ensures
   * example:
       {{1, 2}, 3} == irregular/false
       {{1, 2}, {3}} == irregular/false
       {{1, 2}, {3, 4}} == regular/true
	   {1, 3} == regular/true
	   {} == regular/true
 ```
 
<h3 name = "Duplicate">Duplicate</h3>
 
 `stack.Duplicate(n ...int)`
 ```
 Adds the cards in `stack` to itself `n` - 1 times
  (duplicate 4 means 3 duplicates made; duplicate 1 means don't duplicate; duplicate 0 means empty)
 
 @receiver `stack` type{*Stack}
 @param optional `n` type{int} default 2
 @updates `stack`
 @returns `stack`
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

 `card.Clone(...cloneKey, ...cloneVal)`
 ```
 Returns a clone of the given card

 @receiver `card` type{*Card}
 @returns type{*Card} card clone
 @constructs clone of `card`
 ```

 `stack.Clone(...cloneCards, ...cloneKeys, ...cloneVals)`
 ```
 Returns a clone of the given stack

 @receiver `stack` type{*Stack}
 @returns type{*Stack} stack clone
 @constructs type{*Stack} clone of `stack`
 @ensures the stack clone has the same card pointers as `stack`
 ```
 
<h3 name = "Unique">Unique</h3>
 
 `stack.Unique(typeType, ...pointerType, ...deepSearchType, ...depth)`
 ```
 Removes all cards from `stack` which share the same field value as another card before

 @receiver `stack` type{*Stack}
 @param `typeType` type{TYPE}
 @param optional `pointerType` type{POINTER} default POINTER_False
 @returns `stack`
 @updates `stack` to have no repeating values between field `typeType`
 ```
 
<h3 name = "Equals">Equals</h3>
 
 `card.Equals(*Card, ...compareCards, ...pointerTypeCard, ...pointerTypeKey, ...pointerTypeVal, ...compareIdxs)`
 ```
 Returns whether two cards equal one another
 
 @receiver `thisCard` type{*Card}
 @param `otherCard` type{*Card}
 @param optional `pointerTypeKey` type{POINTER} default POINTER_False
 @param optional `pointerTypeVal` type{POINTER} default POINTER_False
 @param optional `compareIdxs` type{bool} default false
 @returns type{bool}
 ```
 
 `stack.Equals(*Stack, ...compareStacks, ...pointerTypeStack, ...deepSearchType, ...compareCards, ...pointerTypeCard, ...pointerTypeKey, ...pointerTypeVal)`
 ```
 Returns whether two stacks equal one another
 
 @receiver `thisStack` type{*Stack}
 @param `otherStack` type{*Stack}
 @param optional `compareStacks` type{COMPARE} default COMPARE_False
	By default, does not compare the stack structs, but rather their cards; can be set true and adjusted with `pointerTypeStack`
 @param optional `pointerTypeStack` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @param optional `pointerTypeKey` type{POINTER} default POINTER_False
 @param optional `pointerTypeVal` type{POINTER} default POINTER_False
 @returns type{bool}
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
 
<h3 name = "Inverse">Inverse</h3>
 
 `stack.Inverse()`
 ```
 Inverses the ordering of `stack.Cards`
 
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
 
<h3 name = "Lambda">Lambda</h3>
 
 `stack.Lambda(lambda function, ...deepSearchType, ...depth)`
 ```
 Iterate through a stack calling your lambda function on each card
 
 @receiver `stack` type{*Stack}
 @param `lambda` type{func(*Card, ...any)}
 @ensures
  * Each card in `stack` is passed into your lambda function
  * `stack` is the first argument passed into your variadic parameter on the first call
 ```

<h2 name = "generalizedFunctions">Generalized Functions</h2>
 
<h3 name = "Add">Add</h3>
 
 `stack.Add(insert, ...orderType, ...findType, ...findData, ...pointerType, ...deepSearchType, ...depth, ...overrideStackConversion)`
 ```
 Adds to a stack of cards or a cards at (each) position(s) and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `insert` type{Card, Stack}
 @param optional `orderType` type{ORDER} default ORDER_Before
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @param optional `overrideStackConversion` type{bool} default false
	if `insert` is of type Stack:
		if not `overrideStackConversion`:
			add to `stack` from `insert.Cards`
		else if `overrideStackConversion`:
			add the `insert` stack to `stack` as the val of a card
 @returns `stack` if cards were added OR nil if no cards were added (due to invalid find)
 @updates `stack` to have new cards before/after each designated position
 @requires `stack.Clone()` has been implemented
 ```
 
<h3 name = "Move">Move</h3>
 
 `stack.Move(findType_from, orderType, findType_to, ...findData_from, ...findData_to, ...pointerType_from, ...pointerType_to)`
 ```
 Moves one element or slice of cards to before or after another element or slice of cards
 
 @receiver `stack` type{*Stack}
 @param `findType_from` type{FIND}
 @param `orderType` type{ORDER}
 @param `findType_to` type{FIND}
 @param optional `findData_from` type{any} default nil
 @param optional `findData_to` type{any} default nil
 @param optional `pointerType_from` type{POINTER} default POINTER_False
 @param optional `pointerType_to` type{POINTER} default POINTER_False
 @param optional `deepSearchType_from` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `deepSearchType_to` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth_from` type{int} default -1 (deepest)
 @param optional `depth_to` type{int} default -1 (deepest)
 @returns `stack` if moved OR nil if no move occurred (due to bad find)
 @requires you are not moving a stack to a location within that own stack
 @ensures a stack of cards, or individual cards, can be targeted
 ```
 
<h3 name = "Swap">Swap</h3>
 
 `stack.Swap(findType_from, findType_to, ...findData_from, ...findData_to, ...pointerType_from, ...pointerType_to, ...deepSearchType, ...depth)`
 ```
 Swaps one element or slice with the position of another element or slice
 
 @receiver `stack` type{*Stack}
 @param `findType_first` type{FIND}
 @param `findType_second` type{FIND}
 @param optional `findData_first` type{any} default nil
 @param optional `findData_second` type{any} default nil
 @param optional `pointerType_first` type{POINTER} default POINTER_False
 @param optional `pointerType_second` type{POINTER} default POINTER_False
 @param optional `deepSearchType_first` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `deepSearchType_second` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth_first` type{int} default -1 (deepest)
 @param optional `depth_second` type{int} default -1 (deepest)
 @returns `stack` if moved OR nil if no move occurred (due second bad find)
 @requires you are not swapping a stack with a location within that own stack
 @ensures a stack of cards, or individual cards, can be targeted
 ```
 
<h3 name = "Has">Has</h3>
 
 `stack.Has(...findType, ...findData, ...pointerType)`
 ```
 Returns a boolean representing whether a search exists in the stack

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @returns true IF successful search, false IF unsuccessful search
 @requires `stack.Get()` has been implemented
 ```
 
<h3 name = "Get">Get</h3>
 
 `stack.Get(...findType, ...findData, ...pointerType, ...clonesType_card, ...clonesType_keys, ...clonesType_vals)`
 ```
 Gets a card from specified parameters in a stack, or nil if does not exist

 @receiver `stack` type{*Stack}
 @param optional `findType` type{FIND} default FIND_First
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
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
 
 `stack.GetMany(findType, ...findData, ...returnType, ...pointerType, ...clonesType, ...clonesType_keys, ...clonesType_vals)`
 ```
 Gets a stack from specified parameters in a stack
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `clonesType` type{CLONES} default CLONE_False
 @param optional `clonesType_keys` type{CLONES} default CLONE_False
 @param optional `clonesType_vals` type{CLONES} default CLONE_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} the new stack (if find fails, then an empty stack)
 @constructs type{*Stack} new stack of specified values from specified cards in `stack`
 @requires
  * `MakeStack()` has been implemented
  * `clonesType_keys` and `clonesType_vals` are only passed if `returnType` == RETURN_Cards
 @ensures
  * CLONE_True means the cards in the returned stack are clones
  * CLONE_True for `clonesType_keys` means the cards in the returned stack keys are clones
  * CLONE_True for `clonesType_vals` means the cards in the returned stack vals are clones
 ```
 
<h3 name = "Replace">Replace</h3>
 
 `stack.Replace(replaceType, replaceData, findType, ...findData, ...pointerType)`
 ```
 Returns a clone of a found card before its respective field is updated to `replaceData` (OR nil if not found)
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @returns type{*Card} a clone of extracted card OR nil if found no cards
 @updates first found card to `replaceData`
 @requires `stack.Get()` has been implemented
 @ensures if `replaceData` is nil and `replaceType is REPLACE_Card`, the card will be removed from `stack`
 ```
 
<h3 name = "ReplaceMany">ReplaceMany</h3>
 
 `stack.Replace(replaceType, replaceData, findType, ...findData, ...returnType, ...pointerType)`
 ```
 Returns a stack whose values are the original fields updated to `replaceData`
 
 @receiver `stack` type{*Stack}
 @param `replaceType` type{REPLACE}
 @param `replaceData` type{any}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} a stack whose values are the extracted cards pre-update (if find fails, then an empty stack)
 @updates all found cards to `replaceData`
 @requires `stack.GetMany()` has been implemented
 @ensures IF `replaceData` is nil and `replaceType is REPLACE_Card`, the cards found will be removed from `stack`
 ```
 
<h3 name = "Update">Update</h3>
 
 `stack.Update(findType, ...findData, ...pointerType)`
 ```
 Updates a card in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @returns `stack`
 @updates the found card in `stack`
 @requires `stack.Replace()` has been implemented
 ```
 
<h3 name = "UpdateMany">UpdateMany</h3>
 
 `stack.UpdateMany(findType, ...findData, ...pointerType)`
 ```
 Updates cards in and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @returns `stack`
 @updates  the found cards in `stack`
 @requires `stack.ReplaceMany()` has been implemented
 ```
 
<h3 name = "Extract">Extract</h3>
 
 `stack.Extract(findType, ...findData, ...pointerType)`
 ```
 Gets and removes a card from `stack`, or returns nil if it does not exist
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @returns type{*Card} the extracted card OR nil if invalid FIND
 @updates `stack` to no longer have found card
 ```
 
<h3 name = "ExtractMany">ExtractMany</h3>
 
 `stack.ExtractMany(findType, ...findData, ...returnType, ...pointerType)`
 ```
 Gets and removes a set of data from `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `returnType` type{RETURN} default RETURN_Cards
 @param optional `pointerType` type{POINTER} default POINTER_False
 @param optional `deepSearchType` type{DEEPSEARCH} default DEEPSEARCH_False
 @param optional `depth` type{int} default -1 (deepest)
 @returns type{*Stack} the extracted card (if find fails, then an empty stack)
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 ```
 
<h3 name = "Remove">Remove</h3>
 
 `stack.Remove(findType, ...findData, ...pointerType)`
 ```
 Removes a card from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @returns `stack`
 @updates `stack` to no longer have found card
 @requires `stack.Replace()` has been implemented
 ```
 
<h3 name = "RemoveMany">RemoveMany</h3>
 
 `stack.RemoveMany(findType, ...findData, ...pointerType)`
 ```
 Removes a set of cards from and returns `stack`
 
 @receiver `stack` type{*Stack}
 @param `findType` type{FIND}
 @param optional `findData` type{any} default nil
 @param optional `pointerType` type{POINTER} default POINTER_False
 @returns `stack`
 @updates `stack` to no longer have found cards
 @requires `stack.ReplaceMany()` has been implemented
 ```
 
<h2 name = "futureUpdates">Future Updates</h2>

 *To be added later*

<h2 name = "footer">Footer</h1>

This library was created by Gabe Tucker.

If you have any suggestions, questions, or comments you would like to make in respect to this project, please email `tucker.854@osu.edu`.  I appreciate any feedback and will usually respond within 1-2 business days.

Feel free to visit me at `https://www.linkedin.com/in/gabetucker2/`!

[Return to Glossary](#glossary)
