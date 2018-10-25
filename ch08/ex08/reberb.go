package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	wtimeout := make(chan string, 1)
	go watchClose(c, wtimeout)
	for input.Scan() {
		text := input.Text()
		wtimeout <- text
		go echo(c, text, 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	close(wtimeout)
	c.Close()
	fmt.Println("timeout!")
}

//!-

func watchClose(c net.Conn, wtimeout chan string) {
	for {
		select {
		case <-time.After(5 * time.Second):
			c.Close()
		case _, ok := <-wtimeout:
			if !ok {
				return
			}
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
