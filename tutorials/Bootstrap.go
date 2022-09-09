package tutorials

import (
	"github.com/gabetucker2/gogenerics"
	. "github.com/gabetucker2/gostack" //lint:ignore ST1001 Ignore warning
)

/** Executes the Bootstrap.go tutorial */
func Bootstrap() {

	gogenerics.RemoveUnusedError(MakeCard()) // temporary

}
