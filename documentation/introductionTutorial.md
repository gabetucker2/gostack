

<h2>N vs NMany</h2>

 Functions with name "N" will often have a corresponding function named "NMany".
 
 In such a case, N will only "perform" the function intent on the first found card, whereas NMany will "perform" the function intent on every found card.  For instance:

 ```
 stack.Get(FIND_First) // gets the first card in the set "first"
 stack.Get(FIND_All) // gets the first card in the set "all" (same as previous)
 stack.GetMany(FIND_All) // gets all cards in the set "all"
 stack.Replace(FIND_All) // replaces the first card in the set "all"
 stack.ReplaceMany(FIND_All) // replaces all cards in the set "all"
 ```