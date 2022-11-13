package testing

import (
	"fmt"
	. "github.com/gabetucker2/gostack"//lint:ignore ST1001, ignore warning
)

func Experiment() {

	fmt.Println("- BEGINNING EXPERIMENTAL TEST")

	cardA := MakeCard()
	cardB := MakeCard()
	cardC := MakeCard()
	cardD := MakeCard()
	mainStack := MakeStack([]*Card {cardA, cardB, cardC}).Print()
	ptrStack := MakeStack([]any {fmt.Sprintf("%p", cardA), fmt.Sprintf("%p", cardB), fmt.Sprintf("%p", cardD)}).Print()
	ptrStack.GetMany(FIND_Card, mainStack, nil, nil, nil, nil, DEREFERENCE_This).Print() // Stack{&cardA, &cardB}

}
