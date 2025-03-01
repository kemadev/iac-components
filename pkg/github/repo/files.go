package repo

import (
	"slices"

	"github.com/go-git/go-git/v5/plumbing/object"
	gdef "github.com/kemadev/iac-components/pkg/github/define"
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FilesArgs struct {
	UpstreamRepo        string
	ExcludedFiles       []string
	ChangesIgnoredFiles []string
}

var FilesDefaultArgs = FilesArgs{
	ExcludedFiles: []string{
		// Files handled by other functions in this module
		".github/CODEOWNERS",
		"go.mod",
		// Repository specific files
		"CHANGELOG.md",
	},
	ChangesIgnoredFiles: []string{
		"README.md",
		"cmd/main/main.go",
		"build/Dockerfile",
		"tool/docker/docker-compose.yaml",
		"config/env/aws/dev.env",
		"config/release-please/release-please-config.json",
		"config/release-please/release-please-manifest.json",
	},
	UpstreamRepo: "https://github.com/kemadev/repository-template",
}

type GitFile struct {
	Name    string
	Content string
}

func createFilesSetDefaults(args *FilesArgs) {
	args.ExcludedFiles = append(args.ExcludedFiles, FilesDefaultArgs.ExcludedFiles...)
	args.ChangesIgnoredFiles = append(args.ChangesIgnoredFiles, FilesDefaultArgs.ChangesIgnoredFiles...)
	if args.UpstreamRepo == "" {
		args.UpstreamRepo = FilesDefaultArgs.UpstreamRepo
	}
}

func getRepoTemplateFilesList() ([]GitFile, error) {
	var files []GitFile
	templateRepo.Tree.Files().ForEach(func(f *object.File) error {
		c, err := f.Contents()
		if err != nil {
			return err
		}
		files = append(files, GitFile{
			Name:    f.Name,
			Content: c,
		})
		return nil
	})
	return files, nil
}

func createFiles(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, args FilesArgs, branch *github.Branch) error {
	filesList, err := getRepoTemplateFilesList()
	if err != nil {
		return err
	}
	for _, file := range filesList {
		if found := slices.Contains(args.ExcludedFiles, file.Name); found {
			continue
		}
		var ignoreChangesOption pulumi.ResourceOption
		if found := slices.Contains(args.ChangesIgnoredFiles, file.Name); found {
			ignoreChangesOption = pulumi.IgnoreChanges([]string{"content"})
		}
		fileName := util.FormatResourceName(ctx, "Repository file "+file.Name)
		_, err := github.NewRepositoryFile(ctx, fileName, &github.RepositoryFileArgs{
			Repository:        repo.Name,
			Branch:            branch.Branch,
			File:              pulumi.String(file.Name),
			Content:           pulumi.String(file.Content),
			CommitMessage:     pulumi.String(gdef.GitDefaultCommitMessage),
			CommitAuthor:      pulumi.String(gdef.GitCommiterName),
			CommitEmail:       pulumi.String(gdef.GitCommiterEmail),
			OverwriteOnCreate: pulumi.Bool(true),
		}, pulumi.Provider(provider), ignoreChangesOption)
		if err != nil {
			return err
		}
	}
	return nil
}
