package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

// Variables
var (
	font76 font.Face
)

// UI will hold it all
type UI struct {
	buttons         []Button
	allowMousePress bool
	panels          []Panel
	currentPanel    Panel
}

// Panel is a user interfaces
type Panel struct {
	buttons []Button
	panels  []Panel
	active  bool
}

// Button is a thing you press
type Button struct {
	image      ebiten.Image
	xPos, yPos float64
	action     func()
	label      string
}

// InitUI initializes the UI
func InitUI() UI {
	UI76 := UI{}
	UI76.allowMousePress = true
	UI76.panels = append(UI76.panels, MenuMainMenu, MenuLanded, MenuFlying)
	InitMenuMainMenu()
	InitMenuLanded()
	InitMenuFlying()
	UI76.currentPanel = MenuMainMenu
	// Fonts
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	font76 = truetype.NewFace(tt, &truetype.Options{
		Size:    8,
		DPI:     96,
		Hinting: font.HintingFull,
	})
	return UI76
}

// UpdateUI updates the UI
func UpdateUI(screen *ebiten.Image) {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		GS.paused = true
		UI76.currentPanel = MenuMainMenu
	}
	for i := range UI76.currentPanel.buttons {
		btn := UI76.currentPanel.buttons[i]
		opUI := &ebiten.DrawImageOptions{}
		opUI.GeoM.Translate(btn.xPos, btn.yPos) // Moves the object
		screen.DrawImage(&btn.image, opUI)
		text.Draw(screen, btn.label, font76, int(btn.xPos+20), int(btn.yPos+30), color.White)
	}
	ebitenutil.DebugPrint(screen, GS.debuginfo)
	// When the "left mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && UI76.allowMousePress == true {
		// Get the x, y position of the cursor from the CursorPosition() function
		x, y := ebiten.CursorPosition()
		//fmt.Println(GS.debuginfo, GS.current)
		// Display the information with "X: xx, Y: xx" format
		ebitenutil.DebugPrint(screen, fmt.Sprintf("\nX: %d, Y: %d", x, y))
		UIIsButtonPressed(x, y, UI76.currentPanel)
	}
}

// UIIsButtonPressed tells if a UI button was pressed and returns the button
func UIIsButtonPressed(x int, y int, currentPanel Panel) bool {
	//ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'LEFT' mouse button.")
	for i := range currentPanel.buttons {
		xButton := currentPanel.buttons[i].xPos
		yButton := currentPanel.buttons[i].yPos
		maxPress := currentPanel.buttons[i].image.Bounds().Max
		if (x > int(xButton) && x < int(xButton)+maxPress.X) && (y > int(yButton) && y < int(yButton)+maxPress.Y) {
			currentPanel.buttons[i].action()
		}
	}
	return true
}
