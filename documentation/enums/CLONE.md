![Banner](../../images/gostack_SmallerTransparent.png)

<h2>CLONE</h2>

 > True, False

Provides information to a function on whether to clone the key/val of a substack or card.  For example, given `stack.Clone(..., ..., cloneCardKeys, cloneCardVals, cloneSubstackKeys, cloneSubstackVals CLONE)`:

```
myStack.Clone(nil, nil, CLONE_True, CLONE_False, CLONE_True, CLONE_True)
// returns a clone of myStack such that each val of a card in the new stack contains the same object val as the original card, but each card's key is a clone of the original key and each substacks key or val is a clone of the original key or val
```

Enumerators:
 > CLONE_True
 
 > CLONE_False

 ---

 [> Return to enumerators](../enumsAPI.md)