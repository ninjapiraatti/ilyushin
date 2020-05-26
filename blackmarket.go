package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Good is a struct for black market items
type Good struct {
	image     ebiten.Image
	name      string
	buyPrice  int
	sellPrice int
	weight    int
	space     int
}

// Market is a struct for individual markets
type Market struct {
	goods []Good
}
