package main

import "fmt"

func main() {
	
	numbers := []int {1,2,3,4,5}
	for i,val := range numbers{
		fmt.Printf("Index of Numbers is %d, and value is %d\n",i,val)
	}

	data := "Hello world"

	for x, char:= range data{
		fmt.Printf("Index of Data is %d, and char is %c\n", x, char)
	}
}