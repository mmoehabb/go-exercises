// generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"

	"math"
	"math/rand"

	"io"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	bgColorIndex = 0
	primeColorIndex = 1
)

type Config struct {
	cycles int // number of complete x oscillator revolutions
	res float64 // angular resolution
	size int // image canvas convers [-size..+size]
	nframes int // total number of frames
	delay int // delay between frames in 10ms units
}

func main() {
	config := Config{
		cycles: 5,
		res: 0.001,
		size: 100,
		nframes: 64,
		delay: 8,
	}
	file, err := os.Create("lissajous.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lissajous(file, &config)
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
			img.SetColorIndex(
				config.size+int(x*size+0.5), 
				config.size+int(y*size+0.5), 
				primeColorIndex,
			)
		}

		phase += 0.1
		mygif.Delay = append(mygif.Delay, config.delay)
		mygif.Image = append(mygif.Image, img)
	}

	gif.EncodeAll(out, &mygif)
}

