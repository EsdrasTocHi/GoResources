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

// What is a struct?
// is a group of values called fields.

// The expressions []color.Color{...} and gif.GIF{...} are composite literals
// []color.Color{color.White, color.Black} is a slice
// gif.GIF{LoopCount: nframes} is a struct
// Composite literal is a compact notation for instantiating any of Go’s composite types from
// a sequence of element values that are collected together in a single object that can be
// treated as a unit.

var palette = []color.Color{color.White, color.Black}

// The value of a constant must be a number, string or boolean
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// To run this code use:
//  go run lissajous.go > out.gif

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
