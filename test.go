package main

import "fmt"

func (haystack interface{}) Test_IndexOfString(needle interface{}, expectedOutput interface{}) (ret bool) {

}

func main() {

	fmt.Println("BEGINNING TESTS")

	// IndexOf test 1
	haystack := []int{4, 1, 6}
	needle := 6
	expected := 2
	haystack.Test_IndexOf(needle, expected)

	// IndexOf test 2
	haystack = []int{4, 1, 6}
	needle = 5
	expected = -1
	haystack.Test_IndexOf(needle, expected)

}
