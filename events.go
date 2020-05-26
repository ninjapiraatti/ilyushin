package main

import (
	"github.com/hajimehoshi/ebiten"
)

// Event is a multipurpose struct for events
type Event struct {
	image       ebiten.Image
	description string
	choices     []Choice
}

// Choice is a thing you might choose
type Choice struct {
	description string
	effect      func()
}
