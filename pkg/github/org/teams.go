package org

import (
	"github.com/kemadev/iac-components/pkg/util"
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TeamMemberArgs struct {
	Username string
	Role     string
}

type TeamArgs struct {
	Name        string
	Description string
	Privacy     string
	ParentTeam  string
	Members     TeamMemberArgs
}

type TeamsArgs struct {
	Teams []TeamArgs
}

const (
	adminTeamName       = "admins"
	maintainersTeamName = "maintainers"
	developersTeamName  = "developers"
)

var (
	TeamsDefaultArgs = TeamsArgs{
		Teams: []TeamArgs{
			{
				Name:        adminTeamName,
				Description: "Full ccess everywhere",
			},
			{
				Name:        maintainersTeamName,
				Description: "Maintain permissions on all repositories",
			},
			{
				Name:        developersTeamName,
				Description: "Parent team for all developers",
			},
		},
	}
)

func createTeamsSetDefaults(args *TeamsArgs) {
	if args.Teams == nil {
		args.Teams = TeamsDefaultArgs.Teams
		return
	}
	if len(args.Teams) == 0 {
		args.Teams = TeamsDefaultArgs.Teams
		return
	}
	for _, team := range TeamsDefaultArgs.Teams {
		found := false
		for _, t := range args.Teams {
			if team.Name == t.Name {
				found = true
			}
		}
		if !found {
			args.Teams = append(args.Teams, team)
		}
	}
}

func createTeams(ctx *pulumi.Context, provider *github.Provider, argsTeams TeamsArgs) error {
	for _, t := range argsTeams.Teams {
		teamName := util.FormatResourceName(ctx, t.Name)
		team, err := github.NewTeam(ctx, teamName, &github.TeamArgs{
			Name:         pulumi.String(t.Name),
			Description:  pulumi.String(t.Description),
			Privacy:      pulumi.String("closed"),
			ParentTeamId: pulumi.String(t.ParentTeam),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
		teamSettingsName := util.FormatResourceName(ctx, t.Name+" settings")
		_, err = github.NewTeamSettings(ctx, teamSettingsName, &github.TeamSettingsArgs{
			TeamId: team.ID(),
			ReviewRequestDelegation: &github.TeamSettingsReviewRequestDelegationArgs{
				MemberCount: pulumi.Int(1),
				Algorithm:   pulumi.String("LOAD_BALANCE"),
				Notify:      pulumi.Bool(true),
			},
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
		if t.Members != (TeamMemberArgs{}) {
			teamMembersName := util.FormatResourceName(ctx, t.Name+" members")
			_, err = github.NewTeamMembers(ctx, teamMembersName, &github.TeamMembersArgs{
				TeamId: team.ID(),
				Members: func() github.TeamMembersMemberArray {
					var members github.TeamMembersMemberArray
					for _, m := range []TeamMemberArgs{t.Members} {
						members = append(members, &github.TeamMembersMemberArgs{
							Username: pulumi.String(m.Username),
							Role:     pulumi.String(m.Role),
						})
					}
					return members
				}(),
			}, pulumi.Provider(provider))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
