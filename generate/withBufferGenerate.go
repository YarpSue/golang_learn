package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA() chan int {
	ch := make(chan int, 10)
	// 启动一个goroutine 用于生成随机数，函数返回一个通道用于获取随机数
	go func(url string) {
		for {
			ch <- rand.Int()
			fmt.Println(url)
		}
	}("haha")
	return ch
}

func main() {
	ch := GenerateIntA()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}