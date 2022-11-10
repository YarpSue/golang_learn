package main
// 带有方向的 channel
import "fmt"

func counter(out chan<- int)  {
	for i:=0;i<=100;i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v*v
	}
	close(out)
	// in 方向是不能关闭的
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}