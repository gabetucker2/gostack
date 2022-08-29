package main

import (
	"github.com/gabetucker2/gostack/casetests"
	"github.com/gabetucker2/gostack/tutorials"
)

func main() {

	// remove import errors when stuff below is commented out
	casetests.RemoveNotUsedError()
	tutorials.RemoveNotUsedError()

	// tests
	casetests.Run(true)

	// tutorials
	//tutorials.Bootstrap()
	//tutorials.Lambda()

}
