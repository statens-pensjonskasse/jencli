package types

import "strconv"

type JPLDeploy struct {
	Swarm            string `yaml:"swarm"`
	Image            string `yaml:"image"`
	Branch           string `yaml:"branch"`
	UseBranchPostfix bool   `yaml:"useBranchPostfix"`
	Tag              string `yaml:"tag"`
	UseBranchTag     bool   `yaml:"useBranchTag"`
	FullImageName    string `yaml:"fullImageName"`
	Environment      string `yaml:"environment"`
	StackConfig      string `yaml:"stackConfig"`
	Slack            string `yaml:"slack"`
}

func (p JPLDeploy) ToParamMap() map[string]string {
	return map[string]string{
		"swarm":         p.Swarm,
		"image":         p.Image,
		"branch":        p.Branch,
		"branchPostfix": strconv.FormatBool(p.UseBranchPostfix),
		"tag":           p.Tag,
		"environment":   p.Environment,
		"stackConfig":   p.StackConfig,
		"slack":         p.Slack,
	}

}
