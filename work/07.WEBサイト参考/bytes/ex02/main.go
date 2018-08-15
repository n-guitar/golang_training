package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	r := bytes.NewReader([]byte("Hello\nbyte.Reader\n"))
	s := bufio.NewScanner(r)
	line := 0
	for s.Scan() {
		line++
		fmt.Println(line, " ", s.Text())
	}
}
