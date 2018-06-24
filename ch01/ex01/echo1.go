// Echo1は、そのコマンドラインの引数を表示します。
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ { // os.Args[0]は0番目の要素には実行したコマンド名が格納されため
		s += sep + os.Args[i] // s = s + sep + os.Args[i]と同じ
		sep = " "
	}
	fmt.Println(s)
}
