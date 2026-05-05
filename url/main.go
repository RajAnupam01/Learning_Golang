package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning the URL")

	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

// 	fmt.Printf("Type of myURL is %T\n",myURL)

  	parsedURL,err := url.Parse(myURL)

	if err != nil {
		fmt.Println("Error regarding parsing of URL",err)
		return 
	}

	// fmt.Printf("Type of parsedURL %T\n ",parsedURL) 

		fmt.Println("Scheme is ",parsedURL.Scheme)
		fmt.Println("Host is ",parsedURL.Host)
		fmt.Println("Path is ",parsedURL.Path)
		fmt.Println("RawQuery is ",parsedURL.RawQuery)

}