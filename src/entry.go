package main

import "fmt"

func addEntry(entry string) {
	data, err := loadStandupData()
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	data.Today = append(data.Today, entry)
	err = saveStandupData(data)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	fmt.Printf("Added entry: %s\n", entry)
}

func removeEntry(entry string) {
	data, err := loadStandupData()
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	newToday := []string{}
	found := false
	for _, e := range data.Today {
		if e == entry {
			found = true
		} else {
			newToday = append(newToday, e)
		}
	}

	if !found {
		fmt.Printf("Entry not found: %s\n", entry)
		return
	}

	data.Today = newToday
	err = saveStandupData(data)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	fmt.Printf("Removed entry: %s\n", entry)
}