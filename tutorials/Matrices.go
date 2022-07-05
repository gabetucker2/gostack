package tutorials

import (
	"fmt"
	//. "github.com/gabetucker2/gostack"
)

func Matrices() {

	fmt.Println("tutorials/Matrics()")

	///*

	//////////////////////////////
	// ORIGINAL METHOD

	// setup
	k1, k2, k3, k4, k5, k6, k7, k8, k9, k10, k11, k12, k13, k14, k15, k16, k17, k18 := [strings representing keys]
	e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16, e17, e18 := [ints representing values]

	x1 := MakeStack([]*string {k1, k2, k3}, []*int {e1, e2, e3})
	x2 := MakeStack([]*string {k4, k5, k6}, []*int {e4, e5, e6})
	x3 := MakeStack([]*string {k7, k8, k9}, []*int {e7, e8, e9})
	y1 := MakeStack([]*Stack {x1, x2, x3})

	x4 := MakeStack([]*string {k10, k11, k12}, []*int {e10, e11, e12})
	x5 := MakeStack([]*string {k13, k14, k15}, []*int {e13, e14, e15})
	x6 := MakeStack([]*string {k16, k17, k18}, []*int {e16, e17, e18})
	y2 := MakeStack([]*Stack {x4, x5, x6})

	matrix := MakeStack([]*Stack {y1, y2})

	// returns
	matrix                                                    // matrix
	matrix.Get(FIND_Idx, 1)                                   // y2
	matrix.Get(FIND_Idx, 1).Get(FIND_Idx, 0)                  // x4
	matrix.Get(FIND_Idx, 1).Get(FIND_Idx, 0).Get(FIND_Idx, 1) // Card{k11, e11}

	//////////////////////////////
	// NOVEL METHOD

	// setup
	var keys []string {k1, k2, k3, k4, k5, k6, k7, k8, k9, k10, k11, k12, k13, k14, k15, k16, k17, k18}
	var vals []int {e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16, e17, e18}

	var ma [][][]string {[][]int{{k1, k2, k3},{k4:e4, k5:e5, k6:e6},{}}}

	// IF keys/vals are just 1D arrays
	matrix := MakeStackMatrix(keys, vals, [3]int {2, 3, 3})
	// OR, if keys/vals are [][][]string/[][][]vals, i.e. already matrices
	matrix := MakeStackMatrix(ma)

	// returns
	// - fast method
	matrix                                  // matrix

	// (cx = card with Val ex and Key kx)

	matrix.StripStackMatrix(1)                   // y2
	matrix.StripStackMatrix(1, 0)                // x4
	matrix.StripStackMatrix(1, 0, 1)             // c11

	matrix.StripStackMatrix([2]int {0, 1})       // Stack.Cards{Card{y1}, Card{y2}}
	matrix.StripStackMatrix([2]int {0, 1}, 0)    // Stack.Cards{Card{x1}, Card{x4}}
	matrix.StripStackMatrix(0, [2]int {0, 1})    // Stack.Cards{Card{x1}, Card{x2}}
	matrix.StripStackMatrix(1, 0, [2]int {0, 2}) // Stack.Cards{c10, c11, c12}

	// - slow method

	// nil
	matrix.Get(FIND_Key, k11)

	// c11
	matrix.Get(FIND_Key, k11, nil, nil, nil, nil, DEEPSEARCH_True)

	// Stack.Cards{c11, c17}
	matrix.GetMany(FIND_Keys, []string{} {k11, k17}, nil, RETURN_Cards, nil, nil, nil, DEEPSEARCH_TRUE)

	// Stack.Cards{c11, c17}
	matrix.GetMany(FIND_Keys, []string{} {k11, k17}, nil, RETURN_Cards, nil, nil, nil, DEEPSEARCH_TRUE, -1)

	// Stack.Cards{}
	matrix.GetMany(FIND_Keys, []string{} {k11, k17}, nil, RETURN_Cards, nil, nil, nil, DEEPSEARCH_TRUE, 1)

	// (assuming x1, y1, etc are cards whose vals are stacks)

	// matrix
	matrix.GetMany(FIND_Cards, MakeStack(matrix, x1, y1), nil, RETURN_Cards, nil, nil, nil, DEEPSEARCH_TRUE, 0)

	// y1
	matrix.GetMany(FIND_Cards, MakeStack(matrix, x1, y1), nil, RETURN_Cards, nil, nil, nil, DEEPSEARCH_TRUE, 1)

	// x1
	matrix.GetMany(FIND_Cards, MakeStack(matrix, x1, y1), nil, RETURN_Cards, nil, nil, nil, DEEPSEARCH_TRUE, 2)
	//*/

}
