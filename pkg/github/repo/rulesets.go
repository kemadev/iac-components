package repo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RulesetsArgs struct {
	RequiredReviewersNext int
	RequiredReviewersProd int
	RequiredStatusChecks  []string
}

var RulesetsDefaultArgs = RulesetsArgs{
	RequiredReviewersNext: 1,
	RequiredReviewersProd: 1,
	RequiredStatusChecks: []string{
		"Global - CI / Scan code",
		"Global - CI / Scan secrets",
	},
}

func createRulesetsSetDefaults(args *RulesetsArgs) {
	if args.RequiredReviewersNext == 0 {
		args.RequiredReviewersNext = RulesetsDefaultArgs.RequiredReviewersNext
	}
	if args.RequiredReviewersProd == 0 {
		args.RequiredReviewersProd = RulesetsDefaultArgs.RequiredReviewersProd
	}
	if len(args.RequiredStatusChecks) == 0 {
		args.RequiredStatusChecks = RulesetsDefaultArgs.RequiredStatusChecks
	}
}

func createRulesets(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, argsRulesets RulesetsArgs, argsBranches BranchesArgs, argsEnvs EnvsArgs) error {
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

	rulesetBranchEnvDev := util.FormatResourceName(ctx, "Repository ruleset branch env dev")
	_, err = github.NewRepositoryRuleset(ctx, rulesetBranchEnvDev, &github.RepositoryRulesetArgs{
		Repository:  repo.Name,
		Name:        pulumi.String("branch-env-" + argsBranches.Dev),
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
			// Repository Maintainer
			github.RepositoryRulesetBypassActorArgs{
				ActorType:  pulumi.String("RepositoryRole"),
				ActorId:    pulumi.Int(2),
				BypassMode: pulumi.String("always"),
			},
			// Repository Writer
			github.RepositoryRulesetBypassActorArgs{
				ActorType:  pulumi.String("RepositoryRole"),
				ActorId:    pulumi.Int(4),
				BypassMode: pulumi.String("pull_request"),
			},
		},
		Conditions: github.RepositoryRulesetConditionsArgs{
			RefName: github.RepositoryRulesetConditionsRefNameArgs{
				Includes: pulumi.ToStringArray([]string{"refs/heads/" + argsBranches.Dev}),
				Excludes: pulumi.ToStringArray([]string{}),
			},
		},
		Rules: github.RepositoryRulesetRulesArgs{
			Creation:              pulumi.Bool(true),
			Deletion:              pulumi.Bool(true),
			NonFastForward:        pulumi.Bool(true),
			RequiredLinearHistory: pulumi.Bool(true),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	rulesetBranchEnvNext := util.FormatResourceName(ctx, "Repository ruleset branch env next")
	_, err = github.NewRepositoryRuleset(ctx, rulesetBranchEnvNext, &github.RepositoryRulesetArgs{
		Repository:  repo.Name,
		Name:        pulumi.String("branch-env-" + argsBranches.Next),
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
			// Repository Maintainer
			github.RepositoryRulesetBypassActorArgs{
				ActorType:  pulumi.String("RepositoryRole"),
				ActorId:    pulumi.Int(2),
				BypassMode: pulumi.String("always"),
			},
		},
		Conditions: github.RepositoryRulesetConditionsArgs{
			RefName: github.RepositoryRulesetConditionsRefNameArgs{
				Includes: pulumi.ToStringArray([]string{"refs/heads/" + argsBranches.Next}),
				Excludes: pulumi.ToStringArray([]string{}),
			},
		},
		Rules: github.RepositoryRulesetRulesArgs{
			Creation:              pulumi.Bool(true),
			Deletion:              pulumi.Bool(true),
			NonFastForward:        pulumi.Bool(true),
			RequiredLinearHistory: pulumi.Bool(true),
			PullRequest: github.RepositoryRulesetRulesPullRequestArgs{
				RequiredApprovingReviewCount:   pulumi.Int(argsRulesets.RequiredReviewersNext),
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
				RequiredDeploymentEnvironments: pulumi.ToStringArray([]string{argsEnvs.Dev}),
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

	rulesetBranchEnvProd := util.FormatResourceName(ctx, "Repository ruleset branch env prod")
	_, err = github.NewRepositoryRuleset(ctx, rulesetBranchEnvProd, &github.RepositoryRulesetArgs{
		Repository:  repo.Name,
		Name:        pulumi.String("branch-env-" + argsBranches.Prod),
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
				Includes: pulumi.ToStringArray([]string{"refs/heads/" + argsBranches.Prod}),
				Excludes: pulumi.ToStringArray([]string{}),
			},
		},
		Rules: github.RepositoryRulesetRulesArgs{
			Creation:              pulumi.Bool(true),
			Deletion:              pulumi.Bool(true),
			NonFastForward:        pulumi.Bool(true),
			RequiredLinearHistory: pulumi.Bool(true),
			PullRequest: github.RepositoryRulesetRulesPullRequestArgs{
				RequiredApprovingReviewCount:   pulumi.Int(argsRulesets.RequiredReviewersProd),
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
				RequiredDeploymentEnvironments: pulumi.ToStringArray([]string{argsEnvs.Next}),
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
