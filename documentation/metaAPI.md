![Banner](../images/gostack_SmallerTransparent.png)

<h1>Meta Documentation</h1>

 <img src="../images/gostack_Packages.png" width="100%" style="margin-top: 10px; margin-bottom: 10px;"/>

<h2>JDoc Documentation</h2>

 We roughly adhere to [design-by-contract principles, among other principles,](http://web.cse.ohio-state.edu/software/2221/web-sw1/extras/slides/09.Design-by-Contract.pdf) established by The Ohio State University's CSE department in writing JDoc-style documentation for our functions.

<h2>Interfaces</h2>

 Throughout this project, "interface" is used interchangeably with "any".

<h2>How to run our case tests/benchmarks</h2>

Say you wanted to double-check that our case tests of the library are working properly, or perhaps you wanted to add your own case tests for debugging purposes.  Or, perhaps you would like to run our benchmarks on your computer.  In either a case:

 1. [Clone](https://github.com/git-guides/git-clone) this repository to your PC
 2. Open a terminal

<h3>Case Tests</h3>

 1. Enter the `executive` folder of your clone via your terminal
 2. Type `go run .`
 3. See the output for each case test!

 If you'd like to configure this output, enter `executive.go`...

 <img src="../images/caseTestsRun.png" width="25%" style="margin-top: -10px;"/>

 * Passing `true` into `casetests.Run` gives a detailed terminal output, useful for debugging and finding which case test function is causing errors
 * Passing `false` into `casetests.Run` gives a concise terminal output, useful for quickly ensuring that everything is working properly

<h3>Benchmarks</h3>

 1. Enter the `benchmark` folder of your clone via your terminal
 2. Type `go test -bench .`
 3. See the output for each benchmark function test!

 <img src="../images/benchmarksRun.png" width="50%" style="margin-bot: 10px;"/>

Look at the `ns/op` to compare speeds between ***gostack*** and ***native Go***.  This tells you the average time (in nanoseconds) each function call took.

 ---

 [> Return to glossary](../README.md)