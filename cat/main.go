package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		reader := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(reader)
		os.Stdin.Close()
		os.Stdout.Close()
	} else {
		files := os.Args[1:]
		for _, filename := range files {
			ReadFile(filename)
		}
	}
}

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		Print("ERROR: open " + filename + ": no such file or directory")
		z01.PrintRune('\n')
		os.Exit(1)
	} else {
		Print(string(data))
	}
}

func Print(str string) {
	for _, val := range str {
		z01.PrintRune(val)
	}
}
