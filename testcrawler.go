package main

import (
	"fmt"
	"log"
	"testmode/links"
)

func main() {
	list, err :=links.Extract("http://gopl.io/")
	if err != nil {
		log.Print(err)
	}
	fmt.Println(list)
}