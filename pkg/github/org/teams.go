package org

import (
	"fmt"

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
	Members     []TeamMemberArgs
}

type TeamsArgs struct {
	Teams []TeamArgs
}

const (
	AdminTeamName       = "admins"
	MaintainersTeamName = "maintainers"
	DevelopersTeamName  = "developers"
)

var (
	TeamsDefaultArgs = TeamsArgs{
		Teams: []TeamArgs{
			{
				Name:        AdminTeamName,
				Description: "Full access everywhere",
			},
			{
				Name:        MaintainersTeamName,
				Description: "Maintain permissions on all repositories",
			},
			{
				Name:        DevelopersTeamName,
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
		for i, t := range args.Teams {
			if t.Name == team.Name {
				if t.Description == "" {
					args.Teams[i].Description = team.Description
				}
				if t.Privacy == "" {
					args.Teams[i].Privacy = team.Privacy
				}
				if t.ParentTeam == "" {
					args.Teams[i].ParentTeam = team.ParentTeam
				}
				if t.Members == nil {
					args.Teams[i].Members = team.Members
				}
			}
		}
	}
}

func checkTeamMembersAreMembers(argsTeams TeamsArgs, argsMembers MembersArgs) error {
	for _, t := range argsTeams.Teams {
		if t.Members != nil {
			for _, m := range t.Members {
				found := false
				for _, member := range argsMembers.Members {
					if m.Username == member.Username {
						found = true
						break
					}
				}
				if !found {
					return fmt.Errorf("Team member %s in team %s is not also set to be an organization member", m.Username, t.Name)
				}
			}
		}
	}
	return nil
}

func createTeams(ctx *pulumi.Context, provider *github.Provider, argsTeams TeamsArgs, argsMembers MembersArgs) error {
	err := checkTeamMembersAreMembers(argsTeams, argsMembers)
	if err != nil {
		return err
	}
	for _, t := range argsTeams.Teams {
		teamName := util.FormatResourceName(ctx, "team "+t.Name)
		team, err := github.NewTeam(ctx, teamName, &github.TeamArgs{
			Name:         pulumi.String(t.Name),
			Description:  pulumi.String(t.Description),
			Privacy:      pulumi.String("closed"),
			ParentTeamId: pulumi.String(t.ParentTeam),
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}
		teamSettingsName := util.FormatResourceName(ctx, "team "+t.Name+" settings")
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
		if t.Members != nil {
			teamMembersName := util.FormatResourceName(ctx, "team "+t.Name+" members")
			_, err = github.NewTeamMembers(ctx, teamMembersName, &github.TeamMembersArgs{
				TeamId: team.ID(),
				Members: func() github.TeamMembersMemberArray {
					var members github.TeamMembersMemberArray
					for _, m := range t.Members {
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
