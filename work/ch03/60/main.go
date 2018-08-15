package main

import (
	"fmt"
)

var x uint8 = 1<<1 | 1<<5

// 00000001 << 1 = 00000010 | 00000001 << 5 = 0010000
var y uint8 = 1<<1 | 1<<2

// 00000001 << 1 = 00000010 | 00000001 << 2 = 0000100

func main() {
	fmt.Printf("%08b\n", x) // 00100010
	fmt.Printf("%08b\n", y) // 00000110

	fmt.Printf("%08b\n", x&y)
	/*	00100010
			00000110
		AND 00000010
		両方が1の場合1　それ以外は0
	*/
	fmt.Printf("%08b\n", x|y)
	/*	00100010
			00000110
		OR  00100110
	*/
	fmt.Printf("%08b\n", x^y)
	/*	00100010
			00000110
		XOR 00100100
		片方が1の時1
	*/
	fmt.Printf("%08b\n", x&^y)
	/*		00100010
			11111001  ←NOT　00000110
	AND NOT 00100000
	x AND (NOT y)
	*/

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

}
