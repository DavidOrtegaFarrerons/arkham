package config

type Config struct {
	BranchPattern  string `json:"branch_pattern"`
	CommitTemplate string `json:"commit_template"`
}
