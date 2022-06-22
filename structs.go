package main

type Card struct {
	key interface{}
	val interface{}
}

type Stack struct {
	cards []*Card
	size  int
}

type Slice struct {
	startIdx int
	endIdx   int
}

type RETURN int
type POSITION int
type TYPE int
type ORDER int
type MATCH int

const (
	RETURN_None RETURN = iota
	RETURN_Idx
	RETURN_Idxs
	RETURN_Key
	RETURN_Keys
	RETURN_Val
	RETURN_Vals
	RETURN_Card
	RETURN_Cards
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

var testCardA = _gostack_back_NewCard("Card A") // in sample stack
var testCardB = _gostack_back_NewCard("Card B") // in sample stack
var testCardC = _gostack_back_NewCard("Card C") // in sample stack
var testCardD = _gostack_back_NewCard("Card D") // out of sample stack
