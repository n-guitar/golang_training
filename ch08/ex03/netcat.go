package main

import (
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

	//参考ページ： http://d.hatena.ne.jp/taknb2nch/20140210/1392044307
	tcpConn := conn.(*net.TCPConn)

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tcpConn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(tcpConn, os.Stdin)
	tcpConn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
