package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)
// main 多次定义在这个包中
func main()  {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d", resp.StatusCode)
		// 使用io.Copy的方式
		_, err = io.Copy(os.Stdout, resp.Body)
		//b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}
