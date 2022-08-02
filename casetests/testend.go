package casetests

import (
	"fmt"
	"reflect"
	"strconv"

	//lint:ignore ST1001 — we would like to dot import gostack
	. "github.com/gabetucker2/gostack"
)

/** I hate Go.
 This exists only so that warnings are no longer suggested for private functions/variables used by other scripts in this package.
 */
func removeWarning(...any) { test_SampleStack() }

// test variables (test with these variables only after MakeCard's test)

var testCardA *Card
var testCardB *Card
var testCardC *Card
var testCardD *Card

/** Initialize test variables */
func test_Setup() {

	testCardA = MakeCard("Card A") // in sample stack
	testCardB = MakeCard("Card B") // in sample stack
	testCardC = MakeCard("Card C") // in sample stack
	testCardD = MakeCard("Card D") // out of sample stack

	removeWarning(testCardA, testCardB, testCardC, testCardD)

}

/** Returns, from any array type, a version of that array which is converted to type []any
 (These functions are repeated in backend.  Likely fix this in the future to remove redundancy.)

 @param `arr` type{any}
 @return type{[]any}
 @requires `arr` is an array
 */
func test_UnpackArray(arr any) []any {
    valType := reflect.ValueOf(arr)
    new := make([]any, valType.Len())
    for i := 0; i < valType.Len(); i++ {
        new[i] = valType.Index(i).Interface()
    }
    return new
}

/** Returns, from any map type, a version of that map which is converted to type map[any]any
 (These functions are repeated in backend.  Likely fix this in the future to remove redundancy.)

 @param `arr` type{any}
 @return type{[]any}
 @requires `arr` is an array
 */
 func test_UnpackMap(s any) map[any]any {
    v := reflect.ValueOf(s)
    m := make(map[any]any, v.Len())
    for _, k := range v.MapKeys() {
		m[k.Interface()] = v.MapIndex(k).Interface()
    }
    return m
}

/** Test whether stack equals array or map
 (Incomplete documentation)

 @requires `_vals` and `_keys` are either nil or of type{[]any}, and 
 */
func test_StackEqualArrayOrMap(stack *Stack, _vals, _keys any, _ma any) bool {

	// init
	var keys, vals []any
	var ma map[any]any
	if _vals != nil { vals = test_UnpackArray(_vals) }
	if _keys != nil { keys = test_UnpackArray(_keys) }
	if _ma != nil {
		ma = test_UnpackMap(_ma);
		mapReflectVal := reflect.ValueOf(ma)
		for _, k := range mapReflectVal.MapKeys() {
			keys = append(keys, k.Interface())
			vals = append(vals, mapReflectVal.MapIndex(k).Interface())
		}
	}

	// return true iff stack does not raise any mismatches
	for i := range stack.Cards {
		c := stack.Cards[i]

		// if testing a map (no need for match history tracking because go map can't have two of same key)
		if _ma != nil {
			matchExists := false
			for j := range keys {
				if keys[j] == c.Key && vals[j] == c.Val {
					matchExists = true
					break
				}
			}
			if !matchExists {
				return false
			}

		// if testing an array
		} else if (_vals != nil && vals[i] != c.Val) || (_keys != nil && keys[i] != c.Key) {
			 return false
		}
		
	}
	return true
}

/** Return whether len(cards) == cards.size */
func test_LenAndSize(stack *Stack, size []int) (test bool) {

	if size[0] == 0 { test = true }
	for _, s := range size { // get size[0]
		fmt.Printf("s: %d\n", s)
		for i := 0; i < s; i++ {
			fmt.Printf("i: %d, %t\n", i, !test && stack.Depth > 0)
			if !test && stack.Depth > 0 {
				fmt.Println("PASS")
				if len(size) > 1 {
					if stack.Size == len(stack.Cards) {
						test = test_LenAndSize(stack, size[1:])	
					}
				} else {
					test = true
					break
				}
			}
		}
		//if true {break}
	}
	fmt.Printf("test: %t\n", test)
	return

}

/** Return whether the indices correspond to their position in a stack */
func test_IdxsAreGood(stack *Stack) bool {

	good := true
	for i := range stack.Cards {
		c := stack.Cards[i]
		switch c.Val.(type) {
		case *Stack:
			good = test_IdxsAreGood(c.Val.(*Stack))
		}
		if c.Idx != i {
			good = false
			break
		}
	}
	return good

}

func test_Start(funcName string, showTestText bool) {

	// print TESTING line only if showTestText var set to true
	if showTestText {
		fmt.Println("-   TESTING " + funcName + "()")
	}

	test_Setup()

}

func test_End(funcName string, conditions []bool) {

	// set test to -1 (true) by default
	test := -1

	// test each condition and update test flag to index of condition which failed
	for i, c := range conditions {
		if !c {
			test = i
			break
		}
	}

	// set SUCCESS/FAILURE based on which condition, if any, failed
	out := "-   "
	if test == -1 {
		out += "SUCCESS"
	} else {
		out += "FAILURE AT CONDITION #" + strconv.Itoa(test+1) + " in"
	}

	// print all the data together
	fmt.Println(out + " " + funcName + "()")

}

/** Make a sample stack of cards */
func test_SampleStack() *Stack {

	// make a sample stack of form <"Card A", "Card B", "Card C">
	stack := MakeStack()

	// create stack (don't use .Add() function, or else you'll have to case test)
	stack.Cards = append(stack.Cards, testCardA)
	stack.Cards = append(stack.Cards, testCardB)
	stack.Cards = append(stack.Cards, testCardC)
	
	return stack

}
