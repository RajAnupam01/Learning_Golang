package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("We are learning Json")
	person := Person{Name:"John",Age:34, IsAdult:true}
	fmt.Println("Person Data is",person)

   	jsonData,err :=	json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling",err)
	}

	fmt.Println("person Data is ",string(jsonData))

	var personData Person
	err = json.Unmarshal(jsonData,&personData)
	if err != nil{
		fmt.Println("Error Unmarshalling",err)
	}

	fmt.Println("Person Data is",personData)
}
