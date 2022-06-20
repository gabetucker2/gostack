# gostack
 `gostack` introduces **Stacks**, ambiguously-typed (mathematical) strings of index/key-element pairs intended to replace arrays and maps in *golang*.  **Stacks** are introduced alongside a variety of helpful functions to ensure programmer ease-of-use, concision, and flexibility.

 By default for generics, people tend to use *golang*'s list package, but this package is optimized only with the essentials for transforming and selecting list elements.  While `gostack` offers a much wider breadth of functions for transforming and selecting elements, it also allows you to turn **Stacks** into maps, quickly convert between arrays and **Stacks**, and—most excitingly—to use functions based on lambda expression including **sort**, **TrueForAll**, and **RemoveAll**.

 Many of the functions in this project were inspired by functions from *JavaScript* Arrays or *C#* Lists.

 For the purposes of this project, we will use the imagery of a stack of cards.  A ***Stack*** will refer to a stack of cards; each element in that **Stack** will be a ***card***.

<h1>Overview</h1>

<h2>Files</h2>

 An overview of the files in this repository

 * **README.md** is this file
 * **caseTests.go** is a script used to run test cases to ensure functionality of this project's functions; for examples on how to use `gostack` functions, see this file; it is recommended to delete this file if it is not commented out at the time of your installation since it uses the main() function; in order to run test cases with ***test.go*** *not* commented out, run `go run .` in the top directory
 * **functions.go** is where novel functions are stored
 * **go.mod** is used to manage directory access
 * **stacks.go** is where the Stack struct is defined

<h2>Links</h2>

 Where post, blog, and API links relevant to this project are stored

 * https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1?view=net-6.0
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
 * http://web.cse.ohio-state.edu/software/common/doc/components/standard/Standard.html

<h2>Implemented Features</h2>
 
 <h3>Constructors</h3>
 
 * **MakeStack()**
 
 <h3>Transformers</h3>
 
 * **stack.Append()**
 * **stack.Push()**
 * **stack.Pop()**
 * **stack.Behead()**
 * **stack.Clear()**

 <h3>Receivers</h3>

 * **stack.Has()**
 * **stack.IndexOf()**
 
<h2>Unimplemented Features</h2>

 * Add **Fill** function
 * Add **Insert** function
 * Add **CombineWith** function
 * Add **Flip** function
 * Add **Entry** function
 * Add **Clone** function
 * Add **Extract** function
 * Add **ReplaceEntries** function
 * Add **GetFlip** function
 * Add **Remove** function
 * Add **Type** function
 * Add **ToArray** function
 * Add **ToStack** function
 * Add **GetKeys** function
 * Add **GetCards** function
 * Add **Sort** function
 * Add **TrueForAll** function
 * Add **RemoveAll** function
 * *...and many more*

<h1>Feature Documentation</h1>

<h2>Stack Properties</h2>

<h3>cards</h3>
 
 > `stack.cards`
 >> Returns an interface array to represent the elements in the Stack
 >
 >> It is highly recommended against accessing this property, as the entire purpose of this project is for you not to have to manage arrays manually
 
<h3>size</h3>
 
 > `stack.size`
 >> Returns the cardinality of the given Stack

<h2>Stack Functions</h2>

CONSTRUCTOR means the function requires no receiver.  !CONSTRUCTOR means the function doesn't have a receiver or, if it does, it is an existing **Stack** object.

TRANSFORMER means the function updates the inputted **Stack**.

RECEIVER means the function returns a value.

Searching with browser utilities (e.g., `ctrl+f`) may be useful in this section.

<h3>MakeStack</h3>

 > `MakeStack()`
 >> CONSTRUCTOR: ***TRUE***
 >
 >> TRANSFORMER: ***FALSE***
 >
 >> RECEIVER: ***TRUE***
 
 > ***Pseudocode***
 >> return new Stack
 
<h3>Clear</h3>
 
 > `stack.Clear()`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> TRANSFORMER: ***TRUE***
 >
 >> RECEIVER: ***TRUE***
 
 > ***Parameters***
 >> **stack** is the Stack to clear
 
 > ***Pseudocode***
 >> removes all cards in the Stack
 >
 >> return the empty stack
 
<h3>Append</h3>
 
 > `stack.Append(card)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> TRANSFORMER: ***TRUE***
 >
 >> RECEIVER: ***TRUE***
 
 > ***Parameters***
 >> **stack** is the Stack to append
 >
 >> **card** is the ambiguously-typed element to add to the end of the Stack
 
 > ***Pseudocode***
 >> add card to i = stack.size of the Stack
 >
 >> return updated Stack
 
<h3>Push</h3>
 
 > `stack.Push(card)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> TRANSFORMER: ***TRUE***
 >
 >> RECEIVER: ***TRUE***
 
 > ***Parameters***
 >> **stack** is the Stack to push
 >
 >> **card** is the ambiguously-typed element to add to the beginning of the Stack
 
 > ***Pseudocode***
 >> add card to i = 0 of the Stack
 >
 >> **FOR EACH CARD THAT ALREADY EXISTED IN THE STACK**
 >>> that card's previous index i is updated to i + 1
 >
 >> return updated Stack
 
<h3>Pop</h3>
 
 > `stack.Pop()`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> TRANSFORMER: ***TRUE***
 >
 >> RECEIVER: ***TRUE***
 
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
 
<h3>Behead</h3>
 
 > `stack.Behead()`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> TRANSFORMER: ***TRUE***
 >
 >> RECEIVER: ***TRUE***
 
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
 
<h3>Has</h3>
 
 > `stack.Has(card)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> TRANSFORMER: ***FALSE***
 >
 >> RECEIVER: ***TRUE***
 
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

<h3>IndexOf</h3>
 
 > `stack.IndexOf(card)`
 >> CONSTRUCTOR: ***FALSE***
 >
 >> TRANSFORMER: ***FALSE***
 >
 >> RECEIVER: ***TRUE***
 
 > ***Parameters***
 >> **stack** is the Stack to search
 >
 >> **card** is the ambigously-typed element whose index to find
 
 > ***Pseudocode***
 >> **IF CARD IS IN STACK**
 >>> return card index [0, stack.len)
 >
 >> **ELSE**
 >>> return -1
 
<h1>Footer</h1>

This project was created by Gabe Tucker.

If there are any changes or comments you would like to have made in respect to this project, please email `tucker.854@osu.edu`.  I will usually respond in 1-2 business days.

Feel free to visit my personal pages at `https://gabetucker.com` or `https://www.linkedin.com/in/gabetucker2/`.
