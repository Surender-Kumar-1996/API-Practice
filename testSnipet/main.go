package main

import (
	"fmt"
	"regexp"
)

var user = regexp.MustCompile(`^\/users\/(\d+)$`)

func main() {
	fmt.Println("Hello, 世界")
	Matches := user.FindStringSubmatch("/users/1234")
	fmt.Println(Matches)
}
