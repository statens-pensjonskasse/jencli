package common

import (
	"regexp"
	"strings"
)

func NormaliseBranchName(branch string) string {
	// Convert to lowercase
	normalisedBranchName := strings.ToLower(branch)

	normalisedBranchName = strings.ReplaceAll(normalisedBranchName, "æ", "ae")
	normalisedBranchName = strings.ReplaceAll(normalisedBranchName, "ø", "oe")
	normalisedBranchName = strings.ReplaceAll(normalisedBranchName, "å", "aa")
	normalisedBranchName = regexp.MustCompile(`[\s\\/,:;|]`).ReplaceAllString(normalisedBranchName, "-")
	normalisedBranchName = regexp.MustCompile(`[\[\](){}]`).ReplaceAllString(normalisedBranchName, "")
	normalisedBranchName = regexp.MustCompile(`[^a-z0-9._-]`).ReplaceAllString(normalisedBranchName, "x")
	normalisedBranchName = regexp.MustCompile(`[._-]+[._]+`).ReplaceAllString(normalisedBranchName, "-")
	normalisedBranchName = regexp.MustCompile(`[._]+[._-]+`).ReplaceAllString(normalisedBranchName, "-")

	if len(normalisedBranchName) > 128 {
		normalisedBranchName = (normalisedBranchName)[:128]
	}

	for strings.HasPrefix(normalisedBranchName, "_") || strings.HasPrefix(normalisedBranchName, "-") || strings.HasPrefix(normalisedBranchName, ".") {
		normalisedBranchName = (normalisedBranchName)[1:]
	}

	for strings.HasSuffix(normalisedBranchName, "_") || strings.HasSuffix(normalisedBranchName, "-") || strings.HasSuffix(normalisedBranchName, ".") {
		normalisedBranchName = (normalisedBranchName)[:len(normalisedBranchName)-1]
	}

	if len(normalisedBranchName) == 0 {
		normalisedBranchName = "EMPTY"
	}

	return normalisedBranchName
}
