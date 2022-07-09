package main

func AddNumStr[T NumStr](params []T) (sum T) {
	for _, param := range params {
		sum += param
	}
	return
}
