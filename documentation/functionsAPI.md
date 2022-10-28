![Banner](../images/gostack_SmallerTransparent.png)

 <h1>Function Documentation</h1>

 *See [meta documentation](metaAPI.md) if you are curious about the format underlying our function documentation.*

 > Initialization functions
 >> [MakeCard()](#MakeCard)
 >
 >> [MakeStack()](#MakeStack)
 > >
 >> [MakeSubstack()](#MakeSubstack)
 > 
 >> [MakeStackMatrix()](#MakeStackMatrix)

 > Conversion functions
 >> [stack.ToArray()](#stack_ToArray)
 >
 >> [stack.ToMap()](#stack_ToMap)
 >
 >> [stack.ToMatrix()](#stack_ToMatrix)
 
 > Matrix functions
 >> [stack.IsRegular()](#stack_IsRegular)
 >
 >> [stack.Shape()](#stack_Shape)
 >
 >> [stack.StripStackMatrix()](#stack_StripStackMatrix)
 
 > Card/Stack management functions
 >> [card.Print()](#card_Print)
 >>
 >> [stack.Print()](#stack_Print)
 >
 >> [card.SwitchKeyVal()](#card_SwitchKeyVal)
 >>
 >> [stack.SwitchKeyVal()](#stack_SwitchKeyVal)
 >
 >> [card.Clone()](#card_Clone)
 >>
 >> [stack.Clone()](#stack_Clone)
 >
 >> [card.Equals()](#card_Equals)
 >>
 >> [stack.Equals()](#stack_Equals)

 > Miscellaneous functions
 >> [stack.Duplicate()](#stack_Duplicate)
 >
 >> [stack.Empty()](#stack_Empty)
 >
 >> [stack.Unique()](#stack_Unique)
 >
 >> [stack.Shuffle()](#stack_Shuffle)
 >
 >> [stack.Transpose()](#stack_Transpose)
 
 > Lambda functions
 >> [stack.Lambda()](#stack_Lambda)
 >
 >> [stack.LambdaThis()](#stack_LambdaThis)
 >
 >> [stack.LambdaStack()](#stack_LambdaStack)
 >
 >> [stack.LambdaCard()](#stack_LambdaCard)
 >
 >> [stack.LambdaVarAdr()](#stack_LambdaVarAdr)
 
 > Generic functions
 >> [stack.Move()](#stack_Move)
 >
 >> [stack.Swap()](#stack_Swap)
 >
 >> [stack.Has()](#stack_Has)
 >
 >> [stack.Add()](#stack_Add)
 >>
 >> [stack.AddMany()](#stack_AddMany)
 >
 >> [stack.Get()](#stack_Get)
 >>
 >> [stack.GetMany()](#stack_GetMany)
 >
 >> [stack.Replace()](#stack_Replace)
 >>
 >> [stack.ReplaceMany()](#stack_ReplaceMany)
 >
 >> [stack.Extract()](#stack_Extract)
 >>
 >> [stack.ExtractMany()](#stack_ExtractMany)
 >
 >> [stack.Remove()](#stack_Remove)
 >>
 >> [stack.RemoveMany()](#stack_RemoveMany)
 >
 >> [stack.Update()](#stack_Update)
 >>
 >> [stack.UpdateMany()](#stack_UpdateMany)

 <h1 name = "MakeCard">MakeCard</h1>

 `MakeCard(input1 any [nil], input2 any [nil], idx int [-1]) (*Card)`

```
 Creates a card with given properties
 
 @ensures
 | IF `input1` OR `input2` are nil:
 |     MakeCard := func(`val`, `key`, `idx`)
 | ELSE:
 |     MakeCard := func(`key`, `val`, `idx`)
```

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

 <h1 name = "MakeSubstack">MakeSubstack</h1>

 `MakeSubstack(input1 []any|map[any]any|*Stack [nil], input2 []any|*Stack [nil], repeats int [1], overrideCards OVERRIDE [OVERRIDE_False]) (newSubstack *Stack)`

```
 An identical implementation to `MakeStack()`
```

 <h1 name = "MakeStackMatrix">MakeStackMatrix</h1>

 `MakeStackMatrix(input1 []any|[]any {[]any, ..., []any}|map[any]any|map[any]...map[any]|*Stack [nil], input2 []any|[]any {[]any, ..., []any}|*Stack [nil], matrixShape []int [[]int {1}], overrideCards OVERRIDE [OVERRIDE_False]) (newStackMatrix *Stack)`

```
 Creates a stack matrix initialized with starting cards
 
 Where all mentions of array are interchangeable with Stack:
 @requires
 | `input1` is a map and `input2` is nil
 |     OR `input1` is an array and `input2` is nil
 |     OR `input1` is an array and `input2` is an array
 |     OR `input1` is nil and `input2` is an array
 |
 | IF `input1` AND `input2` are both passed as arguments
 |      |`input1`| == |`input2`|
 |
 | `matrixShape` must be an int array representing the shape of a regularly-shaped matrix where:
 | * the first int defines `newStackMatrix.Size`
 | * the last int defines the size of each final stack
 | * the product of `matrixShape` is equal to the amount of elements in your input(s)
 @ensures
 |  IF no `matrixShape` is passed
 |    treating `input1`/`input2` as matrices ([]any {[]any {...}, []any {...}, ..., []any {...}})/a map of matrices (map[any]map[any]...map[any]any)/a StackMatrix:
 |    IF `input1` is passed
 |      IF `input1` is a map
 |        unpack the map into matrix of shape `inputx` with corresponding keys and vals
 |      ELSEIF `input1` is an array and `input2` is nil
 |        unpack values from `input1` into matrix of shape `inputx`
 |      ELSEIF `input1` is an array and `input2` is an array
 |        unpack keys from `input1` and values from `input2` into matrix of shape `inputx`
 |      ELSEIF `input1` is nil and `input2` is an array
 |        unpack keys from `input2` into matrix of shape `inputx` 
 |    ELSEIF `input1` and `input2` are nil
 |      the stack is empty
 |    ELSEIF `matrixShape` is passed
 |      treating `input1`/`input2` as 1D structures ([]any, map[any]any, Stack):
 |      IF `input1` is a map
 |        unpack the map into matrix of shape `matrixShape` with corresponding keys and vals
 |      ELSEIF `input1` is an array and `input2` is nil
 |        IF `input1` is an array of cards:
 |         `overrideCards` == OVERRIDE_True:
 |             MakeStackMatrix([]*Card {cardA}) => stack.Cards = []*Card { card {Idx = 0, Key = nil, Val = cardA} }
 |         `overrideCards` == OVERRIDE_False:
 |             MakeStackMatrix([]*Card {cardA}) => stack.Cards = []*Card {cardA}
 |        ELSE:
 |           unpack values from `input1` into new cards
 |      ELSEIF `input1` is an array and `input2` is an array
 |        unpack keys from `input1` and values from `input2` into matrix of shape `matrixShape`
 |      ELSEIF `input1` is nil and `input2` is an array
 |        unpack keys from `input2` into matrix of shape `matrixShape`
 |      ELSEIF `input1` is nil AND `input2` is nil
 |        create a StackMatrix of shape `matrixShape` whose heightest card keys/vals are nil
```

 [> Return to glossary](../README.md)