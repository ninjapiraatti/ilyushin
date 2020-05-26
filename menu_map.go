package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	buttonExitMapImage *ebiten.Image
	buttonFlyImage     *ebiten.Image
	MenuMap            = Panel{}
)

// InitMenuMap initiates the Map menu
func InitMenuMap() {
	MenuMap.active = false
	buttonExitMapImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonFlyImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonExitMap := Button{*buttonExitMapImage, 20, 400, MapExit, "Exit map"}
	buttonFly := Button{*buttonFlyImage, 160, 400, Fly, "Fly here"}
	MenuMap.buttons = append(MenuMap.buttons, buttonExitMap, buttonFly)
}

// Fly Flys
func Fly() {
	fmt.Println("Mapped!")
}

// MapExit exits the Map
func MapExit() {
	UI76.currentPanel = MenuLanded
	fmt.Println("Exited map!")
}
