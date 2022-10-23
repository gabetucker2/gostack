![Banner](../images/gostack_Smaller.png)

<h1>Frequently Asked Questions</h1>

<h2>General Questions</h2>

---

 **Question:** How difficult is it to learn ***gostack***?

 To rate the difficulty of our tutorials' content from 1 (friendly to new programmers) to 5 (friendly to Go veterans):
 * [Installation](installation.md)
   * Conceptually: 2
   * In practice: 1
 * [Introductory](introduction.md)
   * Conceptually: 1
   * In practice: 1
 * [Post-Introductory](postIntroduction.md)
   * Conceptually: 2
   * In practice: 2
 * [StackMatrices](matrices.md)
   * Conceptually: 2
   * In practice: 3
 * [Lambda](lambda.md)
   * Conceptually: 4
   * In practice: 4
 * [Pointers](lambda.md)
   * Conceptually: 4
   * In practice: 5

---

<h2>Technical Questions</h2>

---

 **Question:** Why have a card.Idx field when the card's index can be found via the array of the stack it is in?

 Well, there are two reasons:

 * If you have a card and don't already know its index, you would have to call an iterative function to get its index.  Being able to just do `card.Idx` is simpler and more optimized.
 * When you're referencing a card's index, doing `card.Idx` gets a card's current position so that you don't have to worry about whether the card's index has changed since the last time you referenced it.

 ---

 [> Return to Glossary](../README.md)