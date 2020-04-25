package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Savedata is a struct for characters
type Savedata struct {
	Gamestate State
	Name      string
}

// Save saves the game
func Save() {
	data := Savedata{
		Name:      "test",
		Gamestate: GS.current,
	}

	f, err := json.MarshalIndent(data, "", " ")
	check(err)
	ioutil.WriteFile("data/savegame.json", f, 0644)
	/*
		f, err := os.Create("data/savegame.json")
		check(err)
		defer f.Close()
	*/

}

// Load loads the game
func Load() {
	//f, err := os.Open("data/savegame.json")
	//check(err)
	fmt.Println("Loading. Not really tho.")
}
