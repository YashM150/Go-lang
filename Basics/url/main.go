package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("This is url handling demo:")

	myurl := "https://example.com:8080/path/resource/to?key1=value1&key2=value2"

	parsedurl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("There is error while parsing:", err)
		return
	}
	fmt.Println("Scheme:", parsedurl.Scheme)
	fmt.Println("Host:", parsedurl.Host)
	fmt.Println("Path:", parsedurl.Path)
	fmt.Println("Params:", parsedurl.RawQuery)

	fmt.Println("Modified query:")
	parsedurl.Path = "/v1/resource"
	parsedurl.RawQuery = "apikey=key1&kvm=key2"

	newurl := parsedurl.String()

	fmt.Println("New url is:", newurl)
}
