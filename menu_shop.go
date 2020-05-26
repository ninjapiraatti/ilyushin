package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	buttonExitImage *ebiten.Image
	buttonBuyImage  *ebiten.Image
	MenuShop        = Panel{}
)

// InitMenuShop initiates the shop menu
func InitMenuShop() {
	MenuShop.active = false
	buttonExitImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonBuyImage, _, err = ebitenutil.NewImageFromFile("assets/button.png", ebiten.FilterDefault)
	buttonExit := Button{*buttonExitImage, 20, 400, ShopExit, "Exit shop"}
	buttonBuy := Button{*buttonBuyImage, 160, 400, Buy, "Buy this"}
	MenuShop.buttons = append(MenuShop.buttons, buttonExit, buttonBuy)
}

// Buy buys
func Buy() {
	fmt.Println("Shopped!")
}

// ShopExit exits the shop
func ShopExit() {
	UI76.currentPanel = MenuLanded
	fmt.Println("Exited shop!")
}
