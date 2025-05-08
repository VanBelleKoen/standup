package main

import (
	"os"
)

func main() {
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