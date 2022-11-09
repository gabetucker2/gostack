package main

import (
	"github.com/gabetucker2/gogenerics"
	"github.com/gabetucker2/gostack/testing"
)

func main() {

	// remove import errors when stuff is commented out
	gogenerics.RemoveUnusedError(testing.Run)

	testing.Run(true)

}
