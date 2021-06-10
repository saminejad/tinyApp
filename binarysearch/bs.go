package main

import (
	"fmt"
)

func main() {

	x := []int{10, 50, 60, 90, 100, 200, 455, 520, 688, 777, 785, 799, 856, 889, 921}
	top, counter := len(x), 0
	bottom := 0
	search := 756
	for {
		middle := (top + bottom) / 2

		if x[middle-1] == search {
			fmt.Printf("index is %d", middle-1)
			break
		} else if x[middle-1] > search {
			top = middle - 1

		} else {
			bottom = middle + 1
		}

		counter++

		if counter >= len(x) {

			fmt.Println("not found")
			break

		}

	}

}
