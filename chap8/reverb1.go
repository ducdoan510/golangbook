package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn3(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn3(c net.Conn) {
	sc := bufio.NewScanner(c)
	for sc.Scan() {
		// this is not a go routine so that the shouts are printed in the sequence of input
		// to make this a goroutine, extra handle needs to be done to make sure all the shouts are finished using wait group
		// refer to reverb2
		echo(c, sc.Text(), 1 * time.Second)
	}
	c.Close()
}
