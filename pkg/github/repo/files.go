package repo

import (
	"slices"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	gdef "github.com/kemadev/iac-components/pkg/github/define"
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
		".github/CODEOWNERS", // Already
		"CHANGELOG.md",
		"README.md",

		// Dockerfile
		// cmd
		// env/
		// release please
		// docker compose
		// templater le go.mod / go.sum

		// pas enable les workflows au bootstrap ca fout le zbeul ca se trigger a chaque commit i.e. chaque file sync
		// pas run sur main mais next, dependsOn sur la branche et argsbranch a passer

		// move deploy rien a foutre dans repo template
		// project wiki
		// merge queue

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
			Repository: repo.Name,
			File:       pulumi.String(file.Name),
			Content:           pulumi.String(file.Content),
			CommitMessage:     pulumi.String(gdef.GitDefaultCommitMessage),
			CommitAuthor:      pulumi.String(gdef.GitCommiterName),
			CommitEmail:       pulumi.String(gdef.GitCommiterEmail),
			OverwriteOnCreate: pulumi.Bool(true),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
	}
	return nil
}
