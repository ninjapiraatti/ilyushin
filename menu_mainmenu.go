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
	buttonResumeImage   *ebiten.Image
	buttonQuitImage     *ebiten.Image
	MenuMainMenu        = Panel{}
)

// InitMenuMainMenu initiates the main menu
func InitMenuMainMenu() {
	MenuMainMenu.active = true
	buttonStartNewImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonLoadImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonResumeImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonQuitImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonStartNew := Button{*buttonStartNewImage, 80, 400, StartNew, "New Game"}
	buttonLoad := Button{*buttonLoadImage, 200, 400, Load, "Load Game"}
	buttonQuit := Button{*buttonQuitImage, 320, 400, Quit, "Save&Quit"}
	buttonResume := Button{*buttonResumeImage, 440, 400, Resume, "Resume"}
	MenuMainMenu.buttons = append(MenuMainMenu.buttons, buttonStartNew, buttonLoad, buttonQuit, buttonResume)
}

// StartNew starts the game
func StartNew() {
	GS.paused = false
	GS.current = landed
	UI76.currentPanel = MenuLanded
}

// Resume resumes the game
func Resume() {
	GS.paused = false
	MenuMainMenu.active = false
	UI76.currentPanel = UI76.prevPanel
}

// Quit quits
func Quit() {
	Save()
	os.Exit(1)
}
