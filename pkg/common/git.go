package common

import (
	"github.com/go-git/go-git/v5"
)

func GetCurrentBranch(dir string) (string, error) {
	gitRepo, err := git.PlainOpen(dir)
	if err != nil {
		return "", err
	}
	head, err := gitRepo.Head()
	if err != nil {
		return "", err
	}
	return head.Name().Short(), nil
}
