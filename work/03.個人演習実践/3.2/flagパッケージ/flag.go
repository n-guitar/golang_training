package main

import (
	"flag"
	"fmt"
)

var version = "1.0.0"

func main() {
	var showVersion bool
	// -v -version が指定された場合に変数が真になる

	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse() //引数からオプションをパースする

	if showVersion {
		fmt.Println("version", version)
		return
	}
}
