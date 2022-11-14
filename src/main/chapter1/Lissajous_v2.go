package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var paletteV2 = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	whiteIndexV2 = 0
	blackIndexV2 = 1
)

/**
 * Usage: go run Lissajous_v2.go > LissajousV2.gif - then open the file
 * Alternatively you can go build the Lissajous_v2.go file and run the
 * executable, redirecting the output to a file - as mentioned in the
 * README.md in the root of this project.
 */
func main() {
	LissajousV2(os.Stdout)
}

func LissajousV2(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
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
