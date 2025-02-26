package org

import (
	"fmt"

	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SettingsArgs struct {
	BillingEmail string
	Blog         string
	Company      string
	Description  string
	Email        string
	Location     string
}

var (
	SettingsDefaultArgs = SettingsArgs{}
)

func createSettingsSetDefaults(args *SettingsArgs) error {
	if args.BillingEmail == "" {
		return fmt.Errorf("BillingEmail is required")
	}
	if args.Blog == "" {
		return fmt.Errorf("Blog is required")
	}
	if args.Company == "" {
		return fmt.Errorf("Company is required")
	}
	if args.Description == "" {
		return fmt.Errorf("Description is required")
	}
	if args.Email == "" {
		return fmt.Errorf("Email is required")
	}
	if args.Location == "" {
		return fmt.Errorf("Location is required")
	}
	return nil
}

func createSettings(ctx *pulumi.Context, provider *github.Provider, argsSettings SettingsArgs) error {
	err := createSettingsSetDefaults(&argsSettings)
	if err != nil {
		return err
	}
	settingsName := util.FormatResourceName(ctx, "Settings")
	_, err = github.NewOrganizationSettings(ctx, settingsName, &github.OrganizationSettingsArgs{
		// Provider is configured to be the owner of the organization
		Name:                        provider.Owner.Elem(),
		Description:                 pulumi.String(argsSettings.Description),
		Email:                       pulumi.String(argsSettings.Email),
		BillingEmail:                pulumi.String(argsSettings.BillingEmail),
		Blog:                        pulumi.String(argsSettings.Blog),
		Company:                     pulumi.String(argsSettings.Company),
		Location:                    pulumi.String(argsSettings.Location),
		DefaultRepositoryPermission: pulumi.String("read"),
		DependabotAlertsEnabledForNewRepositories:             pulumi.Bool(true),
		DependabotSecurityUpdatesEnabledForNewRepositories:    pulumi.Bool(true),
		DependencyGraphEnabledForNewRepositories:              pulumi.Bool(true),
		HasOrganizationProjects:                               pulumi.Bool(true),
		HasRepositoryProjects:                                 pulumi.Bool(true),
		MembersCanCreatePages:                                 pulumi.Bool(false),
		MembersCanCreatePrivatePages:                          pulumi.Bool(false),
		MembersCanCreatePublicPages:                           pulumi.Bool(false),
		MembersCanCreateRepositories:                          pulumi.Bool(true),
		MembersCanCreatePrivateRepositories:                   pulumi.Bool(true),
		MembersCanCreatePublicRepositories:                    pulumi.Bool(true),
		MembersCanForkPrivateRepositories:                     pulumi.Bool(true),
		WebCommitSignoffRequired:                              pulumi.Bool(false),
		SecretScanningEnabledForNewRepositories:               pulumi.Bool(true),
		SecretScanningPushProtectionEnabledForNewRepositories: pulumi.Bool(true),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}
	return nil
}
