package main

// 返回的是一个能返回int类型的匿名函数
func squares() func() int {
	var x int
	// 匿名函数的用法
	return func() int {
		x++
		return x*x
	}
}

//func main() {
//	f := squares()
//	fmt.Println(f())
//	fmt.Println(f())
//	fmt.Println(f())
//}

// closures 闭包技术