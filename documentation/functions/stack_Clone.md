![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.Clone()</h2>

 `stack.Clone(deepSearchType DEEPSEARCH [DEEPSEARCH_True], depth int|[]int|*Stack [-1], cloneCardKeys CLONE [CLONE_True], cloneCardVals CLONE [CLONE_True], cloneSubstackKeys CLONE [CLONE_True], cloneSubstackVals CLONE [CLONE_True]) (newStack stack)`

```
 Returns a clone of `stack`

 @ensures
 | If `cloneSubstackVals` == CLONE_False, then each card holding a substack as its Val will have its Val updated to nil
```

---

 [> Return to functions](../functionsAPI.md)