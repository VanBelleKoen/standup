package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func handleHelp() {
	showHelp()
}

func handleRemove(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing entry to remove.")
		return
	}
	entry := args[0]
	removeEntry(entry)
}

func handleReset() {
	resetTodayList()
}

func handleNote(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Missing note content.")
		return
	}
	note := args[0]
	addNote(note)
}

func handleSyncBranches() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	for _, workspace := range config.Workspaces {
		if strings.HasPrefix(workspace, "~") {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Printf("Error getting home directory: %v\n", err)
				continue
			}
			workspace = filepath.Join(homeDir, workspace[1:])
		}

		absWorkspace, err := filepath.Abs(workspace)
		if err != nil {
			fmt.Printf("Error resolving absolute path for workspace %s: %v\n", workspace, err)
			continue
		}

		if _, err := os.Stat(absWorkspace); os.IsNotExist(err) {
			fmt.Printf("Workspace does not exist: %s\n", absWorkspace)
			continue
		}

		repos, err := findGitRepositories(absWorkspace)
		if err != nil {
			fmt.Printf("Error scanning workspace %s: %v\n", absWorkspace, err)
			continue
		}

		for _, repo := range repos {
			newBranches, err := detectNewBranches(repo)
			if err != nil {
				fmt.Printf("Error detecting branches in repo %s: %v\n", repo, err)
				continue
			}

			if len(newBranches) > 0 {
				logNewBranches(newBranches, repo, absWorkspace)
			} else {
				fmt.Printf("No new branches detected in repo %s.\n", repo)
			}
		}
	}
}

func handleDefault(entries []string) {
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