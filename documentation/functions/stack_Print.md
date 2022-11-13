![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.Print()</h2>

 `stack.Print(name string [""], indent int [0]) (stack)`

```
 Prints information surrounding `stack` to the terminal and returns `stack`
 
 @ensures
 | prints "-" `indent` * 4 times before each line to indicate depth in a stackMatrix
 @examples
 | MakeStack([]string {"Hey", "Hi"}).Print().Remove(FIND_Last).Print() // prints the stack before and after performing the remove function
```

---

 [> Return to functions](../functionsAPI.md)