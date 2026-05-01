package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var num int = 42
	// fmt.Println("num is ",num)
	// fmt.Printf("Type of num is %T\n",num)



	// var data float64 = float64(num)
	// data = data + 1.23
	// fmt.Println("data is ",data)
	// fmt.Printf("Type of data is %T\n",data)


	// num := 123
	// str := strconv.Itoa(num)
	// fmt.Println("num is ",num)
	// fmt.Printf("Type of num is %T\n",str)


	// number_string := "123"
	// number_int,_ := strconv.Atoi(number_string)
	// fmt.Println("str is ",number_int)
	// fmt.Printf("Type of number_int is %T\n",number_int)


	  num_string := "3.14"
	  num_float,_ := strconv.ParseFloat(num_string,64)
	  fmt.Println("number_float is ", num_float)
	  fmt.Printf("Type of number_float is %T\n",num_float)

}