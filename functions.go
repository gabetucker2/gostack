package main

func IndexOf(arr []string, target string) (ret int) {

	ret = -1
	for i := 0; i < len(arr); i++ {
		if target == arr[i] {
			ret = i
			break
		}
	}

	return
}
