# gostack
 `gostack` introduces **Stacks**, ambiguously-typed sets of elements intended to replace arrays and maps in *golang*.  **Stacks** are introduced alongside a variety of helpful functions to ensure programmer ease-of-use, concision, and flexibility.

 For the purposes of this project, we will use the imagery of a stack of cards.  A ***Stack*** will refer to a stack of cards; each element in that stack will be a ***Card***.  This is not to be confused with traditional stack structures (which only push and pop the first element in a stack).

 By default for generics, people tend to use *golang*'s list package, but this package is optimized only with the essentials for transforming and selecting list elements.  While `gostack` offers a much wider breadth of functions for transforming and selecting elements, it also allows you to turn **Stacks** into maps, quickly convert between arrays and **Stacks**, and—most excitingly—to use functions based on lambda expression including **sort**, **TrueForAll**, and **RemoveAll**.

 Many of the functions in this project were inspired by functions from *JavaScript* Arrays or *C#* Lists.

<h1>Overview</h1>

<h2>Files</h2>

 An overview of the files in this repository

 * **README.md** is this file
 * **TODO.txt** is a file with features to be added (significant only to gostack developers)
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

 > ***name*** means it is conventionally acceptable to access this manually
 >
 > **name** means it is highly recommended against accessing this manually; you should instead use our functions

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
 > **Slice**
 >> **startIdx** *int*
 >
 >> **endIdx** *int*

 Position enum:
 > ***Position*** *[enum]*
 >>
 >> Position_First
 >>> *NONE*
 >>
 >> Position_Last
 >>> *NONE*
 >>
 >> Position_Card
 >>> Card
 >>
 >> Position_Idx
 >>> int
 >>
 >> Position_Val
 >>> any type
 >>
 >> Position_Key
 >>> any type
 >>
 >> Position_Slice
 >>> Slice
 >>
 >> Position_All
 >>> *NONE*
 
 <h3>Constructor Functions</h3>

 * **MakeStack()**
 
 <h3>Other Functions</h3>

 * **stack.Add(card, Position_*, _idxData)**
 * **stack.Extract(Position_*, _idxData)**
 * **stack.Replace(Position_*, _idxData)**
 * **stack.Empty()**
 * **stack.Has(Position_*, _idxData)**
 * **stack.Index(Position_*, _idxData)**

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

 > `slice Slice{}`
 >> `slice.startIdx`
 >>> Returns the first index in the desired slice
 >> `slice.endIdx`
 >>> Returns the last index of the desired slice

 <h4>Recommended Uses</h4>
 
 * *None*

<h3>Position</h3>

 This is an enum intended to make it easy to inform functions of the intended target cards.

 Take care to note that some functions do not support certain enum types (supported enum types are documented in function API).  For instance, it wouldn't make sense for you to call `stack.Index()` on a set of cards interspersed throughout a stack, but it would make sense for you to call `stack.Extract()` on set of cards interspersed throughout a stack.

 > ***Position*** *[enum]*
 >> *SampleConstant*
 >>> *The type of the variable (called posData) that needs to be passed into the function utilizing this constant*
 >>> *For instance, if you input POSITION_SLICE, you would need to pass a Slice struct to your posData parameter*
 >>
 >> Position_First
 >>> *NONE*
 >>
 >> Position_Last
 >>> *NONE*
 >>
 >> Position_Card
 >>> Card
 >>
 >> Position_Idx
 >>> int
 >>
 >> Position_Val
 >>> any type
 >>
 >> Position_Key
 >>> any type
 >>
 >> Position_Slice
 >>> Slice
 >>
 >> Position_All
 >>> *NONE*

 <h4>Recommended Uses</h4>
 
 * `stack.Add(card, Position_First)`
 * `stack.Replace(newCard, Position_Key, "This string represents the key of one or multiple cards to target in the replace function")`
 * `stack.Has(Position_Card, card)`
 * `stack.Extract(Position_All)`
 * *...and so on*

<h2>Stack Functions</h2>

 CONSTRUCTOR means the function requires no receiver (i.e., don't need the `thing.` in `thing.function()`); !CONSTRUCTOR means the function's receiver is an existing **Stack** object.

 GETTER means the function returns a value.

 SETTER means the function updates the inputted **Stack**.

 Searching with browser utilities (e.g., `ctrl+f`) may be useful in this section.

<h3>MakeStack</h3>

 > `MakeStack()`
 >> CONSTRUCTOR: ***TRUE***
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***FALSE***
 
 > ***Pseudocode***
 >> return new Stack
 
<h3>Empty</h3>
 
 > `stack.Empty()`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***TRUE***
 
 > ***Parameters***
 >> **stack** is the Stack to Empty
 
 > ***Pseudocode***
 >> removes all cards in the Stack
 >
 >> return the empty stack
 
<h3>AddFirst</h3>
 
 > `stack.AddFirst(card)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***TRUE***
 
 > ***Parameters***
 >> **stack**
 >
 >> **card** is the ambiguously-typed element to add to the beginning of the Stack
 
 > ***Pseudocode***
 >> add card to i = 0 of the Stack
 >
 >> **FOR EACH CARD THAT ALREADY EXISTED IN THE STACK**
 >>> that card's previous index i is updated to i + 1
 >
 >> return updated Stack
 
<h3>AddLast</h3>
 
 > `stack.AddLast(card)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***TRUE***
 
 > ***Parameters***
 >> **stack**
 >
 >> **card** is the ambiguously-typed element to add to the end of the Stack
 
 > ***Pseudocode***
 >> add card to i = stack.size of the Stack
 >
 >> return updated Stack
 
<h3>ExtractFirst</h3>
 
 > `stack.ExtractFirst()`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***TRUE***
 
 > ***Parameters***
 >> **stack** is the Stack from which to remove the first card
 
 > ***Pseudocode***
 >> **IF STACK IS NOT EMPTY**
 >>> remove the first card from the stack
 >>
 >>> return the removed card
 >
 >> **ELSE**
 >>> return nil
 
<h3>ExtractLast</h3>
 
 > `stack.ExtractLast()`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***TRUE***
 
 > ***Parameters***
 >> **stack** is the Stack from which to remove the last card
 
 > ***Pseudocode***
 >> **IF STACK IS NOT EMPTY**
 >>> remove the last card from the stack
 >>
 >>> return the removed card
 >
 >> **ELSE**
 >>> return nil
 
<h3>Has</h3>
 
 > `stack.Has(card)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***FALSE***
 
 > ***Parameters***
 >> **stack** is the Stack to search
 >
 >> **card** is the ambigously-typed element whom to check if exists
 
 > ***Pseudocode***
 >> **IF CARD IS IN STACK**
 >>> return true
 >
 >> **ELSE**
 >>> return false

<h3>IndexCard</h3>
 
 > `stack.IndexCard(card)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> GETTER: ***TRUE***
 >
 >> SETTER: ***FALSE***
 
 > ***Parameters***
 >> **stack** is the Stack to search
 >
 >> **card** is the ambigously-typed element whose index to find
 
 > ***Pseudocode***
 >> **IF CARD IS IN STACK**
 >>> return card index [0, stack.size)
 >
 >> **ELSE**
 >>> return -1
 
<h1>Unimplemented Features</h1>

 <h2>General Features</h2>

 * Add **Fill** function
 * Add **CombineWith** function
 * Add **Flip** function
 * Add **Shuffle** function
 * Add **Entry** function
 * Add **Clone** function
 * Add **GetFlip** function
 * Add **Type** function
 * Add **ToArray** function
 * Add **ToStack** function
 * Add **Get** function

 <h2>Lambda Function Support</h2>

 * Add **GetCards** function
 * Add **Sort** function
 * Add **TrueForAll** function

 <h2>Tensor Function Support</h2>

<h1>Footer</h1>

This project was created by Gabe Tucker.

If there are any changes or comments you would like to have made in respect to this project, please email `tucker.854@osu.edu`.  I appreciate any feedback and will usually respond within 1-2 business days.

Feel free to visit my personal pages at `https://gabetucker.com` or `https://www.linkedin.com/in/gabetucker2/`.
