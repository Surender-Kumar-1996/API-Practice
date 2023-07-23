package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("LCO web request....")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("/t-------------------------")
	fmt.Printf(">*Resp : %v its type  :%T\n", resp, resp)
	fmt.Println(">*Resp status: ", resp.Status)
	fmt.Println(">*Response Status code: ", resp.StatusCode)
	fmt.Println(">*Response Header: ", resp.Header)
	fmt.Println(">*Response TLS: ", resp.TLS)
	fmt.Println(">*Response length: ", resp.ContentLength)
	fmt.Println("/t-------------------------")
	loc, err := resp.Location()
	if err != nil {
		panic(err)
	}
	fmt.Println(">*Location URL value: ", loc)
	fmt.Println(">*URL Host: ", loc.Host)
	fmt.Println(">*URL Path: ", loc.Path)
	fmt.Println("/t-------------------------")

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println("Content of the respnse", content)

}
