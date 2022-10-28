 <h2 name = "MakeCard">MakeCard</h2>

 `MakeCard(input1 any [nil], input2 any [nil], idx int [-1]) (*Card)`

```
 Creates a card with given properties
 
 @ensures
 | IF `input1` OR `input2` are nil:
 |     MakeCard := func(`val`, `key`, `idx`)
 | ELSE:
 |     MakeCard := func(`key`, `val`, `idx`)
```

 [> Return to functions](../functionsAPI.md)