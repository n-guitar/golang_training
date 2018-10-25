// https://dev.classmethod.jp/server-side/language/golang-4/
package main

import "fmt"

//関数の型を宣言
type funcTemplate func(string) string

func greet(name string) string {
	return "hello," + name
}

func x(f funcTemplate) {
	fmt.Println(fmt.Sprintf("%T", f))
}

func main() {
	//main.funcTemplateが出力される
	x(greet)
}
