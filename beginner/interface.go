package main

import "fmt"

type Programmer interface {
	WriteCode() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteCode() string {
	return "fmt.Println()"
}

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
	}
	if i, ok := p.(string); ok {
		fmt.Println("String", i)
	}
}

func main() {
	programmer := GoProgrammer{}
	fmt.Println(programmer.WriteCode())
}
