package main

import (
	"fmt"
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math"
	"math/rand"
	"time"
)

var (
	ww = 1024.0
	hh = 768.0
)

type circle struct {
	x  float64
	y  float64
	r  float64
	dx float64
	dy float64
	c  pixel.RGBA
}

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

	var circles []*circle
	for len(circles) < 10 {
		r := 5 + rand.Float64()*30
		c := &circle{
			x:  r + rand.Float64()*(win.Bounds().W()-r),
			y:  r + rand.Float64()*(win.Bounds().H()-r),
			r:  r,
			dx: rand.Float64()*10 - 5,
			dy: rand.Float64()*10 - 5,
			c:  pixel.RGB(rand.Float64(), rand.Float64(), rand.Float64()),
		}
		circles = append(circles, c)
	}

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {

		win.Clear(colornames.Antiquewhite)

		for _, c := range circles {
			imd := imdraw.New(nil)
			imd.Color = c.c
			imd.Push(pixel.V(c.x, c.y))
			imd.Circle(c.r, 0)
			imd.Draw(win)
		}

		for i := 0; i < len(circles); i++ {
			c1 := circles[i]
			for j := 0; j < len(circles); j++ {
				if i == j {
					continue
				}
				c2 := circles[j]
				dist := math.Sqrt(math.Pow(c1.x-c2.x, 2) + math.Pow(c1.y-c2.y, 2))
				// Collision!
				if dist < c1.r+c2.r {
					//c1.dx, c2.dx = c2.dx, c1.dx
					//c1.dy, c2.dy = c2.dy, c1.dy
					// Calculate the total mass
					totalMass := c1.r + c2.r

					// Calculate the mass difference
					massDiffC1C2 := c1.r - c2.r
					massDiffC2C1 := c2.r - c1.r

					// Calculate new velocities based on the conservation of momentum and energy
					newDxC1 := (massDiffC1C2*c1.dx + 2*c2.r*c2.dx) / totalMass
					newDyC1 := (massDiffC1C2*c1.dy + 2*c2.r*c2.dy) / totalMass
					newDxC2 := (massDiffC2C1*c2.dx + 2*c1.r*c1.dx) / totalMass
					newDyC2 := (massDiffC2C1*c2.dy + 2*c1.r*c1.dy) / totalMass

					// Assign the new velocities to the circles
					c1.dx, c1.dy = newDxC1, newDyC1
					c2.dx, c2.dy = newDxC2, newDyC2
				}
			}
			if c1.x-c1.r < 0 || c1.x+c1.r > win.Bounds().W() {
				c1.dx = -c1.dx
			}
			if c1.y-c1.r < 0 || c1.y+c1.r > win.Bounds().H() {
				c1.dy = -c1.dy
			}
			c1.x += c1.dx
			c1.y += c1.dy
		}

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
