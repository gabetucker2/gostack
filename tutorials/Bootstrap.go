package tutorials

import (
	"fmt"

	//. "github.com/gabetucker2/gostack"
)

func tutorials_bootstrap_pop() {

	// > `stack.Extract(RETURN_Card, POSITION_First)`
	// >> *removes and returns the first card in the stack*

}

func tutorials_bootstrap_push() {

	// > `stack.Add(insert)`
	// >> *adds a card to the beginning of the stack (see default enum types)*

}

func tutorials_bootstrap_indexOf() {

	// > `stack.Get(RETURN_Idx, POSITION_Card, cardToMatch, MATCH_Object)`
	// >> *returns the index of the first found matching card*

}

func Bootstrap() {

	fmt.Println("tutorials.Bootstrap()")

	tutorials_bootstrap_pop()
	tutorials_bootstrap_push()
	tutorials_bootstrap_indexOf()

}
