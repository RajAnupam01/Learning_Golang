package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("currentTime",currentTime)
	fmt.Printf("Type of currentTime %T\n",currentTime)

	formatted := currentTime.Format("01-02-2006 ,Monday, 15:04:05")
	fmt.Println("Formatted Time:",formatted)

}