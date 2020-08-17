package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "port used for clock server")
var timezones = map[string]string {
	"8010": "US/Eastern",
	"8020": "Asia/Tokyo",
	"8030": "Europe/London",
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn2(conn)
	}
}

func timeIn(t time.Time, tz string) (time.Time, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return t, err
	}
	return t.In(loc), nil
}

func handleConn2(c net.Conn) {
	defer c.Close()
	for {
		location := timezones[*port]
		now, err := timeIn(time.Now(), location)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.WriteString(c, location + ": " + now.Format("15:04:05\n"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)
	}
}
