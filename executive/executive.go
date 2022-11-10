package main

import (
	"github.com/gabetucker2/gogenerics"
	"github.com/gabetucker2/gostack/cases"
	"github.com/gabetucker2/gostack/testing"
)

func main() {

	// remove import errors when stuff is commented out
	gogenerics.RemoveUnusedError(cases.Run, testing.RaceNative, testing.RaceGostack)

	cases.Run(true)
	// testing.RaceNative()
	// testing.RaceGostack()

}
