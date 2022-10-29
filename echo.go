package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command: ", os.Args[0])

	for idx, arg := range os.Args[1:] {
		fmt.Printf("%d. %q\n", idx+1, arg)
	}
}
