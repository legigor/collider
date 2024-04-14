package main

import (
	"fmt"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

var (
	ww = 1024.0
	hh = 768.0
)

func run() {
	cfg := pixelgl.WindowConfig{
		//Title:  "Let's draw!",
		Bounds:    pixel.R(0, 0, ww, hh),
		VSync:     true, // check if we need it on manual update
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {

		win.Clear(colornames.Antiquewhite)
		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
