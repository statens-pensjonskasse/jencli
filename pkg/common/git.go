package common

import (
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"strings"
)

func GetRemoteUrl(dir string) (string, error) {
	gitConfig := filepath.Join(dir, ".git", "config")

	config, err := ini.Load(gitConfig)
	if err != nil {
		return "", err
	}

	origin, err := config.GetSection("remote \"origin\"")
	if err != nil {
		return "", err
	}
	remoteUrl, err := origin.GetKey("url")
	if err != nil {
		return "", err
	}
	return remoteUrl.String(), nil
}

func GetProjectName(dir string) (string, error) {
	remoteUrl, err := GetRemoteUrl(dir)
	if err != nil {
		return "", err
	}
	slice := strings.Split(remoteUrl, "/")
	return slice[len(slice)-2], nil

}

func GetRepoName(dir string) (string, error) {
	remoteUrl, err := GetRemoteUrl(dir)
	if err != nil {
		return "", err
	}
	slice := strings.Split(remoteUrl, "/")
	return strings.TrimSuffix(slice[len(slice)-1], ".git"), nil
}

func GetCurrentBranch(dir string) (string, error) {
	headFile := filepath.Join(dir, ".git", "HEAD")
	file, err := os.ReadFile(headFile)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.TrimPrefix(string(file), "ref: refs/heads/")), nil
}
