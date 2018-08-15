// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", CommaBuffer(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func CommaBuffer(s string) string {
	var buf bytes.Buffer
	rune := []byte(s) //https://qiita.com/seihmd/items/4a878e7fa340d7963feeを参考
	remainder := len(s) % 3

	for _, r := range rune { // P78のruneのrangeloop
		if remainder == 0 {
			buf.WriteByte(',')
			remainder = 3
		}
		buf.WriteByte(r) // 余り0以外の場合そのまま出力
		remainder = remainder - 1

	}
	return buf.String()
}

//!-
