package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 服务端程序改写 echo 程序
func main() {
	// 加参数的方式
	port := flag.String("port", "8000", "add port of server")

	flag.Parse()
	fmt.Println(*port)
	listener, err := net.Listen("tcp", "192.168.5.107:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 加个go 便能够处理并发请求的客户端的请求
		go handleConn(conn)
	}
}

// 声音开始从大到小
// 再传回客户端
func echo(c net.Conn, shout string, delay time.Duration) {
	/*
	_, err1 := io.WriteString(c, "\t" + strings.ToUpper(shout))
	if err1 != nil {
		log.Fatal(err1)
	}
	time.Sleep(delay)
	_, err2 := io.WriteString(c, "\t" + shout)
	if err2 != nil {
		log.Fatal(err2)
	}
	time.Sleep(delay)
	_, err3 := io.WriteString(c, "\t" + strings.ToLower(shout) + "\n")
	if err3 != nil {
		log.Fatal(err3)
	}
	time.Sleep(delay)
	*/
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
}

