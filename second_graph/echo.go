package main

import (
	"flag"
	"fmt"
	"strings"
)
// 创建了两个指针
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main()  {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		s := "Usage of ./echo4: \n\t-n omit trailing newline \n\t-s string separator (default \" \")\n"
		fmt.Println(s)
	}
}
