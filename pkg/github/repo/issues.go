package repo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createIssues(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository) error {
	issueLabelsName := util.FormatResourceName(ctx, "Issue labels")
	_, err := github.NewIssueLabels(ctx, issueLabelsName, &github.IssueLabelsArgs{
		Repository: repo.Name,
		Labels:     github.IssueLabelsLabelArray{},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
