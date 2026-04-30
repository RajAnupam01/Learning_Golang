package main

import "fmt"

func main() {
	// var num int = 2
	// var ptr *int = &num

	// num := 2
	// ptr := &num

	// fmt.Println("Num has value", num)
	// fmt.Println("Pointer contains", ptr)
	// fmt.Println("Data contained through pointer",*ptr)

	// var pointer *int

	// if pointer == nil {
	// 	fmt.Println("Pointer is not assigned.")
	// }

	value := 10
	modifyValueByRefrence(&value)
	fmt.Println("Data contained through pointer",value)

}

func modifyValueByRefrence(num *int){
	*num = *num * 20
}