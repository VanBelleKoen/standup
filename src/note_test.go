package main

import (
	"os"
	"testing"
)

func TestAddNote(t *testing.T) {
	dataFile = "test_standup.json"
	defer func() {
		_ = os.Remove(dataFile)
	}()

	err := saveStandupData(&StandupData{})
	if err != nil {
		t.Fatalf("Failed to initialize test data file: %v", err)
	}

	note := "Test note"
	addNote(note)

	data, err := loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Notes) != 1 || data.Notes[0] != note {
		t.Errorf("Expected note '%s' to be added, but got: %v", note, data.Notes)
	}
}