package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i, math.Sqrt(float64(i)))
	}
}
