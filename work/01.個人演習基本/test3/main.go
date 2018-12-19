package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// パラメータ名は,パラメータ名、デフォルト値、説明文
	//　戻り値はポインタになる
	i := flag.Int("yokoyama", 1000, "-n で整数を指定する")
	s := flag.String("s", "status", "-s で文字列を入力")

	// パラメータの解析
	flag.Parse()

	// 取得したパラメータの出力
	fmt.Println(*i, *s)

	cmd := exec.Command("ls", "*s")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s", result)
}
