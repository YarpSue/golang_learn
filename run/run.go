package main

import (
	"fmt"
	"testmode/links"
)

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println(links.Sum(1, 2, 3, 4, 5))
	// 直接传入多个参数的切片
	fmt.Println(links.Sum(values...))
}
