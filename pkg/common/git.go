package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findGitRoot(dir string) (string, error) {
	root := dir
	if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
		// If no error the current directory contains a .git folder
		return root, err
	}
	if root == "/" || root == "." {
		// Reached the root directory
		return "", errors.New("no git repository found")
	}
	// Move to parent directory
	root = filepath.Dir(root)
	return findGitRoot(root)
}

func GetCurrentBranch(dir string) (string, error) {
	gitRoot, err := findGitRoot(dir)
	if err != nil {
		return "", err
	}

	headFile := filepath.Join(gitRoot, ".git", "HEAD")
	file, err := os.ReadFile(headFile)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.TrimPrefix(string(file), "ref: refs/heads/")), nil
}

func GetProject(dir string) (string, error) {
	gitRoot, err := findGitRoot(dir)
	if err != nil {
		return "", err
	}

	configFile := filepath.Join(gitRoot, ".git", "config")
	project, err := processFile(configFile)
	return project, err
}

func extractPartFromURL(line string) (string, error) {
	// Find the part after "scm/"
	prefix := "scm/"
	startIndex := strings.Index(line, prefix)
	if startIndex == -1 {
		prefix = "7999/"
		startIndex = strings.Index(line, prefix)
	}
	if startIndex == -1 {
		return "", errors.New("did not find scm/ or 7999/ in url")
	}
	startIndex += len(prefix) // Move to the character after prefix

	// Find the next "/"
	endIndex := strings.Index(line[startIndex:], "/")
	if endIndex == -1 {
		return "", errors.New("not find ending /")
	}

	// Extract the part between prefix and the next "/"
	return line[startIndex : startIndex+endIndex], nil
}

func processFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains "url"
		if strings.Contains(line, "url") {
			extracted, err := extractPartFromURL(line)
			if extracted != "" {
				return extracted, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return "", err
}
