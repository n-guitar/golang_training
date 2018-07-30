package main

import (
	"flag"
	"fmt"
)

func main() {
	// コマンドパラメータの取得
	// パラメータ名は,パラメータ名、デフォルト値、説明文
	//　戻り値はポインタ

	i := flag.Int("flag1", 123, "整数")
	s := flag.String("flag2", "abc", "文字列")

	// パラメータの解析
	flag.Parse()

	// 取得したパラメータの出力
	fmt.Println(*i, *s)
}
