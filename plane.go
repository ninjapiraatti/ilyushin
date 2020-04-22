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
	planeImage, _, err = ebitenutil.NewImageFromFile("assets/syvattyil2.png", ebiten.FilterDefault)
}

// UpdatePlane updates the plane
func UpdatePlane(screen *ebiten.Image) {
	opPlane := &ebiten.DrawImageOptions{}
	opPlane.GeoM.Rotate(float64(tick%360) * 2 * math.Pi / 360) // Rotates the object
	screen.DrawImage(planeImage, opPlane)
}
