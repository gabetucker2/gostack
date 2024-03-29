![Banner](../../media/gostack_SmallerTransparent.png)

 <h2>MakeStackMatrix()</h2>

 `MakeStackMatrix(input1 []any (deep/shallow)|map[any]any (deep/shallow)|*Stack [nil], input2 []any (deep/shallow)|*Stack [nil], matrixShape []int [[]int {1}]) (newStackMatrix *Stack)`

```
 Creates a stack matrix initialized with starting cards
 
 Where all mentions of array are interchangeable with Stack:
 @requires
 | `input1` is a map and `input2` is nil
 |     OR `input1` is an array and `input2` is nil
 |     OR `input1` is an array and `input2` is an array
 |     OR `input1` is nil and `input2` is an array
 |
 | IF `input1` AND `input2` are both passed as arguments:
 |      |`input1`| == |`input2`|
 |
 | `matrixShape` must be an int array representing the shape of a regularly-shaped matrix where:
 | * the first int defines `newStackMatrix.Size`
 | * the last int defines the size of each final stack
 | * the product of `matrixShape` is equal to the amount of elements in your input(s)
 @ensures
 | Using the same logic as MakeStack() in deciding which of the first two inputs is a key/val:
 |
 |  IF no `matrixShape` is passed:
 |    treating `input1`/`input2` as matrices ([]any {[]any {...}, []any {...}, ..., []any {...}})/a map of matrices (map[any]map[any]...map[any]any)/a StackMatrix:
 |    IF `input1` is passed:
 |      IF `input1` is a map:
 |        unpack the map into matrix of shape `inputx` with corresponding keys and vals
 |      ELSEIF `input1` is an array and `input2` is nil:
 |        unpack values from `input1` into matrix of shape `inputx`
 |      ELSEIF `input1` is an array and `input2` is an array:
 |        unpack keys from `input1` and values from `input2` into matrix of shape `inputx`
 |      ELSEIF `input1` is nil and `input2` is an array:
 |        unpack keys from `input2` into matrix of shape `inputx` 
 |    ELSEIF `input1` and `input2` are nil:
 |      the stack is empty
 |    ELSEIF `matrixShape` is passed:
 |      treating `input1`/`input2` as 1D structures ([]any, map[any]any, Stack):
 |      IF `input1` is a map:
 |        unpack the map into matrix of shape `matrixShape` with corresponding keys and vals
 |      ELSEIF `input1` is an array and `input2` is nil:
 |        IF `input1` is an array of cards:
 |          unpack cards from `input1` into `stack`
 |        ELSE:
 |          unpack values from `input1` into new cards
 |      ELSEIF `input1` is an array and `input2` is an array:
 |        unpack keys from `input1` and values from `input2` into matrix of shape `matrixShape`
 |      ELSEIF `input1` is nil and `input2` is an array:
 |        unpack keys from `input2` into matrix of shape `matrixShape`
 |      ELSEIF `input1` is nil AND `input2` is nil:
 |        create a StackMatrix of shape `matrixShape` whose heightest card keys/vals are nil
 @examples
 | MakeStackMatrix([]int {1, 2, 3, 4}, nil, []int {2, 2}) => Stack{Stack{1, 2}, Stack{3, 4}}
 | MakeStackMatrix([]int {1, 2, 3, 4, 5, 6}, nil, []int {2, 3}) => Stack{Stack{1, 2, 3}, Stack{4, 5, 6}}
 | MakeStackMatrix([]int {1, 2, 3, 4, 5, 6}, nil, []int {3, 2}) => Stack{Stack{1, 2}, Stack{3, 4}, Stack{5, 6}}
 | MakeStackMatrix([]any {[]any {1, 2}, []any {3, 4}}} =>  Stack{Stack{1, 2}, Stack{3, 4}}
 
```

---

 [> Return to functions](../functionsAPI.md)