package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web request.")
	
	response,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	
	if err != nil {
		fmt.Println("Error while getting response.")
		return
	}
	defer response.Body.Close()

	data,err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while reading response",err)
		return 
	}

	fmt.Println("Response", string(data))
	
}