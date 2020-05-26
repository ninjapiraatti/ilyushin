package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	buttonOkImage     *ebiten.Image
	buttonChoiceImage *ebiten.Image
	MenuEvent         = Panel{}
)

// InitMenuEvent initiates the shop menu
func InitMenuEvent() {
	MenuEvent.active = false
	buttonOkImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonChoiceImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonOk := Button{*buttonOkImage, 20, 400, CloseDialog, "Ok"}
	buttonChoice := Button{*buttonChoiceImage, 160, 400, Choose, "Choose"}
	MenuEvent.buttons = append(MenuEvent.buttons, buttonOk, buttonChoice)
}

// Choose chooses
func Choose() {
	fmt.Println("Made a choice!")
}

// CloseDialog exits the event
func CloseDialog() {
	UI76.currentPanel = UI76.prevPanel
	fmt.Println("Exited dialog!")
}

// EventFunc fires the event
func EventFunc() {
	fmt.Println("Event!")
	UI76.prevPanel = UI76.currentPanel
	UI76.currentPanel = MenuEvent
	MenuEvent.active = true
}
