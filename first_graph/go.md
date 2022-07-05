### golang 

```go
package main

import (
	"fmt"
)

func main()  {
	var a string = "initial"
	fmt.Println(a)
}
```
go build 生成exe执行程序
go run example.go

Go语言的基本类型有：
* bool
* string
* int、int8、int16、int32、int64
* uint、uint8、uint16、uint32、uint64、uintptr
* byte // uint8 的别名
* rune // int32 的别名 代表一个 Unicode 码
* float32、float64
* complex64、complex128
* int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等 初始值
* 骆驼命名法
```go
package main

import "fmt"

func main()  {
	var a string = "initial"
	var (
		b string
		c []float32
		d func() bool
		e struct {
			x int
		}
	)
	fmt.Println(a, b, c, d, e)
}
func main() {
   // 定义变量，同时显式初始化。
   // 只能使用在函数内部
   x:=100
   a,s:=1, "abc"
   fmt.Println(x, a, s)
}
```
```text
排序的多重赋值特性
type IntSlice []int
func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
```
匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。 和python区别是go里面匿名变量不会分配内存。

查看一些标准库和社区写的package
* https://golang.org/pkg
* https://godoc.org
* go doc 这个工具可以让你直接在本地命令行阅读标准库的文档

