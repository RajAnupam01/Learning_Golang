package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning web request.")

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error while getting response.")
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while reading response", err)
		return
	}

	fmt.Println("Response", string(data))

	var myTodo Todo
	err = json.Unmarshal(data, &myTodo)
	if err != nil {
		fmt.Println("Error while UnMarshalling.")
	}

	fmt.Println("myTodo",myTodo.Title)

}
