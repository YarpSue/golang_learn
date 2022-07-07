package main

import "fmt"

func main() {
	ma := map[string]int{"a": 1, "b":2}
	fmt.Println(ma["a"])
	fmt.Println(ma["b"])
	fmt.Println(ma)
}