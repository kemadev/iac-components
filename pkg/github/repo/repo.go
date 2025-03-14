package repo

import (
	"fmt"

	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DirectMember struct {
	Username string
	Role     string
}

type Team struct {
	Name string
	Role string
}

type RepositoryArgs struct {
	Name          string
	Description   string
	HomepageUrl   string
	defaultBranch string
	Topics        []string
	Visibility    string
	Archived      bool
	IsTemplate    bool
	Teams         []Team
	DirectMembers []DirectMember
}

var RepositoryDefaultArgs = RepositoryArgs{
	Visibility:    "private",
	IsTemplate:    false,
	defaultBranch: "main",
}

func createRepositorySetDefaults(args *RepositoryArgs) error {
	if args.Description == "" {
		return fmt.Errorf("Repository Description is required")
	} else if args.Description == "CHANGEME" {
		return fmt.Errorf("Repository Description must be changed from the default value")
	}
	if args.Visibility == "" {
		args.Visibility = RepositoryDefaultArgs.Visibility
	} else if args.Visibility == "CHANGEME" {
		return fmt.Errorf("Repository Visibility must be changed from the default value")
	}
	if args.defaultBranch == "" {
		args.defaultBranch = RepositoryDefaultArgs.defaultBranch
	}
	return nil
}

func createRepo(ctx *pulumi.Context, provider *github.Provider, argsRepo RepositoryArgs) (*github.Repository, error) {
	repoName := util.FormatResourceName(ctx, "Repository")
	repo, err := github.NewRepository(ctx, repoName, &github.RepositoryArgs{
		Name:                     pulumi.String(argsRepo.Name),
		Description: pulumi.String(argsRepo.Description),
		HomepageUrl: pulumi.String(argsRepo.HomepageUrl),
		Topics: func() pulumi.StringArrayInput {
			var topics pulumi.StringArray
			for _, topic := range argsRepo.Topics {
				topics = append(topics, pulumi.String(topic))
			}
			return topics
		}(),
		Visibility: pulumi.String(argsRepo.Visibility),
		IsTemplate: pulumi.Bool(argsRepo.IsTemplate),

		// Prevent accidental deletion
		ArchiveOnDestroy: pulumi.Bool(true),
		// Allow non-admins read access from pulumi
		IgnoreVulnerabilityAlertsDuringRead: pulumi.Bool(true),

		AllowSquashMerge:         pulumi.Bool(true),
		SquashMergeCommitTitle:   pulumi.String("PR_TITLE"),
		SquashMergeCommitMessage: pulumi.String("PR_BODY"),
		AllowMergeCommit:         pulumi.Bool(false),
		AllowRebaseMerge:         pulumi.Bool(false),
		AllowUpdateBranch:        pulumi.Bool(true),
		AllowAutoMerge:           pulumi.Bool(true),
		DeleteBranchOnMerge:      pulumi.Bool(true),
		HasDiscussions:           pulumi.Bool(true),
		HasIssues:                pulumi.Bool(true),
		HasProjects:              pulumi.Bool(true),
		HasWiki:                  pulumi.Bool(true),
		HasDownloads:             pulumi.Bool(false),
		Archived:                 pulumi.Bool(argsRepo.Archived),
		WebCommitSignoffRequired: pulumi.Bool(false),
		AutoInit:                 pulumi.Bool(true),

		VulnerabilityAlerts: func() pulumi.Bool {
			if argsRepo.Visibility == "public" {
				return pulumi.Bool(true)
			}
			// Advanced Security is required for private repositories
			return pulumi.Bool(false)
		}(),
		SecurityAndAnalysis: func() *github.RepositorySecurityAndAnalysisArgs {
			if argsRepo.Visibility == "public" {
				return &github.RepositorySecurityAndAnalysisArgs{
					SecretScanning: github.RepositorySecurityAndAnalysisSecretScanningArgs{
						Status: pulumi.String("enabled"),
					},
					SecretScanningPushProtection: github.RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs{
						Status: pulumi.String("enabled"),
					},
				}
			}
			// Advanced Security is required for private repositories
			return nil
		}(),
	}, pulumi.Provider(provider), pulumi.IgnoreChanges([]string{"template"}))
	if err != nil {
		return nil, err
	}
	defaulBranchName := util.FormatResourceName(ctx, "Repository default branch")
	_, err = github.NewBranch(ctx, defaulBranchName, &github.BranchArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(argsRepo.defaultBranch),
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}
	branchDefaultName := util.FormatResourceName(ctx, "Repository default branch")
	_, err = github.NewBranchDefault(ctx, branchDefaultName, &github.BranchDefaultArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(argsRepo.defaultBranch),
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}

	repoCollaboratorsName := util.FormatResourceName(ctx, "Repository collaborators")
	_, err = github.NewRepositoryCollaborators(ctx, repoCollaboratorsName, &github.RepositoryCollaboratorsArgs{
		Repository: repo.Name,
		Users: func() github.RepositoryCollaboratorsUserArray {
			var members github.RepositoryCollaboratorsUserArray
			for _, m := range argsRepo.DirectMembers {
				members = append(members, &github.RepositoryCollaboratorsUserArgs{
					Username:   pulumi.String(m.Username),
					Permission: pulumi.String(m.Role),
				})
			}
			return members
		}(),
		Teams: func() github.RepositoryCollaboratorsTeamArray {
			var teams github.RepositoryCollaboratorsTeamArray
			for _, t := range argsRepo.Teams {
				teams = append(teams, &github.RepositoryCollaboratorsTeamArgs{
					TeamId:     pulumi.String(t.Name),
					Permission: pulumi.String(t.Role),
				})
			}
			return teams
		}(),
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}
	return repo, nil
}
