package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("We are learning function is Golang")
	ans := add(3, 4)
	fmt.Println("The addition of two number is ", ans)
}
