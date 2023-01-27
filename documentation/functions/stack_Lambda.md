![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.LambdaX()</h2>

 ```
 stack.LambdaX(
    lambda func(
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
        workingMem ...any)
    retStack *Stack [nil],
    retCard *Card [nil],
    retVarPtr any [nil],
    workingMem []any [[]any {nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}],
    deepSearchType DEEPSEARCH [DEEPSEARCH_True],
    depth int|[]int|*Stack [-1],
    passType PASS [PASS_Both],
    otherInfo []any {
        retStackPtr,
        retCardPtr
    } []any [[]any {nil, nil}],
 ) (READ BELOW)
 ```

*There are 5 Lambda functions that can take the place of LambdaX.  These are all identical except for their return values:*
* stack.Lambda() (stack, retStack, retCard, retVarPtr) 
* stack.LambdaThis() (stack)
* stack.LambdaStack() (retStack)
* stack.LambdaCard() (retCard)
* stack.LambdaVarAdr() (retVarPtr)

```
 Iterates through `stack` calling your lambda function on each card

 @ensures
 | IF a version for `lambda` is passed that has fewer parameters than the full function:
 |   the function will abstract away unincluded parameters
 | IF you would like to manage more than 10 variables via `workingMem`:
 |   you must pass an []any array into `workingMem` when you call this function
 | IF you would like to reference the object address of `retStack` or `retCard`:
 |   pass the addresses of `retStack` or `retCard` into `otherInfo`
 @examples
 | myStack := MakeStackMatrix([]int {1, 3, 2, 4}, nil, []int {2, 2}).LambdaThis(func(card *Card) {
 |   if card.Idx == 0 && card.Val.(int) % 2 == 0 {
 |     card.Key = "Marker"	
 |   }
 | }) // Stack{nil:1, nil:3, "Marker":2, nil:4}
```

---

 [> Return to functions](../functionsAPI.md)