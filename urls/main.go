package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbjhbc"

func main() {
	fmt.Println("Welcome to handline URL's in golang")
	fmt.Println("MyUrl: ", myurl)

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("/t---------------------")
	fmt.Println("Url Scheme: ", result.Scheme)
	fmt.Println("Url Host: ", result.Host)
	fmt.Println("Url Path: ", result.Path)
	fmt.Println("Url Port: ", result.Port())
	fmt.Println("Url Query: ", result.Query())
	fmt.Println("/t---------------------")

	queryParams := result.Query()

	for i, j := range queryParams {
		fmt.Println(i, " : ", j)
	}

	//Conscructing a url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	resultUrl := partsOfUrl.String()
	fmt.Println("Constructed url : ", resultUrl)

}
