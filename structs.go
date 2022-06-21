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

type POSITION int

const (
	POSITION_First POSITION = iota
	POSITION_Last
	POSITION_Card
	POSITION_Idx
	POSITION_Val
	POSITION_Key
	POSITION_Slice
	POSITION_All
)

var testCardA = _gostack_back_NewCard("Card A") // in sample stack
var testCardB = _gostack_back_NewCard("Card B") // in sample stack
var testCardC = _gostack_back_NewCard("Card C") // in sample stack
var testCardD = _gostack_back_NewCard("Card D") // out of sample stack
