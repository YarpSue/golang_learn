package main

import "net"

func main() {
	net.Listen("tcp", "localhost:7452")

}