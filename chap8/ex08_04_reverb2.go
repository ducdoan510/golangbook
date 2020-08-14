package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
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
		go handleConn4(conn)
	}
}

func echo2(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	fmt.Println(shout)
	wg.Done()
}

func handleConn4(c net.Conn) {
	sc := bufio.NewScanner(c)
	wg := sync.WaitGroup{} // count the number of active 'echo' goroutine

	for sc.Scan() {
		wg.Add(1)
		go echo2(c, sc.Text(), 1 * time.Second, &wg)
	}
	wg.Wait()
	c.Close()
}