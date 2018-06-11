package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 20})

	// Omitted fields will be zero-valued.
	fmt.Println(person{"Fred", 88})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Alice", age: 25})

	s := person{name: "Roy", age: 29}
	fmt.Println(s.name)

	// Can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)
}
