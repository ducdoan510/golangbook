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
		go handleConn5(conn)
	}
}

func echo3(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	fmt.Println(shout)
	wg.Done()
}

func scan(c net.Conn, ch chan string) {
	sc := bufio.NewScanner(c)
	for sc.Scan() {
		ch <- sc.Text()
	}
}

func handleConn5(c net.Conn) {
	fmt.Printf("Start handling connection: %v\n", c)
	wg := sync.WaitGroup{} // count the number of active 'echo' goroutine
	ticker := time.NewTicker(1 * time.Second)
	defer func() {
		ticker.Stop()
		wg.Wait()
		fmt.Printf("Closing connection %v\n", c)
		c.Close()
	}()
	shout := make(chan string)
	go scan(c, shout)

	for counter := 0; counter < 10; counter++{
		select {
		case s := <-shout:
			counter = 0
			wg.Add(1)
			go echo3(c, s, 1*time.Second, &wg)
		case <-ticker.C:
			fmt.Println(counter)
		}
	}
}