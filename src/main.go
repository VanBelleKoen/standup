package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type StandupData struct {
	Today       []string `json:"today"`
	LastWorkday []string `json:"last_workday"`
	Notes       []string `json:"notes"`
}

var dataFile = "standup.json"

func main() {
	if err := updateWorkdayData(); err != nil {
		fmt.Println("Error updating workday data:", err)
		return
	}

	if len(os.Args) < 2 {
		showStandup()
		return
	}

	command := os.Args[1]
	switch command {
	case "--help":
		showHelp()
	case "--remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing entry to remove.")
			return
		}
		entry := os.Args[2]
		removeEntry(entry)
	case "--reset":
		resetTodayList()
	case "--note":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing note content.")
			return
		}
		note := os.Args[2]
		addNote(note)
	default:
		entries := os.Args[1:]
		data, err := loadStandupData()
		if err != nil {
			fmt.Println("Error loading data:", err)
			return
		}

		existingEntries := make(map[string]bool)
		for _, entry := range data.Today {
			existingEntries[entry] = true
		}

		for _, entry := range entries {
			if existingEntries[entry] {
				fmt.Printf("Duplicate entry ignored: %s\n", entry)
			} else {
				data.Today = append(data.Today, entry)
				existingEntries[entry] = true
				fmt.Printf("Added entry: %s\n", entry)
			}
		}

		err = saveStandupData(data)
		if err != nil {
			fmt.Println("Error saving data:", err)
		}
	}
}

func loadStandupData() (StandupData, error) {
	var data StandupData
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return data, err
	}
	err = json.Unmarshal(file, &data)
	return data, err
}

func saveStandupData(data StandupData) error {
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