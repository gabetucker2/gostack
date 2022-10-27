![Banner](../images/gostack_SmallerTransparent.png)

 <h1>Function Documentation</h1>

 *See [meta documentation](metaAPI.md) if you are curious about the format underlying our function documentation.*

 > [MakeCard](#MakeCard)
 > [MakeStack](#MakeStack)

 <h1 name = "MakeCard"></h1>

 `MakeCard(input1 any [nil], input2 any [nil], idx int [-1]) (*Card)`

```
 Creates a card with given properties
 
 @ensures
 | IF `input1` OR `input2` are not passed:
 |     MakeCard := func(`val`, `key`, `idx`)
 | ELSE:
 |     MakeCard := func(`key`, `val`, `idx`)
```

 <h1 name = "MakeStack"></h1>

 `MakeStack(input1 []any [nil], input2 any/map[any]any [nil], repeats int [1], overrideCards OVERRIDE [OVERRIDE_False]) (newStack *Stack)`

```
 Creates a stack of cards with optional starting cards
 
 @notes
 | * By default, if you do MakeStack([]*Card {cardA}), stack.Cards = []*Card {cardA}.  If you would like your cards to have vals pointing to other cards, where stack.Cards = []*Card { card {Idx = 0, Key = nil, Val = cardA} }, set this variable to true.
 |   * repeats the function filling `repeats` (or, if nil or under 0, 1) amount of times
 @requires
 | `input1` is a map and `input2` is nil
 |     OR `input1` is an array and `input2` is nil
 |     OR `input1` is an array and `input2` is an array
 |     OR `input1` is nil and `input2` is an array
 |
 | IF `input1` AND `input2` are both passed as arguments
 |      |`input1`| == |`input2`|
 @ensures
 | assuming all mentions of array are interchangeable with *Stack:
 |    IF `input1` is passed
 |      IF `input1` is a map
 |        unpack the map into new cards with corresponding keys and vals
 |      ELSEIF `input1` is an array and `input2` is not passed/nil
 | 	    IF `input1` is an array of cards:
 | 		  set `stack.Cards` to `input1`
 | 		ELSE:
 |          unpack values from `input1` into new cards
 |      ELSEIF `input1` is an array and `input2` is an array
 |        unpack keys from `input1` and values from `input2` into new cards
 |      ELSEIF `input1` is nil and `input2` is an array
 |        unpack keys from `input2` into new cards
 | 		make `repeats` cards with nil value and nil key
 | 		ELSEIF `input1` is nil and `input2` is nil and `repeats` is passed
 |    ELSE
 |      the stack is empty
```

 [> Return to glossary](../README.md)