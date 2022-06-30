package tutorials

import (
	"fmt"

	. "github.com/gabetucker2/gostack"
)

func Matrices() {

	fmt.Println("tutorials/Matrics()")

	//////////////////////////////
	// ORIGINAL METHOD

	// setup
	k1, k2, k3, k4, k5, k6, k7, k8, k9, k10, k11, k12, k13, k14, k15, k16, k17, k18 := [strings representing keys]
	e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16, e17, e18 := [ints representing values]

	x1 := MakeStack([]*string {k1, k2, k3}, []*int {e1, e2, e3})
	x2 := MakeStack([]*string {k4, k5, k6}, []*int {e4, e5, e6})
	x3 := MakeStack([]*string {k7, k8, k9}, []*int {e7, e8, e9})
	y1 := MakeStack(*[]Stack {x1, x2, x3})

	x4 := MakeStack([]*string {k10, k11, k12}, []*int {e10, e11, e12})
	x5 := MakeStack([]*string {k13, k14, k15}, []*int {e13, e14, e15})
	x6 := MakeStack([]*string {k16, k17, k18}, []*int {e16, e17, e18})
	y2 := MakeStack(*[]Stack {x4, x5, x6})

	matrix := MakeStack(*[]Stack {y1, y2})

	// returns
	matrix                                                    // matrix
	matrix.Get(FIND_Idx, 1)                                   // y2
	matrix.Get(FIND_Idx, 1).Get(FIND_Idx, 0)                  // x4
	matrix.Get(FIND_Idx, 1).Get(FIND_Idx, 0).Get(FIND_Idx, 2) // Card{k11, e11}

	//////////////////////////////
	// NOVEL METHOD

	// setup
	var keys []string {k1, k2, k3, k4, k5, k6, k7, k8, k9, k10, k11, k12, k13, k14, k15, k16, k17, k18}
	var vals []int {e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16, e17, e18}

	matrix := MakeStackMatrix([3]int {3, 3, 2}, keys, vals)

	// returns
	// - fast method
	matrix                                                                            // matrix

	matrix.GetInMatrix(1)                                                             // y2
	matrix.GetInMatrix(1, 0)                                                          // x4
	matrix.GetInMatrix(1, 0, 1)                                                       // Card{k11, e11}

	matrix.GetInMatrix([2]int {0, 1})                                                 // Stack.Cards{Card{y1}, Card{y2}}
	matrix.GetInMatrix([2]int {0, 1}, 0)                                              // Stack.Cards{Card{x1}, Card{x4}}
	matrix.GetInMatrix(0, [2]int {0, 1})                                              // Stack.Cards{Card{x1}, Card{x2}}
	matrix.GetInMatrix(1, 0, [2]int {0, 2})                                           // Stack.Cards{c10, c11, c12} (cx = card with Val ex and Key kx)

	// - slow method
	matrix.Get(FIND_Key, k11, nil, nil, nil, nil, DEEPSEARCH_True)                    // Card{k11, e11}
	matrix.Get(FIND_Keys, []string{} {k11, k17}, nil, nil, nil, nil, DEEPSEARCH_TRUE) // Stack.Cards{Card{k11, e11}, Card{k17, e17}}

}
