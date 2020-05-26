package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Savedata is a struct for characters
type Savedata struct {
	Gamestate State
	Name      string
}

// Save saves the game
func Save() {
	data := Savedata{
		Name:      GS.current.name,
		Gamestate: GS.current,
	}

	f, err := json.MarshalIndent(data, "", " ")
	check(err)
	ioutil.WriteFile("data/savegame.json", f, 0644)
}

// Load loads the game
func Load() {
	jsonFile, err := os.Open("data/savegame.json")
	check(err)
	fmt.Println("Loaded the file.")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	check(err)
	var data Savedata
	json.Unmarshal(byteValue, &data)

	GS.paused = false
	switch loadedGamestate := data.Name; loadedGamestate {
	case "Flying":
		GS.current = flying
		UI76.currentPanel = MenuFlying
	default:
		GS.current = landed
		UI76.currentPanel = MenuLanded
	}
}
