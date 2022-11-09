![Banner](../../images/gostack_SmallerTransparent.png)

 <h2>stack.Swap()</h2>

 ```
 stack.Swap(
    findType1 FIND [FIND_First],
    findType2 FIND [FIND_Last],
    findData1 any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarAdr any,
      otherInfo []any {
            cardAdr,
            parentStackAdr,
            retStackAdr,
            retCardAdr
      },
      workingMem ...any
    ) [nil],
    findData2 any|[]any|*Stack|func(
      card *Card,
      parentStack *Stack,
      isSubstack bool,
      coords *Stack,
      retStack *Stack,
      retCard *Card,
      retVarAdr any,
      otherInfo []any {
            cardAdr,
            parentStackAdr,
            retStackAdr,
            retCardAdr
      },
      workingMem ...any
    ) [nil],
    deepSearchType1 DEEPSEARCH [DEEPSEARCH_False],
    deepSearchType2 DEEPSEARCH [DEEPSEARCH_False],
    depth1 int [-1],
    depth2 int [-1],
    passType1 PASS [PASS_Both],
    passType2 PASS [PASS_Both],
    dereferenceType1 DEREFERENCE [DEREFERENCE_None],
    dereferenceType2 DEREFERENCE [DEREFERENCE_None],
    overrideFindData1 OVERRIDE [OVERRIDE_False],
    overrideFindData2 OVERRIDE [OVERRIDE_False],
    workingMem1 []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
    workingMem2 []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)
 ```

```
 Swaps one card with another and returns `stack`

 @ensures
 | IF `overrideFindDataX` == OVERRIDE_True:
 |   compare whether each element is equal to `findDataX` itself, rather than each element inside of `findDataX` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
```

---

 [> Return to functions](../functionsAPI.md)