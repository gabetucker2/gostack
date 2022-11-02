![Banner](../../images/gostack_SmallerTransparent.png)

 <h2>stack.StripStackMatrix()</h2>

 `stack.StripStackMatrix(idx ...int|[]int|*Stack [[]int {0, 1, ..., stack.Size - 1}]) (stack)`

```
 Updates a stack to represent a selection within that stack matrix

 @requires
 | `idx` refers to valid index positions from the stack
 @examples
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).StripStackMatrix() => Stack {1, 2, 3, 4}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).StripStackMatrix([]int {0, 1}) => Stack {1, 2, 3, 4}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).StripStackMatrix(0) => Stack {1, 2}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).StripStackMatrix(1) => Stack {3, 4}
```

---

 [> Return to functions](../functionsAPI.md)