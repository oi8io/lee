package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	//defer c.Close()
	OnMessage(c)
	return
}
func OnMessage(conn *net.TCPConn) {
	for i := 0; i < 1000; i++ {
		conn.Write(make([]byte, 2e10))
	}

	//conn.Write(make([]byte, 10240))
	//conn.Write(make([]byte, 10240))
	//conn.Write(make([]byte, 10240))
	//conn.Close()
}

func write(c net.Conn, b []byte) {
	n, err := c.Write(b)
	if n != len(b) || err != nil {
		log.Fatalf("err=%q\n", err)
	}
}
func read(c net.Conn, b []byte) {
	for nr := 0; nr != len(b); {
		n, err := c.Read(b)
		nr += n
		if err != nil {
			log.Fatal(err)
		}
	}
}
