package storage

import (
	"encoding/json"
	"os"

	"linkedin-automation-demo/models"
)

type State struct {
	ConnectionsSent int
	Profiles        []models.Profile
}

func SaveState(state State) error {
	file, err := os.Create("data/state.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(state)
}
