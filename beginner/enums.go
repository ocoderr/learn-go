package main

import "fmt"

func main() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)

	fmt.Println(cpp, java, python, golang)
	enums()
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
