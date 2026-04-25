package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int , b int)int{
	return a-b
}

func multiply(a int, b int)(result int){
	result = a*b;
	return result
}

func main() {
	fmt.Println("We are learning function is Golang")
	ans := add(3, 4)
	fmt.Println("The addition of two number is ", ans)
	val := subtract(10,5)
	fmt.Println("The subtraction of two number is",val)
	res := multiply(3,4)
	fmt.Println("The multiplication of two number is ",res)
}
