package main

import "fmt"

func revstring(str string) {
	arr := []byte(str)

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(string(arr[i]))
	}

}
func main() {

	revstring("hello world")
}
