package main

import (
	"github.com/hajimehoshi/ebiten"
)

type plane struct {
	image    ebiten.Image
	altitude float64
	velocity float64
	name     string
	stalled  bool
	crashed  bool
	landed   bool
}
