package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hits", countsShower)
	http.ListenAndServe("localhost:8081", nil)
}

func rootHandler(httpRespWriter http.ResponseWriter, request *http.Request) {
	log.Print("Received a request. Incrementing the counter")
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(httpRespWriter, "URL.path = %q\n", request.URL)
}

func countsShower(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hits so far : %d\n", count)
}
