# gostack
 `gostack` introduces **Stacks**, ambiguously-typed strings of index/key-element pairs intended to replace arrays and maps in golang.  **Stacks** are introduced alongside a variety of helpful base functions to ensure programmer ease-of-use, concision, and flexibility.

 Many of the functions in this project were inspired by functions from *JavaScript* Arrays or *C#* Lists.

<h1>Overview</h1>

 <h2>Files</h2>

 An overview of the files in this repository

 * **README.md** is this file
 * **stacks.go** is where the Stack struct is defined
 * **functions.go** is where novel functions are stored
 * **test.go** is a script used to run test cases to ensure functionality of this project's functions; for examples on how to use `gostack` functions, see this file; it is recommended to delete this file if it is not commented out at the time of your installation since it uses the main() function; in order to run test cases with ***test.go*** *not* commented out, run `go run .` in the top directory
 * **go.mod** is used to initialize the project to its respective git repository

 <h2>Links</h2>

 Where post, blog, and API links which are relevant to this project are stored

 * https://stackoverflow.com/questions/40145569/how-do-you-make-a-function-accept-multiple-types
 * https://medium.com/geekculture/generics-in-go-5a36b1f978bc
 * https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1?view=net-6.0
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
 * http://web.cse.ohio-state.edu/software/common/doc/components/standard/Standard.html

 <h2>Implemented Features</h2>
 
 * **MakeStack** function
 * **stack.Push** function
 * **stack.IndexOf** function
 
 <h2>Unimplemented Features</h2>

 * Add **Fill** function
 * Add **Insert** function
 * Add **Flip** function
 * Add **Entry** function
 * Add **Extract** function
 * Add **ReplaceEntries** function
 * Add **GetFlip** function
 * Add **Remove** function
 * Add **Pop** function
 * Add **Append** function
 * Add **Behead** function
 * Add **InterfaceType** function
 * Add **GetKeys** function
 * Add **GetVals** function
 * Add **Has** function
 * Add **Clear** function
 * Add lambda array sort-return function like for Lists in C#

<h2>Footer</h2>

This project was created by Gabe Tucker.

If there are any changes or comments you would like to have made in respect to this project, please email `tucker.854@osu.edu`; I will usually respond in 1-2 business days.

Feel free to learn more about me at `https://gabetucker.com` or `https://www.linkedin.com/in/gabetucker2/`!

<h1>Implemented Features</h1>

<h2>Stack Properties</h2>

<h3>vals</h3>

> `stack.vals`
>> Returns an interface array to represent the elements in the Stack
>
>> It is highly recommended against accessing this property, as the entire purpose of this project is for you not to have to manage arrays manually

<h3>len</h3>

> `stack.len`
>> Returns the cardinality of the given Stack

<h2>Stack Functions</h2>

 Constructor means the function receiver is the struct (`Stack`) itself
 Non-constructor means the function doesn't have a receiver or, if it does, it is an existing `Stack` object

<h3>MakeStack</h3>

> `MakeStack()`
>> CONSTRUCTOR: ***TRUE***
>
>> GETS: ***TRUE***
>
>> UPDATES: ***FALSE***

> ***Pseudocode***
>> return new Stack

<h3>Push</h3>

> `stack.Push(element)`
>> CONSTRUCTOR: ***FALSE***
>
>> GETS: ***TRUE***
>
>> UPDATES: ***TRUE***

> ***Parameters***
>> **stack** is the Stack on which to push an element
>
>> **element** is the ambiguously-typed element to add to the beginning of the stack

> ***Pseudocode***
>> add an element to i = 0 of the stack
>
>> **FOR EACH ELEMENT THAT ALREADY EXISTED IN THE STACK**
>>> that element's previous index i is updated to i + 1
>
>> return updated Stack

<h3>IndexOf</h3>

> `stack.IndexOf(element)`
>> CONSTRUCTOR: ***FALSE***
>
>> GETS: ***TRUE***
>
>> UPDATES: ***FALSE***

> ***Parameters***
>> **stack** is the Stack through which to search
>
>> **element** is the ambigously-typed element of which to find the index

> ***Pseudocode***
>> **IF NEEDLE IS IN HAYSTACK**
>>> return needle index [0, |haystack|)
>
>> **ELSE**
>>> return -1
