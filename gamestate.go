package main

// Create our empty vars
var (
	mainmenu = State{}
	landed   = State{}
	takeoff  = State{}
	flying   = State{}
	landing  = State{}
	crashing = State{}
	credits  = State{}
)

// GameState will hold info of the gamestate
type GameState struct {
	current   State
	paused    bool
	debuginfo string
	money     int
}

//State is a gamestate
type State struct {
	name string
	xPos float64
	yPos float64
	rot  float64
}

// InitGamestates initiates state
func InitGamestates() {
	landed.name = "Landed"
	landed.xPos = 200
	landed.yPos = 200
	takeoff.name = "Takeoff"
	takeoff.xPos = 50
	takeoff.yPos = 290
	takeoff.rot = -30
	flying.name = "Flying"
	flying.xPos = 200
	flying.yPos = 50
	landing.name = "Landing"
	landing.xPos = 270
	landing.yPos = 20
	landing.rot = 20
}
