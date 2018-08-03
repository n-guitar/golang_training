package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "alpha"
	s = s + "bate"
	s += "gamma"

	fmt.Println(s)

	arr := []string{"alpha", "bata", "gamma"}

	fmt.Println(strings.Join(arr, ","))
}
