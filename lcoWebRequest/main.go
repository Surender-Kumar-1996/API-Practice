package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb .....")
	// PerformGetRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:8000/get"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status of response: ", resp.Status)
	fmt.Println("Content legth Received: ", resp.ContentLength)

	if resp.StatusCode == http.StatusOK {
		// receiveBytes, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("Received response: ", string(receiveBytes))
		var responseString strings.Builder
		content, _ := io.ReadAll(resp.Body)
		bytecount, _ := responseString.Write(content)

		fmt.Println("Bytecount is: ", bytecount)
		fmt.Println("Response received: ", responseString.String())

	} else {
		fmt.Println("Unexpected status returned")
	}
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)

	resp, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Responce body: ", resp.Body)
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response Data: ", string(content))

}

func PerformFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "surender")
	data.Add("lastname", "kumar")
	data.Add("email", "kumar7208383@gmail.com")

	resp, err := http.PostForm(myurl, data)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)

	fmt.Println("Response: ", string(content))
}
