package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// those values can be retrieved from Args as in the exercise
	// for simplicity, hardcode the port numbers here
	ports := []string{"8010", "8020", "8030"}
	for _, port := range ports {
		go dialPort(port)
	}
	time.Sleep(time.Minute)
}

func dialPort(port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	readFrom(os.Stdout, conn)
}

func readFrom(dst io.Writer, src io.Reader) {
	sc := bufio.NewScanner(src)
	for sc.Scan() {
		fmt.Fprintf(dst, sc.Text() + "\n")
	}
}