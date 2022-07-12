package main

// 管道的 函数的输入参数和输出参数都是相同的 chan 类型，函数自己可以调用自己  形成调用链

func chain(in chan<- int) chan int {
	in
}