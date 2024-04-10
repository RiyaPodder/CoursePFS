package main

import (
	"canvas"
	"image"
)

//place your drawing code here.

func AnimateSystem(timePoints []Sky, canvasWidth, imageFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%imageFrequency == 0 { //only draw if current index of universe
			//is divisible by some parameter frequency
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

func DrawToCanvas(s Sky, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()
	boid_radius := 5.0
	// range over all the bodies and draw them.
	for _, b := range s.boids {
		c.SetFillColor(canvas.MakeColor(255, 255, 255))


		cx := (b.position.x / s.width) * float64(canvasWidth)
		cy := (b.position.y / s.width) * float64(canvasWidth)

		cw := float64(canvasWidth)

		if cx < 0.0{
			cx = cx + cw
		}else if cx > cw{
			cx = cx - cw
		}

		if cy < 0.0{
			cy = cy + cw
		}else if cy > cw{
			cy = cy - cw
		}

		r := (boid_radius / s.width) * float64(canvasWidth)

		c.Circle(cx, cy, r)

		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
