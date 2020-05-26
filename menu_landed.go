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
	buttonMapImage     *ebiten.Image
	buttonLEventImage  *ebiten.Image
	MenuLanded         = Panel{}
)

// InitMenuLanded initiates the landed menu
func InitMenuLanded() {
	MenuLanded.active = true
	buttonTakeoffImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonTakeoff := Button{*buttonTakeoffImage, 160, 400, Takeoff, "Take off"}
	buttonShopImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonShop := Button{*buttonShopImage, 300, 400, Shop, "Shop"}
	buttonMapImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonMap := Button{*buttonMapImage, 440, 400, GotoMap, "Map"}
	buttonLEventImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonEvent := Button{*buttonLEventImage, 580, 400, EventFunc, "Event"}
	MenuLanded.buttons = append(MenuLanded.buttons, buttonTakeoff, buttonShop, buttonMap, buttonEvent)
}

// Takeoff takes off
func Takeoff() {
	if GS.current == landed {
		UI76.allowInterface = false
		fmt.Println("Takeoff!")
		GS.current = takeoff
		fmt.Println("Taking off!")
		time.Sleep(3 * time.Second)
		GS.current = flying
		UI76.currentPanel = MenuFlying
		fmt.Println("Flying!")
		UI76.allowInterface = true
	}
}

// Shop shops
func Shop() {
	UI76.currentPanel = MenuShop
	MenuShop.active = true
	UI76.prevPanel = MenuLanded
}

// GotoMap goes to shop
func GotoMap() {
	UI76.currentPanel = MenuMap
	MenuMap.active = true
	UI76.prevPanel = MenuLanded
}
