package main

import (
	"github.com/gabetucker2/gogenerics"
	"github.com/gabetucker2/gostack/testing"
	"github.com/gabetucker2/gostack/benchmark"
)

func main() {

	// remove import errors when stuff is commented out
	gogenerics.RemoveUnusedError(testing.Run, benchmark.Run)

	testing.Run(true)
	benchmark.Run()

}
