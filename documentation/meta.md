![Banner](../images/gostack_SmallerTransparent.png)

<h1>Meta Documentation</h1>

<h2>How to run our case tests</h2>

Say you wanted to double-check that our case tests of the library are working properly, or perhaps you wanted to add your own case tests for debugging purposes.  In such a case:

 1. [Clone](https://github.com/git-guides/git-clone) this repository to your PC
 2. Open a terminal
 3. Enter the `executive` folder of your clone via your terminal
 4. Type `go run .`
 5. See the output for each case test!

 If you'd like to configure this output, enter `executive.go`...

 <img src="../images/caseTestsRun.png" width="25%" style="margin-top: -10px;"/>

 * Passing `true` into `casetests.Run` gives a detailed terminal output, useful for debugging and finding which case test function is causing errors
 * Passing `false` into `casetests.Run` gives a concise terminal output, useful for quickly ensuring that everything is working properly

---

 [> Return to glossary](../README.md)