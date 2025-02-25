package githubrepo

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProviderArgs struct {
	owner string
}

func createProvider(ctx *pulumi.Context, args ProviderArgs) (*github.Provider, error) {
	provider, err := github.NewProvider(ctx, "github", &github.ProviderArgs{
		Owner: pulumi.String(args.owner),
	})
	if err != nil {
		return nil, err
	}
	return provider, nil
}
