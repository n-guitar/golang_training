package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	ch := make(chan string, 3)
	for _, arg := range os.Args[1:] {
		split := strings.Split(arg, "=")
		timezone := split[0]
		url_port := split[1]

		// デバッグ用
		//fmt.Println(timezone, url_port)

		conn, err := net.Dial("tcp", url_port)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go clockWall(conn, timezone, ch)
	}
	for times := range ch {
		fmt.Print(times)
	}
}

func clockWall(conn net.Conn, timezone string, ch chan<- string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		//\rで改行を
		ch <- fmt.Sprintf("\r%s:%s   ", timezone, scanner.Text())
	}
}
