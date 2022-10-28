![Banner](../images/gostack_SmallerTransparent.png)

 <h1>Function Documentation</h1>

 *See [meta documentation](metaAPI.md) if you are curious about the format underlying our function documentation.*

 > Initialization functions
 >> [MakeCard()](functions/MakeCard.md)
 >
 >> [MakeStack()](functions/MakeStack.md)
 > >
 >> [MakeSubstack()](functions/MakeSubstack.md)
 > 
 >> [MakeStackMatrix()](functions/MakeStackMatrix.md)

 > Conversion functions
 >> [stack.ToArray()](functions/stack_ToArray.md)
 >
 >> [stack.ToMap()](functions/stack_ToMap.md)
 >
 >> [stack.ToMatrix()](functions/stack_ToMatrix.md)
 
 > Matrix functions
 >> [stack.IsRegular()](functions/stack_IsRegular.md)
 >
 >> [stack.Shape()](functions/stack_Shape.md)
 >
 >> [stack.StripStackMatrix()](functions/stack_StripStackMatrix.md)
 
 > Card/Stack management functions
 >> [card.Print()](functions/card_Print.md)
 >>
 >> [stack.Print()](functions/stack_Print.md)
 >
 >> [card.SwitchKeyVal()](functions/card_SwitchKeyVal.md)
 >>
 >> [stack.SwitchKeyVal()](functions/stack_SwitchKeyVal.md)
 >
 >> [card.Clone()](functions/card_Clone.md)
 >>
 >> [stack.Clone()](functions/stack_Clone.md)
 >
 >> [card.Equals()](functions/card_Equals.md)
 >>
 >> [stack.Equals()](functions/stack_Equals.md)

 > Miscellaneous functions
 >> [stack.Duplicate()](functions/stack_Duplicate.md)
 >
 >> [stack.Empty()](functions/stack_Empty.md)
 >
 >> [stack.Unique()](functions/stack_Unique.md)
 >
 >> [stack.Shuffle()](functions/stack_Shuffle.md)
 >
 >> [stack.Transpose()](functions/stack_Transpose.md)
 
 > Lambda functions
 >> [stack.Lambda()](functions/stack_Lambda.md)
 >
 >> [stack.LambdaThis()](functions/stack_LambdaThis.md)
 >
 >> [stack.LambdaStack()](functions/stack_LambdaStack.md)
 >
 >> [stack.LambdaCard()](functions/stack_LambdaCard.md)
 >
 >> [stack.LambdaVarAdr()](functions/stack_LambdaVarAdr.md)
 
 > Generic functions
 >> [stack.Move()](functions/stack_Move.md)
 >
 >> [stack.Swap()](functions/stack_Swap.md)
 >
 >> [stack.Has()](functions/stack_Has.md)
 >
 >> [stack.Add()](functions/stack_Add.md)
 >>
 >> [stack.AddMany()](functions/stack_AddMany.md)
 >
 >> [stack.Get()](functions/stack_Get.md)
 >>
 >> [stack.GetMany()](functions/stack_GetMany.md)
 >
 >> [stack.Replace()](functions/stack_Replace.md)
 >>
 >> [stack.ReplaceMany()](functions/stack_ReplaceMany.md)
 >
 >> [stack.Extract()](functions/stack_Extract.md)
 >>
 >> [stack.ExtractMany()](functions/stack_ExtractMany.md)
 >
 >> [stack.Remove()](functions/stack_Remove.md)
 >>
 >> [stack.RemoveMany()](functions/stack_RemoveMany.md)
 >
 >> [stack.Update()](functions/stack_Update.md)
 >>
 >> [stack.UpdateMany()](functions/stack_UpdateMany.md)

 <h2 name = "MakeSubstack">MakeSubstack</h2>

 `MakeSubstack(input1 []any|map[any]any|*Stack [nil], input2 []any|*Stack [nil], repeats int [1], overrideCards OVERRIDE [OVERRIDE_False]) (newSubstack *Stack)`

```
 An identical implementation to `MakeStack()`
```

 <h2 name = "MakeStackMatrix">MakeStackMatrix</h2>

 `MakeStackMatrix(input1 []any (deep/shallow)|map[any]any (deep/shallow)|*Stack [nil], input2 []any (deep/shallow)|*Stack [nil], matrixShape []int [[]int {1}], overrideCards OVERRIDE [OVERRIDE_False]) (newStackMatrix *Stack)`

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