package org

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
		// Actions from reusable workflows and actions
		"anchore/sbom-action@*",
		"anchore/scan-action@*",
		"aws-actions/configure-aws-credentials@*",
		"DavidAnson/markdownlint-cli2-action@*",
		"docker://rhysd/actionlint@*",
		"golangci/golangci-lint-action@*",
		"googleapis/release-please-action@*",
		"goreleaser/goreleaser-action@*",
		"hadolint/hadolint-action@*",
		"ibiqlik/action-yamllint@*",
		"pulumi/actions@*",
		"semgrep/semgrep@*",
		"trufflesecurity/trufflehog@*",
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

func createActions(ctx *pulumi.Context, provider *github.Provider, args ActionsArgs) error {
	actionsOrganizationPermissionsName := util.FormatResourceName(ctx, "Actions organization permissions")
	_, err := github.NewActionsOrganizationPermissions(ctx, actionsOrganizationPermissionsName, &github.ActionsOrganizationPermissionsArgs{
		AllowedActions: pulumi.String("selected"),
		AllowedActionsConfig: &github.ActionsOrganizationPermissionsAllowedActionsConfigArgs{
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
