package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Print(strings.Join(os.Args[1:], "\n"))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0])
}