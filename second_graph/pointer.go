package main

import "fmt"

func main()  {
	x := 1
	// 获取指针
	p := &x
	// 获取值
	fmt.Println(*p)
	// 指针位置的内存块赋值
	*p = 100
	fmt.Println(x)
}