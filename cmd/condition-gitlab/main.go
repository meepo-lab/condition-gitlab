package main

import (
	gitlabCondition "github.com/ted-vo/condition-gitlab/pkg/condition"
	"github.com/ted-vo/semantic-release/v3/pkg/condition"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		CICondition: func() condition.CICondition {
			return &gitlabCondition.GitLab{}
		},
	})
}
