package repo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RepositoryArgs struct {
	// Repository description
	// - required
	Description string
	// Repository homepage URL
	HomepageUrl string
	// Repository topics
	// - required
	Topics []string
	// Repository visibility
	// - required
	Visibility string
	// Repository is a template
	// - false by default
	IsTemplate bool
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
		SquashMergeCommitMessage: pulumi.String("PR_BODY"),
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

		// Unused fields
		// MergeCommitTitle:   pulumi.String("PR_TITLE"),
		// MergeCommitMessage: pulumi.String("PR_BODY"),
		// AutoInit:           pulumi.Bool(true),
		// GitignoreTemplate:  pulumi.String("Go"),
		// LicenseTemplate:    pulumi.String("mpl-2.0"),
		// Template:           pulumi.Bool(false),
		// Pages:              github.RepositoryPagesArgs{},

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
	return repo, nil
}
