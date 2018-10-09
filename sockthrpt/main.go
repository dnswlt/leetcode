package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	bufSize = 1024
)

func handleConnection(conn net.Conn) {
	started := time.Now()
	buf := bufio.NewReader(conn)
	var count int
	for {
		s, err := buf.ReadString('\n')
		if s == "" || err != nil {
			fmt.Print("No more bytes to read: ", err)
			break
		}
		count++
		fmt.Fprintf(conn, "ACK\n")
	}
	elapsed := time.Since(started)
	fmt.Printf("Received %d messages in %dms (%.1f msg/s)\n", count, elapsed.Nanoseconds()/1000000, float64(count)/elapsed.Seconds())
}

func main() {
	if os.Args[1] == "send" {
		n, _ := strconv.Atoi(os.Args[2])
		conn, err := net.Dial("tcp", "localhost:5000")
		if err != nil {
			log.Fatal("Could not dial: ", err)
		}
		defer conn.Close()
		buf := bufio.NewReader(conn)
		for i := 0; i < n; i++ {
			fmt.Fprintf(conn, "PING\n")
			buf.ReadString('\n')
		}
	} else {
		ln, err := net.Listen("tcp", ":5000")
		if err != nil {
			log.Fatal("Could not listen: ", err)
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal("Could not accept: ", err)
			}
			go handleConnection(conn)
		}
	}
}
