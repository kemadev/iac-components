package githubrepo

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createIssues(ctx *pulumi.Context, provider *github.Provider, repo *github.Repository) error {
	issueLabelsName := util.FormatResourceName(ctx, "Issue labels")
	_, err := github.NewIssueLabels(ctx, issueLabelsName, &github.IssueLabelsArgs{
		Repository: repo.Name,
		Labels: github.IssueLabelsLabelArray{
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/docs"),
				Color:       pulumi.String("1850c9"), // Dark Blue
				Description: pulumi.String("Related to documentation"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/infra"),
				Color:       pulumi.String("ff9900"), // Orange
				Description: pulumi.String("Related to infrastructure"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/core"),
				Color:       pulumi.String("e74c3c"), // Red
				Description: pulumi.String("Related to core functionality"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/workflows"),
				Color:       pulumi.String("9b59b6"), // Purple
				Description: pulumi.String("Related to GitHub workflows"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/dependencies"),
				Color:       pulumi.String("1abc9c"), // Turquoise
				Description: pulumi.String("Related to dependencies"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/external"),
				Color:       pulumi.String("34495e"), // Dark Blue
				Description: pulumi.String("Related to external services"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/frontend"),
				Color:       pulumi.String("83ed5a"), // Light Green
				Description: pulumi.String("Related to frontend"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/backend"),
				Color:       pulumi.String("47a7b2"), // Light Blue
				Description: pulumi.String("Related to backend"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/api"),
				Color:       pulumi.String("27ae60"), // Dark Green
				Description: pulumi.String("Related to API"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("area/data"),
				Color:       pulumi.String("d68068"), // Light Red
				Description: pulumi.String("Related to data"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/needs-triage"),
				Color:       pulumi.String("a9eaf2"), // Light Turquoise
				Description: pulumi.String("Needs triage, labeling, and planning"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/needs-reproduction"),
				Color:       pulumi.String("8b58e2"), // Dark Purple
				Description: pulumi.String("Needs to be reproduced and confirmed"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/needs-investigation"),
				Color:       pulumi.String("f1c40f"), // Yellow
				Description: pulumi.String("Needs investigation and analysis"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/needs-info"),
				Color:       pulumi.String("8e44ad"), // Dark Purple
				Description: pulumi.String("Needs more information from parties involved"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/stale"),
				Color:       pulumi.String("bdc3c7"), // Grey
				Description: pulumi.String("Stale, no activity for a while"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/blocked"),
				Color:       pulumi.String("5c6768"), // Dark Grey
				Description: pulumi.String("Blocked, waiting for something"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/help-wanted"),
				Color:       pulumi.String("2ecc71"), // Light Green
				Description: pulumi.String("Assistance from the community is needed"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/duplicate"),
				Color:       pulumi.String("95a5a6"), // Light Grey
				Description: pulumi.String("Already exists, duplicate"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/wont-fix"),
				Color:       pulumi.String("7f8c8d"), // Dark Grey
				Description: pulumi.String("Won't fix, not going to be addressed"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/work-in-progress"),
				Color:       pulumi.String("f1c40f"), // Yellow
				Description: pulumi.String("Currently being worked on"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/up-for-grabs"),
				Color:       pulumi.String("2ecc71"), // Light Green
				Description: pulumi.String("Ready for someone to take it"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("status/closed"),
				Color:       pulumi.String("95a5a6"), // Light Grey
				Description: pulumi.String("No further action planned"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("impact/low"),
				Color:       pulumi.String("97c4aa"), // Light Green
				Description: pulumi.String("Impact is low"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("impact/medium"),
				Color:       pulumi.String("f1c40f"), // Yellow
				Description: pulumi.String("Impact is quite significant"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("impact/high"),
				Color:       pulumi.String("e74c3c"), // Red
				Description: pulumi.String("Impact is critical and needs immediate attention"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("priority/P0"),
				Color:       pulumi.String("e83c81"), // Pink
				Description: pulumi.String("Critical, needs action immediately"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("priority/P1"),
				Color:       pulumi.String("e74c3c"), // Red
				Description: pulumi.String("High priority, needs action soon"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("priority/P2"),
				Color:       pulumi.String("f39c12"), // Orange
				Description: pulumi.String("Medium priority, needs action"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("type/bug"),
				Color:       pulumi.String("e74c3c"), // Red
				Description: pulumi.String("Something is not working as expected"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("type/feature"),
				Color:       pulumi.String("2ecc71"), // Light Green
				Description: pulumi.String("New functionality or feature"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("type/question"),
				Color:       pulumi.String("3498db"), // Blue
				Description: pulumi.String("Question or inquiry"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("type/security"),
				Color:       pulumi.String("c0392b"), // Dark Red
				Description: pulumi.String("Security related / vulnerability, needs immediate attention"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("type/performance"),
				Color:       pulumi.String("f39c12"), // Orange
				Description: pulumi.String("Performance related"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("type/announcement"),
				Color:       pulumi.String("#f1c40f"), // Yellow
				Description: pulumi.String("Announcement or news"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("release/pending"),
				Color:       pulumi.String("f1c40f"), // Yellow
				Description: pulumi.String("Release is pending"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("release/released"),
				Color:       pulumi.String("2ecc71"), // Light Green
				Description: pulumi.String("Release has been completed"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("release/breaking"),
				Color:       pulumi.String("e74c3c"), // Red
				Description: pulumi.String("Breaking changes, needs special attention"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("platform/ios"),
				Color:       pulumi.String("3498db"), // Blue
				Description: pulumi.String("Concerns iOS platform"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("platform/android"),
				Color:       pulumi.String("2ecc71"), // Light Green
				Description: pulumi.String("Concerns Android platform"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("platform/windows"),
				Color:       pulumi.String("415dc1"), // Dark Blue
				Description: pulumi.String("Concerns Windows platform"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("platform/mac"),
				Color:       pulumi.String("e0c6af"), // Light Brown
				Description: pulumi.String("Concerns Mac platform"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("platform/linux"),
				Color:       pulumi.String("e2e18a"), // Light Yellow
				Description: pulumi.String("Concerns Linux platform"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("platform/web"),
				Color:       pulumi.String("607fb2"), // Dark Turquoise
				Description: pulumi.String("Concerns Web (browser) platform"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("deploy/aws"),
				Color:       pulumi.String("f39c12"), // Orange
				Description: pulumi.String("Deployment is on AWS"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("deploy/azure"),
				Color:       pulumi.String("3498db"), // Blue
				Description: pulumi.String("Deployment is on Azure"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("deploy/gcp"),
				Color:       pulumi.String("2ecc71"), // Light Green
				Description: pulumi.String("Deployment is on GCP"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("deploy/on-prem"),
				Color:       pulumi.String("9b59b6"), // Purple
				Description: pulumi.String("Deployment is on-premises"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("size/XS"),
				Color:       pulumi.String("2ecc71"), // Light Green
				Description: pulumi.String("Estimated amount of work is extra small"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("size/S"),
				Color:       pulumi.String("f1c40f"), // Yellow
				Description: pulumi.String("Estimated amount of work is small"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("size/M"),
				Color:       pulumi.String("e67e22"), // Orange
				Description: pulumi.String("Estimated amount of work is medium"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("size/L"),
				Color:       pulumi.String("e74c3c"), // Red
				Description: pulumi.String("Estimated amount of work is large, might need more review"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("size/XL"),
				Color:       pulumi.String("c0392b"), // Dark Red
				Description: pulumi.String("Estimated amount of work is extra large, needs conscientious review"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("size/tbd"),
				Color:       pulumi.String("95a5a6"), // Light Grey
				Description: pulumi.String("Estimated amount of work is yet to be determined"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("complexity/low"),
				Color:       pulumi.String("2ecc71"), // Light Green
				Description: pulumi.String("Estimated complexity for the task is low"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("complexity/medium"),
				Color:       pulumi.String("f1c40f"), // Yellow
				Description: pulumi.String("Estimated complexity for the task is medium"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("complexity/high"),
				Color:       pulumi.String("e74c3c"), // Red
				Description: pulumi.String("Estimated complexity for the task is high, might need expert review"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("env/dev"),
				Color:       pulumi.String("3498db"), // Blue
				Description: pulumi.String("Concerns development environment"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("env/next"),
				Color:       pulumi.String("f1c40f"), // Yellow
				Description: pulumi.String("Concerns next environment"),
			},
			&github.IssueLabelsLabelArgs{
				Name:        pulumi.String("env/prod"),
				Color:       pulumi.String("e74c3c"), // Red
				Description: pulumi.String("Concerns production environment, treat with care"),
			},
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
