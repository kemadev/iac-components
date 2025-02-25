package githubrepo

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WrapperArgs struct {
	provider   ProviderArgs
	actions    ActionsArgs
	branches   BranchesArgs
	envs       EnvsArgs
	rulesets   RulesetsArgs
	repository RepositoryArgs
}

func Wrapper(ctx *pulumi.Context, args WrapperArgs) error {
	provider, err := createProvider(ctx, args.provider)
	if err != nil {
		return err
	}
	repo, err := createRepo(ctx, provider, args.repository)
	if err != nil {
		return err
	}
	err = createBranches(ctx, provider, repo, args.branches)
	if err != nil {
		return err
	}
	envs, err := createEnvironments(ctx, provider, repo, args.envs, args.branches)
	if err != nil {
		return err
	}
	err = createRulesets(ctx, provider, repo, envs, args.rulesets, args.branches)
	if err != nil {
		return err
	}
	err = createActions(ctx, provider, repo, args.actions)
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
