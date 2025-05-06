package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type StandupData struct {
	Today           []string `json:"today"`
	LastWorkday     []string `json:"last_workday"`
	Notes           []string `json:"notes"`
	LastWorkdayDate string   `json:"last_workday_date"`
}

var dataFile = "standup.json"

func main() {
	initializeStandupData()

	if err := updateWorkdayData(); err != nil {
		println("Error updating workday data:", err)
		return
	}

	if len(os.Args) < 2 {
		showStandup()
		return
	}

	command := os.Args[1]
	switch command {
	case "--help":
		handleHelp()
	case "--remove":
		handleRemove(os.Args[2:])
	case "--reset":
		handleReset()
	case "--note":
		handleNote(os.Args[2:])
	case "--sync-branches":
		handleSyncBranches()
	default:
		handleDefault(os.Args[1:])
	}
}

func loadStandupData() (*StandupData, error) {
	var data StandupData
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &data, nil
		}
		return nil, err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func saveStandupData(data *StandupData) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, file, 0644)
}

func resetTodayList() {
	data, err := loadStandupData()
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	data.Today = []string{}
	err = saveStandupData(data)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	fmt.Println("Today list has been reset.")
}