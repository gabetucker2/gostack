# goplus
 This project introduces useful base functions to Go to ensure programmer ease-of-use, concision, and flexibility.

 <h2>Scripts</h2>

 * **functions.go** is where novel functions are stored
 * **test.go** is a script used to run test cases to ensure functionality of other scripts

<h2>TODO</h2>

<h3>Functions to Add</h3>

* Add **Exists** function
* Add **Heap** functions
* Add **Error** function
* Add **bool-coalescing** operator
* Add **null-coalescing** operator
* Add lambda array sort-return function

<h3>Things to Look Into</h3>

* Test whether instance return methods work in Go

<h2>Implemented Changes</h2>

<h3>IndexOf</h3> function

`haystack.IndexOf(needle)`

**haystack** := ambigously-typed array

**needle** := ambigously-typed element

No requirements

Ensures:
* Error if invalid parameters
* Return -1 if element does not exist
* Return first element's index in haystack if element does exist
