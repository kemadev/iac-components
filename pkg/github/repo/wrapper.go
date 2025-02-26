package repo

import (
	p "github.com/kemadev/iac-components/pkg/github/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WrapperArgs struct {
	Provider   p.ProviderArgs
	Branches   BranchesArgs
	Envs       EnvsArgs
	Rulesets   RulesetsArgs
	Repository RepositoryArgs
}

func setDefaultArgs(args *WrapperArgs) {
	p.SetDefaults(&args.Provider)
	createBranchesSetDefaults(&args.Branches)
	createEnvironmentsSetDefaults(&args.Envs)
	createRulesetsSetDefaults(&args.Rulesets)
	createRepositorySetDefaults(&args.Repository)
}

func Wrapper(ctx *pulumi.Context, args WrapperArgs) error {
	setDefaultArgs(&args)
	provider, err := p.NewProvider(ctx, args.Provider)
	if err != nil {
		return err
	}
	repo, err := createRepo(ctx, provider, args.Repository, args.Branches)
	if err != nil {
		return err
	}
	err = createBranches(ctx, provider, repo, args.Branches)
	if err != nil {
		return err
	}
	_, err = createEnvironments(ctx, provider, repo, args.Envs, args.Branches)
	if err != nil {
		return err
	}
	err = createRulesets(ctx, provider, repo, args.Rulesets, args.Branches)
	if err != nil {
		return err
	}
	err = createDependabot(ctx, provider, repo)
	if err != nil {
		return err
	}
	err = createIssues(ctx, provider, repo)
	if err != nil {
		return err
	}
	return nil
}
