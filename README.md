 <h1 name = "preface">PREFACE</h1>

![Banner](Images/gostack_Smaller.png)

 <h1 name = "introduction">Introduction</h1>

 `gostack` introduces **Stack** structures—generic sets of elements intended as an all-in-one package for datastructure management in *golang*.  The elements in stacks are **Card** structures (like a stack of cards).

 With `gostack`, there is no need for maps or arrays; every possible tool you could need to create, update, or access a set of data is encompassed by a few elegant functions.
 
 Assuming `stack` is a predefined stack of cards:

 > Want to remove the first card in a `stack` and get its key?
 >> `key := stack.Extract(RETURN_Key, POSITION_First)`
 
 > Want to replace all cards in `stack` whose values are even ints between -5 and 3 with two new cards and get a stack representing keys of the cards that were replaced?
 >> `cardsToInsert := MakeStack().Add(newCard1, ORDER_After, POSITION_Last).Add(newCard2, ORDER_After, POSITION_Last)`
 >
 >> `gostack_lambda := func(int) (int) {  }`
 >
 >> `oldCards := stack.Replace(cardsToInsert, RETURN_Keys, POSITION_Lambda, LAMBDAAAA)`
 
 > Want to get a unique stack of values of the cards in `stack` whose keys match the object address of Cards defined as UnitTypes unitType1 or unitType2 (where UnitType is your user-defined struct)?
 >> `keys := MakeStack().Add(MakeCard(unitType1), ORDER_After, POSITION_Last).Add(MakeCard(unitType2), ORDER_After, POSITION_Last)`
 >
 >> `unitIndices := stack.Get(RETURN_Idxs, POSITION_Keys, keys, MATCH_Reference).Unique(TYPE_Val)`

<h1 name = "glossary">Glossary</h1>

 > [Files](#files)

 > [Preface](#preface)
 >> [Introduction](#introduction)
 >
 >> [Glossary](#glossary)
 >
 >> [File Explanations](#fileExplanations)
 >
 >> [Links](#links)

 > [Overview](#overview)
 >> [Examples](#examples)
 >
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
 >>>>> [Slice](#slice)
 >>>>
 >>>> [enums](#enums)
 >>>>> [RETURN](#RETURN)
 >>>>
 >>>>> [POSITION](#POSITION)
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
 >>>> [stack.Has(...)](#Has)x
 >
 >> [To Add](#toAdd)
 >
 >> [Footer](#footer)

<h1 name = "fileExplanations">File Explanations</h1>

 * **README.md** is this file
 * **TODO.txt** is a file with features to be added (significant only to `gostack` developers)
 * **caseTests.go** is a script used to run test cases to ensure functionality of this project's functions; it is recommended to delete this file if it is not commented out at the time of your installation since it uses the main() function; in order to run test cases with ***test.go*** *not* commented out, run `go run .` in the top directory
 * **functions.go** is where novel functions are stored
 * **go.mod** is used to manage directory access
 * **structs.go** is where structs are defined

<h1 name = "links">Links</h1>

 Many of the functions in this project were inspired by functions from the documentations below.

 * https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1?view=net-6.0
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
 * http://web.cse.ohio-state.edu/software/common/doc/components/standard/Standard.html

<h1 name = "overview">OVERVIEW</h1>

<h2 name = "examples">Examples</h2>
 
 <h3>...to Demonstrate Flexibility</h3>

 > `stack.Get(RETURN_Card, POSITION_First)`
 >> *returns the first card in the Stack*
 >
 > `stack.Get(RETURN_Card, POSITION_Val, "String Value", MATCH_Object)`
 >> *goes through the stack, finds the first card with val "String Value", and returns that card*
 >
 > `stack.Get(RETURN_Card, POSITION_Key, "String Key", MATCH_Object)`
 >> *goes through the stack, finds the first card with key "String Key", and returns that card*
 >
 > `stack.Get(RETURN_Cards, POSITION_Lambda, {TODO: add})`
 >> *goes through the stack, finds each card for which the lambda expression is true, and return a stack of these cards*
 >
 > `stack.Get(RETURN_Cards, POSITION_Val, "String Value", MATCH_Object)`
 >> *goes through the stack, finds each card with val "String Value", and returns a Stack of each of those cards*
 >
 > `stack.Get(RETURN_Card, POSITION_Val, stackOfValues, MATCH_Object)`
 >> *goes through the stack, finds the first card with one of the values in stackOfValues, and returns that card*
 >
 > `stack.Get(RETURN_Cards, POSITION_Val, stackOfValues, MATCH_Object)`
 >> *goes through the stack, finds each card with one of the values in stackOfValues, and returns a Stack of each of those cards*
 >
 > `stack.Get(RETURN_Card, POSITION_Val, stackOfValues, MATCH_Reference)`
 >> *goes through the stack, finds the first card with the same memory address as one the values in stackOfValues, and returns that card*
 >
 > `stack.Get(RETURN_Cards, POSITION_Val, stackOfValues, MATCH_Reference)`
 >> *goes through the stack, finds each card with a memory address matching one in stackOfValues, and returns a Stack of each of those cards*

 <h3>stack.Push() Function Equivalent</h3>

 > `stack.Add(insert, ORDER_BEFORE, POSITION_First)`
 >> *adds a card to the beginning of the stack*

 <h3>stack.Pop() Function Equivalent</h3>

 > `stack.Extract(RETURN_Card, POSITION_First)`
 >> *removes and returns the first card in the stack*

 <h3>stack.IndexOf(card) Function Equivalent</h3>
 
 > `stack.Get(RETURN_Idx, POSITION_Card, cardToMatch, MATCH_Object)`
 >> *returns the index of the first found matching card*

<h1 name = "briefDocumentation">Brief Documentation</h1>

<h2 name = "dataStructuresBrief">Data Structures</h2>

 > ***name*** means it is conventionally acceptable to access this value manually
 >
 > **name** means it is highly recommended against accessing  value manually; you should instead use our functions

<h3 name = "structsBrief">Structs</h3>

 > ***stack*** *Stack*
 >> **cards** *[]\*Card*
 >
 >> ***size*** *int*

 > **card** *Card*
 >> **key** *any type*
 >
 >> **val** *any type*
 
 > ***Slice*** *Slice*
 >> ***startIdx*** *int*
 >
 >> ***endIdx*** *int*

<h3 name = "enumsBrief">Enums</h3>

 > ***RETURN***
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

 > ***POSITION***
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
 > * POSITION_Lambda *TODO: figure out Lambda*

 > ***ORDER***
 > * ORDER_Before
 > * ORDER_After

 > ***MATCH***
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

 It is highly advised against using these data structures for reasons other than those listed in the *Example Uses* sections.  The entire purpose of this project is for you not to have to manage arrays manually, but documentation for objects intended to be hidden still exists for those who would like to add their own Stack functions

<h3 name = "structs">structs</h3>

<h4 name = "stack">Stack</h4>

 This is the main struct in the project.

 > `stack` *Stack{}*
 >> `cards` *[]\*Card{}*
 >>> Returns an interface array to represent the elements in the Stack
 >> `size` *int*
 >>> Returns the cardinality (i.e., `len(stack.cards)`) of this Stack

 ***Example Uses***

 > `stack.size`
 >> *returns the cardinality of the stack's cards (i.e., amount of cards in the stack)*

<h4 name = "card">Card</h4>

 This is a struct for our elements/maps within stacks.

 >> `card` *Card{}*
 >>> `card.key` *any type (interface{})*
 >>>> A key for this card-map (or nil if doesn't exist)
 >>>
 >>> `card.val` *any type (interface{})*
 >>>> The val of this card (or nil if doesn't exist)

 ***Example Uses***
 
 > *None*

<h4 name = "slice">Slice</h4>

 This is a struct that makes it easier to pass two int values between functions on the backend.

 > `slice` *Slice{}*
 >> `slice.startIdx`
 >>> The first index in the desired slice
 >>
 >> `slice.endIdx`
 >>> The last index of the desired slice

 ***Example Uses***
 
 > `something = slice.startIdx`
 >
 > `something = slice.endIdx`
 >> *access the first or last indices of a slice*

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

 ***Example Uses***
 
 > `stack.Get(RETURN_Card, POSITION_Val, "String Value", MATCH_Object)`
 >> *goes through the stack, finds the first card with val "String Value", and returns that card*
 >
 > `stack.Get(RETURN_Cards, POSITION_Val, "String Value", MATCH_Object)`
 >> *goes through the stack, finds each card with val "String Value", and returns a Stack of each of those cards*
 >
 > `stack.Get(RETURN_Card, POSITION_Val, stackOfValues, MATCH_Object)`
 >> *goes through the stack, finds the first card with one of the values in stackOfValues, and returns that card*
 >
 > `stack.Get(RETURN_Cards, POSITION_Val, stackOfValues, MATCH_Object)`
 >> *goes through the stack, finds each card with one of the values in stackOfValues, and returns a Stack of each of those cards*

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
 >>> interface{} *TODO: figure out lambda*

 ***Example Uses***
 
 > `stack.Get(RETURN_Card, POSITION_First)`
 >> *returns the first card in the Stack*
 >
 > `stack.Get(RETURN_Card, POSITION_Val, "String Value", MATCH_Object)`
 >> *goes through the stack, finds the first card with val "String Value", and returns that card*
 >
 > `stack.Get(RETURN_Card, POSITION_Key, "String Key", MATCH_Object)`
 >> *goes through the stack, finds the first card with key "String Key", and returns that card*
 >
 > `stack.Get(RETURN_Cards, POSITION_Lambda, {TODO: add})`
 >> *goes through the stack, finds each card for which the lambda expression is true, and return a stack of these cards*
 
<h4 name = "ORDER">ORDER</h4>

 This is an enum intended to make it easy to tell certain functions whether to insert a value before or after the input index.

> ***Order***
>> ORDER_Before
>>> default
>> ORDER_After

 ***Example Uses***
 
 > `stack.Add(cardToAdd, ORDER_Before, POSITION_Last)`
 >> *insert a card at the second-to-last index of the stack*
 >
 > `stack.Add(cardToAdd, ORDER_After, POSITION_Last)`
 >> *insert a card at the last index of the stack*

<h4 name = "MATCH">MATCH</h4>

 This is an enum intended to make it easy to target whether a function searching for a match between input data and data in the stack element is matching by having the same values (MATCH_Object) or the same memory address (MATCH_Reference).

 Matching by reference only works for Val, Key, and Card POSITION types.  It would not make much sense to match an index that's managed on the backend by reference (POSITION_Idx), to match a lambda expression (POSITION_Lambda), or to match using a position that's not even comparing values (POSITION_First, POSITION_Last, POSITION_All).

 Take care to note that all cases where objects are matching by reference will also be matching by object.

 > ***Match***
 >> MATCH_Object
 >>> default
 >> MATCH_Reference

 ***Example Uses***
 
 > `stack.Get(RETURN_Card, POSITION_Card, cardStructureToFind, MATCH_Object)`
 >> *returns the first card that has the same structure (key and value) as cardStructureToFind*
 >
 > stack.Get(RETURN_Card, POSITION_Card, exactCardToFind, MATCH_Reference)`
 >> *returns the first card that IS exactCardToFind as stored in memory—not just that's the same structurally*

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
 >> CONSTRUCTOR: ***FALSE***
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
 >> CONSTRUCTOR: ***FALSE***
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
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> desired card
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
 
<h2 name = "toAdd">To Add</h2>

 <h3>Generalized Functions</h3>

 * Add **Move** function

 <h3>Non-Generalized Functions</h3>

 * Add **CombineWith** function
 * Add **Flip** function
 * Add **Shuffle** function
 * Add **Clone** function
 * Add **ToArray** function
 * Add **ToStack** function

<h2 name = "footer">Footer</h1>

This project was created by Gabe Tucker.

If there are any changes or comments you would like to have made in respect to this project, please email `tucker.854@osu.edu`.  I appreciate any feedback and will usually respond within 1-2 business days.

Feel free to visit my personal pages at `https://gabetucker.com` or `https://www.linkedin.com/in/gabetucker2/`.

[Return to Glossary](#glossary)
