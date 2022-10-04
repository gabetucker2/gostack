package gostack

type Card struct {
	Idx int
	Key any
	Val any
}

type Stack struct {
	Cards []*Card
	Size  int
	Depth int
}

type RETURN int
type FIND int
type REPLACE int
type TYPE int
type ORDER int
type POINTER int
type CLONE int
type DEEPSEARCH int
type COMPARE int
type PASS int
//type PRINT int

const (
	RETURN_Idxs RETURN = iota
	RETURN_Keys
	RETURN_Vals
	RETURN_Cards
)

const (
	FIND_First FIND = iota
	FIND_Last
	FIND_Idx
	FIND_Key
	FIND_Val
	FIND_Card
	FIND_Size
	FIND_Depth
	FIND_Slice
	FIND_All
	FIND_Lambda
)

const (
	REPLACE_Key REPLACE = iota
	REPLACE_Val
	REPLACE_Card
	REPLACE_Stack
	REPLACE_Lambda
)

const (
	TYPE_Key TYPE = iota
	TYPE_Val
)

const (
	ORDER_Before ORDER = iota
	ORDER_After
)

const (
	POINTER_False POINTER = iota
	POINTER_True
)

const (
	CLONE_False CLONE = iota
	CLONE_True
)

const (
	DEEPSEARCH_False DEEPSEARCH = iota
	DEEPSEARCH_True
)

const (
	COMPARE_False COMPARE = iota
	COMPARE_True
)

const (
	PASS_False PASS = iota
	PASS_True
)

/*const (
	PRINT_False PRINT = iota
	PRINT_True
)*/

func setRETURNDefaultIfNil(returnType *any) {
	if *returnType == nil {
		*returnType = RETURN_Vals
	}
}

func setFINDDefaultIfNil(findType *any) {
	if *findType == nil {
		*findType = FIND_First
	}
}

func setORDERDefaultIfNil(orderType *any) {
	if *orderType == nil {
		*orderType = ORDER_Before
	}
}

func setPOINTERDefaultIfNil(pointerType *any) {
	if *pointerType == nil {
		*pointerType = POINTER_False
	}
}

func setCLONEDefaultIfNil(cloneType *any) {
	if *cloneType == nil {
		*cloneType = CLONE_True
	}
}

func setDEEPSEARCHDefaultIfNil(deepSearchType *any) {
	if *deepSearchType == nil {
		*deepSearchType = DEEPSEARCH_True
	}
}

func setDepthDefaultIfNil(depth *any) {
	if *depth == nil {
		*depth = -1
	}
}

func setCOMPAREDefaultIfNil(compareType *any) {
	if *compareType == nil {
		*compareType = COMPARE_True
	}
}

// func setPRINTDefaultIfNil(compareType *any) {
// 	if *compareType == nil {
// 		*compareType = PRINT_False
// 	}
// }
