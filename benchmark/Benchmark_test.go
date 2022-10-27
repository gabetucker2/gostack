package benchmark

import (
	"testing"

	"github.com/gabetucker2/gogenerics"
	. "github.com/gabetucker2/gostack"
)

// ----------------------------------------

func Benchmark_Empty(b *testing.B) {
	for i := 0; i < b.N; i++ { }
}

// ----------------------------------------

func test_Native_CreateArray() {

	myArr := []int {1, 2, 3}
	
	gogenerics.RemoveUnusedError(myArr)

}

func test_Gostack_CreateArray() {

	myStack := MakeStack([]int {1, 2, 3})
		
	gogenerics.RemoveUnusedError(myStack)

}

func Benchmark_Native_CreateArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_Native_CreateArray()
    }
}
func Benchmark_Gostack_CreateArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_Gostack_CreateArray()
    }
}

// ----------------------------------------

func test_Native_CreateMap() {

	myMap := map[int]int {1:1, 2:2, 3:3}
	
	gogenerics.RemoveUnusedError(myMap)

}

func test_Gostack_CreateMap() {

	myStack := MakeStack(map[int]int {1:1, 2:2, 3:3})
		
	gogenerics.RemoveUnusedError(myStack)

}

func Benchmark_Native_CreateMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_Native_CreateMap()
    }
}
func Benchmark_Gostack_CreateMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_Gostack_CreateMap()
    }
}

// ----------------------------------------

func test_Native_GetElement() {

	myArr := []int {1, 2, 3}
	myOutput := -1
	
	for _, elem := range myArr {
		if elem == 2 {
			myOutput = 2
		}
	}
	
	gogenerics.RemoveUnusedError(myArr, myOutput)

}

func test_Gostack_GetElement() {

	myStack := MakeStack([]int {1, 2, 3})
	myOutput := myStack.Get(FIND_Val, 2)
		
	gogenerics.RemoveUnusedError(myStack, myOutput)

}

func Benchmark_Native_GetElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_Native_GetElement()
    }
}
func Benchmark_Gostack_GetElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_Gostack_GetElement()
    }
}

// ----------------------------------------

func test_Native_AddElement() {

	myArr := []int {1, 2, 3}
	myArr = append(myArr, 4)
	
	gogenerics.RemoveUnusedError(myArr)

}

func test_Gostack_AddElement() {

	myStack := MakeStack([]int {1, 2, 3}).Add(4)
		
	gogenerics.RemoveUnusedError(myStack)

}

func Benchmark_Native_AddElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_Native_AddElement()
    }
}
func Benchmark_Gostack_AddElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test_Gostack_AddElement()
    }
}
