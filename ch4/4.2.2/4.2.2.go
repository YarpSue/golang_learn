package main
/*
字符串的切片中除去空的元素
*/

import "fmt"
// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data) // `["one" "three" "three"]`
	data2 := nonempty(data)
	fmt.Printf("%q\n", data2) // `["one" "three" "three"]`
}