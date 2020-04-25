package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten"
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
	InitGamestates()
	GS = GameState{}
	GS.current = landed
	ebiten.SetMaxTPS(30)
}

// Update the screen
func update(screen *ebiten.Image) error {
	GS.debuginfo = strconv.FormatBool(UI76.allowMousePress)
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if GS.paused == false {
		tick++
		UpdatePlane(screen)
		go UpdateUI(screen)
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 800, 530, 1, "Ilyushin"); err != nil {
		log.Fatal(err)
	}
}
