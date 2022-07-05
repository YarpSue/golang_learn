package main

import "fmt"

func main() {
	x := gcd(4, 6)
	fmt.Println(x)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}