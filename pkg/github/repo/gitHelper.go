package repo

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

type ghRepoTree struct {
	Repo *git.Repository
	Tree *object.Tree
}

var templateRepo ghRepoTree

func (ghRepoTree) init(repoURL string) error {
	if templateRepo.Repo == nil {
		repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: repoURL,
		})
		if err != nil {
			return err
		}
		templateRepo.Repo = repo
	}
	if templateRepo.Tree == nil {
		ref, err := templateRepo.Repo.Head()
		if err != nil {
			return err
		}

		commit, err := templateRepo.Repo.CommitObject(ref.Hash())
		if err != nil {
			return err
		}

		tree, err := commit.Tree()
		if err != nil {
			return err
		}
		templateRepo.Tree = tree
	}
	return nil
}
