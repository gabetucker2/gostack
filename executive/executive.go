package main

import (
	"github.com/gabetucker2/gogenerics"
	"github.com/gabetucker2/gostack/casetests"
	"github.com/gabetucker2/gostack/tutorials"
)

func main() {

	// remove import errors when stuff is commented out
	gogenerics.RemoveUnusedError(casetests.Run, tutorials.Introduction)

	// tests
	casetests.Run(false)

	// tutorials
	//tutorials.Bootstrap()
	//tutorials.Lambda()

}
