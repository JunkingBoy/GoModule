package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"io"
)

func handleConnection(conn net.Conn) {
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", data)
		fmt.Fprintf(conn, "OK\n")
	}
	conn.Close()
}

func main() {
	In, err := net.Listen("tcp", ":443")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := In.Accept()
		if err != nil {
			log.Fatal("get client connection error:", err)
		}
		go handleConnection(conn)
	}
}
