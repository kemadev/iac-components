package repo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IssueArgs struct {
	Name        string
	Color       string
	Description string
}

var IssuesDefaultArgs = map[string]IssueArgs{
	"area/docs": {
		Name:        "area/docs",
		Color:       "1850c9", // Dark Blue
		Description: "Related to documentation",
	},
	"area/infra": {
		Name:        "area/infra",
		Color:       "ff9900", // Orange
		Description: "Related to infrastructure",
	},
	"area/core": {
		Name:        "area/core",
		Color:       "e74c3c", // Red
		Description: "Related to core functionality",
	},
	"area/workflows": {
		Name:        "area/workflows",
		Color:       "9b59b6", // Purple
		Description: "Related to GitHub workflows",
	},
	"area/dependencies": {
		Name:        "area/dependencies",
		Color:       "1abc9c", // Turquoise
		Description: "Related to dependencies",
	},
	"area/external": {
		Name:        "area/external",
		Color:       "34495e", // Dark Blue
		Description: "Related to external services",
	},
	"area/frontend": {
		Name:        "area/frontend",
		Color:       "83ed5a", // Light Green
		Description: "Related to frontend",
	},
	"area/backend": {
		Name:        "area/backend",
		Color:       "47a7b2", // Light Blue
		Description: "Related to backend",
	},
	"area/api": {
		Name:        "area/api",
		Color:       "27ae60", // Dark Green
		Description: "Related to API",
	},
	"area/data": {
		Name:        "area/data",
		Color:       "d68068", // Light Red
		Description: "Related to data",
	},
	"status/needs-triage": {
		Name:        "status/needs-triage",
		Color:       "a9eaf2", // Light Turquoise
		Description: "Needs triage, labeling, and planning",
	},
	"status/needs-reproduction": {
		Name:        "status/needs-reproduction",
		Color:       "8b58e2", // Dark Purple
		Description: "Needs to be reproduced and confirmed",
	},
	"status/needs-investigation": {
		Name:        "status/needs-investigation",
		Color:       "f1c40f", // Yellow
		Description: "Needs investigation and analysis",
	},
	"status/needs-info": {
		Name:        "status/needs-info",
		Color:       "8e44ad", // Dark Purple
		Description: "Needs more information from parties involved",
	},
	"status/stale": {
		Name:        "status/stale",
		Color:       "bdc3c7", // Grey
		Description: "Stale, no activity for a while",
	},
	"status/blocked": {
		Name:        "status/blocked",
		Color:       "5c6768", // Dark Grey
		Description: "Blocked, waiting for something",
	},
	"status/help-wanted": {
		Name:        "status/help-wanted",
		Color:       "2ecc71", // Light Green
		Description: "Assistance from the community is needed",
	},
	"status/duplicate": {
		Name:        "status/duplicate",
		Color:       "95a5a6", // Light Grey
		Description: "Already exists, duplicate",
	},
	"status/wont-fix": {
		Name:        "status/wont-fix",
		Color:       "7f8c8d", // Dark Grey
		Description: "Won't fix, not going to be addressed",
	},
	"status/work-in-progress": {
		Name:        "status/work-in-progress",
		Color:       "f1c40f", // Yellow
		Description: "Currently being worked on",
	},
	"status/up-for-grabs": {
		Name:        "status/up-for-grabs",
		Color:       "2ecc71", // Light Green
		Description: "Ready for someone to take it",
	},
	"status/closed": {
		Name:        "status/closed",
		Color:       "95a5a6", // Light Grey
		Description: "No further action planned",
	},
	"impact/low": {
		Name:        "impact/low",
		Color:       "97c4aa", // Light Green
		Description: "Impact is low",
	},
	"impact/medium": {
		Name:        "impact/medium",
		Color:       "f1c40f", // Yellow
		Description: "Impact is quite significant",
	},
	"impact/high": {
		Name:        "impact/high",
		Color:       "e74c3c", // Red
		Description: "Impact is critical and needs immediate attention",
	},
	"priority/P0": {
		Name:        "priority/P0",
		Color:       "e83c81", // Pink
		Description: "Critical, needs action immediately",
	},
	"priority/P1": {
		Name:        "priority/P1",
		Color:       "e74c3c", // Red
		Description: "High priority, needs action soon",
	},
	"priority/P2": {
		Name:        "priority/P2",
		Color:       "f39c12", // Orange
		Description: "Medium priority, needs action",
	},
	"type/bug": {
		Name:        "type/bug",
		Color:       "e74c3c", // Red
		Description: "Something is not working as expected",
	},
	"type/feature": {
		Name:        "type/feature",
		Color:       "2ecc71", // Light Green
		Description: "New functionality or feature",
	},
	"type/question": {
		Name:        "type/question",
		Color:       "3498db", // Blue
		Description: "Question or inquiry",
	},
	"type/security": {
		Name:        "type/security",
		Color:       "c0392b", // Dark Red
		Description: "Security related / vulnerability, needs immediate attention",
	},
	"type/performance": {
		Name:        "type/performance",
		Color:       "f39c12", // Orange
		Description: "Performance related",
	},
	"type/announcement": {
		Name:        "type/announcement",
		Color:       "f1c40f", // Yellow
		Description: "Announcement or news",
	},
	"release/pending": {
		Name:        "release/pending",
		Color:       "f1c40f", // Yellow
		Description: "Release is pending",
	},
	"release/released": {
		Name:        "release/released",
		Color:       "2ecc71", // Light Green
		Description: "Release has been completed",
	},
	"release/breaking": {
		Name:        "release/breaking",
		Color:       "e74c3c", // Red
		Description: "Breaking changes, needs special attention",
	},
	"platform/ios": {
		Name:        "platform/ios",
		Color:       "3498db", // Blue
		Description: "Concerns iOS platform",
	},
	"platform/android": {
		Name:        "platform/android",
		Color:       "2ecc71", // Light Green
		Description: "Concerns Android platform",
	},
	"platform/windows": {
		Name:        "platform/windows",
		Color:       "415dc1", // Dark Blue
		Description: "Concerns Windows platform",
	},
	"platform/mac": {
		Name:        "platform/mac",
		Color:       "e0c6af", // Light Brown
		Description: "Concerns Mac platform",
	},
	"platform/linux": {
		Name:        "platform/linux",
		Color:       "e2e18a", // Light Yellow
		Description: "Concerns Linux platform",
	},
	"platform/web": {
		Name:        "platform/web",
		Color:       "607fb2", // Dark Turquoise
		Description: "Concerns Web (browser) platform",
	},
	"deploy/aws": {
		Name:        "deploy/aws",
		Color:       "f39c12", // Orange
		Description: "Deployment is on AWS",
	},
	"deploy/azure": {
		Name:        "deploy/azure",
		Color:       "3498db", // Blue
		Description: "Deployment is on Azure",
	},
	"deploy/gcp": {
		Name:        "deploy/gcp",
		Color:       "2ecc71", // Light Green
		Description: "Deployment is on GCP",
	},
	"deploy/on-prem": {
		Name:        "deploy/on-prem",
		Color:       "9b59b6", // Purple
		Description: "Deployment is on-premises",
	},
	"size/XS": {
		Name:        "size/XS",
		Color:       "2ecc71", // Light Green
		Description: "Estimated amount of work is extra small",
	},
	"size/S": {
		Name:        "size/S",
		Color:       "f1c40f", // Yellow
		Description: "Estimated amount of work is small",
	},
	"size/M": {
		Name:        "size/M",
		Color:       "e67e22", // Orange
		Description: "Estimated amount of work is medium",
	},
	"size/L": {
		Name:        "size/L",
		Color:       "e74c3c", // Red
		Description: "Estimated amount of work is large, might need more review",
	},
	"size/XL": {
		Name:        "size/XL",
		Color:       "c0392b", // Dark Red
		Description: "Estimated amount of work is extra large, needs conscientious review",
	},
	"size/tbd": {
		Name:        "size/tbd",
		Color:       "95a5a6", // Light Grey
		Description: "Estimated amount of work is yet to be determined",
	},
	"complexity/low": {
		Name:        "complexity/low",
		Color:       "2ecc71", // Light Green
		Description: "Estimated complexity for the task is low",
	},
	"complexity/medium": {
		Name:        "complexity/medium",
		Color:       "f1c40f", // Yellow
		Description: "Estimated complexity for the task is medium",
	},
	"complexity/high": {
		Name:        "complexity/high",
		Color:       "e74c3c", // Red
		Description: "Estimated complexity for the task is high, might need expert review",
	},
	"env/dev": {
		Name:        "env/dev",
		Color:       "3498db", // Blue
		Description: "Concerns development environment",
	},
	"env/next": {
		Name:        "env/next",
		Color:       "f1c40f", // Yellow
		Description: "Concerns next environment",
	},
	"env/prod": {
		Name:        "env/prod",
		Color:       "e74c3c", // Red
		Description: "Concerns production environment, treat with care",
	},
}

func createIssues(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository) error {
	// github.NewIssueLabels is too inconsistent, thus creation is done one by one
	for _, issueLabel := range IssuesDefaultArgs {
		issueLabelName := util.FormatResourceName(ctx, "Issue label "+issueLabel.Name)
		_, err := github.NewIssueLabel(ctx, issueLabelName, &github.IssueLabelArgs{
			Repository:  repo.Name,
			Name:        pulumi.String(issueLabel.Name),
			Color:       pulumi.String(issueLabel.Color),
			Description: pulumi.String(issueLabel.Description),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
	}
	milestoneName := util.FormatResourceName(ctx, "Project init milestone")
	milestone, err := github.NewRepositoryMilestone(ctx, milestoneName, &github.RepositoryMilestoneArgs{
		Repository: repo.Name,
		// Provider is configured to be the owner of the repository
		Owner:       provider.Owner.Elem(),
		Title:       pulumi.String("Project initialization"),
		Description: pulumi.String("Initial setup and configuration of the repository"),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	issueName := util.FormatResourceName(ctx, "Project init issue")
	_, err = github.NewIssue(ctx, issueName, &github.IssueArgs{
		Repository: repo.Name,
		Title:      pulumi.String("Welcome! :confetti_ball:"),
		Body: pulumi.String(`## A few things to get you started

- [ ] Edit [CODEOWNERS](../blob/main/.github/CODEOWNERS) to match project codeowners
- [ ] Edit [CHANGELOG.md](../blob/main/CHANGELOG.md) to match project's changelog (basically just remove the whole content)
- [ ] Edit [release-please's manifest](../blob/main/config/release-please/release-please-manifest.json) to ` + "`{}`" + `
- [ ] Edit [README.md](../blob/main/README.md) with project specific information
- [ ] Set up your project's wiki (just clone it locally, edit, and push back!)
- [ ] Enable [merge queue](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/configuring-pull-request-merges/managing-a-merge-queue) if you expect _quite_ a lot of PRs`),
		MilestoneNumber: milestone.Number,
		Labels: pulumi.StringArray{
			pulumi.String(IssuesDefaultArgs["status/up-for-grabs"].Name),
			pulumi.String(IssuesDefaultArgs["priority/P2"].Name),
			pulumi.String(IssuesDefaultArgs["size/XS"].Name),
			pulumi.String(IssuesDefaultArgs["complexity/low"].Name),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
