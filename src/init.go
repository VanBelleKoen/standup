package main

import (
	"encoding/json"
	"os"
	"time"
)

func initializeStandupData() {
	if _, err := os.Stat("standup.json"); os.IsNotExist(err) {
		defaultData := struct {
			LastWorkdayDate string   `json:"last_workday_date"`
			Today           []string `json:"today"`
			LastWorkday     []string `json:"last_workday"`
			Notes           []string `json:"notes"`
		}{
			LastWorkdayDate: time.Now().AddDate(0, 0, -1).Format("02/01/2006"),
			Today:           []string{},
			LastWorkday:     []string{},
			Notes:           []string{},
		}

		file, err := os.Create("standup.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(defaultData); err != nil {
			panic(err)
		}
		return
	}

	file, err := os.Open("standup.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data struct {
		LastWorkdayDate string `json:"last_workday_date"`
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	if data.LastWorkdayDate == "" {
		data.LastWorkdayDate = time.Now().AddDate(0, 0, -1).Format("02/01/2006")
		file, err := os.Create("standup.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		if err := encoder.Encode(data); err != nil {
			panic(err)
		}
	}
}