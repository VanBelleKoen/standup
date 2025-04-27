package main

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Remove(dataFile)

	os.Exit(exitCode)
}

func TestAddEntry(t *testing.T) {
	os.Remove(dataFile)

	entry := "Test task"
	addEntry(entry)

	data, err := loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 1 || data.Today[0] != entry {
		t.Errorf("Expected entry '%s' in Today, got: %v", entry, data.Today)
	}
}

func TestEmptyData(t *testing.T) {
	os.Remove(dataFile)

	data, err := loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 0 || len(data.LastWorkday) != 0 {
		t.Errorf("Expected empty data, got: Today=%v, LastWorkday=%v", data.Today, data.LastWorkday)
	}
	showStandup()
}

func TestSkipWorkday(t *testing.T) {
	os.Remove(dataFile)

	entry := "Monday task"
	addEntry(entry)

	data, err := loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 1 || data.Today[0] != entry {
		t.Errorf("Expected entry '%s' in Today, got: %v", entry, data.Today)
	}

	monday := time.Now().AddDate(0, 0, -2)
	err = os.Chtimes(dataFile, monday, monday)
	if err != nil {
		t.Fatalf("Failed to change file times: %v", err)
	}

	err = updateWorkdayData()
	if err != nil {
		t.Fatalf("updateWorkdayData() returned an error: %v", err)
	}

	data, err = loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 0 {
		t.Errorf("Expected Today to be empty, got: %v", data.Today)
	}

	if len(data.LastWorkday) != 1 || data.LastWorkday[0] != entry {
		t.Errorf("Expected entry '%s' in Last Workday, got: %v", entry, data.LastWorkday)
	}
}