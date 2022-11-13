package testing

import (
	"fmt"
	. "github.com/gabetucker2/gostack"//lint:ignore ST1001, ignore warning
)

func RaceNative() {

	fmt.Println("- BEGINNING NATIVE GO RACE")

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
	for i := 0; i < 16; i++ { // recursion would be unnecessarily complicated for this data structure
		var a, b, c, d int // convert i into 2x2x2x2 sequence: ([0][0][0][0], [0][0][0][1], [0][0][1][0], ..., [1][1][1][1])
		if i%2 == 1 {d = 1}
		if ((i-d)/2)%2 == 1 {c = 1}
		if (4 <= i && i < 8) || (12 <= i && i < 16) {b = 1}
		if 8 <= i {a = 1}
		taskDKeys[a][b][c][d] = taskCKeys2[i]
		taskDVals[a][b][c][d] = taskCVals2[i]
	}

	// PRINTS
	fmt.Print("A: ")
	fmt.Println(taskA)
	fmt.Print("B Keys: ")
	fmt.Println(taskBKeys)
	fmt.Print("B Vals: ")
	fmt.Println(taskBVals)
	fmt.Print("C Keys: ")
	fmt.Println(taskCKeys)
	fmt.Print("C Vals: ")
	fmt.Println(taskCVals)
	fmt.Print("D Keys: ")
	fmt.Println(taskDKeys)
	fmt.Print("D Vals: ")
	fmt.Println(taskDVals)

}

func RaceGostack() {

	fmt.Println("- BEGINNING GOSTACK RACE")

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

	// PRINTS
	taskA.Print("A")
	taskB.Print("B")
	taskC.Print("C")
	taskD.Print("D")

}
