package common

import (
	"errors"
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
