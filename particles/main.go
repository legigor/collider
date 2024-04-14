package main

import (
	"fmt"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

var (
	ww = 1024.0
	hh = 768.0

	particlesNumber = 300
	radius          = 3.0
	lineWidth       = 1.0
	maxSpeed        = 1.0
	colorBackground = colornames.Black
	colorParticle   = colornames.Blue
)

type particle struct {
	x  float64
	y  float64
	dx float64
	dy float64
}

func run() {
	cfg := pixelgl.WindowConfig{
		Bounds:    pixel.R(0, 0, ww, hh),
		VSync:     true, // check if we need it on manual update
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	particles := make([]*particle, particlesNumber)
	for i := range particles {
		particles[i] = &particle{
			x:  radius + rand.Float64()*(win.Bounds().W()-radius),
			y:  radius + rand.Float64()*(win.Bounds().H()-radius),
			dx: rand.Float64()*maxSpeed*2 - maxSpeed,
			dy: rand.Float64()*maxSpeed*2 - maxSpeed,
		}
	}

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {

		win.Clear(colorBackground)

		imd := imdraw.New(nil)

		imd.Color = colorParticle

		for _, p := range particles {
			imd.Push(pixel.V(p.x, p.y))
			imd.Circle(radius, 0)

			for _, pn := range particles {
				if p == pn {
					continue
				}
				dist := pixel.V(p.x, p.y).To(pixel.V(pn.x, pn.y)).Len()
				if dist < .1*win.Bounds().W() {
					imd.Push(pixel.V(p.x, p.y))
					imd.Push(pixel.V(pn.x, pn.y))
				}
			}
			imd.Line(lineWidth)
		}

		imd.Draw(win)
		win.Update()

		for _, p := range particles {
			p.x += p.dx
			p.y += p.dy

			if p.x < radius || p.x > win.Bounds().W()-radius {
				p.dx = -p.dx
			}
			if p.y < radius || p.y > win.Bounds().H()-radius {
				p.dy = -p.dy
			}
		}

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
