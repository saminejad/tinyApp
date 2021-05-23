package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `
	hello world
this is text in python language
www.google.com
https://digikala.com
http://www.google.com
www.saminejad.ir`

	for _, i := range strings.Split(s, "\n") {
		if strings.Index(i, "www") != -1 || strings.Index(i, "http") != -1 {
			fmt.Println(i)
		}
	}
}
