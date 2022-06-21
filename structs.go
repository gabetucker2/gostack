package main

type Card struct {
	key   interface{}
	value interface{}
}

type Stack struct {
	cards []*Card
	size  int
}

type Position int

const (
	Position_First Position = iota
	Position_Last
	Position_Idx
	Position_Val
	Position_Key
	Position_Slice
)

var testCardA = _gostack_back_NewCard("Card A") // in sample stack
var testCardB = _gostack_back_NewCard("Card B") // in sample stack
var testCardC = _gostack_back_NewCard("Card C") // in sample stack
var testCardD = _gostack_back_NewCard("Card D") // out of sample stack
