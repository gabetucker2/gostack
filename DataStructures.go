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
type MATCHBY int
type CLONE int
type DEEPSEARCH int

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
	FIND_Idxs
	FIND_IdxsStack
	FIND_Key
	FIND_Keys
	FIND_KeysStack
	FIND_Val
	FIND_Vals
	FIND_ValsStack
	FIND_Card
	FIND_Cards
	FIND_CardsStack
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
	TYPE_Card
)

const (
	ORDER_Before ORDER = iota
	ORDER_After
)

const (
	MATCHBY_Object MATCHBY = iota
	MATCHBY_Reference
)

const (
	CLONE_True CLONE = iota
	CLONE_False
)

const (
	DEEPSEARCH_True DEEPSEARCH = iota
	DEEPSEARCH_False
)

/*func setRETURNDefaultIfNil(returnType any) {
	if returnType == nil {
		returnType = RETURN_Cards
	}
}*/

func setFINDDefaultIfNil(findType any) {
	if findType == nil {
		findType = FIND_First
	}
}

func setORDERDefaultIfNil(orderType any) {
	if orderType == nil {
		orderType = ORDER_Before
	}
}

func setMATCHBYDefaultIfNil(matchByType any) {
	if matchByType == nil {
		matchByType = MATCHBY_Object
	}
}

func setCLONEDefaultIfNil(cloneType any) {
	if cloneType == nil {
		cloneType = CLONE_False
	}
}

func setDEEPSEARCHDefaultIfNil(deepSearchType any) {
	if deepSearchType == nil {
		deepSearchType = DEEPSEARCH_False
	}
}

func setDepthDefaultIfNil(depth any) {
	if depth == nil {
		depth = 1
	}
}
