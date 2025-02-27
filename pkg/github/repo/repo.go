package repo

import (
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
	Description   string
	HomepageUrl   string
	Topics        []string
	Visibility    string
	IsTemplate    bool
	Teams         []Team
	DirectMembers []DirectMember
}

var RepositoryDefaultArgs = RepositoryArgs{
	Visibility: "private",
	IsTemplate: false,
}

func createRepositorySetDefaults(args *RepositoryArgs) {
	if args.Visibility == "" {
		args.Visibility = RepositoryDefaultArgs.Visibility
	}
	// actually useless, but for consistency
	if args.IsTemplate == false {
		args.IsTemplate = RepositoryDefaultArgs.IsTemplate
	}
}

func createRepo(ctx *pulumi.Context, provider *github.Provider, argsRepo RepositoryArgs, argsBranches BranchesArgs) (*github.Repository, error) {
	repoName := util.FormatResourceName(ctx, "Repository")
	repo, err := github.NewRepository(ctx, repoName, &github.RepositoryArgs{
		// Keep name from import
		// Name:        pulumi.String("repository-template"),

		// Prevent accidental deletion
		ArchiveOnDestroy: pulumi.Bool(true),
		// Allow non-admins read access from pulumi
		IgnoreVulnerabilityAlertsDuringRead: pulumi.Bool(true),

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
		IsTemplate: pulumi.Bool(argsRepo.IsTemplate == true),

		AllowSquashMerge:         pulumi.Bool(true),
		SquashMergeCommitTitle:   pulumi.String("PR_TITLE"),
		SquashMergeCommitMessage: pulumi.String("COMMIT_MESSAGES"),
		AllowMergeCommit:         pulumi.Bool(false),
		AllowRebaseMerge:         pulumi.Bool(false),
		AllowUpdateBranch:        pulumi.Bool(true),
		AllowAutoMerge:           pulumi.Bool(true),
		DeleteBranchOnMerge:      pulumi.Bool(true),
		HasDiscussions:           pulumi.Bool(true),
		HasIssues:                pulumi.Bool(true),
		HasProjects:              pulumi.Bool(true),
		HasWiki:                  pulumi.Bool(false),
		HasDownloads:             pulumi.Bool(false),
		Archived:                 pulumi.Bool(false),
		WebCommitSignoffRequired: pulumi.Bool(false),

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
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}
	branchDefaultName := util.FormatResourceName(ctx, "Repository default branch")
	_, err = github.NewBranchDefault(ctx, branchDefaultName, &github.BranchDefaultArgs{
		Repository: repo.Name,
		Branch:     pulumi.String(argsBranches.Default),
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
