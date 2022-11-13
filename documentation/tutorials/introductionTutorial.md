![Banner](../../media/gostack_SmallerTransparent.png)

 <h1>Introduction</h1>

 Hi there.  Welcome to the gostack introductory tutorial!  This tutorial will familiarize you with the primary functions and concepts in gostack.  If you haven't yet, please take a look at our page on [how gostack works](../overview.md) to get a conceptual overview of this library.  Let's begin with some intuitive examples of ***gostack*** code and break them down.
 
 First, we will make a stack.  This stack, by default, looks like the following:

 ```		
 myStack Stack {
   Size: 0,
   Height: 1,
   []*Cards: { },
 }
 ```

 `myStack := MakeStack()`

 Great!  We have now made an empty stack.  Now let's make some cards we can add to the stack.  The first card will look like the following:

 ```
 cardA Card {
   Idx: -1,
   Key: nil,
   Val: "Butterfly"
 }
 ```
 
 ```
 cardA := MakeCard("Butterfly")
 cardB := MakeCard("Praying Mantis")
 cardC := MakeCard("Beetle")
 ```

 

---
 [> Return to glossary](../README.md)