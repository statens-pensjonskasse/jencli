package common

import (
	"testing"
)

func TestExtractPartFromURL(t *testing.T) {
	tests := []struct {
		name      string // name of the test case
		directory string // input for NormaliseBranchName
		want      string // expected result
	}{
		{
			name:      "Get PU_OPT from url string",
			directory: "url = ssh://git@git.spk.no:7999/scm/PU_OPT/opptjening-innsyn-webservice.git",
			want:      "PU_OPT",
		},
		{
			name:      "Get dev project from url string",
			directory: "url = ssh://git@git.spk.no:7999/dev/jencli.git",
			want:      "dev",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := extractPartFromURL(tt.directory)

			if got != tt.want {
				t.Errorf("ExtractPartFromURL(%v) = %v; want %v", tt.directory, got, tt.want)
			}
		})
	}
}
