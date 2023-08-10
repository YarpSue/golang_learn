package main

import (
	"io"
	"log"
	"net"
	"os"
)

// golang 模仿nc net.Dial  客户端程序

func main() {
	conn, err := net.Dial("tcp", "192.168.5.107:8000")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}