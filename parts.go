package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Part is a general struct for all parts and upgrades
type Part struct {
	imageIntact ebiten.Image
	imageBroken ebiten.Image
	name        string
	hp          int
	price       int
	broken      bool
}
