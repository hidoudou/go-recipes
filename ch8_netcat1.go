package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"fmt"
	"time"
)

func main() {
	//参数解析
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	for i := 1; i < len(args); i++ {
		go run(args[i])
	}

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}

func run(param string) {
	//解析出主机
	var host string
	var pos int
	pos = strings.Index(param, "=")
	fmt.Println(pos)
	rs := []rune(param)
	host = string(rs[pos+1:len(rs)])
	fmt.Println(host)
	conn, err := net.Dial("tcp", host)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

var Usage = func() {
	fmt.Println("wow,incorrect input")
}