package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

// Check checks for errors. I think.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Create our empty vars
var (
	err  error
	UI76 UI
	GS   GameState
	tick int
)

// Run this code once at startup
func init() {
	UI76 = InitUI()
	InitPlane()
	InitGamestates()
	GS = GameState{}
	GS.current = mainmenu

	f, err := os.Create("data/gamedata.dat")
	check(err)
	defer f.Close()

	testdata, err := f.WriteString("Lorem ipsum")
	check(err)
	f.Sync()

	fmt.Println(testdata)
	ebiten.SetMaxTPS(30)
}

// Update the screen
func update(screen *ebiten.Image) error {
	GS.debuginfo = strconv.FormatBool(UI76.allowInterface)
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if GS.paused == false {
		tick++
		UpdatePlane(screen)
	}
	go UpdateUI(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, 800, 530, 1, "Ilyushin"); err != nil {
		log.Fatal(err)
	}
}
