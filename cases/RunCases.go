package cases

import (
	"fmt"

	"github.com/gabetucker2/gogenerics"
)

/** Executes all case tests */
func Run(_showTestText bool) {

	showTestText = _showTestText
	gogenerics.RemoveUnusedError(case_MakeCard, case_MakeStack, case_MakeStackMatrix, case_stack_DimensionalityReduce, case_stack_Filter, case_stack_ToArray, case_stack_ToMap, case_stack_ToMatrix, case_stack_IsRegular, case_stack_Shape, case_stack_Duplicate, case_stack_Empty, case_card_Clone, case_stack_Clone, case_stack_Unique, case_card_Equals, case_stack_Equals, case_stack_Shuffle, case_stack_Flip, case_card_Print, case_stack_Print, case_stack_Lambdas, case_stack_Get, case_stack_GetMany, case_stack_Add, case_stack_AddMany, case_stack_Has, case_stack_Move, case_stack_Replace, case_stack_ReplaceMany, case_stack_Extract, case_stack_ExtractMany, case_stack_Remove, case_stack_RemoveMany, case_stack_Update, case_stack_UpdateMany, case_stack_Swap, case_CSVToStackMatrix, case_stack_ToCSV, case_stack_Coordinates, case_stack_CoordinatesMany)

	fmt.Println("- BEGINNING TESTS")

	// NON-GENERALIZED FUNCTIONS (NOT DEPENDENT ON GENERALIZED FUNCTIONS)
	case_MakeCard("MakeCard") // GOOD
	case_card_Equals("card.Equals") // GOOD
	case_MakeStack("MakeStack") // GOOD
	case_stack_Equals("stack.Equals") // GOOD
	case_MakeStackMatrix("MakeStackMatrix") // GOOD
	case_stack_Lambdas("stack.Lambdas") // GOOD
	case_stack_ToArray("stack.ToArray") // GOOD
	case_stack_ToMap("stack.ToMap") // GOOD
	case_stack_ToMatrix("stack.ToMatrix") // GOOD
	case_stack_IsRegular("stack.IsRegular") // GOOD
	case_stack_Shape("stack.Shape") // GOOD
	case_stack_Duplicate("stack.Duplicate") // GOOD
	case_stack_Empty("stack.Empty") // GOOD
	case_card_Clone("card.Clone") // GOOD
	case_stack_Clone("stack.Clone") // GOOD
	case_stack_Shuffle("stack.Shuffle") // GOOD
	case_card_Print("card.Print") // GOOD
	case_stack_Print("stack.Print") // GOOD
	case_CSVToStackMatrix("CSVToStackMatrix") // GOOD
	case_stack_ToCSV("stack.ToCSV") // GOOD
	case_stack_Coordinates("stack.Coordinates") // GOOD
	case_stack_CoordinatesMany("stack.CoordinatesMany") // GOOD
	case_stack_Flip("stack.Flip") // GOOD
	case_card_SwitchKeyVal("card.SwitchKeyVal") // GOOD
	case_stack_SwitchKeysVals("stack.SwitchKeysVals") // GOOD
	
	// GENERALIZED FUNCTIONS
	case_stack_Get("stack.Get") // GOOD
	case_stack_GetMany("stack.GetMany") // GOOD
	case_stack_Add("stack.Add") // GOOD
	case_stack_AddMany("stack.AddMany") // GOOD
	case_stack_Replace("stack.Replace") // GOOD
	case_stack_ReplaceMany("stack.ReplaceMany") // GOOD
	case_stack_Update("stack.Update") // GOOD
	case_stack_UpdateMany("stack.UpdateMany") // GOOD
	case_stack_Extract("stack.Extract") // GOOD
	case_stack_ExtractMany("stack.ExtractMany") // GOOD
	case_stack_Remove("stack.Remove") // GOOD
	case_stack_RemoveMany("stack.RemoveMany") // GOOD
	case_stack_Has("stack.Has") // GOOD
	case_stack_Move("stack.Move") // GOOD
	case_stack_Swap("stack.Swap") // GOOD
	case_stack_Filter("stack.Filter") // GOOD
	
	// NON-GENERALIZED FUNCTIONS (DEPENDENT ON GENERALIZED FUNCTIONS)
	case_stack_DimensionalityReduce("stack.DimensionalityReduce") // GOOD
	case_stack_Unique("stack.Unique") // GOOD
	case_stack_Transpose("stack.Transpose") // GOOD

}
