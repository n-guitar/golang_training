package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	if s[:1] == "+" {
		return s[:1] + comma(s[1:])
	}

	if s[:1] == "-" {
		return s[:1] + comma(s[1:])
	}

	/*
		func IndexByte(s []byte, c byte) int
		IndexByteは、s内でcが最初に出現する箇所のインデックスを返します。一致しないときは-1を返します。
	*/
	dotindex := strings.IndexByte(s, '.')
	if dotindex != -1 {
		return comma(s[:dotindex]) + s[dotindex:]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
