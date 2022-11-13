package testing

import (
	"fmt"
	. "github.com/gabetucker2/gostack"//lint:ignore ST1001, ignore warning
)

func Experiment() {

	fmt.Println("- BEGINNING EXPERIMENTAL TEST")

	// INIT
	thisStack, retStack, retCard, retVarAdr := MakeStack([]string {"Save retCard", "Make new substack", "Save retVarAdr"}, []any {1, 2, 3}).Lambda(func (card *Card, _ *Stack, _ bool, _ *Stack, retStack *Stack, retCard *Card, retVarAdr any) {

		switch card.Key {
		case "Make new substack":
			*retStack = *MakeSubstack([]*Card {card.Clone()}).Duplicate(5)
		case "Save retCard":
			*retCard = *card
		case "Save retVarAdr":
			*(retVarAdr.(*any)) = card.Val
		}
	
	})
	// Stack{"Save retCard" ... etc}, Stack{2, 2, 2, 2, 2}, card{1}, *any => 3

	thisStack.Print()
	retStack.Print()
	retCard.Print()
	fmt.Printf("retVarAdr: %v\n", retVarAdr)

}
