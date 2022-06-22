# gostack
 `gostack` introduces **Stacks**, ambiguously-typed sets of elements intended to replace arrays and maps in *golang*.  **Stacks** are introduced alongside a variety of helpful functions to ensure programmer ease-of-use, concision, and flexibility.

 For the purposes of this project, we will use the imagery of a stack of cards.  A ***Stack*** will refer to a stack of cards; each element in that stack will be a ***Card***.  This is not to be confused with traditional stack structures (which only push and pop the first element in a stack).

 By default for generics, people tend to use *golang*'s list package, but this package is optimized only with the essentials for transforming and selecting list elements.  While `gostack` offers a much wider breadth of functions for transforming and selecting elements, it also allows you to turn **Stacks** into maps, quickly convert between arrays and **Stacks**, and—most excitingly—to use functions based on lambda expression including **sort**, **TrueForAll**, and **RemoveAll**.

 Many of the functions in this project were inspired by functions from *JavaScript* Arrays or *C#* Lists.

<h1>Overview</h1>

<h2>Files</h2>

 An overview of the files in this repository

 * **README.md** is this file
 * **TODO.txt** is a file with features to be added (significant only to `gostack` developers)
 * **caseTests.go** is a script used to run test cases to ensure functionality of this project's functions; for examples on how to use `gostack` functions, see this file; it is recommended to delete this file if it is not commented out at the time of your installation since it uses the main() function; in order to run test cases with ***test.go*** *not* commented out, run `go run .` in the top directory
 * **functions.go** is where novel functions are stored
 * **go.mod** is used to manage directory access
 * **structs.go** is where structs are defined

<h2>Links</h2>

 Where API links which inspired this project are stored

 * https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1?view=net-6.0
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
 * http://web.cse.ohio-state.edu/software/common/doc/components/standard/Standard.html

<h2>Brief Documentation</h2>

 <h3>Data Structures</h3>

 > ***name*** means it is conventionally acceptable to access this value manually
 >
 > **name** means it is highly recommended against accessing  value manually; you should instead use our functions

 Stack struct:
 > ***stack*** *Stack*
 >> **cards** *[]\*Card*
 >
 >> ***size*** *int*

 Card struct:
 > **card** *Card*
 >> **key** *any type*
 >
 >> **val** *any type*
 
 Slice struct:
 > ***Slice***
 >> ***startIdx*** *int*
 >
 >> ***endIdx*** *int*

 POSITION enum:
 > ***POSITION*** *[enum]*
 >>
 >> POSITION_First
 >>> *NONE*
 >>
 >> POSITION_Last
 >>> *NONE*
 >>
 >> POSITION_Card
 >>> Card
 >>
 >> POSITION_Idx
 >>> int
 >>
 >> POSITION_Val
 >>> any type
 >>
 >> POSITION_Key
 >>> any type
 >>
 >> POSITION_Slice
 >>> Slice
 >>
 >> POSITION_All
 >>> *NONE*
 
 <h3>Constructor Functions</h3>

 * **MakeStack()**
 
 <h3>Other Functions</h3>

 * **stack.Add(card, POSITION_*, _idxData)**
 * **stack.Extract(POSITION_*, _idxData)**
 * **stack.Replace(POSITION_*, _idxData)**
 * **stack.Has(POSITION_*, _idxData)**
 * **stack.Index(POSITION_*, _idxData)**
 * **stack.Empty()**

<h1>Exhaustive Documentation</h1>

<h2>Data Structures</h2>

 It is highly advised against using these data structures for reasons other than those listed in *Recommended Uses* section.  The entire purpose of this project is for you not to have to manage arrays manually, but documentation for objects intended to be hidden still exists for those who would like to add their own Stack functions

<h3>Stack</h3>

 This is the main struct in the project.

 > `stack` *Stack{}*
 >> `cards` *[]\*Card{}*
 >>> Returns an interface array to represent the elements in the Stack
 >> `size` *int*
 >>> Returns the cardinality (i.e., `len(stack.cards)`) of this Stack

 <h4>Recommended Uses</h4>

 * `stack.size` *instead of `len(stack.cards)`*

<h3>Card</h3>

 This is a struct for our elements/maps within stacks.

 >> `card` *Card{}*
 >>> `card.key` *any type [interface{}]*
 >>>> Returns a key for this card-map (or nil if doesn't exist)
 >>> `card.val` *any type [interface{}]*
 >>>> Returns the val of this card (or nil if doesn't exist)

 <h4>Recommended Uses</h4>
 
 * *None*

<h3>Slice</h3>

 This is a struct that makes it easier to pass two int values between functions on the backend.

 > `slice` *Slice{}*
 >> `slice.startIdx`
 >>> Returns the first index in the desired slice
 >> `slice.endIdx`
 >>> Returns the last index of the desired slice

 <h4>Recommended Uses</h4>
 
 * `something = slice.startIdx`
 * `something = slice.endIdx`

<h3>POSITION</h3>

 This is an enum intended to make it easy to inform functions of the intended target cards.

 Take care to note that some functions do not support certain enum types (supported enum types are documented in function API).  For instance, it wouldn't make sense for you to call `stack.Index()` on a set of cards interspersed throughout a stack, but it would make sense for you to call `stack.Extract()` on set of cards interspersed throughout a stack.

 > ***POSITION*** *[enum]*
 >> *POSITION_\* Sample*
 >>> *The type of the variable (called `posData`) that needs to be passed into the function utilizing this constant*
 >>> *For instance, if you input `POSITION_Slice`, you would need to pass a **Slice** struct to your `posData` parameter*
 >>
 >> POSITION_First
 >>> *NONE*
 >>
 >> POSITION_Last
 >>> *NONE*
 >>
 >> POSITION_Card
 >>> Card
 >>
 >> POSITION_Idx
 >>> int
 >>
 >> POSITION_Val
 >>> any type
 >>
 >> POSITION_Key
 >>> any type
 >>
 >> POSITION_Slice
 >>> Slice
 >>
 >> POSITION_All
 >>> *NONE*

 <h4>Recommended Uses</h4>
 
 * `stack.Add(card, true, POSITION_First)`
 * `stack.Replace(newCard, POSITION_Key, "This string represents the key of one or multiple cards to target in this Replace function")`
 * `stack.Has(POSITION_Card, card)`
 * `stack.Extract(POSITION_All)`
 * *...and so on*

<h2>Stack Functions</h2>

 Searching with browser utilities (e.g., `ctrl+f`) may be useful in this section.
 
<h3>_NotationSample</h3>
 
 > `variable1.function(variable2, THING_*, ...optional)`
 >> CONSTRUCTOR: ***TRUE***
 >>>> means the function requires no receiver (i.e., our sample `variable1` should not exist in this function call)
 >>
 >>> **variable1** is the variable we're constructing
 >>
 >> CONSTRUCTOR: ***FALSE***
 >>>> means the function's receiver is an existing stack (i.e., our sample `variable1` must exist in this function call)
 >
 >> GETTER: ***TRUE***
 >>>> means the function returns a value
 >>
 >>> **variable1** is the variable we're getting
 >
 >> SETTER: ***TRUE***
 >>>> means the function updates the inputted stack
 >>
 >>> **variable1** is a variable we're setting
 >>> **variable2** is a variable we're setting
 
 > ***Parameters***
 >> **variable1** *type* is the receiver for the function
 >
 >> **variable2** *type* is the first argument for the function
 >
 >> **THING_\*** *type* refers to how this input argument can be any variable starting with `THING_` that the function specifies is allowed
 >
 >> **...optional** *type* refers to how this input argument does not have to be inputted in the function (refer to documentation to decide whether to input)
 >>> A sample instance where you would not input an argument in this spot is when you're using POSITION_First, which does not intake any posData.  That said, take care not to input more than 1 argument to optional parameters; everything will compile if you do, but this action is not supported by `gostack`.

 > ***Supported POSITIONS***
 >
 > Each of the below positions are supported POSITION_* arguments
 >
 >> POSITION_First
 >
 >> POSITION_Last
 >
 >> POSITION_Card
 >
 >> POSITION_Idx
 >
 >> POSITION_Val
 >
 >> POSITION_Key
 >
 >> POSITION_Slice
 >
 >> POSITION_All
 
 > ***Pseudocode***
 >> This section outlines what the function does in simplistic terms
 >
 >> When pseudocode says a Stack's cards are updated, it is implied that stack.size is updated correspondingly

<h3>MakeStack</h3>

 > `MakeStack()`
 >> CONSTRUCTOR: ***TRUE***
 >>> Stack
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***FALSE***
 
 > ***Pseudocode***
 >> return a new Stack
 
<h3>Add</h3>
 
 > `stack.Add(toAdd, beforeNotAfter, POSITION_*, ...posData)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> **stack**
 >
 >> SETTER: ***TRUE***
 >>> **stack**
 
 > ***Parameters***
 >> **stack** *Stack*
 >
 >> **toAdd** *Card* or *Stack* is either a Card or a Stack of cards to insert at POSITION
 >
 >> **beforeNotAfter** *bool* is used to control whether **card** is added before or after the position
 >
 >> **POSITION_\*** *POSITION* is used to provide the function relevant POSITION data to find the correct position
 >
 >> **...posData** *any type [interface{}]* is used to provide the function relevant additional data to find the correct position

 > ***Supported POSITIONS***
 >> POSITION_First
 >
 >> POSITION_Last
 >
 >> POSITION_Card
 >
 >> POSITION_Idx
 >
 >> POSITION_Val
 >
 >> POSITION_Key
 
 > ***Pseudocode***
 >> **IF VALID POSITION**
 >>> **IF beforeNotAfter**
 >>>> add card before POSITION in the stack
 >>>
 >>> **ELSE**
 >>>> add card after POSITION in the stack
 >>>
 >>> **FOR EACH CARD THAT ALREADY EXISTED IN THE STACK**
 >>>> that card's previous index i is updated to i + 1
 >>>
 >>> return updated stack
 >>
 >> **ELSE**
 >>> return nil
 
<h3>Extract</h3>
 
 > `stack.Extract(POSITION_*, ...posData)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> extracted card *or* nil
 >
 >> SETTER: ***TRUE***
 >>> **stack**
 
 > ***Parameters***
 >> **stack** *Stack* is the Stack from which to remove the first card
 >
 >> **POSITION_\*** *POSITION* is used to provide the function relevant POSITION data to find the correct card to extract
 >
 >> **...posData** *any type [interface{}]* is used to provide the function relevant additional data to find the correct card to extract

 > ***Supported POSITIONS***
 >
 >> POSITION_First
 >
 >> POSITION_Last
 >
 >> POSITION_Card
 >
 >> POSITION_Idx
 >
 >> POSITION_Val
 >
 >> POSITION_Key
 >
 >> POSITION_Slice
 >
 >> POSITION_All
 
 > ***Pseudocode***
 >> **IF STACK IS NOT EMPTY**
 >>> remove cards from the stack based on provided POSITION data
 >>
 >>> return the removed card(s)
 >
 >> **ELSE**
 >>> return nil
 
<h3>Replace</h3>
 
 > `stack.Replace(toInsert, POSITION_*, ...posData)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> card that was replaced *or* nil
 >
 >> SETTER: ***TRUE***
 >>> **stack**
 
 > ***Parameters***
 >> **stack** *Stack* is the Stack from which to remove the first card
 >
 >> **toInsert** *Card* or *Stack* is either a Card or a Stack of cards to insert at POSITION(S) as the replacement
 >
 >> **POSITION_\*** *POSITION* is used to provide the function relevant POSITION data to find the correct card to replace
 >
 >> **...posData** *any type [interface{}]* is used to provide the function relevant additional data to find the correct card to replace

 > ***Supported POSITIONS***
 >
 >> POSITION_First
 >
 >> POSITION_Last
 >
 >> POSITION_Card
 >
 >> POSITION_Idx
 >
 >> POSITION_Val
 >
 >> POSITION_Key
 >
 >> POSITION_Slice
 >
 >> POSITION_All
 
 > ***Pseudocode***
 >> **IF STACK IS NOT EMPTY**
 >>> replace cards from the stack with toInsert based on provided POSITION data
 >>
 >>> return the removed card(s) in a new stack
 >
 >> **ELSE**
 >>> return nil
 
<h3>Has</h3>
 
 > `stack.Has(lookFor)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> bool
 >
 >> SETTER: ***FALSE***
 
 > ***Parameters***
 >> **stack** *Stack* is the Stack to search for **lookFor**
 >
 >> **lookFor** *Card* or *Stack* is either a Card or a Stack of cards to find in **stack**
 
 > ***Pseudocode***
 >> **IF lookFor IS IN STACK**
 >>> return true
 >
 >> **ELSE**
 >>> return false

<h3>Index</h3>
 
 > `stack.Index(lookFor)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> index or Slice (interface{}) of **lookFor** in **stack**
 >
 >> SETTER: ***FALSE***
 
 > ***Parameters***
 >> **stack** *Stack* is the Stack to search for **lookFor**
 >
 >> **lookFor** *Card* or *Stack* is either a Card or a Stack of cards to find in **stack**
 
 > ***Pseudocode***
 >> **IF lookFor IS IN STACK**
 >>> **IF lookFor IS A CARD**
 >>>> return lookFor's index in **stack**
 >>>
 >>> **ELSE IF lookFor IS A STACK**
 >>>> return a Slice representing where lookFor starts and ends in **stack**
 >
 >> **ELSE**
 >>> **IF lookFor IS A CARD**
 >>>> return -1
 >>>
 >>> **ELSE IF lookFor IS A STACK**
 >>>> return Slice with values {-1, -1}
 
<h3>Empty</h3>
 
 > `stack.Empty()`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >>> **stack**
 >
 >> SETTER: ***TRUE***
 >>> **stack**
 
 > ***Parameters***
 >> **stack** is the Stack to by emptied
 
 > ***Pseudocode***
 >> removes all cards in the stack
 >
 >> return the empty stack
 
<h1>Unimplemented Features</h1>

 <h2>General Functions</h2>

 * Add **Get** function
 * Add **Fill** function
 * Add **CombineWith** function
 * Add **Flip** function
 * Add **Shuffle** function
 * Add **Clone** function
 * Add **GetFlip** function
 * Add **Type** function
 * Add **ToArray** function
 * Add **ToStack** function

 <h2>Lambda Functions</h2>

 * Add **GetCards** function
 * Add **Sort** function
 * Add **TrueForAll** function

 <h2>Tensor Support</h2>

<h1>Footer</h1>

This project was created by Gabe Tucker.

If there are any changes or comments you would like to have made in respect to this project, please email `tucker.854@osu.edu`.  I appreciate any feedback and will usually respond within 1-2 business days.

Feel free to visit my personal pages at `https://gabetucker.com` or `https://www.linkedin.com/in/gabetucker2/`.
