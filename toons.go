package main

import "github.com/hajimehoshi/ebiten"

// Toon is a struct for characters
type Toon struct {
	image      ebiten.Image
	xPos, yPos float64
	speed      float64
	name       string
	age        int
}

// Crew is a struct to hold the characters
type Crew struct {
	members []Toon
}
