package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.co.za")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(lw, resp.Body)
}

func (l logWriter) Write(bs []byte) (int, error) {
	// return 1, nil
	fmt.Println(string(bs))
	fmt.Printf("\n\nI just wrote %d bytes \n", len(bs))
	fmt.Println("\n\n", bs)
	return len(bs), nil
}
