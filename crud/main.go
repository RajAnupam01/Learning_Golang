package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest(){
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Request failed ! ", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Server returned error ! ", res.Status)
		return
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error while reading data")
	}

	var myTodo Todo
	err = json.Unmarshal(data, &myTodo)
	if err != nil {
		fmt.Println("Error while unmarshalling the json")
	}

	fmt.Println("The title of myTodo", myTodo.Title)
}





func main() {
	performGetRequest()
}
