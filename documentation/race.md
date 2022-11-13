![Banner](../media/gostack_SmallerTransparent.png)

<h1>Race</h1>

<h2>Introduction</h2>

 Welcome!  This consists of a race between ***native Go*** and ***gostack*** to see who can perform the same data management tasks faster.

 <h4>Jump To</h4>

 > [pseudocode outline](#pseudocode)

 > [native Go](#native)

 > [gostack](#gostack)

 <h3>Assuming you would like to...</h3>

 > A) ...make an array structure `taskA` representing a non-duplicating set of values from a map structure whose keys are in `searchKeys`
 >
 > B) ...create a map structure `taskB` whose keys are `taskA`'s values and whose values are the corresponding indices from the original map structure
 >
 > C) ...in `taskB`'s map, replace pairs whose values are between 1 and 3 with a new slice of key-value pairs, making this new map `taskC`
 >
 > D) ...and appending a clone of `taskC` to itself 3 times, thereafter fitting its key-value pairs to 2x2x2x2 matrix `taskD`

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

// TASK D
=> taskD < < << 40 : 0, "I" : "Am new" >, < "To" : "This set", 520 : 4 >>,
             << 40 : 0, "I" : "Am new" >, < "To" : "This set", 520 : 4 >> >,
           < << 40 : 0, "I" : "Am new" >, < "To" : "This set", 520 : 4 >>,
             << 40 : 0, "I" : "Am new" >, < "To" : "This set", 520 : 4 >> > >
```

---

<h2>Let's see how quickly we can do this using...</h2>

<h3 name = "native">...Go native (completed as quickly as possible, even if not "the best" way)</h3>

```
// INIT
startKeys := []any {"Key A", "Bad Key", "Key A", 2.5, "Michael Keaton"} // must separate maps like so since native Go maps are unordered
startVals := []any {40, "Bad Value", "Hello", 40, "520"}
searchKeys := []any {"Key A", 2.5, "Michael Keaton"}
keysToInsert := []any {"I", "To"}
valsToInsert := []any {"Am new", "This set"}

// TASK A
var taskA []any
for i := 0; i < len(startKeys); i++ {
    k := startKeys[i]
    v := startVals[i]
    cont := false
    for _, searchKey := range searchKeys {
        if k == searchKey {
            cont = true
            break
        }
    }
    if cont {
        for _, otherV := range taskA {
            if otherV == v {
                cont = false
                break
            }
        }
        if cont {
            taskA = append(taskA, v)
        }
    }
}

// TASK B
var taskBKeys, taskBVals []any
for _, taskAVal := range taskA {
    for i := 0; i < len(startVals); i++ {
        v := startVals[i]
        if taskAVal == v {
            taskBKeys = append(taskBKeys, v)
            taskBVals = append(taskBVals, i)
            break
        }
    }
}

// TASK C
var taskCKeys, taskCVals []any
for i := 0; i < len(taskBKeys); i++ {
    k := taskBKeys[i]
    v := taskBVals[i]
    if 1 < v.(int) && v.(int) < 3 {
        for j := 0; j < len(keysToInsert); j++ {
            inK := keysToInsert[j]
            inV := valsToInsert[j]
            taskCKeys = append(taskCKeys, inK)
            taskCVals = append(taskCVals, inV)
        }
    } else {
        taskCKeys = append(taskCKeys, k)
        taskCVals = append(taskCVals, v)
    }
}

// TASK D
taskCKeys2 := append(taskCKeys, taskCKeys...)
taskCKeys2 = append(taskCKeys2, taskCKeys2...)
taskCVals2 := append(taskCVals, taskCVals...)
taskCVals2 = append(taskCVals2, taskCVals2...)
taskDKeys := [][][][]any{{{{nil, nil}, {nil, nil}}, {{nil, nil}, {nil, nil}}}, {{{nil, nil}, {nil, nil}}, {{nil, nil}, {nil, nil}}}}
taskDVals := [][][][]any{{{{nil, nil}, {nil, nil}}, {{nil, nil}, {nil, nil}}}, {{{nil, nil}, {nil, nil}}, {{nil, nil}, {nil, nil}}}}
for i := 0; i < 16; i++ { // recursion would only work until 3 layers down
    var a, b, c, d int // convert i into 2x2x2x2 sequence: ([0][0][0][0], [0][0][0][1], [0][0][1][0], ..., [1][1][1][1])
    if i%2 == 1 {d = 1}
    if ((i-d)/2)%2 == 1 {c = 1}
    if (4 <= i && i < 8) || (12 <= i && i < 16) {b = 1}
    if 8 <= i {a = 1}
    taskDKeys[a][b][c][d] = taskCKeys2[i]
    taskDVals[a][b][c][d] = taskCVals2[i]
}
```

`lines: 69`

<h3 name = "gostack">...gostack</h3>

```
// INIT
start := MakeStack([]any {"Key A", "Bad Key", "Key A", 2.5, "Michael Keaton"}, []any {40, "Bad Value", "Hello", 40, "520"})
searchKeys := MakeStack([]any {"Key A", 2.5, "Michael Keaton"})
pairsToInsert := MakeStack([]any {"I", "To"}, []any {"Am new", "This set"})

// TASK A
taskA := start.GetMany(FIND_Key, searchKeys).Unique()

// TASK B
taskB := MakeStack(taskA, start.GetMany(FIND_KeyVal, taskA, RETURN_Idxs))

// TASK C
taskC := taskB.Clone().Update(REPLACE_Card, pairsToInsert, FIND_Lambda, func(card *Card) bool {
    return 1 < card.Val.(int) && card.Val.(int) < 3
})

// TASK D
taskD := MakeStackMatrix(taskC.Clone().Duplicate(4), taskC.Clone().Duplicate(4).SwitchKeysVals(), []int{2, 2, 2, 2})
```

`lines: 9`

---

<h2>Conclusion</h2>

***gostack*** won the race!  Excluding comments and empty lines, it took over 7 times fewer lines than ***native Go*** (9 compared to 69), in turn saving 60 lines of space.  See our [races script here](../testing/races.go), and feel free to test it out yourself!

*Note that I wrote the native Go comparison code to be as clear and concise as possible, but if there are any places which can be made made more concise, then please let me know in our [Discord server](https://discord.gg/NmxxcKBVBU).*

---

 [> How fast is gostack compared to native Go?](benchmark.md)

 [> Return to glossary](../README.md)