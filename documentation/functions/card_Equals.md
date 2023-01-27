![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>card.Equals()</h2>

 ```
 card.Equals(
    otherCard *Card,
    compareIdxs COMPARE [COMPARE_False],
    compareKeys COMPARE [COMPARE_True],
    compareVals COMPARE [COMPARE_True],
    compareCardPtrs COMPARE [COMPARE_False],
    pointerKeys DEREFERENCE [DEREFERENCE_None],
    pointerVals DEREFERENCE [DEREFERENCE_None]
 ) (cardEqualsOtherCard bool)
 ```

```
 Returns whether one card equals another

 @examples
 | card1 := MakeCard("Hey")
 | card2 := MakeCard("Hey")
 | myStr := "Hey"
 | card1.Equals(card2, nil, nil, nil, COMPARE_False) // True
 | card1.Equals(card2, nil, nil, nil, COMPARE_True) // False
 | card1.Equals(MakeCard(&myStr), nil, nil, nil, nil, nil, DEREFERENCE_This) // True
```

---

 [> Return to functions](../functionsAPI.md)