package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("Example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating a file", err)
	// }
	// defer file.Close()

	// content := "Hello world by prince"

	// byte, error := io.WriteString(file, content+"\n")
	// fmt.Println("byte written", byte)
	// if error != nil {
	// 	fmt.Println("Error while writing a file", error)


	file,err := os.Open("Example.txt")
	if err != nil{
		fmt.Println("Error while opening the file",err)
	}
	defer file.Close()

	buffer := make([]byte,1024)

	for{
		n,err := file.Read(buffer)
		if err == io.EOF{
			break
		}
		if err != nil {
			fmt.Println("Error while reading the file",err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}


	}


