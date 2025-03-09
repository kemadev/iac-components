package main

import (
	"github.com/kemadev/iac-components/pkg/github/org"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := org.Wrapper(ctx, org.WrapperArgs{
			Settings: org.SettingsArgs{
				Company:      "kemadev",
				Description:  "Making cloud infrastructure a breeze!",
				Email:        "contact@kema.dev",
				BillingEmail: "billing@kema.dev",
				Blog:         "https://www.kema.dev",
				Location:     "France",
			},
			Members: org.MembersArgs{
				Members: []org.User{
					{
						Username: "kema-dev",
						Role:     "admin",
					},
				},
			},
			Teams: org.TeamsArgs{
				Teams: []org.TeamArgs{
					{
						Name: org.AdminTeamName,
						Members: []org.TeamMemberArgs{
							{
								Username: "kema-dev",
								Role:     "maintainer",
							},
						},
					},
					{
						Name: org.MaintainersTeamName,
						Members: []org.TeamMemberArgs{
							{
								Username: "kema-dev",
								Role:     "maintainer",
							},
						},
					},
					{
						Name: org.DevelopersTeamName,
						Members: []org.TeamMemberArgs{
							{
								Username: "kema-dev",
								Role:     "maintainer",
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
	return
}
