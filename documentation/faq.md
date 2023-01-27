![Banner](../media/gostack_SmallerTransparent.png)

<h1>Frequently Asked Questions</h1>

<h2>General Questions</h2>

 **Question:** How difficult is it to learn ***gostack***?

 Rating the difficulty of our tutorials' contents from 1 (beginner programmer) to 5 (advanced programmer):
 * [Installation](tutorials/installationTutorial.md)
   * Conceptually: 2
   * In practice: 1
 * [Introductory](tutorials/introductionTutorial.md)
   * Conceptually: 2
   * In practice: 2
 * [StackMatrices](tutorials/matricesTutorial.md)
   * Conceptually: 2
   * In practice: 3
 * [Lambda](tutorials/lambdaTutorial.md)
   * Conceptually: 3
   * In practice: 4
 * [Pointers](tutorials/pointersTutorial.md)
   * Conceptually: 5
   * In practice: 5
 * [Testing](tutorials/testingTutorial.md)
   * Conceptually: 2
   * In practice: 2

---

<h2>Technical Questions</h2>

 **Question:** Why have a card.Idx field when the card's index can be found via the array of the stack it is in?

 There are two reasons:

 * If you have a card and don't already know its index, you would have to call an iterative function to get its index.  Being able to just do `card.Idx` is simpler and more optimized.
 * When you're referencing a card's index, doing `card.Idx` gets a card's current position so that you don't have to worry about whether the card's index has changed since the last time you referenced it.

 ---

 [> Return to glossary](../README.md)