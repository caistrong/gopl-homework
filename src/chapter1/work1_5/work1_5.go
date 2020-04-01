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

var palette = []color.Color{color.White, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	whiteIndex = 0
	greenIndex = 1
)

// run command:
// go run ./src/chapter1/work1_5/work1_5.go > ./src/chapter1/work1_5/out.gif
func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8 //delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // generate a random float range (0.0, 3.0) for relative frequecy of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		// loop 64 times
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) // (0, 0, 201, 201)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// 0.0 -> 10Pi   step: 0.001
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
