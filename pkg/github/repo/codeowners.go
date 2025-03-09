package repo

import (
	"bytes"
	"fmt"

	gdef "github.com/kemadev/iac-components/pkg/github/define"
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
			Entity: "@kemadev/maintainers",
		},
	},
}

var CodeownersDefaultContent = `# File managed by repo-as-code, do not edit manually!
# Read more about CODEOWNERS [here](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners)

`

func createCodeownersSetDefaults(args *CodeownersArgs) error {
	if len(args.Codeowners) == 0 {
		args.Codeowners = CodeownersDefaultArgs.Codeowners
	}
	for _, codeowner := range args.Codeowners {
		if codeowner.Path == "" {
			return fmt.Errorf("Codeowner Path must be set")
		} else if codeowner.Path == "CHANGEME" {
			return fmt.Errorf("Codeowner Path must be changed from the default value")
		}
		if codeowner.Entity == "" {
			return fmt.Errorf("Codeowner Entity must be set")
		} else if codeowner.Entity == "CHANGEME" {
			return fmt.Errorf("Codeowner Entity must be changed from the default value")
		}
	}
	return nil
}

func createCodeownersContent(args *CodeownersArgs) string {
	var codeownersContent bytes.Buffer
	codeownersContent.WriteString(CodeownersDefaultContent)
	for _, codeowner := range args.Codeowners {
		codeownersContent.WriteString(fmt.Sprintf("%s %s\n", codeowner.Path, codeowner.Entity))
	}
	return codeownersContent.String()
}

func createCodeowners(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, args CodeownersArgs, branch *github.Branch) error {
	codeownersGithubFileName := util.FormatResourceName(ctx, "Repository codeowners file")
	_, err := github.NewRepositoryFile(ctx, codeownersGithubFileName, &github.RepositoryFileArgs{
		Repository:        repo.Name,
		File:              pulumi.String(".github/CODEOWNERS"),
		Branch:            branch.Branch,
		Content:           pulumi.String(createCodeownersContent(&args)),
		CommitMessage:     pulumi.String(gdef.GitDefaultCommitMessage),
		CommitAuthor:      pulumi.String(gdef.GitCommiterName),
		CommitEmail:       pulumi.String(gdef.GitCommiterEmail),
		OverwriteOnCreate: pulumi.Bool(true),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
