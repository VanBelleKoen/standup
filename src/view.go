package main

import (
	"fmt"

	"github.com/fatih/color"
)

func showStandup() {
	data, err := loadStandupData()
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Println(green("Standup notes:"))
	for _, note := range data.Notes {
		fmt.Printf("    - %s\n", note)
	}

	fmt.Println(blue("Today:"))
	for _, entry := range data.Today {
		fmt.Printf("    - %s\n", entry)
	}

	fmt.Println(red("Last workday:"))
	for _, entry := range data.LastWorkday {
		fmt.Printf("    - %s\n", entry)
	}
}