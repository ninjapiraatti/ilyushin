package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Plane will hold all the info of the plane
type Plane struct {
	image    ebiten.Image
	altitude float64
	velocity float64
	name     string
	stalled  bool
	crashed  bool
	landed   bool
	parts    []Part
	pitch    float64
	yaw      float64
}
