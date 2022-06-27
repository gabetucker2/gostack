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
type POSITION int
type TYPE int
type ORDER int
type MATCH int

const (
	RETURN_Idx RETURN = iota
	RETURN_Key
	RETURN_Val
	RETURN_Card
)

const (
	POSITION_First POSITION = iota
	POSITION_Last
	POSITION_Idx
	POSITION_Idxs
	POSITION_Key
	POSITION_Keys
	POSITION_Val
	POSITION_Vals
	POSITION_Card
	POSITION_Cards
	POSITION_All
	POSITION_Lambda
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
	MATCH_Object MATCH = iota
	MATCH_Reference
)

func setRETURNDefaultIfNil(returnType interface{}) {
	if returnType == nil {
		returnType = RETURN_Cards
	}
}

func setPOSITIONDefaultIfNil(positionType interface{}) {
	if positionType == nil {
		positionType = POSITION_First
	}
}

func setORDERDefaultIfNil(orderType interface{}) {
	if orderType == nil {
		orderType = ORDER_After
	}
}

func setMATCHDefaultIfNil(matchType interface{}) {
	if matchType == nil {
		matchType = MATCH_Object
	}
}
