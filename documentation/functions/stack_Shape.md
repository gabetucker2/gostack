![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.Shape()</h2>

 `stack.Shape() (stackShape []int)`

```
 Returns an array representing the shape of `stack`

 stack.Shape() (newStack stack)

 @ensures
 | returns nil if it's not regular and thus doesn't have a shape
 @examples
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2, 3}), MakeSubstack([]int {4, 5, 6})}).Shape() => []int {2, 3}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4}), MakeSubstack([]int {5, 6})}).Shape() => []int {3, 2}
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4, 5}), MakeSubstack([]int {6, 7})}) => nil
```

---

 [> Return to functions](../functionsAPI.md)