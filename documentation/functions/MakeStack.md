 <h1 name = "MakeStack">MakeStack</h1>

 `MakeStack(input1 []any|map[any]any|*Stack [nil], input2 []any|*Stack [nil], repeats int [1], overrideCards OVERRIDE [OVERRIDE_False]) (newStack *Stack)`

```
 Creates a stack initialized with starting cards
 
 Where all mentions of array are interchangeable with Stack:
 @notes
 | Makes `repeats` repeats of `input1`/`input2`
 @requires
 | `input1` is a map and `input2` is nil
 |     OR `input1` is an array and `input2` is nil
 |     OR `input1` is an array and `input2` is an array
 |     OR `input1` is nil and `input2` is an array
 |
 | IF `input1` AND `input2` are both passed as arguments
 |      |`input1`| == |`input2`|
 @ensures
 |     IF `input1` is passed
 |       IF `input1` is a map
 |         unpack the map into new cards with corresponding keys and vals
 |       ELSEIF `input1` is an array and `input2` is not passed/nil
 |  	   IF `input1` is an array of cards:
 |           `overrideCards` == OVERRIDE_True:
 |               MakeStack([]*Card {cardA}) => stack.Cards = []*Card { card {Idx = 0, Key = nil, Val = cardA} }
 |           `overrideCards` == OVERRIDE_False:
 |               MakeStack([]*Card {cardA}) => stack.Cards = []*Card {cardA}
 |  	   ELSE:
 |           unpack values from `input1` into new cards
 |       ELSEIF `input1` is an array and `input2` is an array
 |         unpack keys from `input1` and values from `input2` into new cards
 |       ELSEIF `input1` is nil and `input2` is an array
 |         unpack keys from `input2` into new cards
 |  		make `repeats` cards with nil value and nil key
 |  		ELSEIF `input1` is nil and `input2` is nil and `repeats` is passed
 |     ELSE
 |       the stack is empty
```

 [> Return to functions](../functionsAPI.md)