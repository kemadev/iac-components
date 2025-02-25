package repo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProviderArgs struct {
	Owner string
}

func createProvider(ctx *pulumi.Context, args ProviderArgs) (*github.Provider, error) {
	providerName := util.FormatResourceName(ctx, "Provider")
	provider, err := github.NewProvider(ctx, providerName, &github.ProviderArgs{
		Owner: pulumi.String(args.Owner),
	})
	if err != nil {
		return nil, err
	}
	return provider, nil
}
