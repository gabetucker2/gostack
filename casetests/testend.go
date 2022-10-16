package casetests

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gabetucker2/gogenerics"
	. "github.com/gabetucker2/gostack" //lint:ignore ST1001 — we would like to dot import gostack
)

/** Test whether stack equals array or map
 (Incomplete documentation)

 @requires `_vals` and `_keys` are either nil or of type{[]any}, and 
 */
func test_StackEqualArrayOrMap(stack *Stack, _vals, _keys any, _ma any) bool {

	// init
	if _vals == nil && _keys == nil && _ma == nil {
		for i, c := range stack.Cards {
			if c.Key != nil || c.Val != nil {
				fmt.Print("-     DETAIL: ")
				fmt.Printf(gogenerics.IfElse(c.Key != nil, "stack.Cards[%v].Key != nil", "stack.Cards[%v].Val != nil").(string), i)
				return false
			}
		}
	} else {
		var keys, vals []any
		var ma map[any]any
		if _vals != nil { vals = gogenerics.UnpackArray(_vals) }
		if _keys != nil { keys = gogenerics.UnpackArray(_keys) }
		if _ma != nil {
			ma = gogenerics.UnpackMap(_ma);
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
			} else if (_vals != nil && c.Val != vals[i]) || (_keys != nil && c.Key != keys[i]) {
				fmt.Print("-     DETAIL: ")
				fmt.Printf(
					gogenerics.IfElse(_vals != nil && vals[i] != c.Val, "stack.Cards[%v].Val (%v) != expected Val ", "stack.Cards[%v].Key (%v) != expected Key ").(string),
					i,
					gogenerics.IfElse(_vals != nil && vals[i] != c.Val, c.Val, c.Key),
				)
				if _vals != nil && vals[i] != c.Val {
					fmt.Printf("(%v)\n", vals[i])
				} else {
					fmt.Printf("(%v)\n", keys[i])
				}
				return false
			}
			
		}
	}
	
	return true
	
}

/** Return whether len(cards) == cards.Size and whether depth measures are accurate
 NOTE: depth is one variable that's just optional
 */
 func test_StackProperties(stack *Stack, size []int, depth ...int) (test bool) {

	/*

	PSEUDOCODE OUTLINE:

	test = true

	if depth == 0
		depth = 1 // initialize
	
	if stack.Size != len(stack.Cards) || stack.Depth != depth
		test = false
	else
		for each card in stack
			if card has a stack
				test = recursive call(card.Val.(*Stack), size[1:], depth - 1)
			if !test
				break
	
	return

	*/

	test = true

	// initialize depth on first call
	if len(depth) == 0 {
		depth = append(depth, 1)
	}

	if size[0] != stack.Size || stack.Size != len(stack.Cards) || stack.Depth != depth[0] {
		// if invalid condition on test
		fmt.Print("-     DETAIL: ")
		fmt.Printf(
			gogenerics.IfElse(size[0] != stack.Size, "size[0] (%v) does not match stack.Size (%v)\n", gogenerics.IfElse(stack.Size != len(stack.Cards), "stack.Size (%v) does not match len(stack.Cards) (%v)\n", "stack.Depth (%v) does not match expected depth (%v)\n")).(string),
			gogenerics.IfElse(size[0] != stack.Size, size[0], gogenerics.IfElse(stack.Size != len(stack.Cards), stack.Size, stack.Depth)),
			gogenerics.IfElse(size[0] != stack.Size, stack.Size, gogenerics.IfElse(stack.Size != len(stack.Cards), len(stack.Cards), depth[0])),
		)
		test = false
	} else {
		// else iterate through cards in stack, test on each of them

		for i := range stack.Cards {
			c := stack.Cards[i]

			switch c.Val.(type) {
			case *Stack:
				test = test_StackProperties(c.Val.(*Stack), size[1:], depth[0] - 1) // TODO: this seems rather sus... double-check this
			}

			if !test { break }
		}

	}
	
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
			fmt.Printf("-     DETAIL: c.Idx (%v) != i (%v)", c.Idx, i)
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
	fmt.Println(out + " " + funcName + "()            ~ [" + strconv.Itoa(len(conditions)) + " cases]")

}
