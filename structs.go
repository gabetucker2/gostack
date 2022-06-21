package main

type Card struct {
	key   interface{}
	value interface{}
}

type Stack struct {
	cards []*Card
	size  int
}

var testCard1 = _gostack_back_NewCard("Card A") // in sample stack
var testCard2 = _gostack_back_NewCard("Card B") // in sample stack
var testCard3 = _gostack_back_NewCard("Card C") // in sample stack
var testCard4 = _gostack_back_NewCard("Card D") // out of sample stack
