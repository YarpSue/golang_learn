package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

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

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}