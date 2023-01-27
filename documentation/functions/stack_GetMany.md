![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.GetMany()</h2>

 ```
 stack.GetMany(
    findType FIND [FIND_All],
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
    returnType RETURN [RETURN_Cards],
    deepSearchType DEEPSEARCH [DEEPSEARCH_False],
    depth int|[]int|*Stack [-1],
    passType PASS [PASS_Both],
    dereferenceType DEREFERENCE [DEREFERENCE_None],
    overrideFindData OVERRIDE [OVERRIDE_False],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (newStack *Stack)
 ```

```
 Gets a stack of cards from specified parameters in `stack`

 @ensures
 | IF `overrideFindData` == OVERRIDE_True:
 |   compare whether each element is equal to `findData` itself, rather than each element inside of `findData` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
```

---

 [> Return to functions](../functionsAPI.md)