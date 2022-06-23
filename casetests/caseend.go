package casetests

// variables
var showTestText bool

// case functions
// TEMPLATE:
/*
func gostack_case_MyFunction(funcName string) {

	gostack_test_Start(funcName, showTestText)

	// call MyFunction here

	// make miscellaneous variables to test whether MyFunction worked properly here if necessary

	// add failure conditions here for the case test, ensure LenAndSize is always the first
	conditions := []bool{
		gostack_test_LenAndSize(stack, ?),
	}

	gostack_test_End(funcName, conditions)

}
*/

func gostack_case_MakeCard(funcName string) {

	gostack_test_Start(funcName, showTestText)

	stack := MakeCard()

	conditions := []bool{
		gostack_test_LenAndSize(stack, 0),
	}

	gostack_test_End(funcName, conditions)

}

// main function
func Run(_showTestText bool) {

	showTestText = _showTestText

	println("- BEGINNING TESTS (fix failures/errors in descending order)")

	gostack_case_MakeCard("MakeCard")

}
