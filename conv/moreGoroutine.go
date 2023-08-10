package main

type task struct {
	begin int
	end int
	result chan int
}

// 计算单元  计算规定范围内求和
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}
