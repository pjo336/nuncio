package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/pjo336/nuncio/cmd"
)

const port = "8888"

var (
	// VERSION is set during build
	VERSION = "0.0.1"
)

func main() {
	fmt.Println(`
 _   _                  _
| \ | |_   _ _ __   ___(_) ___
|  \| | | | | '_ \ / __| |/ _ \
| |\  | |_| | | | | (__| | (_) |
|_| \_|\__,_|_| |_|\___|_|\___/

		`)

	cmd.Execute(VERSION)
	// listen()
}

func listen() {
	l, err := net.Listen("tcp", "*localhost:5656")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go proxy(conn)
	}
}

func proxy(incoming net.Conn) {
	defer incoming.Close()

	destination, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Println(err)
		return
	}

	defer destination.Close()

	go io.Copy(destination, incoming)
	io.Copy(incoming, destination)
}
