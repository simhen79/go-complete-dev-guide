package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		fileName := args[0]
		// Try to open and read the file
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("An error occurred while reading the file %v", err)
		}
		defer f.Close()
		var buf = make([]byte, 1024)
		for {
			n, err := f.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			if n > 0 {
				fmt.Println(string(buf[:n]))
			}
		}
	} else {
		fmt.Println("Give me a file name please!")
	}

}
