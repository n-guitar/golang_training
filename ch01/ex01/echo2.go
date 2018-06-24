// Echo2は、そのコマンドラインの引数を表示します。
package main

import (
	"fmt"
	"os"
)

// i++ で無理やりインデックスを調整してしまっているのが良くないか？
func main() {
	for i, arg := range os.Args[1:] {
		i++
		fmt.Println(i, arg)
	}
}
