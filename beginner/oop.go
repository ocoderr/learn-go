package main

import "fmt"

type Employee struct {
	Id   string
	Name string
	Age  int
}

func main() {
	employee := Employee{"0", "bob", 20}
	e := new(Employee)
	e.Name = "Alice"
	e.Id = "2"
	e.Age = 30

	fmt.Println(employee)
	fmt.Println(e)
	fmt.Printf("%T\n", employee)
	fmt.Printf("%T", e)
}
