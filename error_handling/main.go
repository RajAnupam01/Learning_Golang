package main

import "fmt"

func divide(a float32, b float32) (float32, error) {
	if b == 0 {
	return 0,fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error Handling in the function")
	ans, _ := divide(10, 0)
	fmt.Println("Divison of two number is ", ans)

}
