package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.nike.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to hit the URL : %v", err)
		os.Exit(1)
	}
	_, err1 := io.Copy(os.Stdout, resp.Body)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "error while copying : %v", err1)
	}
}
