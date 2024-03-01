package types

type JPLDeploy struct {
	Swarm         string
	Image         string
	BranchPostfix bool
	Branch        string
	Tag           string
	Environment   string
	StackConfig   string
	Slack         string
}
