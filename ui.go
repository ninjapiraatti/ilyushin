package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	buttonTakeoffImage *ebiten.Image
	buttonLandImage    *ebiten.Image
	buttonQuitImage    *ebiten.Image
)

// UI will hold it all
type UI struct {
	buttons         []Button
	allowMousePress bool
}

// Button is a thing you press
type Button struct {
	image      ebiten.Image
	xPos, yPos float64
	action     func()
}

// InitUI initializes the UI
func InitUI() UI {
	UI76 := UI{}
	UI76.allowMousePress = true
	buttonTakeoffImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonLandImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonQuitImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonTakeoff := Button{*buttonTakeoffImage, 60, 400, Takeoff}
	buttonLand := Button{*buttonLandImage, 180, 400, Land}
	buttonQuit := Button{*buttonQuitImage, 300, 400, Quit}
	UI76.buttons = append(UI76.buttons, buttonTakeoff, buttonLand, buttonQuit)
	return UI76
}

// UpdateUI updates the UI
func UpdateUI(screen *ebiten.Image) {
	for i := range UI76.buttons {
		opUI := &ebiten.DrawImageOptions{}
		opUI.GeoM.Translate(UI76.buttons[i].xPos, UI76.buttons[i].yPos) // Moves the object
		screen.DrawImage(&UI76.buttons[i].image, opUI)
	}
	ebitenutil.DebugPrint(screen, GS.debuginfo)
}

// Takeoff is here for testing
func Takeoff() {
	fmt.Println("Takeoff!")
	//fmt.Println(testbutton.image.Bounds())
}

// Land is here for testing
func Land() {
	fmt.Println("Land!")
	//fmt.Println(testbutton.image.Bounds())
}

// Quit is here for testing
func Quit() {
	os.Exit(1)
	//fmt.Println(testbutton.image.Bounds())
}

// UIIsButtonPressed tells if a UI button was pressed and returns the button
func UIIsButtonPressed(x int, y int) bool {
	//ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'LEFT' mouse button.")
	for i := range UI76.buttons {
		xButton := UI76.buttons[i].xPos
		yButton := UI76.buttons[i].yPos
		maxPress := UI76.buttons[i].image.Bounds().Max
		if (x > int(xButton) && x < int(xButton)+maxPress.X) && (y > int(yButton) && y < int(yButton)+maxPress.Y) {
			UI76.buttons[i].action()
		}
	}
	return true
}
