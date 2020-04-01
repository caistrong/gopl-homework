package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

// run command:
// go run ./src/chapter1/work1_12/work1_12.go

// browser open http://localhost:8888/?cycles=10.0
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := r.FormValue("cycles")
		lissajous(cycles, w)
	})
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func lissajous(cycles string, out http.ResponseWriter) {
	cyclesInt, err := strconv.ParseFloat(cycles, 64)
	if err != nil {
		out.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(out, "<html><body>You need a param 'cycles' and that must be  a number, you can visit like<a href=\"http://localhost:8888/?cycles=10.0\">http://localhost:8888/?cycles=10.0</a></body></html>")
		return
	}
	const (
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
		for t := 0.0; t < cyclesInt*2*math.Pi; t += res {
			// 0.0 -> 10Pi   step: 0.001
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// NOTE: Mac OS use comman `lsof -nP -iTCP:8888` to check which pid listen port 8888
