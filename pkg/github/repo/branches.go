package githubrepo

import (
	"github.com/kemadev/iac-components/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BranchesArgs struct {
	Dev     string
	Next    string
	Prod    string
	Default string
}

func createBranches(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, args BranchesArgs) error {
	branchDevName := util.FormatResourceName("Branch dev")
	_, err := github.NewBranch(ctx, branchDevName, &github.BranchArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(args.Dev),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	branchNextName := util.FormatResourceName("Branch next")
	_, err = github.NewBranch(ctx, branchNextName, &github.BranchArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(args.Next),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	branchProdName := util.FormatResourceName("Branch prod")
	_, err = github.NewBranch(ctx, branchProdName, &github.BranchArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(args.Prod),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	if args.Default != args.Prod && args.Default != args.Next && args.Default != args.Dev {
		branchDefaultName := util.FormatResourceName("Branch default")
		_, err = github.NewBranchDefault(ctx, branchDefaultName, &github.BranchDefaultArgs{
			Repository: repo.Name,
			Branch:     pulumi.String(args.Default),
			Rename:     pulumi.Bool(false),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
	}

	return nil
}
