![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.ToMatrix()</h2>

 `stack.ToMatrix(returnType RETURN [RETURN_Vals], depth int|[]int|*Stack [-1]) (newMatrix []any {elem/[]any{}})`

```
 Creates a new matrix structure from `stack`

 @examples
 | MakeStack([]int {1, 2, 3, 4}).ToMatrix() => []any {1, 2, 3, 4}
 | MakeStack(*Stack{MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}).ToMatrix() => []any {[]any {1, 2}, []any {3, 4}}
```

---

 [> Return to functions](../functionsAPI.md)