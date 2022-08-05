package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "galaxy 01" || arg == "galaxy" || arg == "01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
