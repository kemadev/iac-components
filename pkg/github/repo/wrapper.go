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
	Codeowners CodeownersArgs
	Files      FilesArgs
}

func setDefaultArgs(args *WrapperArgs) error {
	p.SetDefaults(&args.Provider)
	createBranchesSetDefaults(&args.Branches)
	createEnvironmentsSetDefaults(&args.Envs)
	createRulesetsSetDefaults(&args.Rulesets)
	err := createRepositorySetDefaults(&args.Repository)
	if err != nil {
		return err
	}
	createCodeownersSetDefaults(&args.Codeowners)
	createFilesSetDefaults(&args.Files)
	return nil
}

func Wrapper(ctx *pulumi.Context, args WrapperArgs) error {
	err := setDefaultArgs(&args)
	if err != nil {
		return err
	}
	provider, err := p.NewProvider(ctx, args.Provider)
	if err != nil {
		return err
	}
	repo, err := createRepo(ctx, provider, args.Repository, args.Branches)
	if err != nil {
		return err
	}
	branches, err := createBranches(ctx, provider, repo, args.Branches)
	if err != nil {
		return err
	}
	_, err = createEnvironments(ctx, provider, repo, args.Envs, args.Branches)
	if err != nil {
		return err
	}
	err = createRulesets(ctx, provider, repo, args.Rulesets, args.Branches, args.Envs)
	if err != nil {
		return err
	}
	err = createDependabot(ctx, provider, repo)
	if err != nil {
		return err
	}
	err = createCodeowners(ctx, provider, repo, args.Codeowners, branches.Next)
	if err != nil {
		return err
	}
	err = templateRepo.init(args.Files.UpstreamRepo)
	if err != nil {
		return err
	}
	err = createFiles(ctx, provider, repo, args.Files, branches.Next)
	if err != nil {
		return err
	}
	err = createTemplatedFiles(ctx, provider, repo, args.Files, args.Repository, branches.Next)
	if err != nil {
		return err
	}
	err = createIssues(ctx, provider, repo)
	if err != nil {
		return err
	}
	return nil
}
