![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>MakeCard()</h2>

 `MakeCard(input1 any [nil], input2 any [nil], idx int [-1]) (*Card)`

```
 Creates a card with given properties
 
 @ensures
 | IF `input1` OR `input2` are nil:
 |     MakeCard := func(`val`, `key`, `idx`)
 | ELSE:
 |     MakeCard := func(`key`, `val`, `idx`)
 @examples
 | MakeCard("Hello") => Card{Val: "Hello"}
 | MakeCard(nil, "Hello") => Card{Key: "Hello"}
 | MakeCard(1, 2) => Card{Key: 1, Val: 2}
```

---

 [> Return to functions](../functionsAPI.md)