package main

import "fmt"

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	m1 := map[string]int{}

	fmt.Println(len(m))
	fmt.Println(len(m1))

	for k, v := range m {
		fmt.Printf("%v  %d\n", k, v)
	}

	if i, ok := m["a"]; !ok {
		fmt.Println("not exist")
	} else {
		fmt.Println(i)
	}
}
