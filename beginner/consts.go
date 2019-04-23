package main

import (
	"fmt"
	"math"
)

func main() {
	consts()
}
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
