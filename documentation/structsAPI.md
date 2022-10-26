![Banner](../images/gostack_Smaller.png)

<h1>Structs Documentation</h1>

 In ***gostack***, there are simply two structs:

 <img src="../images/gostack_StackAndCard.png" width="50%" style="margin-bottom: 10px;"/>

 These are the only structs you have to understand in order to master ***gostack***!

 In a stack matrix structure, "substack" will refer to cards whose val is another stack, and "card" will refer to cards whose val is not another stack.  For instance, given some stack matrix named "myMatrix":

 <img src="../images/gostack_StackSample3.png" width="70%" style="margin-bottom: 10px;"/>
 
 The cards whose vals are 1, 3, 2, and 4, respectively (the grandchildren of myMatrix), are called "cards" in this context.  Conversely, the cards whose vals are other stacks (the children of myMatrix) are called "substacks", even though their data type is `*Card`.

 ---

 [> Return to glossary](../README.md)