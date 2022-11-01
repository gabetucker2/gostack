![Banner](../../images/gostack_SmallerTransparent.png)

 <h2>stack.ToArray()</h2>

 `stack.ToArray(returnType RETURN [RETURN_Vals]) (newArray []any)`

```
 Creates a new any array whose elements are the values of the cards in `stack`

 @examples
 | MakeStack([]int {1, 2, 3}, []string {"a", "b", "c"}).ToArray() => []any {1, 2, 3}
 | MakeStack([]int {1, 2, 3}, []string {"a", "b", "c"}).ToArray(RETURN_Keys) => []any {"a", "b", "c"}
 | MakeStack([]int {1, 2, 3}, []string {"a", "b", "c"}).ToArray(RETURN_Idxs) => []any {0, 1, 2}
 | MakeStack([]*Card {cardA, cardB, cardC}).ToArray(RETURN_Cards) => []any {cardA, cardB, cardC}
 | MakeStack([]*Stack {substackA, substackB}).ToArray(RETURN_Cards) => []any {Card{Val:substackA}, Card{Val:substackA}}
 | MakeStack([]*Stack {substackA, substackB}).ToArray(RETURN_Stacks) => []any {substackA, substackB}
```

---

 [> Return to functions](../functionsAPI.md)