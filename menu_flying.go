package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	buttonLandImage  *ebiten.Image
	buttonEventImage *ebiten.Image
	MenuFlying       = Panel{}
)

// InitMenuFlying initiates the flight menu
func InitMenuFlying() {
	MenuFlying.active = true
	buttonLandImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonEventImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonLand := Button{*buttonLandImage, 20, 400, Land, "Land"}
	buttonEvent := Button{*buttonEventImage, 160, 400, Event, "Event"}
	MenuFlying.buttons = append(MenuFlying.buttons, buttonLand, buttonEvent)
}

// Land lands
func Land() {
	if GS.current == flying {
		UI76.allowMousePress = false
		fmt.Println("Land!")
		GS.current = landing
		fmt.Println("Landing!")
		time.Sleep(3 * time.Second)
		GS.current = landed
		UI76.currentPanel = MenuLanded
		fmt.Println("Landed!")
		UI76.allowMousePress = true
	}
}

// Event events!
func Event() {
	fmt.Println("Event!")
}
