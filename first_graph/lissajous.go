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
	//"os"
)


var (
	red = color.RGBA{0xff,0x00,0x00,0xff}
	green = color.RGBA{0x00,0xff,0x00,0xff}
	blue = color.RGBA{0x00,0x00,0xff,0xff}
	cyan = color.RGBA{0x00,0xff,0xff,0xff}
	magenta = color.RGBA{0xff,0x00,0xff,0xff}
	yellow = color.RGBA{0xff,0xff,0x00,0xff}
)
var palette = []color.Color{color.White, red, red, yellow, yellow, green, green, cyan, cyan, blue, blue, magenta, magenta}

// 前者生成的是一个slice切片
//var palette = []color.Color{color.White, green}

// const 数字 字符串 固定的Boolean值
const (
	blackIndex = 0
	greenIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	// 生成的是一个struct结构体
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i<=nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(i%(len(palette)-1)+1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	//lissajous(os.Stdout)
	log.Fatal(http.ListenAndServe("localhost:8999", nil))
}