package repo

import (
	p "github.com/kemadev/iac-components/pkg/github/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WrapperArgs struct {
	Provider   p.ProviderArgs
	Envs       EnvsArgs
	Rulesets   RulesetsArgs
	Repository RepositoryArgs
	Codeowners CodeownersArgs
	Files      FilesArgs
}

func setDefaultArgs(args *WrapperArgs) error {
	p.SetDefaults(&args.Provider)
	createEnvironmentsSetDefaults(&args.Envs)
	createRulesetsSetDefaults(&args.Rulesets)
	err := createRepositorySetDefaults(&args.Repository)
	if err != nil {
		return err
	}
	err = createCodeownersSetDefaults(&args.Codeowners)
	if err != nil {
		return err
	}
	createFilesSetDefaults(&args.Files)
	return nil
}

func Wrapper(ctx *pulumi.Context, args WrapperArgs) error {
	targetBranch := "repo-as-code-update"
	err := setDefaultArgs(&args)
	if err != nil {
		return err
	}
	provider, err := p.NewProvider(ctx, args.Provider)
	if err != nil {
		return err
	}
	repo, err := createRepo(ctx, provider, args.Repository)
	if err != nil {
		return err
	}
	envs, err := createEnvironments(ctx, provider, repo, args.Envs)
	if err != nil {
		return err
	}
	err = createRulesets(ctx, provider, repo, args.Rulesets, envs.prod.Environment.ElementType().Name())
	if err != nil {
		return err
	}
	err = createDependabot(ctx, provider, repo)
	if err != nil {
		return err
	}
	err = createCodeowners(ctx, provider, repo, args.Codeowners, targetBranch)
	if err != nil {
		return err
	}
	err = templateRepo.init(args.Files.UpstreamRepo)
	if err != nil {
		return err
	}
	err = createFiles(ctx, provider, repo, args.Files, targetBranch, repo.DefaultBranch.ElementType().Name())
	if err != nil {
		return err
	}
	err = createTemplatedFiles(ctx, provider, repo, args.Files, args.Repository, targetBranch, repo.DefaultBranch.ElementType().Name())
	if err != nil {
		return err
	}
	err = createIssues(ctx, provider, repo)
	if err != nil {
		return err
	}
	return nil
}
