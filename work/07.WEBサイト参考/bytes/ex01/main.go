package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteString(" ")
	buf.WriteString("bytes.Buffer")

	fmt.Println(buf.String())
}
