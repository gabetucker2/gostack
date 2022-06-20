# gostack
 `gostack` is a project introducing **Stacks**, ambiguously-typed strings of index/key-element pairs intended to replace arrays and maps in golang.  **Stacks** are introduced alongside a variety of helpful base functions to ensure programmer ease-of-use, concision, and flexibility.

 Many of the functions in this project were inspired by JavaScript or C# List functions.

<h1>Implemented Changes</h1>

<h2>Functions</h2>

<h3>IndexOf</h3>

> `haystack.IndexOf(needle)`

> Parameters:
> * **haystack** is an ambigously-typed array
> * **needle** is an ambigously-typed element

> Ensures:
> * Error if there exists invalid arguments
> * Return -1 if element does not exist
> * Return first element's index in haystack if element does exist

<h1>Overview</h1>

 <h2>Scripts</h2>

 * **functions.go** is where novel functions are stored
 * **test.go** is a script used to run test cases to ensure functionality of other scripts

<h2>TODO</h2>

<h3>Functions to Add</h3>

* Add **InterfaceType** function
* Add **GetKeys** function
* Add **GetVals** function
* Add **Exists** function
* Add lambda array sort-return function like for Lists in C#
