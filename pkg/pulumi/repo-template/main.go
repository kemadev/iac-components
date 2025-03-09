package main

import (
	"github.com/kemadev/iac-components/pkg/github/repo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := repo.Wrapper(ctx, repo.WrapperArgs{
			Repository: repo.RepositoryArgs{
				// Set description
				Description: "CHANGEME",
				// Set visibility
				Visibility: "CHANGEME",
				// Optionally set other properties such as teams, direct members, etc.
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
