package repo

import (
	"regexp"
	"strings"

	gdef "github.com/kemadev/iac-components/pkg/github/define"
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func getRepoTemplateFileContent(file string) (string, error) {
	f, err := templateRepo.Tree.File(file)
	if err != nil {
		return "", err
	}
	c, err := f.Contents()
	if err != nil {
		return "", err
	}
	return c, nil
}

func createTemplatedFiles(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, argsFiles FilesArgs, argsRepo RepositoryArgs, targetBranch string, baseBranch string) error {
	goModFileOriginalContent, err := getRepoTemplateFileContent("go.mod")
	if err != nil {
		return err
	}
	upstreamRepoPathParts := strings.Split(argsFiles.UpstreamRepo, "/")
	upstreamRepoName := upstreamRepoPathParts[len(upstreamRepoPathParts)-1]
	// replace repository part of the go.mod file with the new repository name
	re := regexp.MustCompile(`(module\s+github\.com/[^/]+/)` + regexp.QuoteMeta(upstreamRepoName) + `(/.*)?`)
	goModFileTemplatedContent := re.ReplaceAllString(goModFileOriginalContent, "${1}"+argsRepo.Name+"${2}")
	goModFileName := util.FormatResourceName(ctx, "Repository file go.mod")
	_, err = github.NewRepositoryFile(ctx, goModFileName, &github.RepositoryFileArgs{
		Repository:                   repo.Name,
		Branch:                       pulumi.String(targetBranch),
		AutocreateBranch:             pulumi.Bool(true),
		AutocreateBranchSourceBranch: pulumi.String(baseBranch),
		File:                         pulumi.String("go.mod"),
		Content:                      pulumi.String(goModFileTemplatedContent),
		CommitMessage:                pulumi.String(gdef.GitDefaultCommitMessage),
		CommitAuthor:                 pulumi.String(gdef.GitCommiterName),
		CommitEmail:                  pulumi.String(gdef.GitCommiterEmail),
		OverwriteOnCreate:            pulumi.Bool(true),
	}, pulumi.Provider(provider), pulumi.IgnoreChanges([]string{"content"}))
	if err != nil {
		return err
	}
	return nil
}
