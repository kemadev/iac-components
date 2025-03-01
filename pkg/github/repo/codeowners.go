package repo

import (
	"bytes"
	"fmt"

	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CodeownerParam struct {
	Path   string
	Entity string
}

type CodeownersArgs struct {
	Codeowners []CodeownerParam
}

var CodeownersDefaultArgs = CodeownersArgs{
	Codeowners: []CodeownerParam{
		{
			Path:   "*",
			Entity: "@kemadev/maintainers", // TODO make this a reference to org IaC
		},
	},
}

var CodeownersDefaultContent = `# Read more about CODEOWNERS [here](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)

# This files is managed by repository-as-code! Do not edit manually!

`

func createCodeownersSetDefaults(args *CodeownersArgs) {
	if len(args.Codeowners) == 0 {
		args.Codeowners = CodeownersDefaultArgs.Codeowners
	}
}

func createCodeownersContent(args *CodeownersArgs) string {
	var codeownersContent bytes.Buffer
	codeownersContent.WriteString(CodeownersDefaultContent)
	for _, codeowner := range args.Codeowners {
		codeownersContent.WriteString(fmt.Sprintf("%s %s\n", codeowner.Path, codeowner.Entity))
	}
	return codeownersContent.String()
}

func createCodeowners(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, args CodeownersArgs) error {
	codeownersGithubFileName := util.FormatResourceName(ctx, "Repository codeowners file")
	_, err := github.NewRepositoryFile(ctx, codeownersGithubFileName, &github.RepositoryFileArgs{
		Repository:        repo.Name,
		File:              pulumi.String(".github/CODEOWNERS"),
		Content:           pulumi.String(createCodeownersContent(&args)),
		CommitMessage:     pulumi.String("feat(codeowners): update CODEOWNERS file"),
		CommitAuthor:      pulumi.String("pulumi[bot]"), // TODO make this a reference to global settings
		CommitEmail:       pulumi.String("kemadev+pulumi[bot]@users.noreply.github.com"), // TODO make this a reference to global settings
		OverwriteOnCreate: pulumi.Bool(true),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
