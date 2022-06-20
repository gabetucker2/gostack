# gostack
 `gostack` introduces **Stacks**, ambiguously-typed strings of index/key-element pairs intended to replace arrays and maps in golang.  **Stacks** are introduced alongside a variety of helpful base functions to ensure programmer ease-of-use, concision, and flexibility.

 Many of the functions in this project were inspired by functions from *JavaScript* Arrays or *C#* Lists.

<h1>Implemented Changes</h1>

<h2>Functions</h2>

<h3>IndexOf</h3>

> `gostack.IndexOf(needle, haystack)`
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

 * **README.md** is this file with an overview of this project
 * **stacks.go** is where 
 * **functions.go** is where novel functions are stored
 * **test.go** is a script used to run test cases to ensure functionality of this project's functions; for examples on how to use `gostack` functions, see this file

 <h2>Links</h2>

 Where post, blog, and API links which are relevant to this project are stored

 * https://stackoverflow.com/questions/40145569/how-do-you-make-a-function-accept-multiple-types
 * https://medium.com/geekculture/generics-in-go-5a36b1f978bc
 * https://docs.microsoft.com/en-us/dotnet/api/system.collections.generic.list-1?view=net-6.0
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array

<h2>TODO</h2>

<h3>Functions to be Added</h3>

* Add **MakeStack** function
* Add **InterfaceType** function
* Add **GetKeys** function
* Add **GetVals** function
* Add **Exists** function
* Add lambda array sort-return function like for Lists in C#
