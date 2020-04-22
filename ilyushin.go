package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	err  error
	UI76 UI
	GS   GameState
	tick int
)

// Run this code once at startup
func init() {

	if err != nil {
		log.Fatal(err)
	}
	UI76 = InitUI()
	InitPlane()
	GS = GameState{}
	GS.current = mainmenu
	GS.debuginfo = "Debug info"
	ebiten.SetMaxTPS(30)
}

// Update the screen
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if GS.paused == false {
		tick++
		UpdatePlane(screen)
		UpdateUI(screen)
	}

	// When the "left mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && UI76.allowMousePress == true {
		// Get the x, y position of the cursor from the CursorPosition() function
		x, y := ebiten.CursorPosition()
		// Display the information with "X: xx, Y: xx" format
		ebitenutil.DebugPrint(screen, fmt.Sprintf("\nX: %d, Y: %d", x, y))
		UIIsButtonPressed(x, y)
		UI76.allowMousePress = false
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 800, 530, 1, "Ilyushin"); err != nil {
		log.Fatal(err)
	}
}
