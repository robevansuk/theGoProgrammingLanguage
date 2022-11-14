package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/**
 * Usage: go run HttpWithChannels.go https://golang.org http://gopl.io https://godoc.org
 * Output:
 * 0.69s     55097   https://golang.org
 * 0.92s     30539   https://godoc.org
 * 1.84s      4154   http://gopl.io
 *
 * Channel - a communication mechanism for communicating between goroutines
 * goroutine - new thread / concurrent function execution
 */
func main() {
	start := time.Now()
	ch := make(chan string) // create a channel for strings
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a go Routine - asynchronous function
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send error to channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // returns the number of bytes
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs   %7d   %s", secs, nbytes, url) // send a summary line to the channel ch
}
