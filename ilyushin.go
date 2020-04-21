package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	err        error
	background *ebiten.Image
	//engine     = Part{}
)

// Run this code once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/syvattyil.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	//engine.broken = false
}

// Update the screen
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)
	ebitenutil.DebugPrint(screen, "Hello, World!")

	// Get the x, y position of the cursor from the CursorPosition() function
	x, y := ebiten.CursorPosition()

	// Display the information with "X: xx, Y: xx" format
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nX: %d, Y: %d", x, y))

	// When the "left mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'LEFT' mouse button.")
	}
	// When the "right mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		ebitenutil.DebugPrint(screen, "\n\n\nYou're pressing the 'RIGHT' mouse button.")
	}
	// When the "middle mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		ebitenutil.DebugPrint(screen, "\n\n\n\nYou're pressing the 'MIDDLE' mouse button.")
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, 1600, 1060, 0.5, "Ilyushin"); err != nil {
		log.Fatal(err)
	}
	partsTest()
}
