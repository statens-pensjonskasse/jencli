package common

import (
	"strings"
	"testing"
)

func TestNormaliseBranchName(t *testing.T) {
	tests := []struct {
		name   string // name of the test case
		branch string // input for NormaliseBranchName
		want   string // expected result
	}{
		{
			name:   "Regular branch name",
			branch: "feature/SPK-123_abc",
			want:   "feature-spk-123_abc",
		},
		{
			name:   "Special Characters",
			branch: "æøå",
			want:   "aeoeaa",
		},
		{
			name:   "Spaces To Hypens",
			branch: "Dev Branch",
			want:   "dev-branch",
		},
		{
			name:   "Multiple Consequent Dots Replaced By One Hyphen",
			branch: "devel..Version",
			want:   "devel-version",
		},
		{
			name:   "Length Greater Than 128",
			branch: strings.Repeat("a", 200),
			want:   strings.Repeat("a", 128),
		},
		{
			name:   "Empty String In Input",
			branch: "",
			want:   "EMPTY",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormaliseBranchName(tt.branch)

			if got != tt.want {
				t.Errorf("NormaliseBranchName(%v) = %v; want %v", tt.branch, got, tt.want)
			}
		})
	}
}
