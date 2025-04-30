package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func findGitRepositories(root string) ([]string, error) {
	var repos []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}
		if info.IsDir() && strings.HasSuffix(path, ".git") {
			repos = append(repos, filepath.Dir(path))
			fmt.Printf("Found repository: %s\n", filepath.Dir(path))
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the file tree: %v\n", err)
	}

	if len(repos) == 0 {
		fmt.Println("No repositories found.")
	}

	return repos, err
}

func detectNewBranches(repo string) ([]string, error) {
	cmd := exec.Command("git", "for-each-ref", "--sort=-creatordate", "--format=%(refname:short) %(creatordate:iso8601)", "refs/heads/")
	cmd.Dir = repo
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running git command: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	newBranches := []string{}

	data, err := loadStandupData()
	if err != nil {
		return nil, fmt.Errorf("error loading standup data: %w", err)
	}
	lastWorkday, err := time.Parse("02/01/2006", data.LastWorkdayDate)
	if err != nil {
		return nil, fmt.Errorf("error parsing last workday date: %w", err)
	}

	lastWorkday = lastWorkday.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}
		branchName := parts[0]
		creationDate, err := time.Parse("2006-01-02 15:04:05 -0700", parts[1])
		if err != nil {
			fmt.Printf("Error parsing date for branch %s: %v\n", branchName, err)
			continue
		}
		if creationDate.After(lastWorkday) {
			newBranches = append(newBranches, branchName)
		}
	}

	return newBranches, nil
}

func logNewBranches(newBranches []string, repo string, workspace string) {
	data, err := loadStandupData()
	if err != nil {
		fmt.Println("Error loading standup data:", err)
		return
	}

	repoRelativePath, err := filepath.Rel(workspace, repo)
	if err != nil {
		fmt.Printf("Error determining relative path for repo %s: %v\n", repo, err)
		return
	}
	mainDirectory := strings.Split(repoRelativePath, string(os.PathSeparator))[0]

	for _, branch := range newBranches {
		entry := fmt.Sprintf("%s - %s", mainDirectory, branch)
		data.LastWorkday = append(data.LastWorkday, entry)
		fmt.Println("Logged entry to last_workday:", entry)
	}

	err = saveStandupData(data)
	if err != nil {
		fmt.Println("Error saving standup data:", err)
	}
}