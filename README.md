 <h1 name = "preface">PREFACE</h1>

![Banner](images/gostack_Smaller.png)

 <h1 name = "introduction">Introduction</h1>

 `gostack` introduces **Stack** structures—generic sets of elements intended as an all-in-one package for datastructure management in *golang*.  The elements in stacks are **Card** structures (like a stack of cards).

 With `gostack`, there is no need for maps or arrays; every possible tool you could need to create, update, or access a set of data is encompassed by a few elegant functions.
 
 Assuming `stack` is a predefined stack of cards:

 > Want to remove the first card in a `stack` and get its key?
 >> `key := stack.Extract(RETURN_Key, POSITION_First)`
 
 > Want to get a unique stack of values of the cards in `stack` whose keys match the object address of Cards defined as UnitTypes unitType1 or unitType2 (where UnitType is your user-defined struct)?
 >> `keys := MakeStack().Add(MakeCard(unitType1), ORDER_After, POSITION_Last).Add(MakeCard(unitType2), ORDER_After, POSITION_Last)`
 >
 >> `unitIndices := stack.Get(RETURN_Idxs, POSITION_Keys, keys, MATCH_Reference).Unique(TYPE_Val)`
 
 > Want to replace all cards in `stack` whose values are even ints between -5 and 3 with two new cards and get a stack representing keys of the cards that were replaced?
 >> `cardsToInsert := MakeStack().Add(newCard1, ORDER_After, POSITION_Last).Add(newCard2, ORDER_After, POSITION_Last)`
 >
 >> func gostack_ValInRange(stack *Stack, card *Card, workingMemory ...*Stack) (ret bool) {
 >>> v := card.val.(int)
 >>> return -3 < v && v 5 14 && v%2 == 0
 >> }
 >
 >> `oldCards := stack.Replace(cardsToInsert, RETURN_Keys, POSITION_Lambda, gostack_ValInRange)`

<h1 name = "glossary">Glossary</h1>

 > [Files](#files)

 > [Preface](#preface)
 >> [Introduction](#introduction)
 >
 >> [Glossary](#glossary)
 >
 >> [File Explanations](#fileExplanations)

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
 >>>>> [POSITION](#POSITION)
 >>>>
 >>>>> [TYPE](#TYPE)
 >>>>
 >>>>> [ORDER](#ORDER)
 >>>>
 >>>>> [MATCH](#MATCH)
 >>>
 >>> [Non-Generalized Functions](#nonGeneralizedFunctions)
 >>>> [MakeStack()](#MakeStack)
 >>>
 >>>> [MakeCard(...)](#MakeCard)
 >>>
 >>>> [stack.Empty()](#Empty)
 >>>
 >>> [Generalized Functions](#generalizedFunctions)
 >>>> [stack.Add(...)](#Add)
 >>>
 >>>> [stack.Replace(...)](#Replace)
 >>>
 >>>> [stack.Extract(...)](#Extract)
 >>>
 >>>> [stack.Unique(...)](#Unique)
 >>>
 >>>> [stack.Get(...)](#Get)
 >>>
 >>>> [stack.Has(...)](#Has)
 >
 >> [Future Updates](#futureUpdates)
 >
 >> [Footer](#footer)

<h1 name = "fileExplanations">File Explanations</h1>

 > *aorta*
 >> **backend.go** contains the functions to implement **frontend.go** and **caseend.go** functions
 >
 >> **datastructures.go** initializes structs and enums
 >
 >> **frontend.go** contains the functions that the programmer using `gostack` will be calling
 >
 > *casetests*
 >> **caseend.go** contains case tests for **frontend.go** functions
 >
 >> **testend.go** contains functions to implement **caseend.go** functions
 >
 > *images*
 >> **gostack_Smaller.png** is the banner image for this project
 >
 > *tutorials*
 >> **bootstrap.go** is a tutorial on how to implement some common functions using golang
 >
 >> **lambda.go** is a tutorial on how to implement lambda functions
 >
 >> **unaddedData.txt** is where data to be added into future tutorials is stored
 >
 > **go.mod** is to initialize the directories
 >
 > **main.go** is to call functions in this project, either for the case of case testing or executing tutorials
 >
 > **README.md** is this file
 >
 > **TODO.txt** is a task list relevant to developers

<h1 name = "overview">OVERVIEW</h1>

<h1 name = "briefDocumentation">Brief Documentation</h1>

<h2 name = "dataStructuresBrief">Data Structures</h2>

<h3 name = "structsBrief">Structs</h3>

 > **stack** *Stack*
 >> **cards** *[]\*Card*
 >
 >> **size** *int*

 > **card** *Card*
 >> **key** *any type*
 >
 >> **val** *any type*

<h3 name = "enumsBrief">Enums</h3>

 > **RETURN**
 > * _RETURN_NotationSample *type that's returned*
 > * RETURN_None *NONE*
 > * RETURN_Idx *int*
 > * RETURN_Idxs *Stack of ints*
 > * RETURN_Key *any type*
 > * RETURN_Keys *Stack of any type*
 > * RETURN_Val *any type*
 > * RETURN_Vals *Stack of any type*
 > * RETURN_Card *Card*
 > * RETURN_Cards *Stack of Cards*

 > **POSITION**
 > * _POSITION_NotationSample *POSITIONDATA argument type*
 > * POSITION_First *NONE*
 > * POSITION_Last *NONE*
 > * POSITION_Idx *int*
 > * POSITION_Idxs *Stack of ints*
 > * POSITION_Val *any type*
 > * POSITION_Vals *Stack of any type*
 > * POSITION_Key *any type*
 > * POSITION_Keys *Stack of any type*
 > * POSITION_Card *Card*
 > * POSITION_Cards *Stack of Cards*
 > * POSITION_All *NONE*
 > * POSITION_Lambda *lambda function*

 > **ORDER**
 > * ORDER_Before
 > * ORDER_After

 > **MATCH**
 > * MATCH_Object
 > * MATCH_Reference

<h2 name = "nonGeneralizedFunctionsBrief">Non-Generalized Functions</h2>

 * **MakeStack()**
 * **MakeCard(...val, ...key)**
 * **stack.Empty()**

<h2 name = "generalizedFunctionsBrief">Generalized Functions</h2>

 * **stack.Add(insert, ORDER_\*, POSITION_\*, ...POSITIONDATA)**
 * **stack.Replace(insert, RETURN_\*, POSITION_\*, ...POSITIONDATA, ...MATCH_\*)**
 * **stack.Extract(RETURN_\*, POSITION_\*, ...POSITIONDATA, ...MATCH_\*)**
 * **stack.Unique(TYPE_\*)**
 * **stack.Get(RETURN_\*, POSITION_\*, ...POSITIONDATA, ...MATCH_\*)**
 * **stack.Has(RETURN_\*, POSITION_\*, ...POSITIONDATA, ...MATCH_\*)**

<h1 name = "exhaustiveDocumentation">Exhaustive Documentation</h1>

<h2 name = "dataStructures">Data Structures</h2>

<h3 name = "structs">structs</h3>

<h4 name = "stack">Stack</h4>

 This is the main struct in the project.

 > `stack` *Stack{}*
 >> `cards` *[]\*Card{}*
 >>> Returns an interface array to represent the elements in the Stack
 >> `size` *int*
 >>> Returns the cardinality (i.e., `len(stack.cards)`) of this Stack

<h4 name = "card">Card</h4>

 This is a struct for our elements/maps within stacks.

 >> `card` *Card{}*
 >>> `card.idx` *int*
 >>>> The index of this card
 >>>
 >>> `card.key` *any type (interface{})*
 >>>> The key of this card (or nil if doesn't exist)
 >>>
 >>> `card.val` *any type (interface{})*
 >>>> The val of this card (or nil if doesn't exist)

<h3 name = "enums">enums</h3>

<h4 name = "RETURN">RETURN</h4>

 This is an enum intended to make it easy for the user to control the output type from a getter function.

 > ***RETURN***
 >> *_RETURN_NotationSample*
 >>> *The type of variable returned by the function you're calling*
 >>
 >>> *Although the type may say int or Card, the true return type will always be interface{} or nil*
 >>
 >>> *For instance, if you inputted RETURN_Key, you would get a single key interface{} (or nil if doesn't exist).  Alternatively, if you inputted RETURN_Keys, you would get a stack of keys.*
 >>
 >> RETURN_None
 >>> nil
 >>
 >> RETURN_Idx
 >>> int
 >>
 >> RETURN_Idxs
 >>> Stack of ints
 >>
 >> RETURN_Key
 >>> interface{}
 >>
 >> RETURN_Keys
 >>> Stack of interfaces{}
 >>
 >> RETURN_Val
 >>> interface{}
 >>
 >> RETURN_Vals
 >>> Stack of interfaces{}
 >>
 >> RETURN_Card
 >>> Card
 >>
 >> RETURN_Cards
 >>> Stack of Cards

<h4 name = "POSITION">POSITION</h4>

 This is an enum intended to make it easy to flexibly inform functions what the intended target is.

 > ***POSITION***
 >> *_POSITION_NotationSample*
 >>> *The type of the variable (called `data`) that needs to be passed into the function utilizing this constant*
 >>
 >>> *For instance, if you input `POSITION_Keys`, you would need to pass a Stack whose values are the keys you want to find to your `data` parameter*
 >>
 >> POSITION_First
 >>> *NONE*
 >>
 >> POSITION_Last
 >>> *NONE*
 >>
 >> POSITION_Idx
 >>> int
 >>
 >> POSITION_Idxs
 >>> Stack of ints
 >>
 >> POSITION_Val
 >>>  any type (interface{})
 >>
 >> POSITION_Vals
 >>> Stack of any type (interface{})
 >>
 >> POSITION_Key
 >>>  any type (interface{})
 >>
 >> POSITION_Keys
 >>> Stack of  any type (interface{})
 >>
 >> POSITION_Card
 >>> Card
 >>
 >> POSITION_Cards
 >>> Stack of Cards
 >>
 >> POSITION_All
 >>> *NONE*
 >>
 >> POSITION_Lambda
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
>> ORDER_After

<h4 name = "MATCH">MATCH</h4>

 This is an enum intended to make it easy to target whether a function searching for a match between input data and data in the stack element is matching by having the same values (MATCH_Object) or the same memory address (MATCH_Reference).

 Matching by reference only works for Val, Key, and Card POSITION types.  It would not make much sense to match an index that's managed on the backend by reference (POSITION_Idx), to match a lambda expression (POSITION_Lambda), or to match using a position that's not even comparing values (POSITION_First, POSITION_Last, POSITION_All).

 Take care to note that all cases where objects are matching by reference will also be matching by object.

 > ***MATCH***
 >> MATCH_Object
 >>> default
 >> MATCH_Reference

<h2 name = "nonGeneralizedFunctions">Non-Generalized Functions</h2>

<h3 name = "MakeStack">MakeStack</h3>

 > `MakeStack()`
 >> CONSTRUCTOR: ***TRUE***
 >>> Stack
 >
 >> GETTER: ***TRUE***
 >>> Stack
 >
 >> SETTER: ***FALSE***
 
 > ***Pseudocode***
 >> returns a new Stack

<h3 name = "MakeCard">MakeCard</h3>

 > `MakeCard(...val, ...key)`
 >> CONSTRUCTOR: ***TRUE***
 >>> Card
 >
 >> GETTER: ***TRUE***
 >>> Card
 >
 >> SETTER: ***FALSE***
 
 > ***Special Parameters***
 >> **val** *any type* is the new Card's starting val
 >
 >> **key** *any type* is the new Card's starting key
 
 > ***Pseudocode***
 >> Creates a new Card card
 >
 >> Set card.val = **val** (or nil if empty) and card.key == **key** (or nil if empty)
 >
 >> returns card
 
<h3 name = "Empty">Empty</h3>
 
 > `stack.Empty()`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> **stack**
 >
 >> SETTER: ***TRUE***
 >>> **stack**
 
 > ***Pseudocode***
 >> remove all cards in the stack
 >
 >> returns the empty stack
 
<h2 name = "generalizedFunctions">Generalized Functions</h2>
 
<h3 name = "Add">Add</h3>
 
 > `stack.Add(insert, ORDER_*, POSITION_*, ...POSITIONDATA)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> **stack** *or* nil
 >
 >> SETTER: ***TRUE***
 >>> **stack**
 
 > ***Special Parameters***
 >> **insert** *Card* or *Stack* is either a Card or a Stack of cards to insert at POSITION

 > ***Pseudocode***
 >> **IF VALID POSITION**
 >>> **IF beforeNotAfter**
 >>>> add **insert** before (each) POSITION in **stack**
 >>>
 >>> **ELSE**
 >>>> add **insert** after (each) POSITION in **stack**
 >>>
 >>> returns updated stack
 >>
 >> **ELSE**
 >>> return nil
 
<h3 name = "Replace">Replace</h3>
 
 > `stack.Replace(insert, RETURN_*, POSITION_*, ...POSITIONDATA, ...MATCH_*)`
 >> CONSTRUCTOR: ***SOMETIMES***
 >>> Make Stack if RETURNing multiple types
 >
 >> GETTER: ***TRUE***
 >>> RETURN objects that were removed *or* nil
 >
 >> SETTER: ***TRUE***
 >>> **stack**
 
 > ***Special Parameters***
 >> **insert** *Card* or *Stack* is either a Card or a Stack of cards to insert at POSITION(S) as the replacement

 > ***Pseudocode***
 >> **IF VALID POSITION**
 >>> replace cards at each POSITION in **stack** with **insert**
 >>
 >>> returns the selected RETURNS
 >
 >> **ELSE**
 >>> return nil
 
<h3 name = "Extract">Extract</h3>
 
 > `stack.Extract(RETURN_*, POSITION_*, ...POSITIONDATA, ...MATCH_*)`
 >> CONSTRUCTOR: ***SOMETIMES***
 >>> Make Stack if RETURNing multiple types
 >
 >> GETTER: ***TRUE***
 >>> RETURN objects that were removed *or* nil
 >
 >> SETTER: ***TRUE***
 >>> **stack**

 > ***Pseudocode***
 >> **IF VALID POSITION**
 >>> remove cards from the stack based on provided POSITION data
 >>
 >>> return the RETURNS of the old cards
 >
 >> **ELSE**
 >>> return nil
 
<h3 name = "Unique">Unique</h3>
 
 > `stack.Unique(TYPE_*)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> **stack**
 >
 >> SETTER: ***TRUE***
 >>> **stack**

 > ***Pseudocode***
 >> removes cards from the stack whose TYPE are the same value as others in the stack's TYPE values
 >
 >> return **stack**
 
<h3 name = "Get">Get</h3>
 
 > `stack.Get(RETURN_*, POSITION_*, ...POSITIONDATA, ...MATCH_*)`
 >> CONSTRUCTOR: ***SOMETIMES***
 >>> Make Stack if RETURNing multiple types
 >
 >> GETTER: ***TRUE***
 >>> RETURN objects that were gotten *or* nil
 >
 >> SETTER: ***FALSE***

 > ***Pseudocode***
 >> **IF VALID POSITION**
 >>> return the selected RETURNS
 >
 >> **ELSE**
 >>> return nil
 
<h3 name = "Has">Has</h3>
 
 > `stack.Has(RETURN_*, POSITION_*, ...POSITIONDATA, ...MATCH_*)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> bool
 >
 >> SETTER: ***FALSE***
 
 > ***Pseudocode***
 >> **IF STACK HAS TARGETED DATA**
 >>> return true
 >
 >> **ELSE**
 >>> return false
 
<h2 name = "futureUpdates">Future Updates</h2>

 <h3>Generalized Functions</h3>

 * Add **Move** function
 * Add **Fill** function
 * Add **Set(newData, TYPE_*, POSITION_*, ...POSITIONDATA)** function for more efficient replacement as opposed to replace... implement all search functions for individual cards, so card.set, card.replace, card.extract, etc

 <h3>Non-Generalized Functions</h3>

 * Add **CombineWith** function
 * Add **Flip** function
 * Add **Shuffle** function
 * Add **Clone** function
 * Add **ToArray** function
 * Add **ToStack** function

<h2 name = "footer">Footer</h1>

This project was created by Gabe Tucker with the help of Andy Chen.

If there are any changes or comments you would like to have made in respect to this project, please email `tucker.854@osu.edu`.  I appreciate any feedback and will usually respond within 1-2 business days.

Feel free to visit my personal pages at `https://gabetucker.com` or `https://www.linkedin.com/in/gabetucker2/`.

[Return to Glossary](#glossary)
