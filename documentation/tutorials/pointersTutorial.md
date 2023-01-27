![Banner](../../media/gostack_SmallerTransparent.png)

 <h1>Pointers Tutorial</h1>

 <h2>Overview</h2>

 Hi there.  Welcome to the pointers tutorial!  We touched on pointers briefly in the [lambda tutorial](lambdaTutorial.md), but here, we will expand on how to use pointers more generally in ***gostack***.

 When are pointers useful?  Well, one substantial restraint of ***gostack*** is that a card can only be in one stack at a time.  As a result, if you would like to have a stack referencing a set of cards in another stack, then you are out of luck.  In turn, ***gostack*** has a system for dealing with cases similar to these, and it also allows you to easily deal with pointers and addresses more generally.

 The parameter tailored for dealing with dereferencing is the `DEREFERENCE` parameter.  It allows you to determine whether you dereference each variable in the receiver stack or the set of found cards when you are comparing between them.  For instance, given you have a stack of pointers:

```
var str1, str2, str3 any
str1 = "Guy"
str1Ptr := &str1
str2 = "Guy"
str2Ptr := &str2
str3 = "Girl"
str3Ptr := &str3
myStackDeref := MakeStack([]any {str1Ptr, str2Ptr, str3Ptr}).Unique(nil, DEREFERENCE_This) // Stack{str1Ptr, str3Ptr}
myStackNoDeref := MakeStack([]any {str1Ptr, str2Ptr, str3Ptr}).Unique(nil, DEREFERENCE_None) // Stack{str1Ptr, str2Ptr, str3Ptr}
```

Using `DEREFERENCE_This` dereferences each pointer in the receiver stack before comparing whether they are equal to another pointer, thereby determining that even though `str1Ptr != str2Ptr`, the objects to which they point are equal; thus, only one is added to the final stack (viz. Unique's functionality).  However, for the instance with `DEREFERENCE_None`, it sees that `str1Ptr != str2Ptr` shallowly; thus, both are kept in the original stack.

If you do `DEREFERENCE_Found` coupled with `FIND_Val`, you would instead dereference the Val of each found card.  If you instead did so with `FIND_Card`, you would dereference each val, treating each val as a pointer to a card pointer:

```
cardA := MakeCard()
cardB := MakeCard()
cardC := MakeCard()
cardD := MakeCard()
mainStack := MakeStack([]*Card {cardA, cardB, cardC}).Print()
ptrStack := MakeStack([]any {fmt.Sprintf("%p", cardA), fmt.Sprintf("%p", cardB), fmt.Sprintf("%p", cardD)}).Print()
ptrCardsInCommon := ptrStack.GetMany(FIND_Card, mainStack, nil, nil, nil, nil, DEREFERENCE_This) // Stack{&cardA, &cardB}
```

Ensure that you reference the address of each card using `fmt.Sprintf("%p", card)`, or by passing the `RETURN_Adrs` argument to returnType in the `GetMany` function.

Congratulations!  That's the gist of what you need to know to manage pointers in ***gostack***.

---
 [> Return to glossary](../../README.md)