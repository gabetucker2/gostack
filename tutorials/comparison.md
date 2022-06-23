 <h1>Comparison</h1>

 Assume you would A) like to make a list representing a non-duplicating set of values from a map where its keys are either "Key A", 2.5, or "Michael Keaton".  You would B) then like to create a new map such that the list's values are its keys and its values are the corresponding indices from the original list.  Finally, you would C) like, in a copy of B's map, to replace pairs whose values are between 1 and 3 with a new array of key-value pairs.  Ensure no object is cloned in the process.  In pseudocode...

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
 ```

 Now, let's see how quickly we can do this using...

...***classical go***
 ```
 // INIT
 start := map[interface{}]interface{} {"Key A" : 40, "Bad Key" : "Bad Value", "Key A" : "Hello", 2.5 : 40, "Michael Keaton" : 520}
 searchKeys := []interface{} {"Key A", 2.5, "Michael Keaton"}
 pairsToInsert := map[interface{}]interface{} {"I" : "Am new", "To" : "This set"}
 
 // TASK A
 var taskA []interface{}
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
 var taskB map[interface{}]interface{}
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
 var taskC map[interface{}]interface{}
 for k, v := range taskB {
    k2, v2 := taskB[k] // circumvent for loop cloning
    if 1 < v && v < 4 {
        for k3 := range pairsToInsert {
            k4, v4 := pairsToInsert[k3] // circumvent for loop cloning
            taskC[k4] = v4
        }
    } else {
        taskC[k2] = v2
    }
 }
 ```
 `lines: 45`

...***gostack***
 ```
 // INIT
 start := MakeStack(STRUCTURE_Map, map[interface{}]interface{} {"Key A" : 40, "Bad Key" : "Bad Value", "Key A" : "Hello", 2.5 : 40, "Michael Keaton" : 520})
 searchKeys := MakeStack(STRUCTURE_Arr, []interface{} {"Key A", 2.5, "Michael Keaton"})
 pairsToInsert := MakeStack(STRUCTURE_Map, map[interface{}]interface{} {"I" : "Am new", "To" : "This set"})

 // TASK A
 taskA := start.Get(RETURN_Vals, POSITION_Keys, searchKeys).Unique(TYPE_Val)

 // TASK B
 taskB := MakeStack(STRUCTURE_Map, taskA, start.Get(RETURN_Idxs, POSITION_Vals, taskA).Unique(TYPE_Val))

 // TASK C
 func gostack_ValInRange(stack *Stack, card *Card) {
    v := card.val.(int)
	return 1 < v && v < 3
 }

 taskC := taskB.Clone().Replace(pairsToInsert, RETURN_Stack, POSITION_Lambda, gostack_ValInRange)
 ```
 `lines: 10`