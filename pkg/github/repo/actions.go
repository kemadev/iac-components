package githubrepo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionsArgs struct {
	Actions []string
}

func createActions(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, args ActionsArgs) error {
	actionsRepositoryPermissionsName := util.FormatResourceName(ctx, "Actions repository permissions")
	_, err := github.NewActionsRepositoryPermissions(ctx, actionsRepositoryPermissionsName, &github.ActionsRepositoryPermissionsArgs{
		Repository:     repo.Name,
		Enabled:        pulumi.Bool(true),
		AllowedActions: pulumi.String("selected"),
		AllowedActionsConfig: &github.ActionsRepositoryPermissionsAllowedActionsConfigArgs{
			GithubOwnedAllowed: pulumi.Bool(true),
			VerifiedAllowed:    pulumi.Bool(false),
			PatternsAlloweds: func() pulumi.StringArray {
				var patterns pulumi.StringArray
				for _, action := range args.Actions {
					patterns = append(patterns, pulumi.String(action))
				}
				return patterns
			}(),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
