![Banner](../../images/gostack_SmallerTransparent.png)

 <h2>MakeSubstack()</h2>

 `MakeSubstack(input1 []any|map[any]any|*Stack [nil], input2 []any|*Stack [nil], repeats int [1], overrideCards OVERRIDE [OVERRIDE_False]) (newSubstack *Stack)`

 *An identical implementation to [MakeStack()](MakeStack.md) with a different function name for improved conceptual clarity while coding.*

```
@examples
 | MakeStack([]*Stack {MakeSubstack([]int {1, 2}), MakeSubstack([]int {3, 4})}) => Stack{Stack{1, 2}, Stack{3, 4}}
```

---

 [> Return to functions](../functionsAPI.md)