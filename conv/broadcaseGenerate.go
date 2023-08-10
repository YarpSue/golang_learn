package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <- done:
				break Label
			}
		}
		// 这块如果不关闭通道，就会报错
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(done)
	// fatal error : all goroutines are asleep deadlock!  when ch is not with buffer
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// when ch is with buffer 请求过多超过Buffer还是会报错
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}