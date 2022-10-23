package main

import (
	"github.com/gabetucker2/gogenerics"
	"github.com/gabetucker2/gostack/casetests"
)

func main() {

	// remove import errors when stuff is commented out
	gogenerics.RemoveUnusedError(casetests.Run)

	// tests
	casetests.Run(true)

	// tutorials
	//tutorials.Bootstrap()
	//tutorials.Lambda()

}
