package main

import (
	"fmt"
)

const (
	KiB = 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
	TiB = GiB * 1024
	PiB = TiB * 1024
	EiB = PiB * 1024
	ZiB = EiB * 1024
	YiB = ZiB * 1024
)

func main() {
	fmt.Println("KiB:", KiB)
	fmt.Println("MiB:", MiB)
	fmt.Println("TiB:", TiB)
	fmt.Println("PiB:", PiB)
	fmt.Println("EiB:", EiB)
	// fmt.Println("ZiB:", ZiB)
	// fmt.Println("YiB:", YiB)
}
