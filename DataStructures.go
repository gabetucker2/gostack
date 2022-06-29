package gostack

import "reflect"

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
type FIND int
type REPLACE int
type TYPE int
type ORDER int
type MATCHBY int
type CLONE int

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
	FIND_Key
	FIND_Keys
	FIND_Val
	FIND_Vals
	FIND_Card
	FIND_Cards
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

func setRETURNDefaultIfNil(returnType interface{}) {
	if returnType == nil {
		returnType = RETURN_Cards
	}
}

func setFINDDefaultIfNil(findType interface{}) {
	if findType == nil {
		findType = FIND_First
	}
}

func setORDERDefaultIfNil(orderType interface{}) {
	if orderType == nil {
		orderType = ORDER_Before
	}
}

func setMATCHBYDefaultIfNil(matchByType interface{}) {
	if matchByType == nil {
		matchByType = MATCHBY_Object
	}
}

func setCLONEDefaultIfNil(cloneType interface{}) {
	if cloneType == nil {
		cloneType = CLONE_False
	}
}

func cloneInterface(toClone interface{}) interface{} {
	return reflect.New(reflect.ValueOf(toClone).Elem().Type()).Interface()
}
