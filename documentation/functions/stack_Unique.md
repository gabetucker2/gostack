![Banner](../../images/gostack_SmallerTransparent.png)

 <h2>stack.Unique()</h2>

 `stack.Unique(uniqueType TYPE [TYPE_Val]) (stack)`

```
 Removes all cards from `stack` which share a given property as another card in that stack

 @examples
 | MakeStack([]int {1, 2, 3, 1, 2, 4}).Unique() // Stack{1, 2, 3, 4}
 | MakeStack([]int {0, 1, 0, 0, 1, 0}, []int {1, 2, 3, 1, 2, 4}).Unique(TYPE_Key) // Stack{1, 2}
```

---

 [> Return to functions](../functionsAPI.md)