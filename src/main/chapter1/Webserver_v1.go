package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	count = 0
	mu    sync.Mutex
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler func
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
