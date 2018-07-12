package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// 文字列スライスの結合
// スライスの各要素の間を文字列でつなぎ、1つの文字列に結合するにはJoin関数を使用

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func BenchmarkEchoTime(b *testing.B) {

	b.ResetTimer()
}
