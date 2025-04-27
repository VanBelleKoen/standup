package main

import (
	"testing"
)

func TestRemoveEntry(t *testing.T) {
	initialData := StandupData{
		Today:       []string{"Task 1", "Task 2", "Task 3"},
		LastWorkday: []string{},
		Notes:       []string{},
	}
	err := saveStandupData(initialData)
	if err != nil {
		t.Fatalf("Failed to save initial data: %v", err)
	}

	removeEntry("Task 2")
	data, err := loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(data.Today))
	}

	for _, entry := range data.Today {
		if entry == "Task 2" {
			t.Errorf("Entry 'Task 2' was not removed")
		}
	}

	// Test removing a non-existent entry
	removeEntry("Non-existent Task")
	data, err = loadStandupData()
	if err != nil {
		t.Fatalf("Failed to load data: %v", err)
	}

	if len(data.Today) != 2 {
		t.Errorf("Expected 2 entries after attempting to remove non-existent entry, got %d", len(data.Today))
	}
}