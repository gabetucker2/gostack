package gostack

type Card struct {
	Idx int
	Key any
	Val any
}

type Stack struct {
	Cards  []*Card
	Size   int
	Height int
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
type OVERRIDE int
type ACTION int

//type PRINT int

const (
	RETURN_Idxs RETURN = iota
	RETURN_Keys
	RETURN_Vals
	RETURN_Cards
	RETURN_Stacks
)

const (
	FIND_First FIND = iota
	FIND_Last
	FIND_Idx
	FIND_Key
	FIND_Val
	FIND_Card
	FIND_Size
	FIND_Height
	FIND_Slice
	FIND_All
	FIND_Lambda
)

const (
	REPLACE_Key REPLACE = iota
	REPLACE_Val
	REPLACE_Card
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

const (
	OVERRIDE_False OVERRIDE = iota
	OVERRIDE_True
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
		*findType = FIND_Last
	}
}

func setORDERDefaultIfNil(orderType *any) {
	if *orderType == nil {
		*orderType = ORDER_After
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

func setDEEPSEARCHDefaultIfNil(heightSearchType *any) {
	if *heightSearchType == nil {
		*heightSearchType = DEEPSEARCH_True
	}
}

func setHeightDefaultIfNil(height *any) {
	if *height == nil {
		*height = -1
	}
}

func setCOMPAREDefaultIfNil(compareType *any) {
	if *compareType == nil {
		*compareType = COMPARE_True
	}
}

func setOVERRIDEDefaultIfNil(overrideType *any) {
	if *overrideType == nil {
		*overrideType = OVERRIDE_False
	}
}

// func setPRINTDefaultIfNil(compareType *any) {
// 	if *compareType == nil {
// 		*compareType = PRINT_False
// 	}
// }
