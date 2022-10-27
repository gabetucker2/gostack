![Banner](../images/gostack_SmallerTransparent.png)

 <h1>Function Documentation</h1>

 [MakeCard](#MakeCard)

 <h2 name = "MakeCard">MakeCard</h2>
  
 <h3>variadic: input1 any, input2 any, idx int</h3>

```
 Creates a card with given initial parameters
 
 @param optional `input1` type{any} default nil
 @param optional `input2` type{any} default nil
 @param optional `idx` type{int} default -1
 @returns type{*Card} the newly-constructed card
 @constructs type{*Card} a newly-constructed card
 @ensures
   * variable order will vary:
     IF `input1` is passed:
 		IF `input2` is nil/not passed:
 			MakeCard := func(`val`, /, `idx`)
 		IF `input2` is passed:
 			MakeCard := func(`key`, `val`, `idx`)
 	ELSE `input1` is nil/not passed:
 		MakeCard := func(/, `key`, `idx`)
```

 [> Return to glossary](../README.md)