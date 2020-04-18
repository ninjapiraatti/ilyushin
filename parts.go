package main

import (
	"github.com/hajimehoshi/ebiten"
)

type part struct {
	imageIntact ebiten.Image
	imageBroken ebiten.Image
	name        string
	hp          int
	price       int
	broken      bool
}
