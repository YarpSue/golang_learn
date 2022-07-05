package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/count/", counter)
	http.HandleFunc("/", handler2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Print("/")
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Print("/countx")
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()

}