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

type FIND int
type REPLACE int
type RETURN int
type TYPE int
type ORDER int
type DEREFERENCE int
type CLONE int
type DEEPSEARCH int
type COMPARE int
type PASS int
type OVERRIDE int
type REPEAT int

const (
	FIND_First FIND = iota
	FIND_Last
	FIND_Idx
	FIND_Key
	FIND_Val
	FIND_KeyVal
	FIND_Card
	FIND_Coords
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
	RETURN_Idxs RETURN = iota
	RETURN_Keys
	RETURN_Vals
	RETURN_Cards
	RETURN_Adrs
	RETURN_Stacks
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
	DEREFERENCE_None DEREFERENCE = iota
	DEREFERENCE_Both
	DEREFERENCE_Found
	DEREFERENCE_This
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
	PASS_Cards PASS = iota
	PASS_Substacks
	PASS_Both
)

const (
	OVERRIDE_False OVERRIDE = iota
	OVERRIDE_True
)

const (
	REPEAT_False REPEAT = iota
	REPEAT_True
)

func setFINDDefaultIfNil(findType *any) {
	if *findType == nil {
		*findType = FIND_Last
	}
}

func setRETURNDefaultIfNil(returnType *any) {
	if *returnType == nil {
		*returnType = RETURN_Vals
	}
}

func setORDERDefaultIfNil(orderType *any) {
	if *orderType == nil {
		*orderType = ORDER_After
	}
}

func setDEREFERENCEDefaultIfNil(dereferenceType *any) {
	if *dereferenceType == nil {
		*dereferenceType = DEREFERENCE_None
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

func setOVERRIDEDefaultIfNil(overrideType *any) {
	if *overrideType == nil {
		*overrideType = OVERRIDE_False
	}
}
