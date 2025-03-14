package repo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RulesetsArgs struct {
	RequiredReviewersMain int
	RequiredStatusChecks  []string
}

var RulesetsDefaultArgs = RulesetsArgs{
	RequiredReviewersMain: 1,
	RequiredStatusChecks: []string{
		"Global - CI / Scan code",
		"Global - CI / Scan secrets",
		"Check PR title / Check PR title",
	},
}

func createRulesetsSetDefaults(args *RulesetsArgs) {
	if args.RequiredReviewersMain == 0 {
		args.RequiredReviewersMain = RulesetsDefaultArgs.RequiredReviewersMain
	}
	if len(args.RequiredStatusChecks) == 0 {
		args.RequiredStatusChecks = RulesetsDefaultArgs.RequiredStatusChecks
	}
}

func createRulesets(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, argsRulesets RulesetsArgs, argsEnvs EnvsArgs) error {
	rulesetBranchGlobalName := util.FormatResourceName(ctx, "Repository branch ruleset global")
	_, err := github.NewRepositoryRuleset(ctx, rulesetBranchGlobalName, &github.RepositoryRulesetArgs{
		Repository:  repo.Name,
		Name:        pulumi.String("branch-global"),
		Target:      pulumi.String("branch"),
		Enforcement: pulumi.String("active"),
		// @ref https://registry.terraform.io/providers/integrations/github/latest/docs/resources/repository_ruleset#bypass_actors
		BypassActors: github.RepositoryRulesetBypassActorArray{
			// Organization Admin
			github.RepositoryRulesetBypassActorArgs{
				ActorType:  pulumi.String("OrganizationAdmin"),
				ActorId:    pulumi.Int(1),
				BypassMode: pulumi.String("always"),
			},
			// Repository Admin
			github.RepositoryRulesetBypassActorArgs{
				ActorType:  pulumi.String("RepositoryRole"),
				ActorId:    pulumi.Int(5),
				BypassMode: pulumi.String("always"),
			},
		},
		Conditions: github.RepositoryRulesetConditionsArgs{
			RefName: github.RepositoryRulesetConditionsRefNameArgs{
				Includes: pulumi.ToStringArray([]string{"~ALL"}),
				Excludes: pulumi.ToStringArray([]string{}),
			},
		},
		Rules: github.RepositoryRulesetRulesArgs{
			RequiredSignatures: pulumi.Bool(true),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	rulesetTagGlobalName := util.FormatResourceName(ctx, "Repository tag ruleset global")
	_, err = github.NewRepositoryRuleset(ctx, rulesetTagGlobalName, &github.RepositoryRulesetArgs{
		Repository:  repo.Name,
		Name:        pulumi.String("tag-global"),
		Target:      pulumi.String("tag"),
		Enforcement: pulumi.String("active"),
		Conditions: github.RepositoryRulesetConditionsArgs{
			RefName: github.RepositoryRulesetConditionsRefNameArgs{
				Includes: pulumi.ToStringArray([]string{"~ALL"}),
				Excludes: pulumi.ToStringArray([]string{}),
			},
		},
		Rules: github.RepositoryRulesetRulesArgs{
			RequiredSignatures: pulumi.Bool(true),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	rulesetBranchEnvMain := util.FormatResourceName(ctx, "Repository ruleset branch env main")
	_, err = github.NewRepositoryRuleset(ctx, rulesetBranchEnvMain, &github.RepositoryRulesetArgs{
		Repository:  repo.Name,
		Name:        pulumi.String("branch-env-main"),
		Target:      pulumi.String("branch"),
		Enforcement: pulumi.String("active"),
		// @ref https://registry.terraform.io/providers/integrations/github/latest/docs/resources/repository_ruleset#bypass_actors
		BypassActors: github.RepositoryRulesetBypassActorArray{
			// Organization Admin
			github.RepositoryRulesetBypassActorArgs{
				ActorType:  pulumi.String("OrganizationAdmin"),
				ActorId:    pulumi.Int(1),
				BypassMode: pulumi.String("always"),
			},
			// Repository Admin
			github.RepositoryRulesetBypassActorArgs{
				ActorType:  pulumi.String("RepositoryRole"),
				ActorId:    pulumi.Int(5),
				BypassMode: pulumi.String("always"),
			},
		},
		Conditions: github.RepositoryRulesetConditionsArgs{
			RefName: github.RepositoryRulesetConditionsRefNameArgs{
				Includes: pulumi.ToStringArray([]string{"refs/heads/main"}),
				Excludes: pulumi.ToStringArray([]string{}),
			},
		},
		Rules: github.RepositoryRulesetRulesArgs{
			Creation:              pulumi.Bool(false),
			Deletion:              pulumi.Bool(true),
			NonFastForward:        pulumi.Bool(true),
			RequiredLinearHistory: pulumi.Bool(true),
			PullRequest: github.RepositoryRulesetRulesPullRequestArgs{
				RequiredApprovingReviewCount:   pulumi.Int(argsRulesets.RequiredReviewersMain),
				DismissStaleReviewsOnPush:      pulumi.Bool(true),
				RequireCodeOwnerReview:         pulumi.Bool(true),
				RequireLastPushApproval:        pulumi.Bool(true),
				RequiredReviewThreadResolution: pulumi.Bool(true),
			},
			MergeQueue: github.RepositoryRulesetRulesMergeQueueArgs{
				MergeMethod:                  pulumi.String("SQUASH"),
				GroupingStrategy:             pulumi.String("ALLGREEN"),
				MaxEntriesToBuild:            pulumi.Int(10),
				MinEntriesToMerge:            pulumi.Int(1),
				MinEntriesToMergeWaitMinutes: pulumi.Int(5),
				MaxEntriesToMerge:            pulumi.Int(5),
				CheckResponseTimeoutMinutes:  pulumi.Int(5),
			},
			RequiredDeployments: github.RepositoryRulesetRulesRequiredDeploymentsArgs{
				RequiredDeploymentEnvironments: pulumi.ToStringArray([]string{argsEnvs.Prod}),
			},
			RequiredStatusChecks: github.RepositoryRulesetRulesRequiredStatusChecksArgs{
				StrictRequiredStatusChecksPolicy: pulumi.Bool(true),
				DoNotEnforceOnCreate:             pulumi.Bool(false),
				RequiredChecks: func() github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArray {
					var checks github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArray
					for _, check := range argsRulesets.RequiredStatusChecks {
						checks = append(checks, github.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs{
							Context: pulumi.String(check),
						})
					}
					return checks
				}(),
			},
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
