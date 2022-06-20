# gostack
 `gostack` introduces **Stacks**, ambiguously-typed strings of index/key-element pairs intended to replace arrays and maps in golang.  **Stacks** are introduced alongside a variety of helpful base functions to ensure programmer ease-of-use, concision, and flexibility.

 Many of the functions in this project were inspired by functions from *JavaScript* Arrays or *C#* Lists.

<h1>Implemented Changes</h1>

<h2>Functions</h2>

<h3>IndexOf</h3>

> `gostack.IndexOf(needle, haystack)`
>
> GETS: ***TRUE***
>
> SETS: ***FALSE***

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
