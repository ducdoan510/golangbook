package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster area
type client struct {
	ch chan<- string
	name string
}
var (
	entering = make(chan *client)
	leaving = make(chan *client)
	messages = make(chan string)
)

func buildClientList(m map[*client]bool) string {
	builder := bytes.Buffer{}
	builder.WriteString("Current clients: ")
	for cli := range m {
		builder.WriteString(cli.name)
		builder.WriteByte(',')
	}
	return builder.String()
}

func broadcaster() {
	clients := make(map[*client]bool)
	for {
		select {
		case cli := <-entering:
			clients[cli] = true
			cli.ch <- buildClientList(clients)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		case m := <-messages:
			for cli := range clients {
				cli.ch <- m
			}
		}
	}
}

// handle individual connection

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)
	input.Scan()
	who := input.Text()

	ch <- "You are " + who
	cli := &client{ch, who}
	entering <- cli
	messages <- who + " has arrived"

	// Scan input and wait for 5 mins to close connection if idle
	timeout := 5 * time.Minute
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	go func() {
		<-timer.C
		conn.Close()
	}()
	for input.Scan() {
		timer.Reset(timeout)
		messages <- who + ": " + input.Text()
	}
	leaving <- cli
	messages <- who + " has left"
}

func clientWriter(c net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(c, msg)
	}
}
