package main

import (
	"fmt"
	"log"
	"os"
	"testmode/links"
)

// 无限的并行 不受控制
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var tokens = make(chan struct{}, 20)

func crawler2(url string) []string{
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<- tokens
	if err != nil{
		log.Print(err)
	}
	return list
}

func mainBack() {
	worklist := make(chan []string)
	go func() {worklist <- os.Args[1:]}()
	//go func() {fmt.Println(len(worklist))}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawler2(link)
				}(link)
			}
		}
	}
}

func main() {
	worklist := make(chan []string, 20)
	unseenLinks := make(chan string)

	go func() {worklist <- os.Args[1:]}()

	// 分配20个goroutine 来干活  阻塞等待unseenLinks chan中的内容
	for i :=0; i<20; i++ {
		go func(i int) {
			// 当没有收到任务时候 进行阻塞等待
			fmt.Printf("%d wait for unseenLinks\n", i)
			for link := range unseenLinks {
				fmt.Printf("%d get for unseenLinks\n", i)
				foundLinks := crawl(link)
				// 避免死锁  为什么会死锁呢？
				//worklist <- foundLinks
				// 队列满了就需要阻塞 阻塞就意味着消费者少了  消费的少了意味着队列的阻塞更加严重 最后瘫痪 造成死锁
				go func() {
					worklist <- foundLinks
				}()
			}
		}(i)
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}