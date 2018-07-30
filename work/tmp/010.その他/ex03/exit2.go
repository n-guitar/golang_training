package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.OpenFile("test", os.O_RDONLY, 0)
	if err != nil {
		log.Fatalln("エラー", err) //それぞれ対応する Print 関連関数にしたがってログメッセージを出力した後、os.Exit(1) を呼び出してプログラムを終了させます。
	}
}
