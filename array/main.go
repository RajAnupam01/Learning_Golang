package main

import "fmt"

func main() {
	fmt.Println("We are learning array in golang.")

	// var name[5] string 
	// name[0] = "Mani"
	// name[1] = "Anupam"
	// name[2] = "Akash"
	// name[3] = "Maddy"
	// name[4] = "Anuradha"

	// fmt.Println("Names of the person is: ", name)

	// var numbers = [5] int{1,2,3,4,5}
	// fmt.Println("Numbers are:",numbers)

	// fmt.Println("length of character",len(numbers))

	// fmt.Println("Value of characer at second index",numbers[1])

	var price[5] int 
	fmt.Println(price)

	var decided[5]bool
	fmt.Println(decided)

	var values[5] string
	fmt.Println(values)
	fmt.Printf("values are %q\n",values)
}
