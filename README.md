# gostack
 `gostack` introduces **Stacks**, ambiguously-typed strings of index/key-element pairs intended to replace arrays and maps in golang.  **Stacks** are introduced alongside a variety of helpful base functions to ensure programmer ease-of-use, concision, and flexibility.

 Many of the functions in this project were inspired by functions from *JavaScript* Arrays or *C#* Lists.

<h1>Implemented Changes</h1>

<h2>Properties</h2>

<h3>vals</h3>

> `stack.vals`
>> Returns an interface array to represent the elements in the Stack
>
>> It is highly recommended against accessing this property, as the entire purpose of this project is for you to avoid having to manage this property manually as opposed to by using our functions

<h3>len</h3>

> `stack.len`
>> Returns the cardinality (amount of elements) of the given Stack

<h2>Functions</h2>

<h3>MakeStack</h3>

> `gs_MakeStack()`
>> GETS: ***TRUE***
>
>> SETS: ***TRUE***

> ***Pseudocode***
>> return new Stack

<h3>IndexOf</h3>

> `gostack_IndexOf(needle, haystack)`
>> GETS: ***TRUE***
>
>> SETS: ***FALSE***

> ***Parameters***
>> **needle** is an ambigously-typed element
>
>> **haystack** is a Stack

> ***Pseudocode***
>> **IF NEEDLE IS IN HAYSTACK**
>>> return needle index [0, |haystack|)
>
>> **ELSE**
>>> return -1

<h1>Overview</h1>

 <h2>Files</h2>

 An overview of the files in this repository

 * **README.md** is this file
 * **stacks.go** is where the Stack struct is defined
 * **functions.go** is where novel functions are stored
 * **test.go** is a script used to run test cases to ensure functionality of this project's functions; for examples on how to use `gostack` functions, see this file

 <h2>Links</h2>

 Where post, blog, and API links which are relevant to this project are stored

 * https://stackoverflow.com/questions/40145569/how-do-you-make-a-function-accept-multiple-types
 * https://medium.com/geekculture/generics-in-go-5a36b1f978bc
 * https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1?view=net-6.0
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
 * http://web.cse.ohio-state.edu/software/common/doc/components/standard/Standard.html

<h2>Functions to be Added</h2>

* Add **MakeStack** function
* Add **Fill** function
* Add **Insert** function
* Add **Flip** function
* Add **Entry** function
* Add **Extract** function
* Add **ReplaceEntries** function
* Add **GetFlip** function
* Add **Remove** function
* Add **Pop** function
* Add **Push** function
* Add **Append** function
* Add **Behead** function
* Add **InterfaceType** function
* Add **GetKeys** function
* Add **GetVals** function
* Add **Has** function
* Add **Clear** function
* Add lambda array sort-return function like for Lists in C#
