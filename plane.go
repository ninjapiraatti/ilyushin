package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	planeImage *ebiten.Image
	plane      = Plane{}
)

// Plane will hold all the info of the plane
type Plane struct {
	image    ebiten.Image
	altitude float64
	velocity float64
	name     string
	stalled  bool
	crashed  bool
	landed   bool
	parts    []Part
	pitch    float64
	yaw      float64
}

// InitPlane puts the plane on screen
func InitPlane() {
	engine := Part{}
	engine.name = "Konekone"
	plane.parts = append(plane.parts, engine)
	planeImage, _, err = ebitenutil.NewImageFromFile("assets/syvattyil.png", ebiten.FilterDefault)
}

// UpdatePlane updates the plane
func UpdatePlane(screen *ebiten.Image) {
	if GS.current != mainmenu && GS.current != credits {
		opPlane := &ebiten.DrawImageOptions{}
		opPlane.GeoM.Translate(GS.current.xPos, GS.current.yPos)
		opPlane.GeoM.Scale(0.5, 0.5)
		opPlane.GeoM.Rotate(float64(int(GS.current.rot)%360) * 2 * math.Pi / 360) // Rotates the object
		screen.DrawImage(planeImage, opPlane)
	}
}
