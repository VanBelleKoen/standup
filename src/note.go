package main

import "fmt"

func addNote(note string) {
	data, err := loadStandupData()
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	if data.Notes == nil {
		data.Notes = []string{}
	}

	data.Notes = append(data.Notes, note)
	err = saveStandupData(data)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	fmt.Printf("Added note: %s\n", note)
}