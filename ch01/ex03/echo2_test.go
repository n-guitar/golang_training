// Echo2は、そのコマンドラインの引数を表示します。
package main

import (
	"fmt"
	"os"
	"testing"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

func BenchmarkEchoTime(b *testing.B) {

	b.ResetTimer()
}
