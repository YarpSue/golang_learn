package main

func square(n int) int {return n*n}
func negative(n int) int {return -n}
//func product(m,n int) int {return m*n}
func add1(r rune) rune {return r+3}

//func main() {
//	f := square
//	fmt.Println(f(3))
//
//	f = negative
//	fmt.Println(f(3))
//
//	fmt.Printf("%T\n", f)
//	// compile error: can't assign func(int, int) int to func(int) int
//	//f = product
//	//fmt.Println(f(2))
//	var g func(int) int
//	if g != nil {
//		g(3)
//	} else {
//		fmt.Println("g function is null")
//	}
//	fmt.Println(strings.Map(add1, "HAL-9000"))
//}
//
