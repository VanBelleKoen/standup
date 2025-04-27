package main

import (
	"io"
	"os"
	"strings"
	"syscall"
	"testing"
)

func TestShowStandup(t *testing.T) {
	os.Remove(dataFile)

	data := StandupData{
		Today:       []string{"Task 1", "Task 2"},
		LastWorkday: []string{"Task A", "Task B"},
	}
	err := saveStandupData(data)
	if err != nil {
		t.Fatalf("Failed to save data: %v", err)
	}

	r, w, _ := os.Pipe()
	os.Stdout = w

	showStandup()

	w.Close()
	os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")

	output, _ := io.ReadAll(r)

	if !strings.Contains(string(output), "Task 1") || !strings.Contains(string(output), "Task 2") {
		t.Errorf("Output does not contain today's tasks: %s", output)
	}

	if !strings.Contains(string(output), "Task A") || !strings.Contains(string(output), "Task B") {
		t.Errorf("Output does not contain last workday's tasks: %s", output)
	}
}