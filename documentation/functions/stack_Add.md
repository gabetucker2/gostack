![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.Add()</h2>

 ```
 stack.Add(
    insert any|[]any|*Card|*Stack,
    orderType ORDER [ORDER_After],
    findType FIND [FIND_Last],
    findData any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarPtr any,
      otherInfo []any {
            cardPtr,
            parentStackPtr,
            retStackPtr,
            retCardPtr
      },
      workingMem ...any
    ) [nil],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int|[]int|*Stack [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideInsert OVERRIDE [OVERRIDE_False],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)
 ```

```
 Adds `insert` to `stack` before/after first found card and returns `stack`, or nil if invalid find

 @ensures
 | IF `overrideInsert` == OVERRIDE_True:
 |   insert `insert` itself, rather than the elements within `insert` (assuming it is a stack or array)
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
```

---

 [> Return to functions](../functionsAPI.md)