// Echo2は、そのコマンドラインの引数を表示します。
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	/* for文の構文は
		for initialization; condition; post{
			0個以上の文
		}
	initializationはループが開始される前に実行される
	conditionはブーリアン式でループの繰り返し前に評価され、
	trueになればループ文で制御されている文が実行される
	post文はループの本体のあとに実行される。→その後conditionが評価される。
	conditionがfalseの場合にループが終了
	↑が基本。
	↓は文字列やスライスなどのデータを値を範囲(range)を繰り返す書き方。
	os.Argsは以下の特徴を持つ
		os.Argsはstring型のスライス
		0番目の要素には実行したコマンド名が格納される
		1番目以降の各要素にコマンドに渡された各引数が格納される

	os.Args[1:]は
	スライス式のs[m:n]のnが省略された形
	s[m:n]はm番目からn-1番目の要素を参照するスライスを生成する
	mが省略された場合は0
	nが省略された場合はlen(s)
	*/
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
