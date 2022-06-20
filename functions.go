package main

func IndexOf(haystack interface{}, needle interface{}) (ret int) {

	ret = -1
	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			ret = i
			break
		}
	}

	return

}
