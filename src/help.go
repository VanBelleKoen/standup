package main

import "fmt"

func showHelp() {
	fmt.Println("Usage: standup [command] [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  --help          Show this help message")
	fmt.Println("  --remove [entry] Remove an entry from today's list")
	fmt.Println("  --reset         Reset today's list")
	fmt.Println("  [entries...]    Add one or more entries to today's list")
	fmt.Println("Description:")
	fmt.Println("  The standup application helps you manage your daily standup notes.")
	fmt.Println("  Entries added are categorized into 'Today' and 'Last Workday'.")
}