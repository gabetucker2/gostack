 <h1>Race</h1>

<h2>Introduction</h2>

 Welcome!  This consists of a race between ***classical Go*** and ***gostack*** to see who can perform the same data management tasks faster.

 Keep in mind that ***gostack*** code has been updated since this was written and ***classical go*** code is untested, meaning this script will be slightly outdated until ***gostack*** is finished.

 <h4>Jump To</h4>

 > [pseudocode outline](#pseudocode)

 > [classical go](#classical)

 > [gostack](#gostack)

 <h3>Assuming you would like to...</h3>

 > A) ...make a list representing a non-duplicating set of values from a map where its keys are either "Key A", 2.5, or "Michael Keaton"...
 >
 > B) ...create a new map such that the list's values are its keys and its values are the corresponding indices from the original list...
 >
 > C) ...in a copy of B's map, replace pairs whose values are between 1 and 3 with a new slice of key-value pairs...
 >
 > D) ...and create a copy of C, concatenating the array to itself 4 times, and putting its key-value pairs in a 2x2x2x2 matrix...

 ...all the while ensuring no object is cloned in the process, you could use ***classical go*** or ***gostack***:

<h3 name = "pseudocode">...pseudocode outline</h3>

```
// INIT
start <"Key A" : 40, "Bad Key" : "Bad Value", "Key A" : "Hello", 2.5 : 40, "Michael Keaton" : 520>
searchKeys <"Key A", 2.5, "Michael Keaton">
pairsToInsert <"I" : "Am new", "To" : "This set">
 
// TASK A
=> taskA <40, "Hello", 520>
 
// TASK B
=> taskB <40 : 0, "Hello" : 2, 520 : 4>

// TASK C
=> taskC <40 : 0, "I" : "Am new", "To" : "This set", 520 : 4>

// TASK D ('k' represents a key, just an abstraction for legibility)
=> taskD < < << k, k >, < k, k >>, << k, k >, < k, k >> >,
           < << k, k >, < k, k >>, << k, k >, < k, k >> > >
```

<h2>Let's see how quickly we can do this using...</h2>

<h3 name = "classical">...classical go</h3>

```
// INIT
start := map[any]any {"1_Key A" : 40, "Bad Key" : "Bad Value", "2_Key A" : "Hello", 2.5 : 40, "Michael Keaton" : 520} // can't have same key twice so need x_KEYNAME
searchKeys := []any {"Key A", 2.5, "Michael Keaton"}
pairsToInsert := map[any]any {"I" : "Am new", "To" : "This set"}
 
// TASK A
var taskA []any
for i := range len(start) {
    k := start[i] // circumvent for loop cloning of k
    for _, search := range searchKeys {
        if k == search {
            alreadyInArray := false
            for _, v := range taskA {
                if v == k {
                    alreadyInArray = true
                    break
                }
            }
            if !alreadyInArray {
                taskA = append(taskA, k)
            }
            break
        }
    }
}
 
// TASK B
var taskB map[any]any
i = 0
for k, v := range start {
    for j := range len(taskA) {
        a := taskA[j] // circumvent for loop cloning of a
        if a == v {
            taskB[a] = i
        }
    }
    i++
}

// TASK C
var taskC map[any]any
for k, v := range taskB {
    if 1 < v.(int) && v.(int) < 4 {
        for k3 := range pairsToInsert {
            k4, v4 := pairsToInsert[k3] // circumvent for-loop cloning of keys/vals
            taskC[k4] = v4
        }
    } else {
        taskC[k] = v
    }
}

// TASK D
var taskD [][][][]any
type kvPair struct {
    key any
    val any
}
for a := 0; a < 2; a++ { // using for-loops rather than recursion to make the code more legible 
    for b := 0; b < 2; b++ {
        for c := 0; c < 2; c++ {
            for d := 0; d < 2; d++ {
                k, v := taskC[a+(b*2)]
                kClone := reflect.New(reflect.ValueOf(k).Elem().Type()).Interface() // clone our key
                vClone := reflect.New(reflect.ValueOf(v).Elem().Type()).Interface() // clone our value
                taskD[a][b][c][d] = kvPair {key: kClone, val: vClone}
            }
        }
    }
}
```

`lines: 61`

<h3 name = "gostack">...gostack</h3>

```
// INIT
start := MakeStack(map[any]any {"Key A" : 40, "Bad Key" : "Bad Value", "Key A" : "Hello", 2.5 : 40, "Michael Keaton" : 520})
searchKeys := MakeStack([]any {"Key A", 2.5, "Michael Keaton"})
pairsToInsert := MakeStack(map[any]any {"I" : "Am new", "To" : "This set"})

// TASK A
taskA := start.GetMany(FIND_Keys, searchKeys, RETURN_Vals).Unique(TYPE_Val)

// TASK B
taskB := MakeStack(taskA, start.GetMany(FIND_Vals, taskA, RETURN_Vals).Unique(TYPE_Val))

// TASK C
taskC := taskB.Clone().Replace(RETURN_Cards, pairsToInsert, FIND_Lambda, func(stack *Stack, card *Card) bool {
    return 1 < card.Val.(int) && card.Val.(int) < 3
})

// TASK D
taskD := MakeStackMatrix(taskC.Clone().Duplicate(4), nil, []int{2, 2, 2, 2})
```

`lines: 9`

<h2>Conclusion</h2>

***gostack*** won the race!  It took 6.7 times fewer lines than ***classical go*** (9 compared to 61), in turn saving 52 lines of space (excluding comments and empty lines).

 [> Return to **Glossary**](../README.md)