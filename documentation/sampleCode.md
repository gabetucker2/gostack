![Banner](../images/gostack_Smaller.png)

 <h1>Sample Code</h1>

 Let's get a sense of what ***gostack*** looks like.  In order to make a **Stack** from an existing map structure, you would do:

 ```
 existingMap := map[string]string {"Kid":"Tommy", "Adult":"Chuck", "Adult":"Joey"}
 kidsAndAdults := MakeStack(existingMap)
 ```

 In order to get all people in that **Stack** who are adults, you would do:

 ```
 adults := kidsAndAdults.GetMany(FIND_Key, "Adult")
 ```

 Now, to update the key of every adult from "Adult" to 4, we would do:

 ```
 adults.UpdateMany(REPLACE_Key, 4, FIND_All)
 ```

 And to multiply the key of an adult by 3 if and only if its order (1, 2, ..., n) in the adults **Stack** is even:

 ```
 adults.UpdateMany(REPLACE_Lambda, func(card *Card) {
  card.Key = card.Key.(int) * 3
 }, FIND_Lambda, func(card *Card) (bool) {
  return (card.Idx + 1) % 2 == 0
 })
 ```

 [> Return to **Glossary**](../README.md)