package main

// Create our empty vars
var (
	mainmenu = State{}
	landed   = State{}
	takeoff  = State{}
	flying   = State{}
	landing  = State{}
	credits  = State{}
)

// GameState will hold info of the gamestate
type GameState struct {
	current   State
	paused    bool
	debuginfo string
}

//State is a gamestate
type State struct {
	name string
}
