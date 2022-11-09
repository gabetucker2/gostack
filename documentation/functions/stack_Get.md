![Banner](../../images/gostack_SmallerTransparent.png)

 <h2>stack.Get()</h2>

 ```
 stack.Get(
    findType FIND [FIND_Last],
    findData any [nil],
    findCompareRaw COMPARE [COMPARE_False],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int [-1],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    passSubstacks PASS [PASS_True],
    passCard PASS [PASS_True],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)
 ```

```
 Gets the first card from specified parameters in a stack, or nil if does not exist
```

---

 [> Return to functions](../functionsAPI.md)