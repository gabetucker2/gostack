![Banner](../../media/gostack_SmallerTransparent.png)

<h2>DEREFERENCE</h2>

 > None, Both, Found, This

This enumerator which decides whether to dereference the key/val of the found cards data or the input data before checking for equality.  For instance, given `stack.Get(..., ..., ..., ..., ..., ..., dereferenceType DEREFERENCE)`:

```
init1 := 1
intValA = &init1
init2 := 2
intValB = &init2
init3 := 3
intValC = &init3
init4 := 2
intValD = &init4

mainStack := MakeStack([]any {intValA, intValB, intValC})

mainStack.Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_None) // gets intValB since intValB == intValB
mainStack.Get(FIND_Val, intValD, nil, nil, nil, DEREFERENCE_None) // doesn't get intValB since intValB != intValD

mainStack.Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_Both) // gets intValB since *intValB == *intValB
mainStack.Get(FIND_Val, intValD, nil, nil, nil, DEREFERENCE_Both) // gets intValB since 2 == 2

mainStack.Get(FIND_Val, 2, nil, nil, nil, DEREFERENCE_Found) // gets intValB since *intValB == 2

MakeStack([]any {1, 2, 3}).Get(FIND_Val, intValB, nil, nil, nil, DEREFERENCE_This) // gets 2 since 2 == *intValB
```

Enumerators:
 > DEREFERENCE_None
 
 > DEREFERENCE_Both

 > DEREFERENCE_Found
 >> for `stack.Func(findType, findData)`, stuff in `stack` is dereferenced

 > DEREFERENCE_This
 >> for `stack.Func(findType, findData)`, stuff in `findData` is dereferenced

 ---

 [> Return to enumerators](../enumsAPI.md)