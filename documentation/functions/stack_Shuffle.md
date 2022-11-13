![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>stack.Shuffle()</h2>

 `stack.Shuffle(newOrder bool [true]) (stack)`

```
 Shuffles the order of `stack` cards

 @ensures
 | IF `newOrder` == true AND stack.Size > 1:
 |   shuffles `stack` until it is no longer in its previous order
 | rand.Seed is updated to time.Now().UnixNano()
```

---

 [> Return to functions](../functionsAPI.md)