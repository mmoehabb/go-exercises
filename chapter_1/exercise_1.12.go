/// An http server that serves lissajous gifs parameterized
// with 'cycles' url parameter value
package main

import (
	"image"
	"image/color"
	"image/gif"

	"math"
	"math/rand"
	"strconv"

	"fmt"
	"io"
	"log"
	"net/http"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x55, 0x47, 0x9A, 0xFF},
	color.RGBA{0x26, 0x20, 0x46, 0xFF},
	color.RGBA{0x55, 0x47, 0x9A, 0xFF},
	color.RGBA{0x26, 0x20, 0x46, 0xFF},
	color.RGBA{0x55, 0x47, 0x9A, 0xFF},
}

type Config struct {
	cycles  int     // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas convers [-size..+size]
	nframes int     // total number of frames
	delay   int     // delay between frames in 10ms units
}

func main() {
	config := Config{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		ncycles := getFormInt(r, "cycles")
		if ncycles > 0 {
			config.cycles = ncycles
		}
		lissajous(w, &config)
	})

	fmt.Println("Server is listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getFormVal(r *http.Request, key string) string {
	for k, v := range r.Form {
		if k != key {
			continue
		}
		return v[0]
	}
	return ""
}

func getFormInt(r *http.Request, key string) int {
	strval := getFormVal(r, key)
	if strval == "" {
		return -1
	}
	intval, err := strconv.Atoi(strval)
	if err != nil {
		return -1
	}
	return intval
}

func lissajous(out io.Writer, config *Config) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	mygif := gif.GIF{LoopCount: config.nframes}
	phase := 0.0

	for i := 0; i < config.nframes; i++ {
		rect := image.Rect(0, 0, 2*config.size+1, 2*config.size+1)
		img := image.NewPaletted(rect, palette)
		cycles := float64(config.cycles)

		for t := 0.0; t < cycles*2.0*math.Pi; t += config.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			size := float64(config.size)
			color_index := int(t*2)%(len(palette)-1) + 1

			img.SetColorIndex(
				config.size+int(x*size+0.5),
				config.size+int(y*size+0.5),
				uint8(color_index),
			)
		}

		phase += 0.1
		mygif.Delay = append(mygif.Delay, config.delay)
		mygif.Image = append(mygif.Image, img)
	}

	gif.EncodeAll(out, &mygif)
}
