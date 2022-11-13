![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.Move()</h2>

 ```
 stack.Move(
    orderType ORDER [ORDER_After],
    findTypeFrom FIND [FIND_First],
    findTypeTo FIND [FIND_Last],
    findDataFrom any|[]any|*Stack|func(
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
    findDataTo any|[]any|*Stack|func(
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
    deepSearchTypeFrom DEEPSEARCH [DEEPSEARCH_False],
    deepSearchTypeTo DEEPSEARCH [DEEPSEARCH_False],
    depthFrom int [-1],
    depthTo int [-1],
    passTypeFrom PASS [PASS_Both],
    passTypeTo PASS [PASS_Both],
    dereferenceTypeFrom DEREFERENCE [DEREFERENCE_None],
    dereferenceTypeTo DEREFERENCE [DEREFERENCE_None],
    overrideFindDataFrom OVERRIDE [OVERRIDE_False],
    overrideFindDataTo OVERRIDE [OVERRIDE_False],
    workingMemFrom []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
    workingMemTo []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}]
 ) (stack)
 ```

```
 Moves one card to before/after another card and returns `stack`

 @ensures
 | IF `overrideFindDataX` == OVERRIDE_True:
 |   compare whether each element is equal to `findDataX` itself, rather than each element inside of `findDataX` (assuming it is a stack or array)
 | IF a version for func input data is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
```

---

 [> Return to functions](../functionsAPI.md)