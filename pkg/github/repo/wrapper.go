package repo

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WrapperArgs struct {
	Provider   ProviderArgs
	Actions    ActionsArgs
	Branches   BranchesArgs
	Envs       EnvsArgs
	Rulesets   RulesetsArgs
	Repository RepositoryArgs
}

func Wrapper(ctx *pulumi.Context, args WrapperArgs) error {
	provider, err := createProvider(ctx, args.Provider)
	if err != nil {
		return err
	}
	repo, err := createRepo(ctx, provider, args.Repository)
	if err != nil {
		return err
	}
	err = createBranches(ctx, provider, repo, args.Branches)
	if err != nil {
		return err
	}
	envs, err := createEnvironments(ctx, provider, repo, args.Envs, args.Branches)
	if err != nil {
		return err
	}
	err = createRulesets(ctx, provider, repo, envs, args.Rulesets, args.Branches)
	if err != nil {
		return err
	}
	err = createActions(ctx, provider, repo, args.Actions)
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
