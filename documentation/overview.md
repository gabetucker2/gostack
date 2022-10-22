![Banner](../images/gostack_Smaller.png)

 <h1>Overview</h1>

 Let's get a sense of what ***gostack*** looks like.  We'll start with the structs it introduces:
 
 <img src="../images/gostack_StackAndCard.png" width="50%" style="margin-bottom: 10px;"/>

 Great!  Now, let's make a map structure the way we would in ***native Go***:

 ```
 existingMap := map[string]string {"Kid":"Tommy", "Adult":"Chuck", "Adult":"Joey"}
 ``` 
 Now that we have our map structure, let's make it into a **Stack** so that we can perform ***gostack*** functions on it:

 ```
 kidsAndAdults := MakeStack(existingMap)
 ```

 Nice!  Now, let's take a peek inside `kidsAndAdults`:
 
 <img src="../images/gostack_StackSample1.png" width="50%" style="margin-bottom: 10px;"/>

 Great!  But this seems a little redundant.  Why do we have an **Idx** property when the card's index can already be found by looking at the **Cards** array?  There are two reasons:

 * If you have a card and don't already know its index, you would have to call an iterative function to find its index.  Being able to just do `card.Idx` is simpler and more optimized.
 * If you grab a card in a stack,

 In order to get all people in `kidsAndAdults` who are adults, you would do:

 ```
 adults := kidsAndAdults.GetMany(FIND_Key, "Adult")
 ```

 Now, to update the key of every adult from "Adult" to 4, we would do:

 ```
 adults.UpdateMany(REPLACE_Key, 4, FIND_All)
 ```

 And to multiply the key of an adult by 3 if and only if its order (1, 2, ..., n) in `adults` is even:

 ```
 adults.UpdateMany(REPLACE_Lambda, func(card *Card) {
  card.Key = card.Key.(int) * 3
 }, FIND_Lambda, func(card *Card) (bool) {
  return (card.Idx + 1) % 2 == 0
 })
 ```


---

 [> How much time does gostack save?](race.md)

 [> Return to **Glossary**](../README.md)
 
---