package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// ignore errors highlighted by IDE - these files are not intended to be run as
// part of a package but as individual files.
const (
	whiteIndexV2 = 0
	blackIndexV2 = 1
)

var (
	count     = 0
	mu        sync.Mutex
	paletteV2 = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler func
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		raw := r.URL.RawQuery
		fmt.Printf(raw)
		cycles := strings.Split(raw, "=")
		lissajousV2(w, cycles[1])
	})
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
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func lissajousV2(out io.Writer, cyclesParam string) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	cycles, _ := strconv.ParseFloat(cyclesParam, 64)
	if len(cyclesParam) == 0 || cycles == 0 {
		cycles = 1
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, paletteV2)
		for t := float64(0); t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndexV2)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
