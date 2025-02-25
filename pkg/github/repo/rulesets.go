package repo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RulesetsArgs struct {
	requiredReviewersNext int
	requiredReviewersProd int
}

func createRulesets(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository, environments TEnvironmentsCreated, argsRulesets RulesetsArgs, argsBranches BranchesArgs) error {
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
				RequiredApprovingReviewCount:   pulumi.Int(argsRulesets.requiredReviewersNext),
				DismissStaleReviewsOnPush:      pulumi.Bool(true),
				RequireCodeOwnerReview:         pulumi.Bool(true),
				RequireLastPushApproval:        pulumi.Bool(true),
				RequiredReviewThreadResolution: pulumi.Bool(true),
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
				RequiredApprovingReviewCount:   pulumi.Int(argsRulesets.requiredReviewersProd),
				DismissStaleReviewsOnPush:      pulumi.Bool(true),
				RequireCodeOwnerReview:         pulumi.Bool(true),
				RequireLastPushApproval:        pulumi.Bool(true),
				RequiredReviewThreadResolution: pulumi.Bool(true),
			},
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
