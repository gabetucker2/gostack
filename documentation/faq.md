![Banner](../images/gostack_Smaller.png)

<h1>Frequently Asked Questions</h1>

<h2>General Questions</h2>

 **Question:** How difficult is it to learn ***gostack***?

 Rating the difficulty of our tutorials' contents from 1 (beginner programmer) to 5 (advanced programmer):
 * [Installation](installationTutorial.md)
   * Conceptually: 2
   * In practice: 1
 * [Introductory](introductionTutorial.md)
   * Conceptually: 1
   * In practice: 1
 * [Post-Introductory](postIntroductionTutorial.md)
   * Conceptually: 2
   * In practice: 2
 * [StackMatrices](matricesTutorial.md)
   * Conceptually: 2
   * In practice: 3
 * [Lambda](lambdaTutorial.md)
   * Conceptually: 4
   * In practice: 4
 * [Pointers](pointersTutorial.md)
   * Conceptually: 4
   * In practice: 5

---

<h2>Technical Questions</h2>

 **Question:** Why have a card.Idx field when the card's index can be found via the array of the stack it is in?

 There are two reasons:

 * If you have a card and don't already know its index, you would have to call an iterative function to get its index.  Being able to just do `card.Idx` is simpler and more optimized.
 * When you're referencing a card's index, doing `card.Idx` gets a card's current position so that you don't have to worry about whether the card's index has changed since the last time you referenced it.

 ---

 [> Return to glossary](../README.md)