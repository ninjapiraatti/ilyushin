package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	buttonTakeoffImage *ebiten.Image
	buttonShopImage    *ebiten.Image
	MenuLanded         = Panel{}
)

// InitMenuLanded initiates the landed menu
func InitMenuLanded() {
	MenuLanded.active = true
	buttonTakeoffImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonTakeoff := Button{*buttonTakeoffImage, 160, 400, Takeoff}
	buttonShopImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonShop := Button{*buttonShopImage, 300, 400, Shop}
	MenuLanded.buttons = append(MenuLanded.buttons, buttonTakeoff, buttonShop)
}

// Takeoff takes off
func Takeoff() {
	if GS.current == landed {
		UI76.allowMousePress = false
		fmt.Println("Takeoff!")
		GS.current = takeoff
		fmt.Println("Taking off!")
		time.Sleep(3 * time.Second)
		GS.current = flying
		UI76.currentPanel = MenuFlying
		fmt.Println("Flying!")
		UI76.allowMousePress = true
	}
}

// Shop shops
func Shop() {
	fmt.Println("Going to store!")
}
