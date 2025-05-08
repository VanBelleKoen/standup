package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
)

type StandupData struct {
	Today           []string `json:"today"`
	LastWorkday     []string `json:"last_workday"`
	Notes           []string `json:"notes"`
	LastWorkdayDate string   `json:"last_workday_date"`
}

func getHomeDirectory() string {
	usr, err := user.Current()
	if err != nil {
		panic("Unable to determine the user's home directory")
	}
	return usr.HomeDir
}

var dataFile = getHomeDirectory() + "/standup.json"

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