package main

import (
	"os"
	"testing"
	"time"
)

func TestUpdateWorkdayData(t *testing.T) {
	os.Remove(dataFile)

	entry := "Test task for today"
	addEntry(entry)

	timeInPast := time.Now().Add(-24 * time.Hour)
	os.Chtimes(dataFile, timeInPast, timeInPast)

	err := updateWorkdayData()
	if err != nil {
		t.Fatalf("updateWorkdayData() returned an error: %v", err)
	}

	data, err := loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 0 {
		t.Errorf("Expected Today to be empty, got: %v", data.Today)
	}

	if len(data.LastWorkday) != 1 || data.LastWorkday[0] != entry {
		t.Errorf("Expected entry '%s' in LastWorkday, got: %v", entry, data.LastWorkday)
	}
}

func TestDateTransition(t *testing.T) {
	os.Remove(dataFile)

	entry := "Test task for today"
	addEntry(entry)

	data, _ := loadStandupData()
	data.LastWorkday = data.Today
	data.Today = []string{}
	saveStandupData(data)

	err := updateWorkdayData()
	if err != nil {
		t.Fatalf("Failed to update workday data: %v", err)
	}

	data, err = loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 0 {
		t.Errorf("Expected Today to be empty after date transition, got: %v", data.Today)
	}

	if len(data.LastWorkday) != 1 || data.LastWorkday[0] != entry {
		t.Errorf("Expected entry '%s' in LastWorkday after date transition, got: %v", entry, data.LastWorkday)
	}
}

func TestThirdDayTransition(t *testing.T) {
	os.Remove(dataFile)

	entry := "Test task for today"
	addEntry(entry)

	data, _ := loadStandupData()
	data.LastWorkday = data.Today
	data.Today = []string{}
	saveStandupData(data)

	err := updateWorkdayData()
	if err != nil {
		t.Fatalf("Failed to update workday data: %v", err)
	}

	data, err = loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 0 {
		t.Errorf("Expected Today to be empty on the third day, got: %v", data.Today)
	}

	if len(data.LastWorkday) != 1 || data.LastWorkday[0] != entry {
		t.Errorf("Expected entry '%s' in LastWorkday on the third day, got: %v", entry, data.LastWorkday)
	}
}