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

var (
	IssuesDefaultArgs = []IssueArgs{
		{
			Name:        "area/docs",
			Color:       "1850c9", // Dark Blue
			Description: "Related to documentation",
		},
		{
			Name:        "area/infra",
			Color:       "ff9900", // Orange
			Description: "Related to infrastructure",
		},
		{
			Name:        "area/core",
			Color:       "e74c3c", // Red
			Description: "Related to core functionality",
		},
		{
			Name:        "area/workflows",
			Color:       "9b59b6", // Purple
			Description: "Related to GitHub workflows",
		},
		{
			Name:        "area/dependencies",
			Color:       "1abc9c", // Turquoise
			Description: "Related to dependencies",
		},
		{
			Name:        "area/external",
			Color:       "34495e", // Dark Blue
			Description: "Related to external services",
		},
		{
			Name:        "area/frontend",
			Color:       "83ed5a", // Light Green
			Description: "Related to frontend",
		},
		{
			Name:        "area/backend",
			Color:       "47a7b2", // Light Blue
			Description: "Related to backend",
		},
		{
			Name:        "area/api",
			Color:       "27ae60", // Dark Green
			Description: "Related to API",
		},
		{
			Name:        "area/data",
			Color:       "d68068", // Light Red
			Description: "Related to data",
		},
		{
			Name:        "status/needs-triage",
			Color:       "a9eaf2", // Light Turquoise
			Description: "Needs triage, labeling, and planning",
		},
		{
			Name:        "status/needs-reproduction",
			Color:       "8b58e2", // Dark Purple
			Description: "Needs to be reproduced and confirmed",
		},
		{
			Name:        "status/needs-investigation",
			Color:       "f1c40f", // Yellow
			Description: "Needs investigation and analysis",
		},
		{
			Name:        "status/needs-info",
			Color:       "8e44ad", // Dark Purple
			Description: "Needs more information from parties involved",
		},
		{
			Name:        "status/stale",
			Color:       "bdc3c7", // Grey
			Description: "Stale, no activity for a while",
		},
		{
			Name:        "status/blocked",
			Color:       "5c6768", // Dark Grey
			Description: "Blocked, waiting for something",
		},
		{
			Name:        "status/help-wanted",
			Color:       "2ecc71", // Light Green
			Description: "Assistance from the community is needed",
		},
		{
			Name:        "status/duplicate",
			Color:       "95a5a6", // Light Grey
			Description: "Already exists, duplicate",
		},
		{
			Name:        "status/wont-fix",
			Color:       "7f8c8d", // Dark Grey
			Description: "Won't fix, not going to be addressed",
		},
		{
			Name:        "status/work-in-progress",
			Color:       "f1c40f", // Yellow
			Description: "Currently being worked on",
		},
		{
			Name:        "status/up-for-grabs",
			Color:       "2ecc71", // Light Green
			Description: "Ready for someone to take it",
		},
		{
			Name:        "status/closed",
			Color:       "95a5a6", // Light Grey
			Description: "No further action planned",
		},
		{
			Name:        "impact/low",
			Color:       "97c4aa", // Light Green
			Description: "Impact is low",
		},
		{
			Name:        "impact/medium",
			Color:       "f1c40f", // Yellow
			Description: "Impact is quite significant",
		},
		{
			Name:        "impact/high",
			Color:       "e74c3c", // Red
			Description: "Impact is critical and needs immediate attention",
		},
		{
			Name:        "priority/P0",
			Color:       "e83c81", // Pink
			Description: "Critical, needs action immediately",
		},
		{
			Name:        "priority/P1",
			Color:       "e74c3c", // Red
			Description: "High priority, needs action soon",
		},
		{
			Name:        "priority/P2",
			Color:       "f39c12", // Orange
			Description: "Medium priority, needs action",
		},
		{
			Name:        "type/bug",
			Color:       "e74c3c", // Red
			Description: "Something is not working as expected",
		},
		{
			Name:        "type/feature",
			Color:       "2ecc71", // Light Green
			Description: "New functionality or feature",
		},
		{
			Name:        "type/question",
			Color:       "3498db", // Blue
			Description: "Question or inquiry",
		},
		{
			Name:        "type/security",
			Color:       "c0392b", // Dark Red
			Description: "Security related / vulnerability, needs immediate attention",
		},
		{
			Name:        "type/performance",
			Color:       "f39c12", // Orange
			Description: "Performance related",
		},
		{
			Name:        "type/announcement",
			Color:       "#f1c40f", // Yellow
			Description: "Announcement or news",
		},
		{
			Name:        "release/pending",
			Color:       "f1c40f", // Yellow
			Description: "Release is pending",
		},
		{
			Name:        "release/released",
			Color:       "2ecc71", // Light Green
			Description: "Release has been completed",
		},
		{
			Name:        "release/breaking",
			Color:       "e74c3c", // Red
			Description: "Breaking changes, needs special attention",
		},
		{
			Name:        "platform/ios",
			Color:       "3498db", // Blue
			Description: "Concerns iOS platform",
		},
		{
			Name:        "platform/android",
			Color:       "2ecc71", // Light Green
			Description: "Concerns Android platform",
		},
		{
			Name:        "platform/windows",
			Color:       "415dc1", // Dark Blue
			Description: "Concerns Windows platform",
		},
		{
			Name:        "platform/mac",
			Color:       "e0c6af", // Light Brown
			Description: "Concerns Mac platform",
		},
		{
			Name:        "platform/linux",
			Color:       "e2e18a", // Light Yellow
			Description: "Concerns Linux platform",
		},
		{
			Name:        "platform/web",
			Color:       "607fb2", // Dark Turquoise
			Description: "Concerns Web (browser) platform",
		},
		{
			Name:        "deploy/aws",
			Color:       "f39c12", // Orange
			Description: "Deployment is on AWS",
		},
		{
			Name:        "deploy/azure",
			Color:       "3498db", // Blue
			Description: "Deployment is on Azure",
		},
		{
			Name:        "deploy/gcp",
			Color:       "2ecc71", // Light Green
			Description: "Deployment is on GCP",
		},
		{
			Name:        "deploy/on-prem",
			Color:       "9b59b6", // Purple
			Description: "Deployment is on-premises",
		},
		{
			Name:        "size/XS",
			Color:       "2ecc71", // Light Green
			Description: "Estimated amount of work is extra small",
		},
		{
			Name:        "size/S",
			Color:       "f1c40f", // Yellow
			Description: "Estimated amount of work is small",
		},
		{
			Name:        "size/M",
			Color:       "e67e22", // Orange
			Description: "Estimated amount of work is medium",
		},
		{
			Name:        "size/L",
			Color:       "e74c3c", // Red
			Description: "Estimated amount of work is large, might need more review",
		},
		{
			Name:        "size/XL",
			Color:       "c0392b", // Dark Red
			Description: "Estimated amount of work is extra large, needs conscientious review",
		},
		{
			Name:        "size/tbd",
			Color:       "95a5a6", // Light Grey
			Description: "Estimated amount of work is yet to be determined",
		},
		{
			Name:        "complexity/low",
			Color:       "2ecc71", // Light Green
			Description: "Estimated complexity for the task is low",
		},
		{
			Name:        "complexity/medium",
			Color:       "f1c40f", // Yellow
			Description: "Estimated complexity for the task is medium",
		},
		{
			Name:        "complexity/high",
			Color:       "e74c3c", // Red
			Description: "Estimated complexity for the task is high, might need expert review",
		},
		{
			Name:        "env/dev",
			Color:       "3498db", // Blue
			Description: "Concerns development environment",
		},
		{
			Name:        "env/next",
			Color:       "f1c40f", // Yellow
			Description: "Concerns next environment",
		},
		{
			Name:        "env/prod",
			Color:       "e74c3c", // Red
			Description: "Concerns production environment, treat with care",
		},
	}
)

func createIssues(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository) error {
	for _, issueLabel := range IssuesDefaultArgs {
		issueLabelName := util.FormatResourceName(ctx, issueLabel.Name)
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
	return nil
}
