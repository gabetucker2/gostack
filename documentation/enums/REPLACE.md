![Banner](../../images/gostack_SmallerTransparent.png)

<h2>REPLACE</h2>

 > Key, Val, Card, Lambda

Many functions have "replace functionality".  This means they A) have a REPLACE parameter called `replaceType` and B) have an interface parameter called `replaceWith` whose dynamic type will vary based on the input for `replaceType`.  Given `stack.Replace(..., ..., ..., replaceType)`:

```
stack.ReplaceMany(FIND_All, nil, nil, REPLACE_Val, "NewVal") // replaces the val of all cards with "NewVal"
stack.ReplaceMany(FIND_All, nil, nil, REPLACE_Lambda, func (*Card card) {card.Val = card.Val.(int) * 3}) // (assuming each found card's val is an int) multiplies the val of every card by 3
```

Sample:
 > enumerator type's name
 >> the accepted data type(s), one of which must be passed into `replaceWith`
 >
 >> the target(s) being replaced by `replaceWith`

Enumerators:
 > REPLACE_Key
 >> `any`
 >
 >> each found card's Key
 >
 > REPLACE_Val
 >> `any`
 >
 >> each found card's Val
 >
 > REPLACE_Card
 >> `*Card`
 >
 >> each found card
 >
 > FIND_Lambda
 >>> `func(*Card card, *Stack parentStack, bool isSubstack, ...any workingMemAdrs)`
 >>
 >>> where you can pass a function containing between 0 and all of these parameters (assuming no parameter's order is changed)
 >
 >> there is no build-in target; the lambda function passed in is responsible for updating whatever target it would like responsibly

 ---

 [> Return to enumerators](../enumsAPI.md)