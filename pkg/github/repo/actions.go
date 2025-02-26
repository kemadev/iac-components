package repo

import (
	"slices"

	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionsArgs struct {
	Actions []string
}

var (
	ActionsDefaultActions = []string{
		// Internal workflows and actions
		"kemadev/workflows-and-actions/.github/workflows/*",
		"kemadev/workflows-and-actions/.github/actions/*",
	}
)

func createActionsSetDefaults(args *ActionsArgs) {
	if args.Actions == nil {
		args.Actions = ActionsDefaultActions
		return
	}
	if len(args.Actions) == 0 {
		args.Actions = ActionsDefaultActions
		return
	}
	for _, action := range ActionsDefaultActions {
		found := slices.Contains(args.Actions, action)
		if !found {
			args.Actions = append(args.Actions, action)
		}
	}
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
