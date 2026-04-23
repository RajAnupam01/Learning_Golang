package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.8298

	fmt.Println("name:", name, "age :", age, "height:", height)
	fmt.Println("Hello world")

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type is name is %s\n", name)
	fmt.Printf("Name: %s,Age: %d, Height: %.2f", name, age, height)
}
