package main

import "fmt"

type Vector struct {
	X int
	Y int
}

func main() {
	var v Vector

	v.X = 2
	v.Y = 5
	fmt.Println(v)
}
