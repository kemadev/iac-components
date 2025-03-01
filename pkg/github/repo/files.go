package repo

import (
	"slices"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FilesArgs struct {
	UpstreamRepo string
	IgnoredGlobs []string
}

var FilesDefaultArgs = FilesArgs{
	IgnoredGlobs: []string{
		"CHANGELOG.md",
		".github/CODEOWNERS",
		"README.md",
	},
	UpstreamRepo: "https://github.com/kemadev/repository-template",
}

type GitFile struct {
	Name    string
	Content string
}

func createFilesSetDefaults(args *FilesArgs) {
	args.IgnoredGlobs = append(args.IgnoredGlobs, FilesDefaultArgs.IgnoredGlobs...)
	if args.UpstreamRepo == "" {
		args.UpstreamRepo = FilesDefaultArgs.UpstreamRepo
	}
}

func getRepoTemplateFilesList(upstreamRepo string) ([]GitFile, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: upstreamRepo,
	})
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}
	var files []GitFile
	tree.Files().ForEach(func(f *object.File) error {
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

func createFiles(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, args FilesArgs) error {
	filesList, err := getRepoTemplateFilesList(args.UpstreamRepo)
	if err != nil {
		return err
	}
	for _, file := range filesList {
		if found := slices.Contains(args.IgnoredGlobs, file.Name); found {
			continue
		}
		fileName := util.FormatResourceName(ctx, "Repository file "+file.Name)
		_, err := github.NewRepositoryFile(ctx, fileName, &github.RepositoryFileArgs{
			Repository:        repo.Name,
			File:              pulumi.String(file.Name),
			Content:           pulumi.String(file.Content),
			CommitMessage:     pulumi.String("feat(repo-sync): update repository files"),
			CommitAuthor:      pulumi.String("pulumi[bot]"),                                  // TODO make this a reference to global settings
			CommitEmail:       pulumi.String("kemadev+pulumi[bot]@users.noreply.github.com"), // TODO make this a reference to global settings
			OverwriteOnCreate: pulumi.Bool(true),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
	}
	return nil
}
