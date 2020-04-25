package main

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	buttonStartNewImage *ebiten.Image
	buttonLoadImage     *ebiten.Image
	buttonQuitImage     *ebiten.Image
	MenuMainMenu        = Panel{}
)

// InitMenuMainMenu initiates the main menu
func InitMenuMainMenu() {
	MenuMainMenu.active = true
	buttonStartNewImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonLoadImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonQuitImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonStartNew := Button{*buttonStartNewImage, 80, 400, Start, "New Game"}
	buttonLoad := Button{*buttonLoadImage, 200, 400, Load, "Load Game"}
	buttonQuit := Button{*buttonQuitImage, 320, 400, Quit, "Save&Quit"}
	MenuMainMenu.buttons = append(MenuMainMenu.buttons, buttonStartNew, buttonLoad, buttonQuit)
}

// Start starts the game
func Start() {
	GS.current = landed
	UI76.currentPanel = MenuLanded
}

// Quit quits
func Quit() {
	Save()
	os.Exit(1)
}
