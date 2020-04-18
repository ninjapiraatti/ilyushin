package main

import "github.com/hajimehoshi/ebiten"

type toon struct {
	image      ebiten.Image
	xPos, yPos float64
	speed      float64
	name       string
	age        int
}
