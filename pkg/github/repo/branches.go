package repo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BranchesArgs struct {
	Dev     string
	Next    string
	Prod    string
	Default string
}

var BranchesDefaultArgs = BranchesArgs{
	Dev:     "dev",
	Next:    "next",
	Prod:    "main",
	Default: "main",
}

func createBranchesSetDefaults(args *BranchesArgs) {
	if args.Dev == "" {
		args.Dev = BranchesDefaultArgs.Dev
	}
	if args.Next == "" {
		args.Next = BranchesDefaultArgs.Next
	}
	if args.Prod == "" {
		args.Prod = BranchesDefaultArgs.Prod
	}
	if args.Default == "" {
		args.Default = BranchesDefaultArgs.Default
	}
}

type TBranchesCreated struct {
	Dev  *github.Branch
	Next *github.Branch
	Prod *github.Branch
}

func createBranches(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, args BranchesArgs) (TBranchesCreated, error) {
	branchDevName := util.FormatResourceName(ctx, "Branch dev")
	dev, err := github.NewBranch(ctx, branchDevName, &github.BranchArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(args.Dev),
	}, pulumi.Provider(provider))
	if err != nil {
		return TBranchesCreated{}, err
	}

	branchNextName := util.FormatResourceName(ctx, "Branch next")
	next, err := github.NewBranch(ctx, branchNextName, &github.BranchArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(args.Next),
	}, pulumi.Provider(provider))
	if err != nil {
		return TBranchesCreated{}, err
	}

	branchProdName := util.FormatResourceName(ctx, "Branch prod")
	prod, err := github.NewBranch(ctx, branchProdName, &github.BranchArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(args.Prod),
	}, pulumi.Provider(provider))
	if err != nil {
		return TBranchesCreated{}, err
	}

	if args.Default != args.Prod && args.Default != args.Next && args.Default != args.Dev {
		branchDefaultName := util.FormatResourceName(ctx, "Branch default")
		_, err := github.NewBranchDefault(ctx, branchDefaultName, &github.BranchDefaultArgs{
			Repository: repo.Name,
			Branch:     pulumi.String(args.Default),
			Rename:     pulumi.Bool(false),
		}, pulumi.Provider(provider))
		if err != nil {
			return TBranchesCreated{}, err
		}
	}

	return TBranchesCreated{
		Dev:  dev,
		Next: next,
		Prod: prod,
	}, nil
}
