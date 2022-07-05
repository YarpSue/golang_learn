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
