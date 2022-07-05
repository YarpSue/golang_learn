package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// 获取所有的参数
	files := os.Args[1:]
	if len(files) == 0 {
		// 没有加返回 counts 值还是发生了变化  引用的拷贝
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\v", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	// count lines > 1
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	// 以流的形式读取
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}