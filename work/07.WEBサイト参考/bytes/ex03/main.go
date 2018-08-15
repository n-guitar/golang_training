package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(CommaBuffer(os.Args[i]))
	}
}

func CommaBuffer(s string) string {
	var buf bytes.Buffer
	//n := len(s)
	remainder := len(s) % 3
	comma_num := len(s)
	//runes := []byte(s)
	//n := len(s)
	if comma_num <= 3 {
		buf.WriteString(s[:])
	}
	for comma_num > 3 {
		switch remainder {
		case 2:
			buf.WriteString(s[:2])
			buf.WriteString(",")
			comma_num = len(s[2:])
			remainder = len(s[2:]) % 3
		case 1:
			buf.WriteString(s[:1])
			buf.WriteString(",")
			comma_num = len(s[1:])
			remainder = len(s[1:]) % 3
		case 0:
			buf.WriteString(s[:3])
			buf.WriteString(",")
			comma_num = len(s[3:])
			remainder = len(s[3:]) % 3
		}
	}
	return buf.String()

}
