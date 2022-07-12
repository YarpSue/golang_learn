package main

import (
	"fmt"
	"math/rand"
)

// 一个融合了并发、缓存、退出通知等多重特性的生成器

//done 接收通知退出信号
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
		close(ch)
	}()
	return ch
}

func GenerateIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <- done:
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

// 通过select执行散入Fan in 操作
func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Label:
		for {
			select {
			case ch <- <- GenerateIntA(send):
			case ch <- <- GenerateIntB(send):
			case <- done:
				send <- struct{}{}
				send <- struct{}{}
				break Label
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	// 创建一个作为接收退出信号的chan
	done := make(chan struct{})
	// 启动生成器
	ch := GenerateInt(done)
	// 获取生成器资源
	for i := 0; i < 100; i++ {
		fmt.Println(<- ch)
	}
	// 通知生产者停止生产
	done <- struct{}{}
	fmt.Println("Stop Generate")
	fmt.Println(<-ch)
}