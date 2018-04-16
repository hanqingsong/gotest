package main

import (
	"fmt"
	"regexp"
)

func main() {
	isOk,error:=regexp.MatchString("[a-z]{4}","qingsong")
	isOk2,error:=regexp.MatchString("[a-z]{4}","123")
	if error!=nil {
		panic(error)
	}
	fmt.Println(isOk)
	fmt.Println(isOk2)

}