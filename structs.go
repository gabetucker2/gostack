package main

type Card struct {
	indices []int
	keys    []interface{}
	values  []interface{}
}

type Stack struct {
	cards []Card
	size  int
}
