package tutorials

import (
	"fmt"

	. "github.com/gabetucker2/gostack"
)

func tutorials_bootstrap_pop() {

	// > `stack.Extract(RETURN_Card, FINDBY_First)`
	// >> *removes and returns the first card in the stack*

}

func tutorials_bootstrap_push() {

	// > `stack.Add(insert)`
	// >> *adds a card to the beginning of the stack (see default enum types)*

}

func tutorials_bootstrap_indexOf(stack *Stack, val interface{}) int {

	//return stack.Get(FINDBY_Val, val).Idx
	// >> *returns the index of the first found matching card*

}

/** Executes the Bootstrap.go tutorial */
func Bootstrap() {

	fmt.Println("tutorials.Bootstrap()")

	tutorials_bootstrap_pop()
	tutorials_bootstrap_push()
	tutorials_bootstrap_indexOf()

}
