 <h1 name = "preface">PREFACE</h1>

![Banner](images/gostack_Smaller.png)

 *"The purpose of abstraction is not to be vague, but to create a new semantic level in which one can be absolutely precise." - Edsger W. Dijkstra*

 <h1 name = "introduction">Introduction</h1>

 Introducing **Stacks**—sets of **Card** elements (like a stack of cards)—***gostack*** serves as an all-in-one library for flexible, parsimonious, and elegant scripting in ***Go***.

 ***Go***, despite its elegance, is often excruciating.  We often find ourselves writing tedious code to complete seemingly simple tasks, adding an extra layer of complexity that distracts from the actual goal.  ***gostack*** solves this issue by creating a novel framework for data management, abstracting away this layer of complexity with English-like code so that you can focus more on what matters.

 With arrays, array matrices, maps, map matrices, and interfaces to generalize your code, we felt that this system could use a redesign:

<img src="images/gostack_StackAndCard.png" width="50%" style="margin-top: 20px;margin-left: 25%;"/>

 ***gostack***'s stacks...
 * ...replace arrays and maps, and matrices, eliminating the need for translating data between varying data types while supporting smooth conversion between stacks and your existing data structures
 * ...offer the minimum functions needed for unlimited flexibility, allowing the user to seamlessly write what would previously have been a verbose monstrosity of 5 nested for-loops in a single, yet concise, line
 * ...allow the user to get and set based on reference or object with ease, preventing the user from having to worry about convoluted pointer/address management
 * ...support the treatment of stacks as matrices, allowing the user to easily manage data tables, perform linear algebra, and control deep-versus-shallow matrix operations
 * ..., even when our built-in functions aren't enough, allow the user to effortlessly implement their own lambda functions to create novel stack mechanisms of their own design

 Is ***gostack*** really more efficient than ***classical go***?  To put this to the test, we created a race for the two; they each have to complete 3 data management tasks as quickly and efficiently as possible.  Whereas ***classical go*** took 61 lines to make it to the finish, ***gostack*** took merely 9—[see for yourself!](/documentation/race.md)

 To get a better feel of the library, feel free to take a look at some [examples](/tutorials/Bootstrap.go) of how ***gostack*** can substitute commonly-used functions.  Alternatively, take a look at our beginner-friendly [introductory tutorial](/tutorials/Introduction.go)!

<h2 name = "footer">Footer</h1>

 This library was created chiefly by Gabe Tucker [[email](mailto:tucker.854@osu.edu)/[LinkedIn](https://www.linkedin.com/in/gabetucker2/)] with contributions from Patrick Da Silva [[email](mailto:dasilva.30@osu.edu)/[LinkedIn](https://www.linkedin.com/in/patrick-da-silva-871833225/)] and Andy Chen [[email](mailto:chenandy@usc.edu)] to facilitate coding a neural network model based in Go.

 If you have any suggestions, questions, or comments in respect to this project, please email me at [tucker.854@osu.edu](mailto:tucker.854@osu.edu).  I will usually respond within 1-2 business days.  Alternatively (and preferably), please join our [official Discord community](https://discord.gg/NmxxcKBVBU) devoted to helping users with gostack issues, announcing exciting updates, and streamlining communication between developers and users!

[Return to Glossary](#glossary)
