package main

import (
	"fmt"
	"regexp"
)

func main() {

	num, _ := regexp.MatchString(`\D`, "hello123")
	if num {
		fmt.Println("only number")
	}
}
