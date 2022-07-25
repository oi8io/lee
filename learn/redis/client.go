package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

var cli *net.TCPConn

func main() {
	var ip = "127.0.0.1:6379"
	addr, err := net.ResolveTCPAddr("tcp", ip)
	if err != nil {
		log.Fatalln(err)
	}
	tcp, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}
	cli = tcp
	err = SendCMD("ping")
	if err != nil {
		panic(err)
	}
	SendCMD("set", "name", "Oj8K黄黄的")
	SendCMD("get", "name")
	SendCMD("mget", "name", "wangbandan")

}

const (
	respTypeString    = '+'
	respTypeError     = '-'
	respTypeNumber    = ':'
	respTypeMultiline = '$'
	respTypeAll       = '*'
)

func readResp() error {
	var resp = make([]byte, 0)
	for {
		b := make([]byte, 1024)
		r, err := cli.Read(b)
		if err != nil {
			return err
		}
		resp = append(resp, b...)
		if r < 1024 {
			break
		}
	}
	fields := strings.Fields(string(resp))
	fmt.Println(fields)
	return nil
	switch resp[0] {
	case respTypeString:
		fmt.Println("READ STRING", string(resp))
	case respTypeError:
		fmt.Println("read err ", string(resp))
	case respTypeMultiline:
		fmt.Println("read lines ", string(resp))
	case respTypeAll:
		fmt.Println("read all", string(resp))
	case respTypeNumber:
		fmt.Println("read number", string(resp))
	default:
		return errors.New("解析错误")
	}
	return nil
}

func SendCMD(strs ...string) error {
	x := fmt.Sprintf("*%d\r\n", len(strs))
	for i := 0; i < len(strs); i++ {
		x += fmt.Sprintf("$%d\r\n%s\r\n", len([]byte(strs[i])), strs[i])
	}
	write, err := cli.Write([]byte(x))
	if err != nil {
		return err
	}
	fmt.Println("sent:", write, err)
	return readResp()
}
