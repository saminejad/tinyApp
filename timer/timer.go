package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	j := 0
	for i := 0; i >= 0; i++ {
		time.Sleep(time.Second)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("%d m : %d s", j, i)
		if i == 59 {
			i = 0
			j++
		}

	}
}
