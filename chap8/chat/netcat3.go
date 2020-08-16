package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- true
	}()
	mustCopy3(conn, os.Stdin)
	switch conn := conn.(type) {
	case *net.TCPConn:
		log.Println("TCPConn")
		err := conn.CloseWrite()
		if err != nil {
			log.Println(err)
		}
	default:
		conn.Close()
	}
	fmt.Println("release done")
	<-done
}

func mustCopy3(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
