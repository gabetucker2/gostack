package gostack

type Card struct {
	Idx int
	Key interface{}
	Val interface{}
}

type Stack struct {
	Cards []*Card
	Size  int
}

type RETURN int
type FINDBY int
type TYPE int
type ORDER int
type MATCHBY int

const (
	RETURN_Idxs RETURN = iota
	RETURN_Keys
	RETURN_Vals
	RETURN_Cards
)

const (
	FINDBY_First FINDBY = iota
	FINDBY_Last
	FINDBY_Idx
	FINDBY_Idxs
	FINDBY_Key
	FINDBY_Keys
	FINDBY_Val
	FINDBY_Vals
	FINDBY_Card
	FINDBY_Cards
	FINDBY_All
	FINDBY_Lambda
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

func setRETURNDefaultIfNil(returnType interface{}) {
	if returnType == nil {
		returnType = RETURN_Cards
	}
}

func setFINDBYDefaultIfNil(findByType interface{}) {
	if findByType == nil {
		findByType = FINDBY_First
	}
}

func setORDERDefaultIfNil(orderType interface{}) {
	if orderType == nil {
		orderType = ORDER_After
	}
}

func setMATCHBYDefaultIfNil(matchByType interface{}) {
	if matchByType == nil {
		matchByType = MATCHBY_Object
	}
}
