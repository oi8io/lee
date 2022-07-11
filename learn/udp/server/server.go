package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const N = 1e9 + 7

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		c, err := ln.AcceptTCP()
		if err != nil {
			return
		}
		go handle(c)
		//var b [8]byte
		//read(c, b[:])
		//write(c, b[:])
	}

}
func handle(conn *net.TCPConn) {
	for {
		_, err := conn.Read(make([]byte, 10240))
		if err != nil {
			break
		}
		time.Sleep(10 * time.Second)
	}
	//conn.Close()
}
func write(c net.Conn, b []byte) {
	n, err := c.Write(b)
	if n != len(b) || err != nil {
		log.Println("err=%q\n", err)
	}
}
func read(c net.Conn, r []byte) {
	i, err := c.Read(r)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(i))

}
