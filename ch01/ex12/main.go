package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

//var cycles int

const (
	whiteINdex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	//http.HandleFunc("/?cycles=", set_num)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		log.Panicf("予期せぬエラーが発生しました: %v\n", err)
	}
	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
	const (
		// cycles  = 5
		res     = 0.001
		size    = 100
		nfrmaes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nfrmaes}
	phase := 0.0
	for i := 0; i < nfrmaes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
