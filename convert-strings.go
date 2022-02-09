package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Begin by importing the Go standard library strings"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.Title(str))

	str = "     52    "
	fmt.Printf("\n%v Type: %T,Length:%v\n", str, str, len(str))
}
