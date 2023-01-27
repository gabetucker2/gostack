![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.DimensionalityReduce()</h2>

 `stack.DimensionalityReduce(idx ...int|[]int|*Stack [[]int {0, 1, ..., stack.Size - 1}]) (stack)`

```
 Updates a stack to represent a selection of elements within that stack matrix

 @notes
 | This can be used to reduce an ND array structure to a 1D vector structure
 | This can be used to select a subset of a matrix
 | This can be used to select a card at given coordinates
 @requires
 | `idx` refers to valid index positions from the stack
 @examples
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).DimensionalityReduce() => Stack {1, 2, 3, 4}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).DimensionalityReduce([]int {0, 1}) => Stack {1, 2, 3, 4}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).DimensionalityReduce(0) => Stack {1, 2}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).DimensionalityReduce(1) => Stack {3, 4}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).DimensionalityReduce(0, 1) => Stack {2}
```

---

 [> Return to functions](../functionsAPI.md)