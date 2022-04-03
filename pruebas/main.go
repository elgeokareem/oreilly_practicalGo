package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchRemoteResource(url string) ([]byte, error) {
	r, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

func main() {
	lol := os.Args[1:]
	fmt.Println("kekerinos", lol)
}
