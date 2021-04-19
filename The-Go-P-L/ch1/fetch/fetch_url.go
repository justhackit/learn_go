package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		fetch_url(url)
	}
}

func fetch_url(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get : %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	bytesRead, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch : error reading %s : %v", url, err)
		os.Exit(1)
	}
	fmt.Printf("Total Bytes read : %d\n", bytesRead)
}
